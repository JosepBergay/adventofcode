use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day2 {}

impl Day2 {
    fn parse_input(&self, input: String) -> Vec<Vec<i32>> {
        input
            .split('\n')
            .filter(|line| !line.is_empty())
            .map(|line| line.split(' ').map(|l| l.parse().unwrap()).collect())
            .collect()
    }

    fn part1(&self, parsed: &Vec<Vec<i32>>) -> i32 {
        parsed.iter().fold(0, |acc, report| {
            let is_asc = report[0] < report[1];

            let is_safe = report.windows(2).all(|levels| {
                let diff = levels[0] - levels[1];
                if is_asc {
                    -3 <= diff && diff <= -1
                } else {
                    1 <= diff && diff <= 3
                }
            });

            acc + if is_safe { 1 } else { 0 }
        })
    }

    fn part2(&self, _parsed: Vec<Vec<i32>>) -> &str {
        "TODO"
    }
}

impl Day for Day2 {
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
fn test_day2_p1() {
    let input = String::from(
        "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
",
    );

    let day = Day2::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 2)
}

#[test]
fn test_day2_p2() {
    let input = String::from(
        "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
",
    );

    let day = Day2::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
