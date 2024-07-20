// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
    Iso = "Iso"// isosceles
    Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else {
		k = Sca
	}


	return k
}
