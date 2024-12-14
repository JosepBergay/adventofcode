use std::error::Error;

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day9 {}

impl Day9 {
    fn parse_input(&self, input: String) -> String {
        input
    }

    fn part1(&self, parsed: &str) -> usize {
        let aux = parsed
            .chars()
            .filter(|c| *c != '\n')
            .enumerate()
            .flat_map(|(i, c)| {
                let block_len = c.to_digit(10).unwrap();

                let id = if i % 2 == 0 { Some(i / 2) } else { None };

                (0..block_len).map(move |_| id)
            })
            .collect::<Vec<_>>();

        let mut ordered = vec![];
        let mut idx = 0;
        let mut last_idx = aux.len();

        while idx < last_idx {
            let opt = aux[idx];
            if let Some(c) = opt {
                ordered.push(c);
            } else if let Some((i, o)) = aux[idx..last_idx]
                .iter()
                .enumerate()
                .rfind(|(_, &o)| o.is_some())
            {
                ordered.push(o.unwrap());
                last_idx = idx + i;
            }
            idx += 1;
        }

        ordered
            .iter()
            .enumerate()
            .fold(0, |acc, (i, v)| acc + i * (*v as usize))
    }

    fn part2(&self, _parsed: String) -> usize {
        0
    }
}

impl Day for Day9 {
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
fn test_day9_p1() {
    let input = String::from("2333133121414131402");

    let day = Day9::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed);

    assert_eq!(res, 1928)
}

#[test]
fn test_day9_p2() {
    let input = String::from("");

    let day = Day9::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 2858)
}
