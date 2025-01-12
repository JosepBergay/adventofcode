use std::{
    cmp::Ordering,
    collections::{HashMap, VecDeque},
    error::Error,
};

use super::{
    baseday::DayResult,
    point2d::{get_orthogonal_directions, Point2D},
    Day,
};

#[derive(Default)]
pub struct Day21 {}

type Input = Vec<String>;

impl Day21 {
    fn parse_input(&self, input: String) -> Input {
        input
            .lines()
            .filter(|l| !l.is_empty())
            .map(|l| l.to_string())
            .collect()
    }

    fn part1(&self, parsed: &Input) -> usize {
        let numeric_start = Point2D::new(2, 3);
        let directional_start = Point2D::new(2, 0);

        let mut complexity_sum = 0;

        for code in parsed {
            let mut min_moves = usize::MAX;

            for moves in get_numeric_keypad_moves(numeric_start, code) {
                // print_moves(&moves);
                // println!("{moves:?}");
                for moves in get_directional_keypad_moves(directional_start, moves) {
                    // print_moves(&moves);
                    // println!("{moves:?}");
                    let moves = get_directional_keypad_moves(directional_start, moves);

                    min_moves = min_moves.min(moves.iter().map(|m| m.len()).min().unwrap());
                }
                // print_moves(&moves);
                // println!("{moves:?}");
            }

            let numeric_part = code[..code.len() - 1].parse::<usize>().unwrap();

            // println!("Code [{code}] num: {}, len: {}", numeric_part, min_moves);

            complexity_sum += min_moves * numeric_part;

            // panic!("fok")
        }

        complexity_sum
    }

    fn part2(&self, _parsed: Input) -> &str {
        "TODO"
    }
}

fn get_unitary_moves(
    from: Point2D,
    to: Point2D,
    skip: Point2D,
    max_x: i32,
    max_y: i32,
) -> Vec<Vec<Point2D>> {
    let mut all_moves = vec![];

    let dist = to - from;

    let mut y_moves = match dist.y.cmp(&0) {
        Ordering::Equal => vec![],
        Ordering::Greater => vec![Point2D::new(0, 1); dist.y as usize],
        Ordering::Less => vec![Point2D::new(0, -1); dist.y.abs() as usize],
    };

    let mut x_moves = match dist.x.cmp(&0) {
        Ordering::Equal => vec![],
        Ordering::Greater => vec![Point2D::new(1, 0); dist.x as usize],
        Ordering::Less => vec![Point2D::new(-1, 0); dist.x.abs() as usize],
    };

    if dist.y == 0 {
        x_moves.push(Point2D::new(0, 0)); // (0,0) Means 'A' press
        all_moves.push(x_moves);
    } else if dist.y == 0 {
        y_moves.push(Point2D::new(0, 0)); // (0,0) Means 'A' press
        all_moves.push(y_moves);
    } else {
        let paths = get_possible_paths(from, to, skip, max_x, max_y);

        let mut min_paths: Vec<Vec<Point2D>> = vec![];

        for p in paths {
            if min_paths.first().is_none_or(|mp| mp.len() > p.len()) {
                min_paths = vec![];
            }
            min_paths.push(p);
        }

        for path in min_paths {
            let mut moves = vec![];
            for window in path.windows(2) {
                if let [p1, p2] = window {
                    moves.push(*p2 - *p1);
                }
            }

            moves.push(Point2D::new(0, 0)); // (0,0) Means 'A' press
            all_moves.push(moves);
        }
    }

    all_moves
}

fn get_possible_paths(
    from: Point2D,
    to: Point2D,
    skip: Point2D,
    max_x: i32,
    max_y: i32,
) -> Vec<Vec<Point2D>> {
    let mut paths = vec![];

    let dirs = get_orthogonal_directions();

    let mut costs = HashMap::new();
    costs.insert(from, 0);
    let mut q = VecDeque::new();
    q.push_back((from, vec![]));

    while let Some((curr, mut path)) = q.pop_front() {
        path.push(curr);

        if curr == to {
            paths.push(path);
            continue;
        }

        let curr_cost = *costs.get(&curr).unwrap();

        for next in dirs
            .iter()
            .map(|d| *d + curr)
            .filter(|&n| n != skip && (0..=max_x).contains(&n.x) && (0..=max_y).contains(&n.y))
        {
            if costs.get(&next).is_none_or(|c| curr_cost + 1 <= *c) {
                costs.insert(next, curr_cost + 1);
                q.push_back((next, path.clone()));
            }
        }
    }

    paths
}

