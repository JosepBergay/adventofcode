use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day7 {}

impl Day7 {
    fn parse_input(&self, input: String) -> Vec<(i64, Vec<i64>)> {
        input
            .lines()
            .filter_map(|l| {
                l.split_once(": ").and_then(|(res, ops)| {
                    Some((
                        res.parse::<i64>().unwrap(),
                        ops.split(" ")
                            .map(|op| op.parse::<i64>().unwrap())
                            .collect(),
                    ))
                })
            })
            .collect()
    }

    fn part1(&self, parsed: &Vec<(i64, Vec<i64>)>) -> i64 {
        let mut sum = 0;

        for (res, ops) in parsed {
            // Input works both for 0 and 1.
            if evaluate_p1(res, ops, 0) {
                sum += res;
            }
        }

        sum
    }

    fn part2(&self, parsed: &Vec<(i64, Vec<i64>)>) -> i64 {
        let mut sum = 0;

        for (res, ops) in parsed {
            // Acc start value is arbitrary and should be determined by each operation:
            // (+ 0, * 1, || 0). Input in this case just works for 1.
            if evaluate_p2(res, ops, 1) {
                sum += res;
            }
        }

        sum
    }
}

fn evaluate_p1(res: &i64, ops: &[i64], acc: i64) -> bool {
    if ops.is_empty() {
        return acc == *res;
    }

    let op = ops.first().unwrap();

    evaluate_p1(res, &ops[1..], acc + op) || evaluate_p1(res, &ops[1..], acc * op)
}

fn evaluate_p2(res: &i64, ops: &[i64], acc: i64) -> bool {
    if ops.is_empty() {
        return acc == *res;
    }

    let op = ops.first().unwrap();

    evaluate_p2(res, &ops[1..], acc + op)
        || evaluate_p2(res, &ops[1..], acc * op)
        || evaluate_p2(
            res,
            &ops[1..],
            (acc.to_string() + &op.to_string()).parse::<i64>().unwrap(),
        )
}

impl Day for Day7 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&parsed);
        let p2 = self.part2(&parsed);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[test]
fn test_day7_p1() {
    let input = String::from(
        "190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
",
    );

    let day = Day7::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 3749)
}

#[test]
fn test_day7_p2() {
    let input = String::from(
        "190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
",
    );

    let day = Day7::default();
    let parsed = day.parse_input(input);
    let res = day.part2(&parsed);

    assert_eq!(res, 11387)
}
