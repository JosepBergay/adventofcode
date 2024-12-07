use std::error;

pub struct DayResult {
    pub part1: String,
    pub part2: String,
}

pub trait Day {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn error::Error>>;
}
