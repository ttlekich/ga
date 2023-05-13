package hexgrid

import "errors"

type hex struct {
	q int
	r int
	s int
}

func NewHex(q, r, s int) (*hex, error) {
	if q+r+s != 0 {
		return nil, errors.New("q + r + s must be 0")
	}
	return &hex{q: q, r: r, s: s}, nil
}

// Equality
func Eq(a *hex, b *hex) bool {
	return a.q == b.q && a.r == b.r && a.s == b.s
}

func NotEq(a *hex, b *hex) bool {
	return !Eq(a, b)
}

// Arithmetic
func Add(a *hex, b *hex) *hex {
	return &hex{a.q + b.q, a.r + b.r, a.s + b.s}
}

func Subtract(a *hex, b *hex) *hex {
	return &hex{a.q - b.q, a.r - b.r, a.s - b.s}
}

func Scale(a *hex, k int) *hex {
	return &hex{a.q * k, a.r * k, a.s * k}
}

// Distance
func Length(h *hex) int {
	return int((Abs(h.q) + Abs(h.r) + Abs(h.s)) / 2)
}

func Distance(a *hex, b *hex) int {
	return Length(Subtract(a, b))
}
