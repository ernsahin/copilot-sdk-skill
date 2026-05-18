import { CopilotClient } from "@github/copilot-sdk";
import type { PermissionRequest, PermissionRequestResult } from "@github/copilot-sdk";

const prompt = process.argv.slice(2).join(" ") || "Say hello in one short sentence.";

function conservativePermission(request: PermissionRequest): PermissionRequestResult {
  if (request.kind === "read") {
    return { kind: "approved" };
  }
  console.error(`Denied ${request.kind} request from ${request.toolName ?? "unknown tool"}`);
  return { kind: "denied-no-approval-rule-and-could-not-request-from-user" };
}

const client = new CopilotClient({ logLevel: "error" });
await client.start();

try {
  const session = await client.createSession({
    model: "gpt-5",
    onPermissionRequest: conservativePermission,
  });

  try {
    const done = new Promise<void>((resolve) => {
      session.on("assistant.message", (event) => {
        process.stdout.write(event.data.content);
      });
      session.on("session.idle", () => resolve());
    });

    await session.send({ prompt });
    await done;
    process.stdout.write("\n");
  } finally {
    await session.disconnect();
  }
} finally {
  await client.stop();
}
