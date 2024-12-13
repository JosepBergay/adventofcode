use std::error::Error;

use super::map2d::Map2D;
use super::point2d::Point2D;
use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day4 {}

impl Day4 {
    fn parse_input(&self, input: String) -> Map2D<char> {
        Map2D::<char>::from_string(input)
    }

    fn part1(&self, parsed: &Map2D<char>) -> usize {
        let dirs: Vec<Point2D> = (-1..=1)
            .flat_map(|x| (-1..=1).map(move |y| Point2D { x, y }))
            .filter(|p| p.x != 0 || p.y != 0)
            .collect();

        let mut total = 0;

        for curr in parsed.iter() {
            let c = parsed.get(curr).unwrap();

            if *c != 'X' {
                continue;
            }

            let count = dirs
                .iter()
                .filter(|dir| {
                    "MAS".chars().enumerate().all(|(i, c)| {
                        let idx = (1 + i).try_into().unwrap();
                        parsed
                            .get(**dir * idx + curr)
                            .filter(|l| **l == c)
                            .is_some()
                    })
                })
                .count();

            total += count;
        }

        total
    }

    fn part2(&self, parsed: Map2D<char>) -> usize {
        let dir1 = vec![Point2D { x: -1, y: -1 }, Point2D { x: 1, y: 1 }];
        let dir2 = vec![Point2D { x: -1, y: 1 }, Point2D { x: 1, y: -1 }];

        let mut total = 0;

        for p in parsed.iter() {
            let c = parsed.get(p).unwrap();

            if *c != 'A' {
                continue;
            }

            let candidate = dir1
                .iter()
                .filter_map(|d| parsed.get(p + *d))
                .collect::<String>();

            let aux = candidate.as_str();

            if aux != "MS" && aux != "SM" {
                continue;
            }

            let candidate = dir2
                .iter()
                .filter_map(|d| parsed.get(p + *d))
                .collect::<String>();

            let aux2 = candidate.as_str();

            if aux2 == "MS" || aux2 == "SM" {
                total += 1;
            }
        }

        total
    }
}

impl Day for Day4 {
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
fn test_day4_p1() {
    let input = String::from(
        "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
",
    );

    let day = Day4::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 18)
}

#[test]
fn test_day4_p2() {
    let input = String::from(
        "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
",
    );

    let day = Day4::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 9)
}
