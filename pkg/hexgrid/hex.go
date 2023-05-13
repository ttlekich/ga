package hexgrid

type Hex struct {
	q int
	r int
	s int
}

func NewHex(q, r int) Hex {
	return Hex{q: q, r: r, s: -q - r}
}

// Equality
func Eq(a Hex, b Hex) bool {
	return a.q == b.q && a.r == b.r && a.s == b.s
}

func NotEq(a Hex, b Hex) bool {
	return !Eq(a, b)
}

// Arithmetic
func Add(a Hex, b Hex) Hex {
	return NewHex(a.q+b.q, a.r+b.r)
}

func Subtract(a Hex, b Hex) Hex {
	return NewHex(a.q-b.q, a.r-b.r)
}

func Multiply(a Hex, b Hex) Hex {
	return NewHex(a.q*b.q, a.r*b.r)
}

// Distance
func Length(h Hex) int {
	return int((Abs(h.q) + Abs(h.r) + Abs(h.s)) / 2)
}

func Distance(a Hex, b Hex) int {
	return Length(Subtract(a, b))
}
