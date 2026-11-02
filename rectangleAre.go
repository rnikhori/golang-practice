package main

import "fmt"

// A struct groups related fields into one named type — here, a rectangle
// is just its width and height bundled together. Lowercase field/type
// names are UNEXPORTED: only code inside this package ("main") can see them.
// That's Go's version of "private" — there's no private keyword, just casing.
type Rectangle struct {
	width  float64
	height float64
}

// This is a METHOD, not a plain function — the receiver "(r *Rectangle)"
// between "func" and the method name is what attaches it to the Rectangle type.
// It's a POINTER receiver (*Rectangle, not Rectangle) because that's the
// production convention: once you might need to mutate a struct via some
// method, ALL methods on that type use pointer receivers for consistency,
// even read-only ones like this. It also avoids copying the struct on
// every call.
func (r *Rectangle) Area() float64 {
	// r is a pointer, but Go lets you write r.width instead of (*r).width —
	// the dot syntax automatically dereferences the pointer for you.
	return r.width * r.height
}

// NewRectangle is a CONSTRUCTOR by convention, not by language feature.
// Go has no constructor keyword — this is just a regular function named
// "NewX", which is the idiomatic pattern every Go programmer recognizes.
// It exists so we can VALIDATE inputs before a Rectangle is ever created,
// something a bare struct literal (Rectangle{width: 5, height: 3}) can't do.
func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		// Returning nil for the pointer and a non-nil error is the standard
		// Go idiom for "construction failed" — there's no exception mechanism.
		return nil, fmt.Errorf("width and height must be positive, got width=%v height=%v", width, height)
	}
	// &Rectangle{...} creates a Rectangle value AND takes its address in one
	// step, so NewRectangle returns a *Rectangle (pointer) as promised.
	return &Rectangle{width: width, height: height}, nil
}

func main() {
	r, err := NewRectangle(5, 3)
	if err != nil {
		// Go has no try/catch — errors are just ordinary return values you
		// check explicitly, immediately after the call that might produce one.
		fmt.Println("failed to create rectangle:", err)
		return
	}

	// r is a *Rectangle here. Calling r.Area() works the same whether r is
	// a pointer or a value — Go automatically handles the pointer dereference.
	fmt.Println("area:", r.Area())
}
