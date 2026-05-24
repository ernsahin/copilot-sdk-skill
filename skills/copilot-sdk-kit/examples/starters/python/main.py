import asyncio
import sys

from copilot import CopilotClient
from copilot.generated.session_events import AssistantMessageData, SessionIdleData
from copilot.session import PermissionRequestResult


def conservative_permission(request, invocation):
    if request.kind.value == "read":
        return PermissionRequestResult(kind="approved")
    print(f"Denied {request.kind.value} request", file=sys.stderr)
    return PermissionRequestResult(kind="denied-no-approval-rule-and-could-not-request-from-user")


async def main():
    prompt = " ".join(sys.argv[1:]) or "Say hello in one short sentence."

    async with CopilotClient() as client:
        async with await client.create_session(
            model="gpt-5",
            on_permission_request=conservative_permission,
        ) as session:
            done = asyncio.Event()

            def on_event(event):
                match event.data:
                    case AssistantMessageData() as data:
                        print(data.content, end="")
                    case SessionIdleData():
                        done.set()

            session.on(on_event)
            await session.send(prompt)
            await done.wait()
            print()


if __name__ == "__main__":
    asyncio.run(main())
