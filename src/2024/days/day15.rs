use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day15 {}

impl Day15 {
    fn parse_input(&self, input: String) -> (Vec<Vec<char>>, String) {
        let mut split = input.split("\n\n");

        let map = split
            .next()
            .unwrap()
            .split('\n')
            .filter(|line| !line.is_empty())
            .map(|line| line.chars().collect())
            .collect::<Vec<Vec<char>>>();

        (map, String::from(split.next().unwrap()))
    }

    fn part1(&self, parsed: &(Vec<Vec<char>>, String)) -> usize {
        let (map, moves) = parsed;
        let mut map = map.clone();
        let mut curr = (0, 0);

        for y in 0..map.len() {
            for x in 0..map[y].len() {
                if map[y][x] == '@' {
                    curr = (x, y)
                }
            }
        }

        for m in moves.chars() {
            let dir: (isize, isize) = match m {
                '>' => (1, 0),
                'v' => (0, 1),
                '<' => (-1, 0),
                '^' => (0, -1),
                _ => {
                    continue;
                }
            };

            // Find next empty space
            let mut p = curr;
            while map[p.1][p.0] != '#' {
                if map[p.1][p.0] == '.' {
                    // Found empty space. Shift boxes and robot.
                    // If curr + dir == p then we will overwrite twice so order matters.
                    map[p.1][p.0] = 'O';
                    map[curr.1][curr.0] = '.';

                    curr = (
                        ((curr.0 as isize) + dir.0) as usize,
                        ((curr.1 as isize) + dir.1) as usize,
                    );

                    map[curr.1][curr.0] = '@';

                    break;
                }
                p = (
                    ((p.0 as isize) + dir.0) as usize,
                    ((p.1 as isize) + dir.1) as usize,
                );
            }
        }

        return compute_gps_coords(&map);
    }

    fn part2(&self, _parsed: (Vec<Vec<char>>, String)) -> &str {
        "TODO"
    }
}

fn compute_gps_coords(map: &Vec<Vec<char>>) -> usize {
    let mut total = 0;

    for y in 0..map.len() {
        for x in 0..map[y].len() {
            if map[y][x] == 'O' {
                total += 100 * y + x;
            }
        }
    }

    total
}

fn print_map(map: &Vec<Vec<char>>) {
    let mut s = String::with_capacity(map[0].len() * map.len());

    for row in map {
        for c in row {
            s.push(*c);
        }
        s.push('\n');
    }

    println!("{s}");
}

impl Day for Day15 {
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
fn test_day15_p1() {
    let input = String::from(
        "##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
",
    );

    let day = Day15::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 2028)
}

#[test]
fn test_day15_p2() {
    let input = String::from("");

    let day = Day15::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