/**
 *     +---+---+
 *     | ^ | A |
 * +---+---+---+
 * | < | v | > |
 * +---+---+---+
 */
fn get_directional_keypad_moves(start_pos: Point2D, presses: Vec<Point2D>) -> Vec<Vec<Point2D>> {
    let empty_space = Point2D::new(0, 0);

    let mut curr = start_pos;

    let mut moves = vec![];

    for press in presses {
        let dest = match press {
            Point2D { x: 0, y: -1 } => Point2D::new(1, 0),
            Point2D { x: 0, y: 0 } => Point2D::new(2, 0),
            Point2D { x: -1, y: 0 } => Point2D::new(0, 1),
            Point2D { x: 0, y: 1 } => Point2D::new(1, 1),
            Point2D { x: 1, y: 0 } => Point2D::new(2, 1),
            _ => panic!("Unknown button press"),
        };

        let u_moves = get_unitary_moves(curr, dest, empty_space, 2, 1);

        moves = merge_moves(moves, u_moves);

        curr = dest;
    }

    moves
}

fn _print_moves(dirs: &Vec<Point2D>) {
    let str: String = dirs
        .iter()
        .map(|d| match d {
            Point2D { x: 0, y: 0 } => 'A',
            Point2D { x: 1, y: 0 } => '>',
            Point2D { x: -1, y: 0 } => '<',
            Point2D { x: 0, y: 1 } => 'v',
            Point2D { x: 0, y: -1 } => '^',
            _ => panic!("unknown dir"),
        })
        .collect();

    println!("{str}");
}

/**
 * +---+---+---+
 * | 7 | 8 | 9 |
 * +---+---+---+
 * | 4 | 5 | 6 |
 * +---+---+---+
 * | 1 | 2 | 3 |
 * +---+---+---+
 *     | 0 | A |
 *     +---+---+
 */
fn get_numeric_keypad_moves(start_pos: Point2D, code: &String) -> Vec<Vec<Point2D>> {
    let empty_space = Point2D::new(0, 3);

    let mut curr = start_pos;

    let mut moves = vec![];

    for c in code.chars() {
        let dest = match c {
            '7' => Point2D::new(0, 0),
            '8' => Point2D::new(1, 0),
            '9' => Point2D::new(2, 0),
            '4' => Point2D::new(0, 1),
            '5' => Point2D::new(1, 1),
            '6' => Point2D::new(2, 1),
            '1' => Point2D::new(0, 2),
            '2' => Point2D::new(1, 2),
            '3' => Point2D::new(2, 2),
            '0' => Point2D::new(1, 3),
            'A' => Point2D::new(2, 3),
            _ => panic!("Unknown keypad button"),
        };

        let u_moves = get_unitary_moves(curr, dest, empty_space, 2, 3);

        moves = merge_moves(moves, u_moves);

        curr = dest;
    }

    moves
}

fn merge_moves(moves: Vec<Vec<Point2D>>, u_moves: Vec<Vec<Point2D>>) -> Vec<Vec<Point2D>> {
    if moves.is_empty() {
        return u_moves;
    } else {
        let mut new_moves = Vec::with_capacity(u_moves.len() * moves.len());

        for um in u_moves {
            for m in &moves {
                new_moves.push([m.clone(), um.clone()].concat());
                // m.extend(um.clone());
            }
        }

        return new_moves;
    }
}

impl Day for Day21 {
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
fn test_day21_p1() {
    let input = String::from(
        "029A
980A
179A
456A
379A
",
    );

    let day = Day21::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    let paths = get_possible_paths(
        Point2D::new(0, 1),
        Point2D::new(2, 0),
        Point2D::new(0, 0),
        2,
        1,
    );

    assert!(paths.contains(&vec![
        Point2D::new(0, 1),
        Point2D::new(1, 1),
        Point2D::new(1, 0),
        Point2D::new(2, 0)
    ]));
    assert!(paths.contains(&vec![
        Point2D::new(0, 1),
        Point2D::new(1, 1),
        Point2D::new(2, 1),
        Point2D::new(2, 0)
    ]));

    let _paths = get_possible_paths(
        Point2D::new(1, 3),
        Point2D::new(2, 0),
        Point2D::new(0, 3),
        2,
        3,
    );
    // assert!(paths.contains(&vec![
    //     Point2D::new(1, 3),
    //     Point2D::new(1, 1),
    //     Point2D::new(2, 1),
    //     Point2D::new(0, 3)
    // ]));

    assert_eq!(res, 126384)
}

#[test]
fn test_day21_p2() {
    let input = String::from("");

    let day = Day21::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
