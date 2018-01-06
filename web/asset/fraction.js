// Fraction is an integer divided by another integer.
class Fraction {

  constructor(numerator, denominator) {
    this.numerator = numerator;
    this.denominator = denominator;
  }

  toString() {
    const f = simplify(this);
    if (f.denominator == 1) {
      return f.numerator.toString();
    }
    return f.numerator.toString() + '/' + f.denominator.toString();
  }

}

// isUndefined returns true if the denominator of the Fraction is 0.
function isUndefined(f) {
  return f.denominator == 0;
}

// isWhole returns true if the simplified Fraction is a whole number.
function isWhole(f) {
  return simplify(f).denominator == 1;
}

// simplify the Fraction by dividing the greatest common divisor of the
// numerator and denominator out.
function simplify(f) {
  let a = f.numerator;
  let b = f.denominator;
  while (b != 0) {
    const t = a;
    a = b;
    b = t % b;
  }
  return new Fraction(f.numerator/a, f.denominator/a);
}

// add the two Fractions and return the simplified form.
function add(a, b) {
  const numerator = b.denominator*a.numerator + a.denominator*b.numerator;
  const denominator = a.denominator * b.denominator;
  return simplify(new Fraction(numerator, denominator));
}
