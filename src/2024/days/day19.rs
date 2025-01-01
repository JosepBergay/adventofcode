use std::{collections::HashMap, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day19 {}

impl Day19 {
    fn parse_input(&self, input: String) -> (usize, usize) {
        let mut lines = input.lines();

        let mut patterns = lines
            .next()
            .unwrap()
            .split(", ")
            .map(|l| l.to_string())
            .collect::<Vec<_>>();

        patterns.sort_by(|a, b| b.len().cmp(&a.len()));

        let mut p1 = 0;
        let mut p2 = 0;
        let cache = &mut HashMap::new();

        for towel in lines.filter(|l| !l.is_empty()) {
            let count = find_patterns(
                &towel,
                &patterns.iter().filter(|p| towel.contains(*p)).collect(),
                cache,
            );
            p2 += count;
            if count > 0 {
                p1 += 1;
            }
        }

        (p1, p2)
    }

    fn part1(&self, solution: &(usize, usize)) -> usize {
        solution.0
    }

    fn part2(&self, solution: (usize, usize)) -> usize {
        solution.1
    }
}

fn find_patterns<'a>(
    towel: &'a str,
    patterns: &Vec<&'a String>,
    cache: &mut HashMap<&'a str, usize>,
) -> usize {
    if towel.is_empty() {
        return 1;
    }

    if cache.contains_key(towel) {
        return cache[towel];
    }

    patterns
        .iter()
        .filter(|&p| towel.starts_with(*p))
        .map(|p| {
            let count = find_patterns(&towel[p.len()..], patterns, cache);

            cache.entry(&towel[p.len()..]).or_insert(count);

            count
        })
        .sum()
}

impl Day for Day19 {
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
fn test_day19_p1() {
    let input = String::from(
        "r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
",
    );

    let day = Day19::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 6)
}

#[test]
fn test_day19_p2() {
    let input = String::from(
        "r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
",
    );

    let day = Day19::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 16)
}
