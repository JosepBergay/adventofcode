use std::{
    collections::{HashSet, VecDeque},
    error::Error,
};

use crate::days::map2d::Map2D;

use super::{baseday::DayResult, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day20 {}

type Input = (Point2D, Point2D, Map2D<char>);

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

        (start, end, map)
    }

    fn part1(&self, parsed: &Input, savings: usize) -> usize {
        let (start, end, map) = parsed;
        let mut path = vec![];
        let mut seen = HashSet::new();

        let mut q = VecDeque::new();

        q.push_back(*start);

        while let Some(curr) = q.pop_front() {
            path.push(curr);
            seen.insert(curr);

            if curr == *end {
                break;
            }

            for neigh in map.get_adjacent(curr) {
                if *map.get(neigh).unwrap() != '#' && !seen.contains(&neigh) {
                    q.push_back(neigh);
                }
            }
        }

        let dists = vec![
            Point2D::new(2, 0),
            Point2D::new(1, 1),
            Point2D::new(0, 2),
            Point2D::new(-1, 1),
            Point2D::new(-2, 0),
            Point2D::new(-1, -1),
            Point2D::new(0, -2),
            Point2D::new(1, -1),
        ];

        path[..path.len() - 2]
            .iter()
            .enumerate()
            .flat_map(|(i, p)| {
                let path_left = &path[i + 2..];

                let cheats = dists.iter().map(|d| *p + *d).filter(|&next| {
                    map.get(next).is_some_and(|c| *c != '#')
                        && path_left
                            .iter()
                            .position(|it| *it == next)
                            .is_some_and(|j| j >= savings)
                });

                cheats
            })
            .count()
    }

    fn part2(&self, _parsed: Input) -> &str {
        "TODO"
    }
}

impl Day for Day20 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&parsed, 100);
        let p2 = self.part2(parsed);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[test]
fn test_day20_p1() {
    let input = String::from(
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
    );

    let day = Day20::default();
    let parsed = day.parse_input(input);

    let res = day.part1(&parsed, 2);
    assert_eq!(res, 44);
}

#[test]
fn test_day20_p2() {
    let input = String::from("");

    let day = Day20::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
