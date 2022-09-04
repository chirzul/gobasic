package task

type Calculate2D interface {
	Area() float64
	Perimeter() float64
}

type Calculate3D interface {
	Volume() float64
}

type Calculate interface {
	Calculate2D
	Calculate3D
}

type Cube struct {
	Side float64
}

func (c *Cube) Area() float64 {
	return c.Side * c.Side
}

func (c *Cube) Perimeter() float64 {
	return 4 * c.Side
}

func (c *Cube) Volume() float64 {
	return c.Side * c.Side * c.Side
}
