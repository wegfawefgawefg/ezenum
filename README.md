# ezenum
golang code gen tool for ez enumz

## How do
1. Make an ezenum with consts like you normally would. 
Specify explicit values, and commend each variant with a description.
Mark the type as EZENUM.
```go
package responsecodes

type TestResponseCodes int // EZENUM

const (
	Continue           TestResponseCodes = 100 // Continue: The client can continue with the request.
	SwitchingProtocols TestResponseCodes = 101 // Switching Protocols: The server understands the request and is asking for a protocol switch to proceed.
	Ok                 TestResponseCodes = 200 // OK: The "request" was successful, and the response contains the requested information.
)
```

2. go get github.com/wegfawefgawefg/ezenum

3. go run github.com/wegfawefgawefg/ezenum

It will recursively look through the folder for go files with EZENUMs inside, and generate an *_ezenum_gen.go in the same directory.
Heres some sample output:

```go
package responsecodes

func (r TestResponseCodes) AsCode() int {
	return int(r)
}

func (r TestResponseCodes) GetDescription() string {
	switch r {
	case SwitchingProtocols:
		return "Switching Protocols: The server understands the request and is asking for a protocol switch to proceed."
	case Ok:
		return "OK: The \"request\" was successful, and the response contains the requested information."
	case Continue:
		return "Continue: The client can continue with the request."
	default:
		return "Unknown Response"
	}
	return "Unknown Response"
}

func IsValidTestResponseCodes(code int) bool {
	switch TestResponseCodes(code) {
	case SwitchingProtocols:
	case Ok:
	case Continue:
		return true
	default:
		return false
	}
	return false
}
```

## Features
- will escape strings 4 u

## Unfeatures
- not sure if it can handle string values...
- not sure if it can handle some missing description comments...
