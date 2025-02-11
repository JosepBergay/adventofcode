use std::{collections::HashSet, error::Error};

use super::{baseday::DayResult, map2d::Map2D, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day6 {}

impl Day6 {
    fn parse_input(&self, input: String) -> (Point2D, Map2D<char>, Vec<Point2D>) {
        let map = Map2D::<char>::from_string(input);

        let start = map.iter().find(|p| *map.get(*p).unwrap() == '^').unwrap();

        let dirs = vec![
            Point2D { x: 0, y: -1 },
            Point2D { x: 1, y: 0 },
            Point2D { x: 0, y: 1 },
            Point2D { x: -1, y: 0 },
        ];

        (start, map, dirs)
    }

    fn part1(&self, parsed: &(Point2D, Map2D<char>, Vec<Point2D>)) -> usize {
        let (start, map, dirs) = parsed;

        let mut curr_pos = *start;
        let mut curr_dir = 0; // We start going up
        let mut seen = HashSet::new();

        while !map.is_out_of_bounds(curr_pos) {
            if *map.get(curr_pos).unwrap() == '#' {
                // Get back and turn right
                curr_pos -= dirs[curr_dir];
                curr_dir = (curr_dir + 1) % 4;
            } else {
                seen.insert(curr_pos);
            }

            curr_pos += dirs[curr_dir];
        }

        seen.len()
    }

    fn part2(&self, parsed: (Point2D, Map2D<char>, Vec<Point2D>)) -> i32 {
        let (start, map, dirs) = parsed;

        let mut count = 0;

        let mut obstacles_set = Map2D::new(vec![vec![false; map.map[0].len()]; map.map.len()]);
        let mut start = start;
        let mut start_dir = 0;

        loop {
            let mut curr_pos = start;
            let mut curr_dir = start_dir; // We start going up
            let mut seen = Map2D::new(vec![vec![4; map.map[0].len()]; map.map.len()]);
            let mut obstacle_pos: Option<Point2D> = None;

            while !map.is_out_of_bounds(curr_pos) {
                if seen.get(curr_pos).is_some_and(|v| *v != 4) {
                    // Found loop!
                    count += 1;
                    break;
                }

                let c = map.get(curr_pos).unwrap();

                if obstacle_pos.is_none() && !obstacles_set.get(curr_pos).unwrap() && c != &'#' {
                    // Optimization: save pos+dir so next loop we start from here (4x faster).
                    start = curr_pos;
                    start_dir = curr_dir;
                    obstacle_pos = Some(curr_pos);
                    obstacles_set.map[curr_pos.y as usize][curr_pos.x as usize] = true;
                }

                if obstacle_pos.is_some_and(|p| p == curr_pos) || c == &'#' {
                    // Get back and turn right
                    curr_pos -= dirs[curr_dir];
                    curr_dir = (curr_dir + 1) % 4;
                } else {
                    seen.map[curr_pos.y as usize][curr_pos.x as usize] = curr_dir;
                }

                curr_pos += dirs[curr_dir];
            }

            if obstacle_pos.is_none() {
                break;
            }
        }

        count
    }
}

impl Day for Day6 {
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
fn test_day6_p1() {
    let input = String::from(
        "....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
",
    );

    let day = Day6::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 41)
}

#[test]
fn test_day6_p2() {
    let input = String::from(
        "....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
",
    );

    let day = Day6::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 6)
}
