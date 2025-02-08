use std::{collections::VecDeque, error::Error};

use crate::days::map2d::Map2D;

use super::{baseday::DayResult, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day20 {}

type Input = (Vec<Point2D>, Map2D<isize>);

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
        let mut seen: Map2D<isize> = Map2D::new(vec![vec![-1; map.map[0].len()]; map.map.len()]);

        let mut q = VecDeque::new();

        q.push_back(start);

        while let Some(curr) = q.pop_front() {
            path.push(curr);
            seen.map[curr.y as usize][curr.x as usize] = (path.len() - 1) as isize;

            if curr == end {
                break;
            }

            for neigh in map.get_adjacent(curr) {
                if *map.get(neigh).unwrap() != '#' && !seen.get(neigh).is_some_and(|v| *v != -1) {
                    q.push_back(neigh);
                }
            }
        }

        (path, seen)
    }

    fn part1(&self, parsed: &Input, savings: usize) -> usize {
        let (path, point_to_idx_map) = parsed;

        count_cheats(path, point_to_idx_map, 2, savings)
    }

    fn part2(&self, parsed: Input, savings: usize) -> usize {
        let (path, point_to_idx_map) = parsed;

        count_cheats(&path, &point_to_idx_map, 20, savings)
    }
}

fn count_cheats(
    path: &Vec<Point2D>,
    point_to_idx_map: &Map2D<isize>,
    cheat_max_distance: usize,
    savings: usize,
) -> usize {
    let cheat_dists: Vec<_> = (2..=cheat_max_distance)
        .flat_map(|n| get_points_at(n as i32))
        .collect();

    let mut count = 0;

    for (i, p) in path[..path.len() - 2].iter().enumerate() {
        count += cheat_dists
            .iter()
            .filter(|&d| {
                let dist = (d.x.abs() + d.y.abs()) as usize;

                if dist > path.len() - i {
                    return false;
                }

                let next = *p + *d;

                point_to_idx_map
                    .get(next)
                    .and_then(|v| if *v == -1 { None } else { Some(*v as usize) })
                    .is_some_and(|end_idx| end_idx > i && (end_idx - i - dist) >= savings)
            })
            .count();
    }

    count
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
