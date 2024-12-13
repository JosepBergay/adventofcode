use std::{fmt, ops};

#[derive(Copy, Clone, PartialEq, Eq, Hash)]
pub struct Point2D {
    pub x: i32,
    pub y: i32,
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
