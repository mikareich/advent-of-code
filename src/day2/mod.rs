// https://adventofcode.com/2025/day/2

use std::fs;

type Range = (u64, u64);

const PRIMES: [u8; 54] = [
    2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
    101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193,
    197, 199, 211, 223, 227, 229, 233, 239, 241, 251,
];

const BASE_10: u64 = 10;

fn read_ranges(use_example: bool) -> Option<Vec<Range>> {
    let path = if use_example {
        "src/day2/example-input"
    } else {
        "src/day2/input"
    };

    let input = fs::read_to_string(path).ok()?;
    let input = input.replace("\n", "");

    let mut ranges: Vec<Range> = vec![];
    for raw_range in input.split_terminator(",") {
        let (raw_start, raw_end) = raw_range.split_once("-")?;

        let start = raw_start.parse::<u64>().ok()?;
        let end = raw_end.parse::<u64>().ok()?;

        ranges.push((start, end));
    }

    Some(ranges)
}

fn get_primefactors(product: u8) -> Vec<u8> {
    let mut primefactors: Vec<u8> = vec![];
    for prime in PRIMES {
        if product % prime == 0 {
            primefactors.push(prime)
        }

        if prime > product {
            break;
        }
    }

    primefactors
}

pub fn sum_invalid_ids() -> Option<()> {
    let ranges = read_ranges(false)?;

    let get_length = |w: u64| (w.ilog10() + 1) as u8;
    let get_prefix = |w: u64, i: u8| w / BASE_10.pow((get_length(w) - i) as u32);

    let mut id_sum = 0;

    for range in ranges {
        let (start, end) = range;

        for id in start..=end {
            let length = get_length(id);

            for partition in get_primefactors(length) {
                let prefix = get_prefix(id, length / partition);

                let is_invalid = id
                    == prefix * (BASE_10.pow(length as u32) - 1)
                        / (BASE_10.pow((length / partition) as u32) - 1);

                if is_invalid {
                    id_sum += id;
                    break;
                }
            }
        }
    }

    println!("id sum {id_sum}");

    Some(())
}
