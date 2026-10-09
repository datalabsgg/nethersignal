# NetherSignal

NetherSignal is an experimental signaling backend for game servers using Minecraft Bedrock's NetherNet transport. NetherNet is Minecraft's WebRTC-based transport; NetherSignal is a separate service, not a Minecraft protocol implementation or game server.

For background on NetherNet signaling, see [Mojang's NetherNet HTTP Signaling Partner Onboarding Guide](https://mojang.github.io/bedrock-protocol-docs/guides/nether-net-onboarding-guide/).

## Roadmap

1. **A working server**
   - Implement `Listen` with a configurable address and proper startup errors.
   - Define and validate server configuration.
   - Add graceful shutdown and basic tests.
2. **Core signaling**
   - Implement the required NetherNet signaling endpoints and message flow.
   - Validate requests and return consistent errors.
   - Add integration tests and document how to run a local instance.
3. **API for extra features**
   - Provide a documented extension API for custom routes, middleware, or hooks without changing the core signaling flow.
   - Keep optional features separate from the required protocol behavior.
4. **Proxy-to-proxy connections**
   - Design authenticated communication between NetherSignal instances or proxies.
   - Support forwarding or coordination between proxies, with clear rules for routing and failure handling.
   - Add multi-instance tests and deployment guidance.

The roadmap is directional and may change as the NetherNet requirements and project design become clearer.
