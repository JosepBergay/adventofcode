use std::{
    collections::{HashMap, HashSet},
    error::Error,
};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day24 {}

// w1, op, w2, out
type Gate = (String, String, String, String);

type Input = (HashMap<String, bool>, Vec<Gate>);

impl Day24 {
    fn parse_input(&self, input: String) -> Input {
        let mut split = input.split("\n\n");

        let wires = split
            .next()
            .unwrap()
            .to_owned()
            .lines()
            .map(|l| {
                let mut s = l.split(": ");
                let wire = s.next().unwrap().to_owned();
                let value = s.next().unwrap() == "1";
                (wire, value)
            })
            .collect();

        let gates = split
            .next()
            .unwrap()
            .to_owned()
            .lines()
            .map(|l| {
                // "w1 op w2 -> out"
                let mut s = l.split(" ");
                let w1 = s.next().unwrap().to_owned();
                let op = s.next().unwrap().to_owned();
                let w2 = s.next().unwrap().to_owned();
                s.next();
                let out = s.next().unwrap().to_owned();
                (w1, op, w2, out)
            })
            .collect();

        (wires, gates)
    }

    fn part1(&self, parsed: &Input) -> usize {
        let (mut wires, gates) = parsed.clone();
        let mut seen = HashSet::new();

        while gates.len() != seen.len() {
            for g in &gates {
                if seen.contains(&g) {
                    continue;
                }

                if wires.contains_key(&g.0) && wires.contains_key(&g.2) {
                    let out_value = match g.1.as_str() {
                        "AND" => wires[&g.0] & wires[&g.2],
                        "OR" => wires[&g.0] | wires[&g.2],
                        "XOR" => wires[&g.0] ^ wires[&g.2],
                        _ => panic!("Unknown operator"),
                    };
                    wires.insert(g.3.clone(), out_value);
                    seen.insert(g);
                }
            }
        }

        let mut out = wires
            .iter()
            .filter(|(k, _v)| k.starts_with("z"))
            .collect::<Vec<_>>();
        out.sort();

        let s = out
            .iter()
            .rev()
            .map(|(_k, v)| if **v { "1" } else { "0" })
            .collect::<String>();

        usize::from_str_radix(s.as_str(), 2).unwrap()
    }

    fn part2(&self, _parsed: Input) -> &str {
        "TODO"
    }
}

impl Day for Day24 {
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
        "x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
",
    )
}

#[test]
fn test_day24_p1() {
    let input = get_test_input();

    let day = Day24::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 2024)
}

#[test]
fn test_day24_p2() {
    let input = get_test_input();

    let day = Day24::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, "TODO")
}
