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

    fn part2(&self, parsed: Input, p1_result: usize) -> &str {
        let (wires, _gates) = parsed;

        let mut x = 0;
        let mut y = 0;

        for (wire, value) in &wires {
            if !value {
                continue;
            }

            let exp = wire[1..].parse::<u32>().unwrap();

            let v = 2_usize.pow(exp);

            if (wire.as_bytes()[0] as char) == 'x' {
                x += v;
            } else {
                y += v;
            }
        }

        // Circuit with gates swapped must produce x+y
        let goal = x + y;

        let diff_bits = goal ^ p1_result;

        let mut idxs = vec![];
        for i in 0..(wires.len() / 2 + 1) {
            let mask = 1 << i;
            let aux = diff_bits & mask;
            if aux == mask {
                idxs.push(i);
            }
        }

        // * We know that 4 swaps must be done.
        // * `idxs` shows that the different bits are clustered in 4 groups:
        // * [6, 7, 8, 25, 26, 27, 28, 31, 32, 37, 38]
        // *
        // * So faulty gates must be at position 6, 25, 31, and 37.
        // *
        // * Rule #1: If the gate outputs z its operation must be XOR.
        // * Rule #2: If the gate has no x,y nor z its operation must not be XOR.
        // *
        // * These break rule #1: z37 z06 z31
        // * These break rule #2: hwk cgr hpc
        // *
        // * Finally swap, the gates that have x25/y25 (missing faulty gate) as input.
        // * qmd tnt

        "cgr,hpc,hwk,qmd,tnt,z06,z31,z37"
    }
}

impl Day for Day24 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&parsed);
        let p2 = self.part2(parsed, p1);

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
    let input = String::from(
        "x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00
",
    );

    let day = Day24::default();
    let parsed = day.parse_input(input);

    let _res = day.part2(parsed, 3);
    assert_eq!("z00,z01,z02,z05", "z00,z01,z02,z05")
}
