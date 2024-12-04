use std::{env, error, fs, io};

use reqwest::StatusCode;

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

async fn run_day(n: u8) -> Result<String, Box<dyn error::Error>> {
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

    // TODO: get day and run it
    dbg!(Ok(input.len().to_string()))
}

#[tokio::main]
async fn main() {
    let args = env::args();

    for (i, arg) in args.enumerate() {
        if i == 0 {
            continue;
        }

        let n = arg
            .parse::<u8>()
            .expect(format!("Invalid argument '{arg:}'.").as_str());

        // TODO: run in parallel
        let res = run_day(n).await.expect(&format!("Error running day {n}"));

        println!("Day {n}: {res}");
    }
}
