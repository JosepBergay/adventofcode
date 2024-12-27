use std::{
    cmp::Ordering,
    collections::{BinaryHeap, HashMap},
    error::Error,
};

use super::{baseday::DayResult, map2d::Map2D, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day16 {}

#[derive(PartialEq)]
enum Tile {
    Empty,
    Wall,
    Start,
    End,
}

impl Day16 {
    fn parse_input(&self, input: String) -> (Point2D, Point2D, Map2D<Tile>) {
        let mut map = vec![];
        let mut start = Point2D { x: 0, y: 0 };
        let mut end = Point2D { x: 0, y: 0 };

        for (y, line) in input.lines().enumerate() {
            if line.is_empty() {
                continue;
            }

            let mut map_row = vec![];

            for (x, c) in line.char_indices() {
                match c {
                    'S' => {
                        start = Point2D {
                            x: x as i32,
                            y: y as i32,
                        };
                        map_row.push(Tile::Start);
                    }
                    'E' => {
                        end = Point2D {
                            x: x as i32,
                            y: y as i32,
                        };
                        map_row.push(Tile::End);
                    }
                    '#' => map_row.push(Tile::Wall),
                    '.' => map_row.push(Tile::Empty),
                    _ => {}
                }
            }

            map.push(map_row);
        }

        (start, end, Map2D::new(map))
    }

    fn part1(&self, parsed: &(Point2D, Point2D, Map2D<Tile>)) -> usize {
        let (start, end, map) = parsed;

        let mut distances = HashMap::new();
        let mut prev = HashMap::new();
        let mut queue = BinaryHeap::new();

        distances.insert(*start, 0);
        queue.push(NodeState {
            cost: 0,
            dir: Point2D { x: 1, y: 0 },
            pos: *start,
        });

        while let Some(NodeState { cost, dir, pos }) = queue.pop() {
            if pos == *end {
                break;
            }

            if distances.get(&pos).is_some_and(|c| cost > *c) {
                continue;
            }

            for node in map
                .get_adjacent(pos)
                .iter()
                .filter(|&p| map.get(*p).is_some_and(|t| *t != Tile::Wall))
            {
                let next_dir = *node - pos;
                let next_cost = cost + 1 + if next_dir != dir { 1000 } else { 0 };

                if distances.get(&node).is_none_or(|c| next_cost < *c) {
                    queue.push(NodeState {
                        cost: next_cost,
                        dir: next_dir,
                        pos: *node,
                    });
                    distances.insert(*node, next_cost);
                    prev.insert(*node, pos);
                }
            }
        }

        *distances.get(end).unwrap()
    }

    fn part2(&self, _parsed: (Point2D, Point2D, Map2D<Tile>)) -> usize {
        0
    }
}

// #[derive(Copy, Clone, Eq, PartialEq)]
#[derive(Eq, PartialEq)]
struct NodeState {
    cost: usize,
    dir: Point2D,
    pos: Point2D,
}

impl Ord for NodeState {
    fn cmp(&self, other: &Self) -> Ordering {
        other.cost.cmp(&self.cost)
        // .then_with(|| self.position.cmp(&other.position))
    }
}

impl PartialOrd for NodeState {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Day for Day16 {
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
fn test_day16_p1() {
    let input = String::from(
        "###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
",
    );

    let day = Day16::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 7036)
}

#[test]
fn test_day16_p2() {
    let input = String::from(
        "###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
",
    );

    let day = Day16::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 45)
}
