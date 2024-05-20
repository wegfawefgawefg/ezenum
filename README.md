

# ezenum
golang code gen tool for ez enumz

![Icon](logo/ezenum_logo_icon.png)

## How do

1. Make an ezenum with consts like you normally would.
2. Specify explicit values, and comment each variant with a description.
3. Mark the type as EZENUM with a comment. "// EZENUM"

    ```go
    package responsecodes

    type TestResponseCodes int // EZENUM

    const (
        Continue           TestResponseCodes = 100 // Continue: The client can continue with the request.
        SwitchingProtocols TestResponseCodes = 101 // Switching Protocols: The server understands the request and is asking for a protocol switch to proceed.
        Ok                 TestResponseCodes = 200 // OK: The "request" was successful, and the response contains the requested information.
    )
    ```

4. go get github.com/wegfawefgawefg/ezenum

5. go run github.com/wegfawefgawefg/ezenum

It will recursively look through the folder for go files with EZENUMs inside, and generate an *_ezenum_gen.go in the same directory.
Heres the output from the above code:

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
