use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day2 {}

fn is_report_safe(report: &[i32]) -> bool {
    let is_asc = report[0] < report[1];

    report.windows(2).all(|levels| {
        let diff = levels[0] - levels[1];
        if is_asc {
            -3 <= diff && diff <= -1
        } else {
            1 <= diff && diff <= 3
        }
    })
}

impl Day2 {
    fn parse_input(&self, input: String) -> Vec<Vec<i32>> {
        input
            .split('\n')
            .filter(|line| !line.is_empty())
            .map(|line| line.split(' ').map(|l| l.parse().unwrap()).collect())
            .collect()
    }

    fn part1(&self, parsed: &Vec<Vec<i32>>) -> usize {
        parsed
            .iter()
            .filter(|report| is_report_safe(report))
            .count()
    }

    fn part2(&self, parsed: Vec<Vec<i32>>) -> usize {
        parsed
            .iter()
            .filter(|report| {
                for i in 0..report.len() {
                    let sliced_report = [&report[..i], &report[i + 1..]].concat();
                    if is_report_safe(&sliced_report) {
                        return true;
                    }
                }
                false
            })
            .count()
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

    assert_eq!(res, 4)
}
