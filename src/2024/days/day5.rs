use std::{
    collections::{hash_set, HashMap, HashSet},
    error::Error,
};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day5 {}

impl Day5 {
    fn parse_input(&self, input: String) -> (HashMap<i32, Vec<i32>>, Vec<Vec<i32>>) {
        let mut input_iter = input.split("\n\n");

        let rules = input_iter.next().unwrap().split("\n");

        let mut deps = HashMap::new();

        for line in rules {
            let mut line_iter = line.split('|');

            let dep = line_iter.next().unwrap().parse().unwrap();
            let key = line_iter.next().unwrap().parse().unwrap();

            let v = deps.entry(key).or_insert(vec![]);
            v.push(dep);
        }

        let updates = input_iter
            .next()
            .unwrap()
            .split("\n")
            .filter(|line| *line != "")
            .map(|line| line.split(',').map(|c| c.parse().unwrap()).collect())
            .collect::<Vec<Vec<i32>>>();

        (deps, updates)
    }

    fn part1(&self, parsed: &(HashMap<i32, Vec<i32>>, Vec<Vec<i32>>)) -> i32 {
        let (rules, updates) = parsed;

        let mut printed: Vec<&Vec<i32>> = vec![];

        'outer: for up in updates {
            let mut seen = HashSet::new();

            for (_i, page) in up.iter().enumerate() {
                let deps = rules.get(page);

                if deps.is_none_or(|ds| {
                    ds.iter()
                        .filter(|d| up.contains(d))
                        .all(|d| seen.contains(d))
                }) {
                    seen.insert(page);
                    continue;
                }

                continue 'outer;
            }

            printed.push(up);
        }

        printed.iter().fold(0, |acc, up| {
            let mid = up[(up.len() - 1) / 2];
            acc + mid
        })
    }

    fn part2(&self, _parsed: (HashMap<i32, Vec<i32>>, Vec<Vec<i32>>)) -> &str {
        "TODO"
    }
}

impl Day for Day5 {
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
fn test_day5_p1() {
    let input = String::from(
        "47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
",
    );

    let day = Day5::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 143)
}

#[test]
fn test_day5_p2() {
    let input = String::from("");

    let day = Day5::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
