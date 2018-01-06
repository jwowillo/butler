package recipe

import "fmt"

// Fraction is an integer divided by another integer.
type Fraction struct {
	Numerator, Denominator int
}

// IsWhole returns true if the simplified Fraction is a whole number.
func IsWhole(f Fraction) bool {
	return Simplify(f).Denominator == 1
}

// IsUndefined returns true if the denominator of the Fraction is 0.
func IsUndefined(f Fraction) bool {
	return f.Denominator == 0
}

// Simplify the Fraction by dividing the greatest common divisor of the
// numerator and denominator out.
func Simplify(f Fraction) Fraction {
	a, b := f.Numerator, f.Denominator
	for b != 0 {
		a, b = b, a%b
	}
	return Fraction{
		Numerator:   f.Numerator / a,
		Denominator: f.Denominator / a,
	}
}

func (f Fraction) String() string {
	f = Simplify(f)
	if f.Denominator == 1 {
		return fmt.Sprintf("%d", f.Numerator)
	}
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}
