use days::{get_days, Day};
use reqwest::StatusCode;
use std::{env, error, fs, io, time::Instant};

pub mod days;

async fn fetch_day(n: u8) -> Result<String, reqwest::Error> {
    let cookie =
        env::var("SESSION_COOKIE").expect("Environment variable `SESSION_COOKIE` must be set");

    let client = reqwest::Client::new();

    let res = client
        .get(format!("https://adventofcode.com/2024/day/{n}/input"))
        .header("Cookie", format!("session={cookie}"))
        .send()
        .await?;

    if res.status() != StatusCode::OK {
        panic!("Error fetching day {n}");
    }

    res.text().await
}

async fn get_input(n: u8) -> Result<String, Box<dyn error::Error>> {
    let dir_path = "src/2024/days";

    let file_path = format!("{dir_path}/day{n}.txt");

    let res = fs::read_to_string(&file_path);

    let input = match res {
        Ok(txt) => txt,
        Err(e) => match e.kind() {
            io::ErrorKind::NotFound => {
                let input = fetch_day(n).await?;

                fs::create_dir_all(dir_path)?;

                fs::write(file_path, &input)?;

                input
            }
            other => panic!("error reading file {other}"),
        },
    };

    Ok(input)
}

async fn run_day(n: u8, day: &Box<dyn Day>) -> Result<(), Box<dyn error::Error>> {
    let input = get_input(n).await?;

    let now = Instant::now();

    let result = day.exec(input).expect(&format!("Error running day {n}"));

    let elapsed = now.elapsed();

    println!(
        "Day {n}: [Part1]: {} [Part2]: {} ({:?})", // {:.2?}
        result.part1, result.part2, elapsed
    );

    Ok(())
}

#[tokio::main]
async fn main() {
    let args = env::args()
        .skip(1)
        .filter_map(|n| {
            n.parse::<u8>()
                .ok()
                .and_then(|n| if 1 <= n && n <= 25 { Some(n) } else { None })
        })
        .collect::<Vec<_>>();

    let day_map = get_days();

    let mut day_args: Vec<&u8> = if args.is_empty() {
        day_map.keys().collect()
    } else {
        args.iter().collect()
    };

    day_args.sort();

    println!("Running days: {day_args:?}");

    let now = Instant::now();

    for &day_num in &day_args {
        let entry = day_map.get(day_num);

        let _ = match entry {
            Some(day) => run_day(*day_num, day).await,
            None => panic!("Day {day_num} was not added to map"),
        };
    }

    println!("Ran {} days in {:?}", day_args.len(), now.elapsed());
}
