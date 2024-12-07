use std::collections::HashMap;

pub use baseday::Day;
use day1::Day1;
use day2::Day2;

pub mod baseday;
pub mod day1;
pub mod day2;

pub fn get_days() -> HashMap<u8, Box<dyn Day>> {
    let all_days = HashMap::from([
        (1, Box::new(Day1::default()) as Box<dyn Day>),
        (2, Box::new(Day2::default()) as Box<dyn Day>),
    ]);

    all_days
}
