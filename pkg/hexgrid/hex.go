package hexgrid

type hex struct {
	q int
	r int
	s int
}

func NewHex(q, r int) hex {
	return hex{q: q, r: r, s: -q - r}
}
