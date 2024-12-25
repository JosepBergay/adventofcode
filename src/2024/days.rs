use std::collections::HashMap;

pub use baseday::Day;
use day1::Day1;
use day10::Day10;
use day11::Day11;
use day12::Day12;
use day13::Day13;
use day14::Day14;
use day15::Day15;
use day2::Day2;
use day3::Day3;
use day4::Day4;
use day5::Day5;
use day6::Day6;
use day7::Day7;
use day8::Day8;
use day9::Day9;

pub mod baseday;
pub mod day1;
pub mod day10;
pub mod day11;
pub mod day12;
pub mod day13;
pub mod day14;
pub mod day15;
pub mod day2;
pub mod day3;
pub mod day4;
pub mod day5;
pub mod day6;
pub mod day7;
pub mod day8;
pub mod day9;
pub mod map2d;
pub mod point2d;

pub fn get_days() -> HashMap<u8, Box<dyn Day>> {
    let all_days = HashMap::from([
        (1, Box::new(Day1::default()) as Box<dyn Day>),
        (2, Box::new(Day2::default()) as Box<dyn Day>),
        (3, Box::new(Day3::default()) as Box<dyn Day>),
        (4, Box::new(Day4::default()) as Box<dyn Day>),
        (5, Box::new(Day5::default()) as Box<dyn Day>),
        (6, Box::new(Day6::default()) as Box<dyn Day>),
        (6, Box::new(Day6::default()) as Box<dyn Day>),
        (7, Box::new(Day7::default()) as Box<dyn Day>),
        (7, Box::new(Day7::default()) as Box<dyn Day>),
        (8, Box::new(Day8::default()) as Box<dyn Day>),
        (8, Box::new(Day8::default()) as Box<dyn Day>),
        (9, Box::new(Day9::default()) as Box<dyn Day>),
        (9, Box::new(Day9::default()) as Box<dyn Day>),
        (10, Box::new(Day10::default()) as Box<dyn Day>),
        (10, Box::new(Day10::default()) as Box<dyn Day>),
        (11, Box::new(Day11::default()) as Box<dyn Day>),
        (11, Box::new(Day11::default()) as Box<dyn Day>),
        (12, Box::new(Day12::default()) as Box<dyn Day>),
        (12, Box::new(Day12::default()) as Box<dyn Day>),
        (13, Box::new(Day13::default()) as Box<dyn Day>),
        (13, Box::new(Day13::default()) as Box<dyn Day>),
        (14, Box::new(Day14::default()) as Box<dyn Day>),
        (14, Box::new(Day14::default()) as Box<dyn Day>),
        (15, Box::new(Day15::default()) as Box<dyn Day>),
    ]);

    all_days
}
