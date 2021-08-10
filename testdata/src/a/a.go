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

func NewLit3(n int, s string) struct { // OK
	n int
	S string
} {
	return struct {
		n int
		S string
	}{
		n: n,
		S: s,
	}
}

func NewLit4(n int, s string) struct { // OK
	T
} {
	return struct {
		T
	}{T{
		N: n,
		S: s,
	}}
}

func NewLit5(n int, s string) struct { // OK
	*T
} {
	return struct {
		*T
	}{&T{
		N: n,
		S: s,
	}}
}

type t = T

func NewLit6(n int, s string) struct { // OK
	t
} {
	return struct {
		t
	}{t{
		N: n,
		S: s,
	}}
}

type Pointer struct {
	X int
	Y int
}

func Pt(x, y int) Pointer { // OK - short name
	return Pointer{
		X: x,
		Y: y,
	}
}
