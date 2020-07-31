# go-generate-struct-interface

Generate interface and Get* methods from a struct.


## Install:
```bash
go get -u github.com/GongT/go-generate-struct-interface/cmd/go-generate-struct-interface
```

## Usage:

Add this line at first line:

```go
//go:generate go-generate-struct-interface

package xxxx

type someStruct struct {
	A uint
	B *string
	C *someType
	C someOtherType

	priField string
}
```

and run [go generate](https://blog.golang.org/generate)
