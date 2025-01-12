use std::{
    cmp::Ordering,
    collections::{HashMap, VecDeque},
    error::Error,
    hash::Hash,
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

    fn solve(&self, parsed: &Input, robot_count: usize) -> usize {
        let mut complexity_sum = 0;

        let cache = &mut HashMap::new();

        let numeric_keypad = get_numeric_keypad();

        for code in parsed {
            let min_moves = find_fewest_button_presses_count(
                robot_count,
                &code.chars().collect(),
                cache,
                &numeric_keypad,
            );

            let numeric_part = code[..code.len() - 1].parse::<usize>().unwrap();

            complexity_sum += min_moves * numeric_part;
        }

        complexity_sum
    }

    fn part1(&self, parsed: &Input) -> usize {
        self.solve(parsed, 3)
    }

    fn part2(&self, parsed: Input) -> usize {
        self.solve(&parsed, 26)
    }
}

struct Keypad<T> {
    max_x: i32,
    max_y: i32,
    empty_space: Point2D,
    start: Point2D, // 'A' position
    layout: HashMap<T, Point2D>,
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
fn get_numeric_keypad() -> Keypad<char> {
    Keypad {
        empty_space: Point2D::new(0, 3),
        layout: HashMap::from([
            ('7', Point2D::new(0, 0)),
            ('8', Point2D::new(1, 0)),
            ('9', Point2D::new(2, 0)),
            ('4', Point2D::new(0, 1)),
            ('5', Point2D::new(1, 1)),
            ('6', Point2D::new(2, 1)),
            ('1', Point2D::new(0, 2)),
            ('2', Point2D::new(1, 2)),
            ('3', Point2D::new(2, 2)),
            ('0', Point2D::new(1, 3)),
            ('A', Point2D::new(2, 3)),
        ]),
        max_x: 2,
        max_y: 3,
        start: Point2D::new(2, 3),
    }
}

/**
 *     +---+---+
 *     | ^ | A |
 * +---+---+---+
 * | < | v | > |
 * +---+---+---+
 */
fn get_directional_keypad() -> Keypad<Point2D> {
    Keypad {
        empty_space: Point2D::new(0, 0),
        layout: HashMap::from([
            (Point2D { x: 0, y: -1 }, Point2D::new(1, 0)),
            (Point2D { x: 0, y: 0 }, Point2D::new(2, 0)),
            (Point2D { x: -1, y: 0 }, Point2D::new(0, 1)),
            (Point2D { x: 0, y: 1 }, Point2D::new(1, 1)),
            (Point2D { x: 1, y: 0 }, Point2D::new(2, 1)),
        ]),
        max_x: 2,
        max_y: 1,
        start: Point2D::new(2, 0),
    }
}

fn find_fewest_button_presses_count<T>(
    robot_count: usize,
    path: &Vec<T>,
    cache: &mut HashMap<(Point2D, Point2D, usize), usize>,
    keypad: &Keypad<T>,
) -> usize
where
    T: Eq,
    T: Hash,
{
    if robot_count == 0 {
        return path.len();
    }

    let mut curr = keypad.start; // Always start at 'A'
    let mut count = 0;

    for dir in path {
        let dest = *keypad.layout.get(dir).unwrap();

        let key = (curr, dest, robot_count);

        count += cache.get(&key).cloned().unwrap_or_else(|| {
            let moves = get_unitary_moves(curr, dest, keypad);

            let val = moves
                .iter()
                .map(|next_path| {
                    find_fewest_button_presses_count(
                        robot_count - 1,
                        next_path,
                        cache,
                        &get_directional_keypad(),
                    )
                })
                .min()
                .unwrap();

            cache.insert(key, val);

            val
        });

        curr = dest;
    }

    count
}

fn get_unitary_moves<T: Eq>(from: Point2D, to: Point2D, keypad: &Keypad<T>) -> Vec<Vec<Point2D>> {
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
        let paths = get_possible_paths(from, to, keypad);

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

fn get_possible_paths<T>(from: Point2D, to: Point2D, keypad: &Keypad<T>) -> Vec<Vec<Point2D>> {
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

        for next in dirs.iter().map(|d| *d + curr).filter(|&n| {
            n != keypad.empty_space
                && (0..=keypad.max_x).contains(&n.x)
                && (0..=keypad.max_y).contains(&n.y)
        }) {
            if costs.get(&next).is_none_or(|c| curr_cost + 1 <= *c) {
                costs.insert(next, curr_cost + 1);
                q.push_back((next, path.clone()));
            }
        }
    }

    paths
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
        &Keypad {
            max_x: 2,
            max_y: 1,
            empty_space: Point2D::new(0, 0),
            start: Point2D::new(2, 2),
            layout: HashMap::from([(1, Point2D::new(2, 2))]),
        },
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

    assert_eq!(res, 126384)
}

#[test]
fn test_day21_p2() {
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
    let res = day.part2(parsed);

    assert_eq!(res, 154115708116294)
}
