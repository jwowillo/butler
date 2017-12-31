package recipe

import "fmt"

// Fraction is an integer divided by another integer.
type Fraction struct {
	Numerator, Denominator int
}

// IsWhole returns true if the simplified Fraction is a whole number.
func (f Fraction) IsWhole() bool {
	return f.Simplified().Denominator == 1
}

// IsUndefined returns true if the denominator of the Fraction is 0.
func (f Fraction) IsUndefined() bool {
	return f.Denominator == 0
}

// Simplified returns the same Fraction with the greatest common divisor of the
// numerator and denominator divided out.
func (f Fraction) Simplified() Fraction {
	a, b := f.Numerator, f.Denominator
	for b != 0 {
		a, b = b, a%b
	}
	return Fraction{
		Numerator:   f.Numerator / a,
		Denominator: f.Denominator / a,
	}
}

// String representation of the Fraction.
func (f Fraction) String() string {
	f = f.Simplified()
	if f.Denominator == 1 {
		return fmt.Sprintf("%d", f.Numerator)
	}
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}
