package builder

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	b Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		b: builder,
	}
}

// Construct Product
func (d *Director) Construct() {
	d.b.Part1()
	d.b.Part2()
	d.b.Part3()
}

type Builder1 struct {
	s string
}

func (b *Builder1) Part1() {
	b.s += "1"
}

func (b *Builder1) Part2() {
	b.s += "2"
}

func (b *Builder1) Part3() {
	b.s += "3"
}
