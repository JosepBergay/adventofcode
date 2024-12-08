use std::{error::Error, fmt, ops};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day4 {}

#[derive(Copy, Clone, PartialEq)]
struct Point2D {
    x: i32,
    y: i32,
}

impl ops::Add for Point2D {
    type Output = Self;

    fn add(self, other: Self) -> Self::Output {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

impl ops::Mul<i32> for Point2D {
    type Output = Self;

    fn mul(self, v: i32) -> Self::Output {
        Self {
            x: self.x * v,
            y: self.y * v,
        }
    }
}

impl fmt::Debug for Point2D {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Point2D ({}, {})", self.x, self.y)
    }
}

struct Map2D<T> {
    map: Vec<Vec<T>>,
}

impl<T> Map2D<T> {
    fn is_out_of_bounds(&self, p: Point2D) -> bool {
        if p.y < 0 || self.map.len() <= p.y.try_into().unwrap() {
            return true;
        }

        p.x < 0 || self.map[0].len() <= p.x.try_into().unwrap()
    }

    fn get(&self, p: Point2D) -> Option<&T> {
        if self.is_out_of_bounds(p) {
            None
        } else {
            let v = &self.map[usize::try_from(p.y).expect(format!("y {p:?}").as_str())]
                [usize::try_from(p.x).expect(format!("x {p:?}").as_str())];
            Some(v)
        }
    }
}

impl Day4 {
    fn parse_input(&self, input: String) -> Map2D<char> {
        let map = input
            .split('\n')
            .filter(|line| !line.is_empty())
            .map(|line| line.chars().collect())
            .collect::<Vec<Vec<char>>>();

        Map2D { map }
    }

    fn part1(&self, parsed: &Map2D<char>) -> usize {
        let dirs: Vec<Point2D> = (-1..=1)
            .flat_map(|x| (-1..=1).map(move |y| Point2D { x, y }))
            .filter(|p| p.x != 0 || p.y != 0)
            .collect();

        let mut total = 0;

        for (y, line) in parsed.map.iter().enumerate() {
            for (x, c) in line.iter().enumerate() {
                if *c != 'X' {
                    continue;
                }

                let curr = Point2D {
                    x: x.try_into().unwrap(),
                    y: y.try_into().unwrap(),
                };

                let count = dirs
                    .iter()
                    .filter_map(|dir| {
                        parsed
                            .get(*dir + curr)
                            .take_if(|c| **c == 'M')
                            .and(parsed.get(*dir * 2 + curr).take_if(|c| **c == 'A'))
                            .and(parsed.get(*dir * 3 + curr).take_if(|c| **c == 'S'))
                    })
                    .count();

                total += count;
            }
        }

        total
    }

    fn part2(&self, _parsed: Map2D<char>) -> &str {
        "TODO"
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
    let input = String::from("");

    let day = Day4::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
