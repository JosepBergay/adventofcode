use std::{error::Error, usize};

use regex::Regex;

use super::{baseday::DayResult, Day};

#[derive(Debug)]
struct Machine {
    a: (f64, f64),
    b: (f64, f64),
    goal: (f64, f64),
}

#[derive(Default)]
pub struct Day13 {}

impl Day13 {
    fn parse_input(&self, input: String) -> Vec<Machine> {
        let re = Regex::new(
            r"Button A: X\+(\d+), Y\+(\d+)
Button B: X\+(\d+), Y\+(\d+)
Prize: X=(\d+), Y=(\d+)",
        )
        .unwrap();

        let mut machines = vec![];

        for capture in re.captures_iter(&input) {
            let (_, digits) = capture.extract();

            let [ax, ay, bx, by, px, py] = digits.map(|d| d.parse::<f64>().unwrap());

            machines.push(Machine {
                a: (ax, ay),
                b: (bx, by),
                goal: (px, py),
            });
        }

        machines
    }

    fn part1(&self, parsed: &Vec<Machine>) -> usize {
        let mut total = 0;

        for machine in parsed {
            total += get_buttons_pushes_cost(machine, 0);
        }

        total
    }

    fn part2(&self, parsed: Vec<Machine>) -> usize {
        let diff = 10000000000000;

        let mut total = 0;

        for machine in parsed {
            total += get_buttons_pushes_cost(&machine, diff);
        }

        total
    }
}

fn get_buttons_pushes_cost(machine: &Machine, diff: usize) -> usize {
    let goal_x = machine.goal.0 + diff as f64;
    let goal_y = machine.goal.1 + diff as f64;

    // Through equation solving we know that:
    let b = (goal_y * machine.a.0 - machine.a.1 * goal_x)
        / (machine.b.1 * machine.a.0 - machine.a.1 * machine.b.0);

    if b % 1.0 != 0.0 {
        return 0;
    }

    let a = (goal_x - machine.b.0 * b) / machine.a.0;

    if a % 1.0 != 0.0 {
        return 0;
    }

    (a as usize) * 3 + (b as usize)
}

impl Day for Day13 {
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
fn test_day13_p1() {
    let input = String::from(
        "Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
",
    );

    let day = Day13::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 480)
}

#[test]
fn test_day13_p2() {
    let input = String::from(
        "Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
",
    );

    let day = Day13::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 0)
}
