use std::{collections::HashSet, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day19 {}

impl Day19 {
    fn parse_input(&self, input: String) -> (Vec<String>, Vec<String>) {
        let mut lines = input.lines();

        let patterns = lines
            .next()
            .unwrap()
            .split(", ")
            .map(|l| l.to_string())
            .collect();

        let towels = lines
            .filter(|l| !l.is_empty())
            .map(|l| l.to_string())
            .collect();

        (patterns, towels)
    }

    fn part1(&self, parsed: &(Vec<String>, Vec<String>)) -> usize {
        let (mut patterns, towels) = parsed.clone();

        patterns.sort_by(|a, b| b.len().cmp(&a.len()));

        towels
            .iter()
            .filter(|towel| {
                find_patterns(
                    towel,
                    &patterns.iter().filter(|p| towel.contains(*p)).collect(),
                    &mut HashSet::new(),
                )
            })
            .count()
    }

    fn part2(&self, _parsed: (Vec<String>, Vec<String>)) -> usize {
        0
    }
}

fn find_patterns<'a>(
    towel: &'a str,
    patterns: &Vec<&'a String>,
    not_found: &mut HashSet<&'a str>,
) -> bool {
    if towel.is_empty() {
        return true;
    }

    if not_found.contains(towel) {
        return false;
    }

    patterns.iter().filter(|&p| towel.starts_with(*p)).any(|p| {
        let found = find_patterns(&towel[p.len()..], patterns, not_found);
        if !found {
            not_found.insert(&towel[p.len()..]);
        }
        found
    })
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
    let input = String::from("");

    let day = Day19::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 0)
}
