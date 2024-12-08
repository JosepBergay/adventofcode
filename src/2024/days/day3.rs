use std::error::Error;

use regex::Regex;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day3 {}

impl Day3 {
    fn parse_input(&self, input: String) -> String {
        input
    }

    fn part1(&self, parsed: &String) -> i32 {
        let re = Regex::new(r"mul\(([0-9]{1,3}),([0-9]{1,3})\)").expect("Invalid regex");

        let mut total = 0;

        for c in re.captures_iter(parsed) {
            let (_, digits) = c.extract();

            let [a, b] = digits.map(|d| d.parse::<i32>().unwrap());

            total += a * b
        }

        total
    }

    fn part2(&self, parsed: String) -> i32 {
        let mut total = 0;

        let mut is_enabled = true;
        let mut instructions = parsed.as_str();

        while instructions.len() > 0 {
            let delimiter = if is_enabled { "don't()" } else { "do()" };
            let found = instructions.split_once(delimiter);
            match found {
                None => {
                    if is_enabled {
                        total += self.part1(&String::from(instructions));
                    }
                    break;
                }
                Some((pre, post)) => {
                    if is_enabled {
                        total += self.part1(&String::from(pre));
                    }
                    is_enabled = !is_enabled;
                    instructions = post;
                }
            }
        }

        total
    }
}

impl Day for Day3 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&parsed);
        let p2 = self.part2(parsed);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[test]
fn test_day3_p1() {
    let input =
        String::from("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))");

    let day = Day3::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 161)
}

#[test]
fn test_day3_p2() {
    let input =
        String::from("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))");

    let day = Day3::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 48)
}
