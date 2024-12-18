use std::{
    collections::{HashMap, HashSet},
    error::Error,
};

use super::{baseday::DayResult, map2d::Map2D, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day10 {}

impl Day10 {
    fn parse_input(&self, input: String) -> (Vec<Point2D>, Vec<Point2D>, Map2D<u32>) {
        let mut heads = vec![];
        let mut tails = vec![];

        let map = Map2D::new(
            input
                .lines()
                .filter(|l| !l.is_empty())
                .enumerate()
                .map(|(y, l)| {
                    l.chars()
                        .enumerate()
                        .map(|(x, c)| {
                            let i = c.to_digit(10).unwrap();

                            if i == 0 {
                                heads.push(Point2D {
                                    x: x.try_into().unwrap(),
                                    y: y.try_into().unwrap(),
                                });
                            } else if i == 9 {
                                tails.push(Point2D {
                                    x: x.try_into().unwrap(),
                                    y: y.try_into().unwrap(),
                                });
                            }
                            i
                        })
                        .collect()
                })
                .collect(),
        );

        (heads, tails, map)
    }

    fn part1(&self, parsed: &(Vec<Point2D>, Vec<Point2D>, Map2D<u32>)) -> usize {
        let (heads, _, map) = parsed;

        let mut count = 0;

        for start in heads {
            let mut seen = HashSet::new();

            count += walk_trail(*start, map, &mut seen);
        }

        count
    }

    fn part2(&self, parsed: (Vec<Point2D>, Vec<Point2D>, Map2D<u32>)) -> usize {
        let (heads, _, map) = parsed;

        let mut count = 0;

        for start in heads {
            let mut map_count = HashMap::new();

            walk_trail_2(start, &map, &mut map_count);

            count += map_count.values().fold(0, |acc, i| acc + i);
        }

        count
    }
}

fn walk_trail_2(curr: Point2D, map: &Map2D<u32>, map_count: &mut HashMap<Point2D, usize>) -> bool {
    if map.get(curr).is_some_and(|v| *v == 9) {
        return true;
    }

    for node in get_next_nodes(curr, map) {
        if walk_trail_2(node, map, map_count) {
            let count = map_count.entry(node).or_insert(0);
            *count += 1;
        }
    }

    return false;
}

fn get_next_nodes(curr: Point2D, map: &Map2D<u32>) -> Vec<Point2D> {
    let dirs = vec![
        Point2D { x: 1, y: 0 },
        Point2D { x: -1, y: 0 },
        Point2D { x: 0, y: 1 },
        Point2D { x: 0, y: -1 },
    ];

    let i = map.get(curr).unwrap();

    dirs.iter()
        .filter_map(|&d| {
            let p = curr + d;
            map.get(p).filter(|&j| i + 1 == *j).and(Some(p))
        })
        .collect()
}

fn walk_trail(curr: Point2D, map: &Map2D<u32>, seen: &mut HashSet<Point2D>) -> usize {
    if seen.contains(&curr) {
        return 0;
    }

    seen.insert(curr);

    if map.get(curr).is_some_and(|v| *v == 9) {
        return 1;
    }

    let mut acc = 0;

    for node in get_next_nodes(curr, map) {
        acc += walk_trail(node, map, seen);
    }

    acc
}

impl Day for Day10 {
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
fn test_day10_p1() {
    let input = String::from(
        "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
",
    );

    let day = Day10::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 36)
}

#[test]
fn test_day10_p2() {
    let input = String::from(
        "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
",
    );

    let day = Day10::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 81)
}
