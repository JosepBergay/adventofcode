use std::{collections::HashMap, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day1 {}

impl Day1 {
    fn parse_input(&self, input: String) -> (Vec<i32>, Vec<i32>) {
        let mut left = vec![];
        let mut right = vec![];

        for line in input.split('\n') {
            if line.is_empty() {
                continue;
            }

            for (i, d) in line.split("   ").enumerate() {
                let value = d.parse().expect(format!("Line is {line}").as_str());
                match i {
                    0 => left.push(value),
                    1 => right.push(value),
                    _ => panic!(),
                }
            }
        }

        (left, right)
    }

    fn part1(&self, lists: &(Vec<i32>, Vec<i32>)) -> i32 {
        assert_eq!(lists.0.len(), lists.1.len());

        let mut left = lists.0.clone();
        left.sort();
        let mut right = lists.1.clone();
        right.sort();

        left.iter()
            .enumerate()
            .fold(0, |acc, (i, x)| acc + (x - right[i]).abs())
    }

    fn part2(&self, lists: (Vec<i32>, Vec<i32>)) -> i32 {
        let mut left = HashMap::new();

        for n in lists.0 {
            let count = left.entry(n).or_insert(0);
            *count += 1;
        }

        let mut right = HashMap::new();

        for n in lists.1 {
            let count = right.entry(n).or_insert(0);
            *count += 1;
        }

        left.iter().fold(0, |acc, (k, v)| {
            let value = right.entry(*k).or_default();

            let aux = &*value * v * k;

            acc + aux
        })
    }
}

impl Day for Day1 {
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
fn test_day1_p1() {
    let input = String::from(
        "3   4
4   3
2   5
1   3
3   9
3   3",
    );

    let day = Day1::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 11)
}

#[test]
fn test_day1_p2() {
    let input = String::from(
        "3   4
4   3
2   5
1   3
3   9
3   3",
    );

    let day = Day1::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 31)
}
