use std::error::Error;

use super::{baseday::DayResult, Day};

pub struct Day1 {}

impl Day1 {
    fn parse_input(&self, _input: String) {
        todo!()
    }

    fn part1(&self) -> Result<String, Box<dyn Error>> {
        todo!()
    }

    fn part2(&self) -> Result<String, Box<dyn Error>> {
        todo!()
    }
}

impl Day for Day1 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        self.parse_input(input);

        let p1 = self.part1()?;
        let p2 = self.part2()?;

        Ok(DayResult {
            part1: p1,
            part2: p2,
        })
    }
}

#[test]
fn test_day1_p1() {
    let input = String::from(
        "3   4
4   3
2   5
1   3
3   9
3   3",
    );

    let day = Day1 {};
    day.parse_input(input);
    let p1 = day.part1();

    assert_eq!(p1.ok(), Some(String::from("11")))
}
