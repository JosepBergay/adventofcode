use std::collections::HashMap;

pub use baseday::Day;

pub mod baseday;

pub fn get_days() -> HashMap<u8, Box<dyn Day>> {
    let all_days = HashMap::from([]);

    all_days
}
