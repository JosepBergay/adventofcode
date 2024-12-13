use std::{collections::HashSet, error::Error};

use super::{baseday::DayResult, map2d::Map2D, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day6 {}

impl Day6 {
    fn parse_input(&self, input: String) -> Map2D<char> {
        Map2D::<char>::from_string(input)
    }

    fn part1(&self, map: &Map2D<char>) -> usize {
        let dirs = vec![
            Point2D { x: 0, y: -1 },
            Point2D { x: 1, y: 0 },
            Point2D { x: 0, y: 1 },
            Point2D { x: -1, y: 0 },
        ];

        let mut curr_pos = map.iter().find(|p| *map.get(*p).unwrap() == '^').unwrap();

        // We start going up
        let mut curr_dir = 0;

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

    fn part2(&self, parsed: Map2D<char>) -> &str {
        "TODO"
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
    let input = String::from("");

    let day = Day6::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
