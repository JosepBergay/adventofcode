use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day22 {}

impl Day22 {
    fn parse_input(&self, input: String) -> Vec<Vec<usize>> {
        input
            .lines()
            .filter_map(|l| l.parse::<usize>().ok())
            .map(|n| {
                let steps = 2000;
                let mut v = Vec::with_capacity(steps);
                let mut curr = n;
                for _ in 0..steps {
                    v.push(curr);
                    curr = compute_secret(curr);
                }
                v
            })
            .collect()
    }

    fn part1(&self, parsed: &Vec<Vec<usize>>) -> usize {
        parsed.iter().map(|v| v.last().unwrap()).sum()
    }

    fn part2(&self, monkeys: Vec<Vec<usize>>) -> usize {
        let max_diff_sequences = 19_usize.pow(4);
        // Create global map of sequences to bananas
        // (when using Vec instead of HashMap/HashSet run time went down from ~500ms to ~70ms)
        let mut map = vec![0; max_diff_sequences];

        // Scan each monkey for unique sequences
        for secrets in monkeys {
            let mut seen = vec![false; max_diff_sequences];

            for win in secrets.windows(5) {
                let hash = get_hash(win);

                // When a unique sequence is found update global map with price
                if seen[hash] == false {
                    seen[hash] = true;
                    map[hash] += win.last().unwrap() % 10;
                }
            }
        }

        // Return global map max value
        *map.iter().max().unwrap()
    }
}

fn get_diff(curr: usize, prev: usize) -> usize {
    let diff = (curr % 10) as isize - (prev % 10) as isize + 9;
    diff as usize
}

fn get_hash(win: &[usize]) -> usize {
    let a = get_diff(win[1], win[0]);
    let b = get_diff(win[2], win[1]);
    let c = get_diff(win[3], win[2]);
    let d = get_diff(win[4], win[3]);

    let base = 19_usize;
    let hash = d * base.pow(3) + c * base.pow(2) + b * base + a;

    hash
}

fn compute_secret(curr: usize) -> usize {
    let mut curr = prune_secret_number(mix_secret_number(curr, curr * 64));
    curr = prune_secret_number(mix_secret_number(curr, curr / 32));
    curr = prune_secret_number(mix_secret_number(curr, curr * 2048));
    curr
}

fn mix_secret_number(secret_number: usize, value: usize) -> usize {
    secret_number ^ value
}

fn prune_secret_number(secret_number: usize) -> usize {
    secret_number % 16777216
}

impl Day for Day22 {
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
fn test_day22_p1() {
    let input = String::from(
        "1
10
100
2024
",
    );

    let day = Day22::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 37327623)
}

#[test]
fn test_day22_p2() {
    let input = String::from(
        "1
2
3
2024
",
    );

    let day = Day22::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 23)
}
