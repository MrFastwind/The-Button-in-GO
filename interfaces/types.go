package interfaces
type Float interface{
	~float32 | ~float64
}

type Index interface{
	~uint|~uint8|~uint16|~uint32|~uint64
}