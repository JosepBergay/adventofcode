use std::collections::HashMap;

pub use baseday::Day;
use day1::Day1;
use day2::Day2;
use day3::Day3;
use day4::Day4;
use day5::Day5;

pub mod baseday;
pub mod day1;
pub mod day2;
pub mod day3;
pub mod day4;
pub mod day5;
pub mod map2d;
pub mod point2d;

pub fn get_days() -> HashMap<u8, Box<dyn Day>> {
    let all_days = HashMap::from([
        (1, Box::new(Day1::default()) as Box<dyn Day>),
        (2, Box::new(Day2::default()) as Box<dyn Day>),
        (3, Box::new(Day3::default()) as Box<dyn Day>),
        (4, Box::new(Day4::default()) as Box<dyn Day>),
        (5, Box::new(Day5::default()) as Box<dyn Day>),
    ]);

    all_days
}
