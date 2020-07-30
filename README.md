# go-generate-struct-interface

Generate interface and Get* methods from a struct.


## Usage:
Each file must contains and only contains one struct
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
