use std::{fmt, ops};

#[derive(Copy, Clone, PartialEq, Eq, Hash, PartialOrd, Ord)]
pub struct Point2D {
    pub x: i32,
    pub y: i32,
}

impl Point2D {
    pub fn new(x: i32, y: i32) -> Point2D {
        Point2D { x, y }
    }
}

impl ops::Add for Point2D {
    type Output = Self;

    fn add(self, other: Self) -> Self::Output {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

impl ops::AddAssign for Point2D {
    fn add_assign(&mut self, rhs: Self) {
        *self = Self {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
        };
    }
}

impl ops::Sub for Point2D {
    type Output = Self;

    fn sub(self, rhs: Self) -> Self::Output {
        Self {
            x: self.x - rhs.x,
            y: self.y - rhs.y,
        }
    }
}

impl ops::SubAssign for Point2D {
    fn sub_assign(&mut self, rhs: Self) {
        *self = Self {
            x: self.x - rhs.x,
            y: self.y - rhs.y,
        };
    }
}

impl ops::Mul<i32> for Point2D {
    type Output = Self;

    fn mul(self, v: i32) -> Self::Output {
        Self {
            x: self.x * v,
            y: self.y * v,
        }
    }
}

impl fmt::Debug for Point2D {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Point2D ({}, {})", self.x, self.y)
    }
}

pub fn get_manhattan_distance(p1: &Point2D, p2: &Point2D) -> i32 {
    (p1.x - p2.x).abs() + (p1.y - p2.y).abs()
}

pub fn cmp_by_manhattan_distance(
    p1: Option<&Point2D>,
    p2: Option<&Point2D>,
    goal: &Point2D,
) -> std::cmp::Ordering {
    if p1.is_some() && p2.is_some() {
        let d1 = get_manhattan_distance(p1.unwrap(), goal);
        let d2 = get_manhattan_distance(p2.unwrap(), goal);

        d1.cmp(&d2)
    } else if p1.is_some() {
        std::cmp::Ordering::Greater
    } else if p2.is_some() {
        std::cmp::Ordering::Less
    } else {
        std::cmp::Ordering::Equal
    }
}

pub fn get_orthogonal_directions() -> Vec<Point2D> {
    vec![
        Point2D { x: 1, y: 0 },
        Point2D { x: -1, y: 0 },
        Point2D { x: 0, y: 1 },
        Point2D { x: 0, y: -1 },
    ]
}
