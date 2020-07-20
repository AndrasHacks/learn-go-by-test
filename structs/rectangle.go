package main

type Rectangle struct {
	Width  float64
	Height float64
}

// func (recevierName ReceiverType) MethodName(arguments)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
