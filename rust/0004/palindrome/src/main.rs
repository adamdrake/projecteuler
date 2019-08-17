#![feature(test)]

extern crate test;
use test::Bencher;

// A palindromic number is one which reads the same backwards and forwards.
// Examples of palindromic numbers include 123321 and 12321.
// Problem: find the largest palindromic number that is a product of two 3-digit numbers.

// Solution:

// Start with the answer in order to verify and benchmark possible solutions
const ANSWER: i32 = 906609;

// We will need a helper function to determine if a number is palindromic
//
// is_palindromic_v1() will convert the number to a string, reverse the string, and compare that with the
// string representation of the number.
fn is_palindromic_v1(i: i32) -> bool {
    return i.to_string().chars().rev().collect::<String>() == i.to_string();
}

// is_palindromic_v2() will build up a new number `reverse` by shifting `reverse` one decimal place
// and adding the 10s place of the candidate number until no places are left in the candidate
// number.  In this way, `reverse` is constructed as the successive tens place numbers.
// from the original candidate number.  This version is ~16x faster than is_palindromic_v1().
fn is_palindromic_v2(i: i32) -> bool {
    let mut reversed = 0;
    let mut number = i;

    while number > 0 {
        reversed = reversed * 10 + number % 10;
        number = number / 10;
    }
    return reversed == i || i == number / 10;
}

// v0() is the naive attempt: multiply three-digit numbers until palindrome is found
// if it's larger than the most recent palindrom found then keep it.  Iterate until the end.
fn v0() -> i32 {
    let mut res = 0;
    for i in 100..=999 {
        for j in 100..=999 {
            let prod = i * j;
            if is_palindromic_v1(prod) && prod > res {
                res = prod;
            }
        }
    }
    return res;
}

// v1(): FUN FACT - Rust has short-circuit evaluation.
fn v1() -> i32 {
    let mut res = 0;
    for i in 100..=999 {
        for j in 100..=999 {
            let prod = i * j;
            // Note that this order changed.
            // Consequently, the right-hand side of the && is only evaluated if needed.
            // This is important since string casting and reversing is so slow and the search
            // space is so large.  Many languages have short-circuit/lazy evalution.
            // This matters less as we start restricting the search space.
            if prod > res && is_palindromic_v1(prod) {
                res = prod;
            }
        }
    }
    return res;
}

// v2(): Start from the biggest numbers, which makes a lot more sense.
fn v2() -> i32 {
    let mut res = 0;
    for i in (100..=999).rev() {
        for j in (100..=999).rev() {
            let prod = i * j;

            if prod > res && is_palindromic_v1(prod) {
                res = prod;
                break;
            }
        }
    }
    return res;
}

// v3(): Now stop if result is greater than current product.  If you
// are multiplying and the product is smaller than the current number, you won't be
// able to get around the fact that p * q = q * p, so you can stop.
fn v3() -> i32 {
    let mut res = 0;
    for i in (100..=999).rev() {
        for j in (100..=999).rev() {
            let prod = i * j;

            if prod > res && is_palindromic_v1(prod) {
                res = prod;
                break;
            } else if res > prod {
                break;
            }
        }
    }
    return res;
}

// v4(): Note that the number is probably 6 digits, i.e. abccba = p * q where each letter is some number between 0 and 9
// then we can rewrite as a * 100000 + b * 10000 + c * 1000 + c * 100 + b * 10 + a * 1 = p * q
// which we can rewrite as a * 100001 + b * 10010 + c * 1100 = p * q
// which can be factored to 11 * (a * 9091 + b * 910 + c * 100) = p * q
// therefore p * q must be divisible by 11
// therefore either p or q must be divisible by 11.
//
// Since p or q must be divisible by 11, just start the iteration at 990 because it is
// the largest possible number in the given range that is also divisible by 11 and
// work our way down in steps of 11 versus checking divisibility with modulus operator
// since the modulus operator must do multiple arithmetic operations and is comparatively slow.
fn v4() -> i32 {
    let mut res = 0;
    for i in (100..=990).rev().step_by(11) {
        for j in (100..=999).rev() {
            let prod = i * j;
            if prod > res && is_palindromic_v1(prod) {
                res = prod;
                break;
            } else if res > prod {
                break;
            }
        }
    }
    return res;
}

// v5(): Get rid of the slow string reversing and use the numerical palindrome helper function.
fn v5() -> i32 {
    let mut res = 0;
    for i in (100..=990).rev().step_by(11) {
        for j in (100..=999).rev() {
            let prod = i * j;

            if prod > res && is_palindromic_v2(prod) {
                res = prod;
                break;
            } else if res > prod {
                break;
            }
        }
    }
    return res;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_v0() {
        assert_eq!(v0(), ANSWER);
    }

    #[test]
    fn test_v1() {
        assert_eq!(v1(), ANSWER);
    }

    #[test]
    fn test_v2() {
        assert_eq!(v2(), ANSWER);
    }

    #[test]
    fn test_v3() {
        assert_eq!(v3(), ANSWER);
    }

    #[test]
    fn test_v4() {
        assert_eq!(v4(), ANSWER);
    }

    #[test]
    fn test_v5() {
        assert_eq!(v5(), ANSWER);
    }
}

mod benchmarks {
    use super::*;

    #[bench]
    fn bench_is_palindromic_v1(b: &mut Bencher) {
        b.iter(|| {
            let n = test::black_box(998899);
            is_palindromic_v1(n)
        })
    }

    #[bench]
    fn bench_is_palindromic_v2(b: &mut Bencher) {
        b.iter(|| {
            let n = test::black_box(998899);
            is_palindromic_v2(n)
        })
    }

    #[bench]
    fn bench_v0(b: &mut Bencher) {
        b.iter(|| v0())
    }

    #[bench]
    fn bench_v1(b: &mut Bencher) {
        b.iter(|| v1())
    }

    #[bench]
    fn bench_v2(b: &mut Bencher) {
        b.iter(|| v2())
    }

    #[bench]
    fn bench_v3(b: &mut Bencher) {
        b.iter(|| v3())
    }

    #[bench]
    fn bench_v4(b: &mut Bencher) {
        b.iter(|| v4())
    }

    #[bench]
    fn bench_v5(b: &mut Bencher) {
        b.iter(|| v5())
    }
}
