package responsecodes

type TestResponseCodes int // EZENUM

const (
	Continue           TestResponseCodes = 100 // Continue: The client can continue with the request.
	SwitchingProtocols TestResponseCodes = 101 // Switching Protocols: The server understands the request and is asking for a protocol switch to proceed.
	Ok                 TestResponseCodes = 200 // OK: The "request" was successful, and the response contains the requested information.
)
