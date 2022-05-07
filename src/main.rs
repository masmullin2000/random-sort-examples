fn main() {
    let vec = make_random_vec(1_000_000, 100);
    for _ in 0..1000 {
        let mut v = vec.clone();
        //v.sort_unstable();
        v.sort();
    }
}

use rand::prelude::*;

pub fn make_random_vec(sz: usize, modulus: i64) -> Vec<i64> {
    let mut v: Vec<i64> = Vec::new();
    for _ in 0..sz {
        let x: i64 = random();
        v.push(x % modulus);
    }
    v
}
