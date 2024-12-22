use std::{
    collections::{HashSet, VecDeque},
    error::Error,
};

use crate::days::map2d::Map2D;

use super::{baseday::DayResult, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day12 {}

impl Day12 {
    fn parse_input(&self, input: String) -> (usize, usize) {
        let map = Map2D::<char>::from_string(input);

        let mut visited = HashSet::new();
        let dirs = vec![
            Point2D { x: 1, y: 0 },
            Point2D { x: -1, y: 0 },
            Point2D { x: 0, y: 1 },
            Point2D { x: 0, y: -1 },
        ];
        let mut price_p1 = 0;
        let mut price_p2 = 0;
        let mut to_explore = vec![Point2D { x: 0, y: 0 }];

        while let Some(start) = to_explore.pop() {
            if visited.contains(&start) {
                continue;
            }

            let plant = map.get(start).unwrap();
            let mut area = 0;
            let mut perimeter = 0;
            let mut corners = 0;

            let mut q = VecDeque::from([start]);

            while let Some(curr) = q.pop_front() {
                if visited.contains(&curr) {
                    continue;
                }

                visited.insert(curr);
                area += 1;
                corners += count_corners(curr, &map);

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

            price_p1 += area * perimeter;
            price_p2 += area * corners;
        }

        (price_p1, price_p2)
    }

    fn part1(&self, prices: &(usize, usize)) -> usize {
        prices.0
    }

    fn part2(&self, prices: (usize, usize)) -> usize {
        prices.1
    }
}

fn count_corners(curr: Point2D, map: &Map2D<char>) -> usize {
    let dirs = vec![
        Point2D { x: 0, y: -1 },
        Point2D { x: 1, y: 0 },
        Point2D { x: 0, y: 1 },
        Point2D { x: -1, y: 0 },
        Point2D { x: 0, y: -1 },
    ];

    let plant = map.get(curr);

    let mut count = 0;

    for pairs in dirs.windows(2) {
        let dir1 = pairs[0];
        let dir2 = pairs[1];

        let p1 = map.get(curr + dir1);
        let p2 = map.get(curr + dir2);
        let diag = map.get(curr + dir1 + dir2);

        if (p1 != plant && p2 != plant) || (p1 == plant && p2 == plant && diag != plant) {
            count += 1;
        }
    }

    count
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
    let res = day.part2(parsed);

    assert_eq!(res, 1206)
}
