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

function isUndefined(f) {
  return f.denominator == 0;
}

function isWhole(f) {
  return simplify(f).denominator == 1;
}

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

function add(a, b) {
  const numerator = b.denominator*a.numerator + a.denominator*b.numerator;
  const denominator = a.denominator * b.denominator;
  return simplify(new Fraction(numerator, denominator));
}
