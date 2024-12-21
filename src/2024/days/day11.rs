use std::{collections::HashMap, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day11 {}

impl Day11 {
    fn parse_input(&self, input: String) -> HashMap<usize, usize> {
        input
            .lines()
            .take(1)
            .collect::<String>()
            .split(" ")
            .map(|s| (s.parse::<usize>().unwrap(), 1))
            .collect()
    }

    fn part1(&self, parsed: &HashMap<usize, usize>) -> usize {
        blink(25, parsed.clone())
    }

    fn part2(&self, parsed: HashMap<usize, usize>) -> usize {
        blink(75, parsed)
    }
}

fn change_stone(stone: usize) -> (usize, Option<usize>) {
    if stone == 0 {
        return (1, None);
    }

    let mut digits = 0;
    let mut n = stone;
    while n > 0 {
        n /= 10;
        digits += 1;
    }

    if digits % 2 == 0 {
        let pow = 10_usize.pow(digits / 2);

        return (stone / pow, Some(stone % pow));
    }

    (stone * 2024, None)
}

fn blink(iterations: usize, parsed: HashMap<usize, usize>) -> usize {
    let mut dict = parsed;

    for _n in 0..iterations {
        let mut new_dict = HashMap::new();

        for (stone, count) in dict {
            let (first, second) = change_stone(stone);

            let new_count = new_dict.entry(first).or_default();
            *new_count += count;

            if let Some(n) = second {
                let new_count = new_dict.entry(n).or_default();
                *new_count += count;
            }
        }

        dict = new_dict;
    }

    dict.values().sum()
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

    assert_eq!(res, 0)
}
