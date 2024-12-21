use std::{
    collections::{HashSet, VecDeque},
    error::Error,
};

use crate::days::map2d::Map2D;

use super::{baseday::DayResult, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day12 {}

impl Day12 {
    fn parse_input(&self, input: String) -> Map2D<char> {
        Map2D::<char>::from_string(input)
    }

    fn part1(&self, map: &Map2D<char>) -> usize {
        let mut visited = HashSet::new();
        let dirs = vec![
            Point2D { x: 1, y: 0 },
            Point2D { x: -1, y: 0 },
            Point2D { x: 0, y: 1 },
            Point2D { x: 0, y: -1 },
        ];
        let mut price = 0;
        let mut to_explore = vec![Point2D { x: 0, y: 0 }];

        while let Some(start) = to_explore.pop() {
            if visited.contains(&start) {
                continue;
            }

            let plant = map.get(start).unwrap();
            let mut area = 0;
            let mut perimeter = 0;
            let mut q = VecDeque::from([start]);

            while let Some(curr) = q.pop_front() {
                if visited.contains(&curr) {
                    continue;
                }

                visited.insert(curr);
                area += 1;

                for d in &dirs {
                    let p = *d + curr;

                    match map.get(p) {
                        None => {
                            perimeter += 1;
                        }
                        Some(c) => {
                            if c == plant {
                                q.push_back(p);
                            } else {
                                perimeter += 1;
                                to_explore.push(p);
                            }
                        }
                    }
                }
            }

            price += area * perimeter;
        }

        price
    }

    fn part2(&self, _parsed: Map2D<char>) -> usize {
        0
    }
}

impl Day for Day12 {
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
fn test_day12_p1() {
    let input = String::from(
        "RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
",
    );

    let day = Day12::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 1930)
}

#[test]
fn test_day12_p2() {
    let input = String::from("");

    let day = Day12::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 0)
}
