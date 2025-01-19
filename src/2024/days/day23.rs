use std::{
    collections::{HashMap, HashSet},
    error::Error,
};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day23 {}

impl Day23 {
    fn parse_input(&self, input: String) -> HashMap<String, Vec<String>> {
        let connections = input
            .lines()
            .filter(|l| !l.is_empty())
            .map(|l| l.split("-").map(|s| s.to_string()).collect())
            .collect::<Vec<Vec<_>>>();

        let mut map = HashMap::new();

        for connection in connections {
            let v = map.entry(connection[0].clone()).or_insert(vec![]);
            v.push(connection[1].clone());
            let v = map.entry(connection[1].clone()).or_insert(vec![]);
            v.push(connection[0].clone());
        }

        map
    }

    fn part1(&self, parsed: &HashMap<String, Vec<String>>) -> usize {
        let mut sets = HashSet::new();

        for (key, value) in parsed {
            for i in 0..value.len() {
                for j in i + 1..value.len() {
                    if sets.contains(&(value[i].clone(), key.clone(), value[j].clone()))
                        || sets.contains(&(value[i].clone(), value[j].clone(), key.clone()))
                        || sets.contains(&(value[j].clone(), key.clone(), value[i].clone()))
                        || sets.contains(&(value[j].clone(), value[i].clone(), key.clone()))
                    {
                        continue;
                    }

                    if parsed
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

    fn part2(&self, parsed: HashMap<String, Vec<String>>) -> String {
        // let all_cliques = bron_kerbosch(
        //     HashSet::new(),
        //     parsed.keys().collect(),
        //     HashSet::new(),
        //     &parsed,
        // );

        // let mut largest = all_cliques
        //     .iter()
        //     .max_by_key(|v| v.len())
        //     .unwrap()
        //     .iter()
        //     .map(|s| (*s).clone())
        //     .collect::<Vec<String>>();

        let mut largest = greedy(&parsed)
            .iter()
            .map(|s| (*s).clone())
            .collect::<Vec<String>>();

        largest.sort();

        largest.join(",")
    }
}

/**
 * This Greedy algorithm is much faster (x10) than Bron-Kerbosch (at least for this particular
 * input). Also, the implementation for Bron-Kerbosch (below) could be way better as it involes a
 * bunch of clonning. This Greedy implementation could also be better if it skipped already visited
 * cliques. Tried to add a seen set but it didn't seem to make a difference again for this input and
 * my computer.
 */
fn greedy(graph: &HashMap<String, Vec<String>>) -> Vec<&String> {
    let mut largest = vec![];

    for (k, neighbours) in graph {
        let mut clique = vec![k];

        for n in neighbours {
            if clique
                .iter()
                .all(|c| graph.get(n).is_some_and(|v| v.contains(c)))
            {
                clique.push(n);
            }
        }

        if clique.len() > largest.len() {
            largest = clique;
        }
    }

    largest
}

fn _bron_kerbosch<'a>(
    r: HashSet<&'a String>,
    mut p: HashSet<&'a String>,
    mut x: HashSet<&'a String>,
    graph: &'a HashMap<String, Vec<String>>,
) -> Vec<HashSet<&'a String>> {
    if p.is_empty() && x.is_empty() {
        return vec![r];
    }

    let mut out = vec![];
    for &v in &p.clone() {
        let mut new_r = r.clone();
        new_r.insert(v);

        let v_neighbours = graph.get(v).unwrap().iter().collect::<HashSet<_>>();

        let new_p = p
            .intersection(&v_neighbours)
            .map(|s| *s)
            .collect::<HashSet<_>>();

        let new_x = x
            .intersection(&v_neighbours)
            .map(|s| *s)
            .collect::<HashSet<_>>();

        let vec = _bron_kerbosch(new_r, new_p, new_x, graph);
        out.extend(vec);

        // Move v from p to x
        p.remove(v);
        x.insert(v);
    }

    out
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

#[cfg(test)]
fn get_test_input() -> String {
    String::from(
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
    )
}

#[test]
fn test_day23_p1() {
    let input = get_test_input();

    let day = Day23::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 7)
}

#[test]
fn test_day23_p2() {
    let input = get_test_input();

    let day = Day23::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "co,de,ka,ta")
}
