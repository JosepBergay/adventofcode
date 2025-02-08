use std::{collections::HashSet, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day25 {}

type Input = (Vec<Vec<usize>>, Vec<Vec<usize>>);

impl Day25 {
    fn parse_input(&self, input: String) -> Input {
        let mut keys = vec![];
        let mut locks = vec![];

        for l in input.split("\n\n").filter(|l| !l.is_empty()) {
            let is_lock = l.starts_with("#");

            let lines = l.lines().collect::<Vec<_>>();

            let mut acc = vec![0; lines[0].len()];

            for x in 0..lines[0].len() {
                let count = lines[1..lines.len() - 1]
                    .iter()
                    .map(|line| line.as_bytes()[x] as char)
                    .filter(|c| *c == '#')
                    .count();

                acc[x] = count;
            }

            if is_lock {
                locks.push(acc)
            } else {
                keys.push(acc)
            };
        }

        (keys, locks)
    }

    fn part1(&self, parsed: &Input) -> usize {
        let (keys, locks) = parsed;

        let mut count = 0;

        for key in keys {
            for lock in locks {
                if key.iter().zip(lock).all(|(k, l)| k + l <= 5) {
                    count += 1;
                }
            }
        }

        count
    }

    fn part2(&self, parsed: Input) -> &str {
        "TODO"
    }
}

impl Day for Day25 {
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
fn test_day25_p1() {
    let input = String::from(
        "#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
",
    );

    let day = Day25::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 3)
}

#[test]
fn test_day25_p2() {
    let input = String::from("");

    let day = Day25::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
