use std::{
    cmp::Ordering,
    collections::{BinaryHeap, HashMap, HashSet},
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
    fn parse_input(&self, input: String) -> (usize, usize) {
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

        let map = Map2D::new(map);

        let mut lowest = usize::MAX;
        let mut tiles: HashSet<Point2D> = HashSet::new();
        let mut costs = HashMap::new();
        let mut queue = BinaryHeap::new();

        let mut path = Vec::new();
        path.push(start);
        queue.push(NodeState {
            cost: 0,
            dir: Point2D { x: 1, y: 0 },
            pos: start,
            path,
        });

        while let Some(state) = queue.pop() {
            let NodeState {
                cost,
                dir,
                path,
                pos,
            } = state;

            costs.insert((pos, dir), cost);

            if pos == end {
                if lowest < cost {
                    break;
                }
                lowest = cost;
                tiles.extend(&path);
            }

            if costs.get(&(pos, dir)).is_some_and(|c| cost > *c) {
                continue;
            }

            let to_left = Point2D { x: dir.y, y: dir.x };
            let to_right = Point2D {
                x: -1 * dir.y,
                y: -1 * dir.x,
            };

            let next_nodes = [(dir, 1), (to_left, 1001), (to_right, 1001)];

            for (next_dir, cost_inc) in next_nodes {
                let next_pos = pos + next_dir;

                if map.get(next_pos).is_none_or(|t| *t == Tile::Wall) {
                    continue;
                }

                let next_cost = cost + cost_inc;

                if costs
                    .get(&(next_pos, next_dir))
                    .is_none_or(|c| next_cost < *c)
                {
                    let mut next_path = path.clone();
                    next_path.push(next_pos);

                    queue.push(NodeState {
                        cost: next_cost,
                        dir: next_dir,
                        pos: next_pos,
                        path: next_path,
                    });
                }
            }
        }

        (lowest, tiles.len())
    }

    fn part1(&self, parsed: &(usize, usize)) -> usize {
        let (lowest, _) = parsed;

        *lowest
    }

    fn part2(&self, parsed: (usize, usize)) -> usize {
        let (_, tiles_count) = parsed;

        tiles_count
    }
}

#[derive(Eq, PartialEq)]
struct NodeState {
    cost: usize,
    dir: Point2D,
    pos: Point2D,
    path: Vec<Point2D>,
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
