use std::error::Error;

use regex::Regex;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day17 {}

#[derive(Debug, Clone)]
struct Register {
    a: usize,
    b: usize,
    c: usize,
}

impl Day17 {
    fn parse_input(&self, input: String) -> (Register, Vec<u8>) {
        let re = Regex::new(
            r"Register A: (\d+)
Register B: (\d+)
Register C: (\d+)

Program: (.*)
",
        )
        .expect("Invalid regex");

        for capture in re.captures_iter(input.as_str()) {
            let (_, rest): (&str, [&str; 4]) = capture.extract();
            let mut rest = rest.iter();
            let a = rest.next().unwrap().parse::<usize>().unwrap();
            let b = rest.next().unwrap().parse::<usize>().unwrap();
            let c = rest.next().unwrap().parse::<usize>().unwrap();
            let program = rest
                .next()
                .unwrap()
                .split(",")
                .map(|c| c.parse::<u8>().unwrap())
                .collect();

            return (Register { a, b, c }, program);
        }

        panic!("Invalid input")
    }

    fn part1(&self, reg: &mut Register, program: &Vec<u8>) -> String {
        let output = run_program(reg, program);

        output
            .iter()
            .map(|d| d.to_string())
            .collect::<Vec<_>>()
            .join(",")
    }

    fn part2(&self, program: Vec<u8>) -> usize {
        // Program in reverse
        // 3,0 -> Loop While A != 0
        // 5,5 -> Output B % 8          A,0,C
        // 0,3 -> A = A / 8             8*A,0,C
        // 1,4 -> B = B XOR 4           8*A,4,C
        // 4,7 -> B = B XOR C           8*A,
        // 7,5 -> C = A / 2**B
        // 1,1 -> B = B XOR 1
        // 2,4 -> B = A % 8

        let mut aux = find_register_a(program.len() - 1, 0, &program);

        aux.sort();

        aux[0]
    }
}

fn find_register_a(curr_idx: usize, a: usize, program: &Vec<u8>) -> Vec<usize> {
    let low = a * 8;
    let high = (a + 1) * 8;
    let slice = &program[curr_idx..];

    let mut results = vec![];

    for i in low..high {
        let output = run_program(&mut Register { a: i, b: 0, c: 0 }, &program);

        if output.as_slice() == slice {
            if curr_idx == 0 {
                results.push(i);
            } else {
                let res = find_register_a(curr_idx - 1, i, program);
                if !res.is_empty() {
                    results.extend(res);
                }
            }
        }
    }

    results
}

fn run_program(reg: &mut Register, program: &Vec<u8>) -> Vec<u8> {
    let mut output = vec![];
    let program_len = program.len();

    let mut instruction_ptr = 0;

    while instruction_ptr < program_len {
        let opcode = program[instruction_ptr];
        let operand = program[instruction_ptr + 1];

        match opcode {
            // adv
            0 => {
                let combo_operand = get_combo_operand(reg, operand);
                reg.a = reg.a / 2_usize.pow(combo_operand as u32);
                instruction_ptr += 2;
            }
            // bxl
            1 => {
                reg.b = reg.b ^ operand as usize;
                instruction_ptr += 2;
            }
            // bst
            2 => {
                let combo_operand = get_combo_operand(reg, operand);
                reg.b = combo_operand % 8;
                instruction_ptr += 2;
            }
            // jnz
            3 => {
                if reg.a != 0 {
                    instruction_ptr = operand as usize;
                } else {
                    instruction_ptr += 2;
                }
            }
            // bxc
            4 => {
                reg.b = reg.b ^ reg.c;
                instruction_ptr += 2;
            }
            // out
            5 => {
                let combo_operand = get_combo_operand(reg, operand);
                output.push((combo_operand % 8) as u8);
                instruction_ptr += 2;
            }
            // bdv
            6 => {
                let combo_operand = get_combo_operand(reg, operand);
                reg.b = reg.a / 2_usize.pow(combo_operand as u32);
                instruction_ptr += 2;
            }
            // cdv
            7 => {
                let combo_operand = get_combo_operand(reg, operand);
                reg.c = reg.a / 2_usize.pow(combo_operand as u32);
                instruction_ptr += 2;
            }
            _ => {
                panic!("Unknown instruction")
            }
        }
    }

    output
}

fn get_combo_operand(reg: &Register, operand: u8) -> usize {
    match operand {
        4 => reg.a,
        5 => reg.b,
        6 => reg.c,
        7 => {
            // println!("FOK {reg:?} {opcode} {operand} {it}");
            panic!("Invalid program")
        }
        i => i as usize,
    }
}

impl Day for Day17 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&mut parsed.0.clone(), &parsed.1);
        let p2 = self.part2(parsed.1);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[test]
fn test_day17_p1() {
    let input = String::from(
        "Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
",
    );

    let day = Day17::default();
    let (mut reg, program) = day.parse_input(input);
    let res = day.part1(&mut reg, &program);

    assert_eq!(res, "4,6,3,5,6,3,5,2,1,0");
}

#[test]
fn test_day17_p2() {
    let input = String::from(
        "Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
",
    );

    let day = Day17::default();
    let (_, program) = day.parse_input(input);
    let res = day.part2(program);

    assert_eq!(res, 117440)
}
