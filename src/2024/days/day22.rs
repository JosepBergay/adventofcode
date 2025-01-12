use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day22 {}

impl Day22 {
    fn parse_input(&self, input: String) -> Vec<usize> {
        input.lines().filter_map(|l| l.parse().ok()).collect()
    }

    fn part1(&self, parsed: &Vec<usize>) -> usize {
        parsed.iter().map(|n| compute_secret(*n, 2000)).sum()
    }

    fn part2(&self, _parsed: Vec<usize>) -> &str {
        "TODO"
    }
}

fn compute_secret(base: usize, steps: usize) -> usize {
    let mut curr = base;

    for _ in 0..steps {
        curr = prune_secret_number(mix_secret_number(curr, curr * 64));
        curr = prune_secret_number(mix_secret_number(curr, curr / 32));
        curr = prune_secret_number(mix_secret_number(curr, curr * 2048));
    }

    curr
}

fn mix_secret_number(secret_number: usize, value: usize) -> usize {
    secret_number ^ value
}

fn prune_secret_number(secret_number: usize) -> usize {
    secret_number % 16777216
}

impl Day for Day22 {
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
fn test_day22_p1() {
    let input = String::from(
        "1
10
100
2024
",
    );

    let day = Day22::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 37327623)
}

#[test]
fn test_day22_p2() {
    let input = String::from("");

    let day = Day22::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
