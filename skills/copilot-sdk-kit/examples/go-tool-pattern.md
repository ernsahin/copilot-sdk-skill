# Go Tool Shape

This file is a shape reference only.

It is not source verification. Verify exact tool APIs, schema conventions, invocation metadata, and trace fields against the installed SDK or upstream Go source before writing final code.

## Conceptual Pattern

Use SDK-native typed tools when the SDK provides them. A good tool has:

1. A narrow name and purpose.
2. A typed input contract.
3. Validation for required and risky arguments.
4. Explicit error returns.
5. Runtime permission checks when the operation can read sensitive data, write, execute, call a network, or mutate an external system.
6. Audit metadata that avoids leaking secrets.

## Review Questions

1. Is the tool scope narrow enough?
2. Are arguments validated before execution?
3. Is the permission policy matched to operation risk?
4. Are failures returned explicitly and recoverably?
5. Does telemetry avoid sensitive content?
6. Can another workflow reuse this tool without inheriting hidden assumptions?
