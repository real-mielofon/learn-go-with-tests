package shapes

type Rectangle struct {
	width  float64
	height float64
}

// Perimeter return perimeter of square with width and height
func Perimeter(r Rectangle) float64 {
	return (r.width + r.height) * 2
}

func Area(r Rectangle) float64 {
	return (r.width * r.height)
}
