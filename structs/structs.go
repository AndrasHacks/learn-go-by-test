package main

func GetPerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func GetArea(rectange Rectangle) float64 {
	return rectange.Width * rectange.Width
}
