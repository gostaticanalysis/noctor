package a

type T struct {
	N int
	S string
}

func NewT1(n int, s string) *T { // want "unnecessary constructor like function"
	return &T{
		N: n,
		S: s,
	}
}

func NewT2(n int, s string) T { // want "unnecessary constructor like function"
	return T{
		N: n,
		S: s,
	}
}

func NewT3(n int, s string) T { // OK
	n++
	return T{
		N: n,
		S: s,
	}
}

func NewLit1(n int, s string) struct { // want "unnecessary constructor like function"
	N int
	S string
} {
	return struct {
		N int
		S string
	}{
		N: n,
		S: s,
	}
}

func NewLit2(n int, s string) *struct { // want "unnecessary constructor like function"
	N int
	S string
} {
	return &struct {
		N int
		S string
	}{
		N: n,
		S: s,
	}
}
