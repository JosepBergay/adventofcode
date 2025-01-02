use std::{
    collections::{HashSet, VecDeque},
    error::Error,
};

use crate::days::map2d::Map2D;

use super::{baseday::DayResult, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day20 {}

type Input = (Vec<Point2D>, Map2D<char>);

impl Day20 {
    fn parse_input(&self, input: String) -> Input {
        let map = Map2D::<char>::from_string(input);

        let mut start = Point2D::new(0, 0);
        let mut end = Point2D::new(0, 0);

        for p in map.iter() {
            match *map.get(p).unwrap() {
                'E' => end = p,
                'S' => start = p,
                _ => {}
            }
        }

        let mut path = vec![];
        let mut seen = HashSet::new();

        let mut q = VecDeque::new();

        q.push_back(start);

        while let Some(curr) = q.pop_front() {
            path.push(curr);
            seen.insert(curr);

            if curr == end {
                break;
            }

            for neigh in map.get_adjacent(curr) {
                if *map.get(neigh).unwrap() != '#' && !seen.contains(&neigh) {
                    q.push_back(neigh);
                }
            }
        }

        (path, map)
    }

    fn part1(&self, parsed: &Input, savings: usize) -> usize {
        let (path, map) = parsed;

        let dists = get_points_at(2);

        path[..path.len() - 2]
            .iter()
            .enumerate()
            .flat_map(|(i, p)| {
                dists
                    .iter()
                    .map(|d| *p + *d)
                    .filter(move |&next| {
                        map.get(next).is_some_and(|c| *c != '#')
                            && path[i + 2..]
                                .iter()
                                .position(|it| *it == next)
                                .is_some_and(|j| j >= savings)
                    })
                    .collect::<Vec<Point2D>>()
            })
            .count()
    }

    fn part2(&self, parsed: Input, savings: usize) -> usize {
        let (path, map) = parsed;

        let cheat_dists: Vec<_> = (2..=20).flat_map(|n| get_points_at(n)).collect();

        let mut count = 0;

        for (i, p) in path.iter().enumerate() {
            let cheat_ends = cheat_dists.iter().filter(|&d| {
                let dist = (d.x.abs() + d.y.abs()) as usize;

                if dist > path.len() - i {
                    return false;
                }

                let next = *p + *d;
                map.get(next).is_some_and(|c| *c != '#')
                    && path[i + dist..]
                        .iter()
                        .position(|it| *it == next)
                        .is_some_and(|j| j >= savings)
            });

            count += cheat_ends.count();
        }

        count
    }
}

fn get_points_at(manhattan_dist: i32) -> Vec<Point2D> {
    let manhattan_dist = manhattan_dist.abs();

    let mut out = vec![];

    for x in -manhattan_dist..=manhattan_dist {
        for y in -manhattan_dist..=manhattan_dist {
            if x.abs() + y.abs() == manhattan_dist {
                out.push(Point2D::new(x, y));
            }
        }
    }

    out
}

impl Day for Day20 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&parsed, 100);
        let p2 = self.part2(parsed, 100);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[cfg(test)]
fn get_test_input() -> String {
    String::from(
        "###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
",
    )
}

#[test]
fn test_day20_p1() {
    let input = get_test_input();

    let day = Day20::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed, 2);

    assert_eq!(res, 44);
}

#[test]
fn test_day20_p2() {
    let input = get_test_input();

    let day = Day20::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed, 50);

    assert_eq!(res, 285);
}
