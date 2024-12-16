use std::{cmp, error::Error, usize};

use super::{baseday::DayResult, Day};

#[derive(Default)]
pub struct Day9 {}

impl Day9 {
    fn parse_input(&self, input: String) -> Vec<Option<usize>> {
        input
            .chars()
            .filter(|c| *c != '\n')
            .enumerate()
            .flat_map(|(i, c)| {
                let block_len = c.to_digit(10).unwrap();

                let id = if i % 2 == 0 { Some(i / 2) } else { None };

                (0..block_len).map(move |_| id)
            })
            .collect::<Vec<_>>()
    }

    fn part1(&self, parsed: &Vec<Option<usize>>) -> usize {
        let mut acc = 0;
        let mut idx = 0;
        let mut last_idx = parsed.len();

        while idx < last_idx {
            let opt = parsed[idx];

            if let Some(id) = opt {
                acc += idx * id
            } else if let Some((i, o)) = parsed[idx..last_idx]
                .iter()
                .enumerate()
                .rfind(|(_, o)| o.is_some())
            {
                acc += idx * o.unwrap();
                last_idx = idx + i;
            }

            idx += 1;
        }

        acc
    }

    fn part2(&self, mut parsed: Vec<Option<usize>>) -> usize {
        let mut last_idx = parsed.len() - 1;

        let mut min_space_available = usize::MAX;

        loop {
            let last = parsed[..=last_idx]
                .iter()
                .enumerate()
                .rfind(|(_, o)| o.is_some());

            if let Some((i, Some(id))) = last {
                let opt = parsed[..i]
                    .iter()
                    .rposition(|o| o.is_none_or(|id2| id2 != *id));

                if opt.is_none() {
                    break;
                }

                let j = opt.unwrap();

                let file_size = i - j;
                last_idx = j;

                if file_size >= min_space_available {
                    continue;
                }

                if let Some(insert_at) = find_empty_space(file_size, &parsed[..=j]) {
                    for k in 0..file_size {
                        parsed.swap(insert_at + k, last_idx + 1 + k);
                    }
                } else {
                    min_space_available = cmp::min(file_size, min_space_available);
                }
            }
        }

        parsed
            .iter()
            .enumerate()
            .fold(0, |acc, (i, opt)| match opt {
                Some(id) => acc + i * id,
                None => acc,
            })
    }
}

fn find_empty_space(size: usize, parsed: &[Option<usize>]) -> Option<usize> {
    let mut empty_space = 0;

    for (i, opt) in parsed.iter().enumerate() {
        if opt.is_some() && empty_space != 0 {
            empty_space = 0;
        } else if opt.is_none() {
            empty_space += 1;

            if empty_space == size {
                return Some(1 + i - size);
            }
        }
    }

    None
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
    let input = String::from("2333133121414131402");

    let day = Day9::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed);

    assert_eq!(res, 2858)
}
