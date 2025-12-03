// https://adventofcode.com/2025/day/1

use std::fs;

fn parse_document() -> Result<Vec<i16>, ()> {
    let content = match fs::read_to_string("src/day1/input") {
        Err(_) => Err(()),
        Ok(content) => Ok(Vec::from_iter(
            content
                .split_terminator("\n")
                .map(|raw_instruction| {
                    let (direction, distance) = raw_instruction.split_at(1);
                    let distance = distance.parse::<i16>().unwrap();

                    if direction == "L" {
                        -distance
                    } else if direction == "R" {
                        distance
                    } else {
                        panic!("Bad input read: {}", raw_instruction);
                    }
                })
                .into_iter(),
        )),
    };

    content
}

const START_POSITION: i16 = 50;
const NUMBER_OF_DIGITS: i16 = 100;

pub fn get_password() {
    let document = match parse_document() {
        Ok(document) => document,
        Err(_) => panic!("Could not parse input!"),
    };

    let mut position = START_POSITION;
    let mut password = 0;

    for instruction in document {
        let overflow_position = position + instruction;
        let euclid_factor = overflow_position.div_euclid(NUMBER_OF_DIGITS);

        position = overflow_position - euclid_factor * NUMBER_OF_DIGITS;
        password += euclid_factor.abs();
    }

    println!("The password is {password}")
}
