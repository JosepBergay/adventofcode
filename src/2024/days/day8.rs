use std::{
    collections::{HashMap, HashSet},
    error::Error,
};

use super::{baseday::DayResult, map2d::Map2D, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day8 {}

impl Day8 {
    fn parse_input(&self, input: String) -> (HashMap<char, Vec<Point2D>>, Map2D<char>) {
        let mut antennas = HashMap::new();

        let map2d = Map2D::new(
            input
                .lines()
                .enumerate()
                .map(|(y, l)| {
                    l.chars()
                        .enumerate()
                        .map(|(x, c)| {
                            if c != '.' {
                                let points = antennas.entry(c).or_insert(vec![]);

                                points.push(Point2D {
                                    x: x as i32,
                                    y: y as i32,
                                });
                            }

                            c
                        })
                        .collect()
                })
                .collect(),
        );

        (antennas, map2d)
    }

    fn part1(&self, parsed: &(HashMap<char, Vec<Point2D>>, Map2D<char>)) -> usize {
        let (antennas, map2d) = parsed;

        let mut antinodes = HashSet::new();

        for (_, points) in antennas {
            if points.len() < 2 {
                continue;
            }

            for (i, &p1) in points.iter().enumerate() {
                for &p2 in points[(i + 1)..].iter() {
                    let v = p2 - p1;
                    let a1 = p2 + v;
                    let a2 = p1 - v;

                    for a in [a1, a2] {
                        if !map2d.is_out_of_bounds(a) {
                            antinodes.insert(a);
                        }
                    }
                }
            }
        }

        antinodes.len()
    }

    fn part2(&self, parsed: (HashMap<char, Vec<Point2D>>, Map2D<char>)) -> usize {
        let (antennas, map2d) = parsed;

        let mut antinodes = HashSet::new();

        for (_, points) in antennas {
            if points.len() < 2 {
                continue;
            }

            for (i, &p1) in points.iter().enumerate() {
                for &p2 in points[(i + 1)..].iter() {
                    let v = p2 - p1;

                    let mut curr = p1;
                    while !map2d.is_out_of_bounds(curr) {
                        antinodes.insert(curr);
                        curr += v;
                    }
                    curr = p2;
                    while !map2d.is_out_of_bounds(curr) {
                        antinodes.insert(curr);
                        curr -= v;
                    }
                }
            }
        }

        antinodes.len()
    }
}

impl Day for Day8 {
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
fn test_day8_p1() {
    let input = String::from(
        "............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
",
    );

    let day = Day8::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 14)
}

#[test]
fn test_day8_p2() {
    let input = String::from(
        "............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
",
    );

    let day = Day8::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 34)
}
