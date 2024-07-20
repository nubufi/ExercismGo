package complexnumbers

import "math"

// Define the Number type here.

type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.a+n2.a,n1.b+n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.a-n2.a,n1.b-n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	a,b,c,d := n1.a,n1.b,n2.a,n2.b
	return Number{a*c-b*d,b*c+a*d}
}

func (n Number) Times(factor float64) Number {
	return Number{factor*n.a,factor*n.b}
}

func (n1 Number) Divide(n2 Number) Number {
	a,b,c,d := n1.a,n1.b,n2.a,n2.b

	newA := (a*c+b*d)/(c*c+d*d)
	newB := (b*c-a*d)/(c*c+d*d)

	return Number{newA,newB}
}

func (n Number) Conjugate() Number {
	return Number{n.a,-n.b}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a+n.b*n.b)
}

func (n Number) Exp() Number {
	expA := math.Exp(n.a)*math.Cos(n.b)
	expB := math.Exp(n.a)*math.Sin(n.b)

	return Number{expA,expB}
}
