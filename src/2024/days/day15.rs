use std::{collections::HashSet, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day15 {}

impl Day15 {
    fn parse_input(&self, input: String) -> (Vec<Vec<char>>, String, (usize, usize)) {
        let mut split = input.split("\n\n");

        let map = split
            .next()
            .unwrap()
            .split('\n')
            .filter(|line| !line.is_empty())
            .map(|line| line.chars().collect())
            .collect::<Vec<Vec<char>>>();

        let mut curr = (0, 0);

        for y in 0..map.len() {
            for x in 0..map[y].len() {
                if map[y][x] == '@' {
                    curr = (x, y)
                }
            }
        }

        (map, String::from(split.next().unwrap()), curr)
    }

    fn part1(&self, parsed: &(Vec<Vec<char>>, String, (usize, usize))) -> usize {
        let (map, moves, start) = parsed;
        let mut map = map.clone();
        let mut curr = start.clone();

        for dir in moves.chars().filter_map(get_next_dir) {
            if let Some(p) = find_next_empty_space(curr, dir, &map) {
                // Found empty space. Shift boxes and robot.
                // If curr + dir == p then we will overwrite twice so order matters.
                map[p.1][p.0] = 'O';
                map[curr.1][curr.0] = '.';

                curr = move_towards(curr, dir);

                map[curr.1][curr.0] = '@';
            }
        }

        compute_gps_coords(&map)
    }

    fn part2(&self, parsed: (Vec<Vec<char>>, String, (usize, usize))) -> usize {
        let moves = parsed.1;
        let mut map = vec![vec!['.'; parsed.0[0].len() * 2]; parsed.0.len()];
        let mut curr = (parsed.2 .0 * 2, parsed.2 .1);

        // Expand map
        for (y, row) in parsed.0.iter().enumerate() {
            for (x, c) in row.iter().enumerate() {
                let next = match c {
                    '#' => ('#', '#'),
                    'O' => ('[', ']'),
                    '.' => ('.', '.'),
                    '@' => {
                        curr = (x * 2, y);
                        ('@', '.')
                    }
                    _ => panic!("Unexpected char"),
                };
                map[y][x * 2] = next.0;
                map[y][x * 2 + 1] = next.1;
            }
        }

        let box_str = "[]";
        let start_with = ".@";
        let end_with = "@.";
        for dir in moves.chars().filter_map(get_next_dir) {
            if dir.1 == 0 {
                // Move horizontally like part 1
                if let Some(p) = find_next_empty_space(curr, dir, &map) {
                    let (range, replace_with) = if p.0 > curr.0 {
                        // Move right
                        (
                            curr.0..=p.0,
                            start_with
                                .chars()
                                .chain(box_str.chars().cycle())
                                .take(p.0 - curr.0 + 1)
                                .collect::<Vec<char>>(),
                        )
                    } else {
                        // Move left
                        (
                            p.0..=curr.0,
                            box_str
                                .chars()
                                .cycle()
                                .take(curr.0 - p.0 - 1)
                                .chain(end_with.chars())
                                .collect(),
                        )
                    };

                    map[p.1].splice(range, replace_with);

                    curr = move_towards(curr, dir);
                }

                continue;
            }

            // Move vertically
            if find_next_empty_space_vertically(&vec![curr], dir, &mut map) {
                curr = move_towards(curr, dir);
            }
        }

        // print_map(&map);
        compute_gps_coords(&map)
    }
}

fn find_next_empty_space_vertically(
    curr: &Vec<(usize, usize)>,
    dir: (i32, i32),
    map: &mut Vec<Vec<char>>,
) -> bool {
    let mut count = 0;

    let mut to_check = vec![];

    for p in curr {
        let curr_c = map[p.1][p.0];

        match curr_c {
            '.' => {
                count += 1;
            }
            '#' => {
                return false;
            }
            _ => {
                to_check.push(p);
            }
        }
    }

    if count == curr.len() {
        return true;
    }

    let nexts = to_check
        .iter()
        .map(|&p| {
            let next = move_towards(*p, dir);
            let nexts = match map[next.1][next.0] {
                '[' => vec![next, (next.0 + 1, next.1)],
                ']' => vec![(next.0 - 1, next.1), next],
                _ => vec![next],
            };
            nexts
        })
        .flatten()
        .collect::<HashSet<_>>();

    if find_next_empty_space_vertically(&Vec::from_iter(nexts), dir, map) {
        // Move curr vertically
        for curr in to_check {
            let next = move_towards(*curr, dir);
            map[next.1][next.0] = map[curr.1][curr.0];
            map[curr.1][curr.0] = '.';
        }

        return true;
    }

    false
}

fn find_next_empty_space(
    curr: (usize, usize),
    dir: (i32, i32),
    map: &Vec<Vec<char>>,
) -> Option<(usize, usize)> {
    let mut p = curr;

    while map[p.1][p.0] != '#' {
        if map[p.1][p.0] == '.' {
            return Some(p);
        }

        p = move_towards(p, dir);
    }

    None
}

fn move_towards(curr: (usize, usize), dir: (i32, i32)) -> (usize, usize) {
    (
        ((curr.0 as i32) + dir.0) as usize,
        ((curr.1 as i32) + dir.1) as usize,
    )
}

fn get_next_dir(c: char) -> Option<(i32, i32)> {
    match c {
        '>' => Some((1, 0)),
        'v' => Some((0, 1)),
        '<' => Some((-1, 0)),
        '^' => Some((0, -1)),
        _ => None,
    }
}

fn compute_gps_coords(map: &Vec<Vec<char>>) -> usize {
    let mut total = 0;

    for y in 0..map.len() {
        for x in 0..map[y].len() {
            if map[y][x] == 'O' || map[y][x] == '[' {
                total += 100 * y + x;
            }
        }
    }

    total
}

fn _print_map(map: &Vec<Vec<char>>) {
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

#[cfg(test)]
fn get_test_input() -> String {
    String::from(
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
    )
}

#[test]
fn test_day15_p1() {
    let input = get_test_input();

    let day = Day15::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 2028)
}

#[test]
fn test_day15_p2() {
    let input = get_test_input();

    let day = Day15::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 9021)
}
