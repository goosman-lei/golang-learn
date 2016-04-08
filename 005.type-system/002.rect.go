package main

import "fmt"
import "strconv"

type Rect struct {
	x, y          int
	width, height int
}

func (r Rect) Area() int {
	return r.width * r.height
}

func (r Rect) Position() string {
	return "[" + strconv.Itoa(r.x) + "," + strconv.Itoa(r.y) + "]"
}

func (r *Rect) AddHeight(incr int) {
	r.height += incr
}

func main() {
	r := new(Rect)

	rRefCopy := r

	r.x = 103
	r.y = 372
	r.width = 200
	r.height = 300

	rValCopy := *r

	r.AddHeight(10)

	fmt.Println("original", r, "position:", r.Position(), "area:", r.Area())
	fmt.Println("refcopy", rRefCopy, "position:", rRefCopy.Position(), "area:", rRefCopy.Area())
	fmt.Println("valcopy", rValCopy, "position:", rValCopy.Position(), "area:", rValCopy.Area())
}
