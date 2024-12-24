use std::error::Error;

use regex::Regex;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day14 {}

impl Day14 {
    fn parse_input(&self, input: String) -> Vec<[isize; 4]> {
        let re = Regex::new(r"p=(\d+),(\d+) v=(-?\d+),(-?\d+)").unwrap();

        let mut robots = vec![];

        for capture in re.captures_iter(&input) {
            let (_, digits) = capture.extract();

            robots.push(digits.map(|d| d.parse::<isize>().unwrap()));
        }

        robots
    }

    fn part1(&self, parsed: &mut Vec<[isize; 4]>, width: isize, height: isize) -> usize {
        let seconds = 100;

        let robots = parsed;

        for _ in 0..seconds {
            for robot in &mut *robots {
                let [px, py, vx, vy] = robot;

                *px = (*px + *vx + width) % width;
                *py = (*py + *vy + height) % height;
            }
        }

        let mut quadrant_count = vec![0, 0, 0, 0];

        let mid_x = (width - 1) / 2;
        let mid_y = (height - 1) / 2;

        for [px, py, _, _] in &mut *robots {
            if *px < mid_x && *py < mid_y {
                quadrant_count[0] += 1;
            } else if *px > mid_x && *py < mid_y {
                quadrant_count[1] += 1;
            } else if *px < mid_x && *py > mid_y {
                quadrant_count[2] += 1;
            } else if *px > mid_x && *py > mid_y {
                quadrant_count[3] += 1;
            }
        }

        quadrant_count.iter().product()
    }

    fn part2(&self, _parsed: Vec<[isize; 4]>) -> usize {
        0
    }
}

impl Day for Day14 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let mut parsed = self.parse_input(input);

        let p1 = self.part1(&mut parsed, 101, 103);
        let p2 = self.part2(parsed);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[test]
fn test_day14_p1() {
    let input = String::from(
        "p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
",
    );

    let day = Day14::default();
    let mut parsed = day.parse_input(input);
    let res = day.part1(&mut parsed, 11, 7);

    assert_eq!(res, 12)
}

#[test]
fn test_day14_p2() {
    let input = String::from("");

    let day = Day14::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 0)
}
