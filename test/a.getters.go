// +build !windows

package xxxx

type ReadOnlySomeType interface {
	GetSub() string
}

func (self someType) GetSub() string{
	return self.Sub
}

type ReadOnlySomeStruct interface {
	GetU() uint
	GetU1() uint
	GetU2() uint
	GetU3() uint
	GetStrPtr() *string
	GetOtherInterface() someInterface
	GetOtherStruct() someType
	GetOtherStructPtr() *someType
	GetChannel() <- chan bool
	GetArray() []byte
	GetArrayFixed() [4]byte
	GetMap() map[string]bool
	GetComplexMap1() map[*someType]struct{ i int }
	GetComplexMap2() map[struct{ i int }]*someType
	GetPriField() string
	GetPriFieldNoGet() string
}

func (self someStruct) GetU() uint{
	return self.U
}

func (self someStruct) GetU1() uint{
	return self.U1
}

func (self someStruct) GetU2() uint{
	return self.U2
}

func (self someStruct) GetU3() uint{
	return self.u3
}

func (self someStruct) GetStrPtr() *string{
	return self.StrPtr
}

func (self someStruct) GetOtherInterface() someInterface{
	return self.OtherInterface
}

func (self someStruct) GetOtherStruct() someType{
	return self.OtherStruct
}

func (self someStruct) GetOtherStructPtr() *someType{
	return self.OtherStructPtr
}

func (self someStruct) GetChannel() <- chan bool{
	return self.Channel
}

func (self someStruct) GetArray() []byte{
	return self.Array
}

func (self someStruct) GetArrayFixed() [4]byte{
	return self.ArrayFixed
}

func (self someStruct) GetMap() map[string]bool{
	return self.Map
}

func (self someStruct) GetComplexMap1() map[*someType]struct{ i int }{
	return self.ComplexMap1
}

func (self someStruct) GetComplexMap2() map[struct{ i int }]*someType{
	return self.ComplexMap2
}

func (self someStruct) GetPriField() string{
	return self.priField
}

func (self someStruct) GetPriFieldNoGet() string{
	return self.priFieldNoGet
}

