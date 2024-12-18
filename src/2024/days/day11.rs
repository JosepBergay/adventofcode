use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day11 {}

impl Day11 {
    fn parse_input(&self, input: String) -> Vec<String> {
        input
            .lines()
            .take(1)
            .collect::<String>()
            .split(" ")
            .map(|s| String::from(s))
            .collect()
    }

    fn part1(&self, parsed: &Vec<String>) -> usize {
        let mut curr = parsed.clone();

        for _n in 0..25 {
            // println!("{n} -> {curr:?}");
            curr = blink(curr);
        }

        curr.len()
    }

    fn part2(&self, _parsed: Vec<String>) -> &str {
        "TODO"
    }
}

fn blink(curr: Vec<String>) -> Vec<String> {
    let mut out = vec![];

    for s in curr {
        if s == "0" {
            out.push(String::from("1"));
        } else if s.len() % 2 == 0 {
            let (first, last) = s.split_at(s.len() / 2);

            out.push(first.to_string());

            let mut last = last.trim_start_matches("0");
            if last == "" {
                last = "0"
            }
            out.push(last.to_string());
        } else {
            let i = s.parse::<u64>().unwrap();

            out.push((i * 2024).to_string());
        }
    }

    out
}

impl Day for Day11 {
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
fn test_day11_p1() {
    let input = String::from(
        "125 17
",
    );

    let day = Day11::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 55312)
}

#[test]
fn test_day11_p2() {
    let input = String::from("");

    let day = Day11::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
