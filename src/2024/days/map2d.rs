use super::point2d::Point2D;

pub struct Map2D<T> {
    pub map: Vec<Vec<T>>,
}

impl<T> Map2D<T> {
    pub fn new(map: Vec<Vec<T>>) -> Map2D<T> {
        Map2D { map }
    }

    pub fn from_string(input: String) -> Map2D<char> {
        let map = input
            .split('\n')
            .filter(|line| !line.is_empty())
            .map(|line| line.chars().collect())
            .collect::<Vec<Vec<char>>>();

        Map2D::new(map)
    }

    pub fn is_out_of_bounds(&self, p: Point2D) -> bool {
        if p.y < 0 || self.map.len() <= p.y.try_into().unwrap() {
            return true;
        }

        p.x < 0 || self.map[0].len() <= p.x.try_into().unwrap()
    }

    pub fn get(&self, p: Point2D) -> Option<&T> {
        if self.is_out_of_bounds(p) {
            None
        } else {
            let v = &self.map[usize::try_from(p.y).expect(format!("y {p:?}").as_str())]
                [usize::try_from(p.x).expect(format!("x {p:?}").as_str())];
            Some(v)
        }
    }

    pub fn get_adjacent(&self, curr: Point2D) -> Vec<Point2D> {
        let dirs = vec![
            Point2D { x: 1, y: 0 },
            Point2D { x: -1, y: 0 },
            Point2D { x: 0, y: 1 },
            Point2D { x: 0, y: -1 },
        ];

        dirs.iter()
            .map(|d| curr + *d)
            .filter(|p| !self.is_out_of_bounds(*p))
            .collect()
    }

    pub fn iter(&self) -> Map2DIterator<T> {
        Map2DIterator {
            map2d: self,
            curr: Point2D { x: -1, y: 0 },
        }
    }
}

pub struct Map2DIterator<'a, T> {
    curr: Point2D,
    map2d: &'a Map2D<T>,
}

impl<'a, T> Iterator for Map2DIterator<'a, T> {
    type Item = Point2D;

    fn next(&mut self) -> Option<Self::Item> {
        let left = Point2D {
            x: self.curr.x + 1,
            y: self.curr.y,
        };

        if !self.map2d.is_out_of_bounds(left) {
            self.curr = left;
            return Some(left);
        }

        let start_of_next_line = Point2D {
            x: 0,
            y: self.curr.y + 1,
        };

        if !self.map2d.is_out_of_bounds(start_of_next_line) {
            self.curr = start_of_next_line;
            return Some(start_of_next_line);
        }

        None
    }
}
