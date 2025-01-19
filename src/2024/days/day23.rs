use std::{
    collections::{HashMap, HashSet},
    error::Error,
};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day23 {}

impl Day23 {
    fn parse_input(&self, input: String) -> Vec<Vec<String>> {
        input
            .lines()
            .filter(|l| !l.is_empty())
            .map(|l| l.split("-").map(|s| s.to_string()).collect())
            .collect()
    }

    fn part1(&self, parsed: &Vec<Vec<String>>) -> usize {
        let mut map = HashMap::new();

        for connection in parsed {
            let v = map.entry(connection[0].clone()).or_insert(vec![]);
            v.push(connection[1].clone());
            let v = map.entry(connection[1].clone()).or_insert(vec![]);
            v.push(connection[0].clone());
        }

        let mut sets = HashSet::new();

        for (key, value) in &map {
            for i in 0..value.len() {
                for j in i + 1..value.len() {
                    if sets.contains(&(value[i].clone(), key.clone(), value[j].clone()))
                        || sets.contains(&(value[i].clone(), value[j].clone(), key.clone()))
                        || sets.contains(&(value[j].clone(), key.clone(), value[i].clone()))
                        || sets.contains(&(value[j].clone(), value[i].clone(), key.clone()))
                    {
                        continue;
                    }

                    if map
                        .get(&value[i])
                        .is_some_and(|v| v.iter().any(|v| *v == value[j]))
                    {
                        sets.insert((key.clone(), value[i].clone(), value[j].clone()));
                    }
                }
            }
        }

        sets.iter()
            .filter(|s| s.0.starts_with("t") || s.1.starts_with("t") || s.2.starts_with("t"))
            .count()
    }

    fn part2(&self, _parsed: Vec<Vec<String>>) -> &str {
        "TODO"
    }
}

impl Day for Day23 {
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
fn test_day23_p1() {
    let input = String::from(
        "kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
",
    );

    let day = Day23::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 7)
}

#[test]
fn test_day23_p2() {
    let input = String::from("");

    let day = Day23::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
