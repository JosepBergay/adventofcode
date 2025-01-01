use std::{
    collections::{HashMap, VecDeque},
    error::Error,
};

use super::{baseday::DayResult, point2d::Point2D, Day};

#[derive(Default)]
pub struct Day18 {}

impl Day18 {
    fn parse_input(&self, input: String) -> (Vec<Point2D>, Point2D) {
        let mut end = Point2D { x: 0, y: 0 };

        let coords = input
            .lines()
            .filter(|l| !l.is_empty())
            .map(|l| {
                let mut split = l.split(",");
                let x = split.next().unwrap().parse::<i32>().unwrap();
                let y = split.next().unwrap().parse::<i32>().unwrap();
                end.x = end.x.max(x);
                end.y = end.y.max(y);
                Point2D { x, y }
            })
            .collect();

        (coords, end)
    }

    fn part1(&self, parsed: &(Vec<Point2D>, Point2D), max_len: usize) -> usize {
        let (blocks, end) = parsed;

        find_path_count_bfs(end, &blocks[..max_len]).unwrap()
    }

    fn part2(&self, parsed: (Vec<Point2D>, Point2D), max_len: usize) -> String {
        let (blocks, end) = parsed;

        for i in ((max_len + 1)..blocks.len()).rev() {
            if find_path_count_bfs(&end, &blocks[..i]).is_some() {
                return format!("{},{}", blocks[i].x, blocks[i].y);
            }
        }

        panic!("Result not found")
    }
}

fn find_path_count_bfs(end: &Point2D, blocks: &[Point2D]) -> Option<usize> {
    let mut q = VecDeque::new();
    q.push_back((Point2D { x: 0, y: 0 }, 0));

    let mut costs = HashMap::new();

    let dirs = vec![
        Point2D { x: 1, y: 0 },
        Point2D { x: -1, y: 0 },
        Point2D { x: 0, y: 1 },
        Point2D { x: 0, y: -1 },
    ];

    while let Some((p, acc)) = q.pop_front() {
        if p == *end {
            return Some(acc);
        }

        for d in &dirs {
            let next = *d + p;
            if !blocks.contains(&next)
                && next.x >= 0
                && next.y >= 0
                && next.x <= end.x
                && next.y <= end.y
                && costs.get(&next).is_none_or(|c| *c > acc + 1)
            {
                costs.insert(next, acc + 1);
                q.push_back((next, acc + 1));
            }
        }
    }

    None
}

impl Day for Day18 {
    fn exec(&self, input: String) -> Result<DayResult, Box<dyn Error>> {
        let parsed = self.parse_input(input);

        let p1 = self.part1(&parsed, 1024);
        let p2 = self.part2(parsed, 1024);

        Ok(DayResult {
            part1: p1.to_string(),
            part2: p2.to_string(),
        })
    }
}

#[cfg(test)]
fn get_test_input() -> String {
    String::from(
        "5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
",
    )
}

#[test]
fn test_day18_p1() {
    let input = get_test_input();

    let day = Day18::default();
    let parsed = day.parse_input(input);
    let res = day.part1(&parsed, 12);

    assert_eq!(res, 22)
}

#[test]
fn test_day18_p2() {
    let input = get_test_input();

    let day = Day18::default();
    let parsed = day.parse_input(input);
    let res = day.part2(parsed, 12);

    assert_eq!(res, "6,1")
}
