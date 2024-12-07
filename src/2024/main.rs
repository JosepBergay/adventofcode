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
    let args = env::args();

    let days = get_days();

    for (i, arg) in args.enumerate() {
        if i == 0 || i < 1 || i > 25 {
            continue;
        }

        // TODO: run in parallel
        let n = arg
            .parse::<u8>()
            .expect(format!("Invalid argument '{arg:}'.").as_str());

        let entry = days.get(&n);

        let _ = match entry {
            Some(day) => run_day(n, day).await,
            None => panic!("Day {n} was not added to map"),
        };
    }
}
