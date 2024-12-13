use std::{collections::HashMap, error::Error};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day5 {}

impl Day5 {
    fn parse_input(&self, input: String) -> (HashMap<i32, Vec<i32>>, Vec<Vec<i32>>, Vec<Vec<i32>>) {
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

        let (ordered, unordered): (Vec<Vec<i32>>, Vec<Vec<i32>>) =
            updates.into_iter().partition(|up| {
                let mut seen = vec![];

                for page in up {
                    if all_deps_satisfied(&page, &seen, up, &deps) {
                        seen.push(*page);
                    } else {
                        break;
                    }
                }
                seen.len() == up.len()
            });

        (deps, ordered, unordered)
    }

    fn part1(&self, parsed: &(HashMap<i32, Vec<i32>>, Vec<Vec<i32>>, Vec<Vec<i32>>)) -> i32 {
        let (_, ordered, _) = parsed;

        add_mid_positions(ordered)
    }

    fn part2(&self, parsed: (HashMap<i32, Vec<i32>>, Vec<Vec<i32>>, Vec<Vec<i32>>)) -> i32 {
        let (deps_map, _, unordered) = parsed;

        let mut ordered: Vec<Vec<i32>> = vec![];

        for up in unordered {
            let set = custom_sort(&up, &deps_map);

            ordered.push(Vec::from_iter(set));
        }

        add_mid_positions(&ordered)
    }
}

fn custom_sort(update: &Vec<i32>, deps_map: &HashMap<i32, Vec<i32>>) -> Vec<i32> {
    let mut sorted = vec![];

    let mut to_add = update.clone();

    while !to_add.is_empty() {
        let mut new_to_add = vec![];

        for page in to_add {
            if all_deps_satisfied(&page, &sorted, update, deps_map) {
                sorted.push(page);
            } else {
                new_to_add.push(page);
            }
        }

        to_add = new_to_add;
    }

    sorted
}

fn all_deps_satisfied(
    page: &i32,
    seen: &Vec<i32>,
    update: &Vec<i32>,
    deps_map: &HashMap<i32, Vec<i32>>,
) -> bool {
    deps_map
        .get(page)
        .is_none_or(|deps| deps.iter().all(|d| !update.contains(d) || seen.contains(d)))
}

fn add_mid_positions(updates: &Vec<Vec<i32>>) -> i32 {
    updates.iter().fold(0, |acc, up| {
        let mid = up[(up.len() - 1) / 2];
        acc + mid
    })
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

#[cfg(test)]
fn get_test_input() -> String {
    String::from(
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
    )
}

#[test]
fn test_day5_p1() {
    let input = get_test_input();

    let day = Day5::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 143)
}

#[test]
fn test_day5_p2() {
    let input = get_test_input();

    let day = Day5::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 123)
}
