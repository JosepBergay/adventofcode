import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 5;

type Coords = [number, number];
type Line = [Coords, Coords];
type ParsedInput = Line[];
type FullMap<T> = {
  [x: number]: T;
};
type Map = Partial<FullMap<Partial<FullMap<number>>>>;

const parser = (input: string): ParsedInput => {
  const rows = input.split("\n");
  rows.splice(-1);
  const coords = rows.map((row) =>
    row.split(" -> ").map((coords) => coords.split(",").map((c) => parseInt(c)))
  );
  return coords as Line[];
};

const isVertical = (line: Line) => line[0][0] === line[1][0];
const isHorizontal = (line: Line) => line[0][1] === line[1][1];

const increaseMapValue = (map: Map, [x, y]: Coords) => {
  if (!map[x]) map[x] = {};
  const value = map[x]![y];
  map[x]![y] = value ? value + 1 : 1;
};

const countOverlappingPoints = (map: Map) => {
  let overlappingPoints = 0;
  for (const x in map) {
    for (const y in map[x]) {
      const value = map[x]![y as unknown as number]!;
      if (value >= 2) overlappingPoints++;
    }
  }
  return overlappingPoints;
};

const executePart1 = (input: ParsedInput): string => {
  const map: Map = {};

  for (const [start, end] of input) {
    if (isVertical([start, end])) {
      increaseMapValue(map, start);
      increaseMapValue(map, end);

      const height = start[1] - end[1];
      if (height > 0) {
        for (let y = 1; y < height; y++) {
          increaseMapValue(map, [start[0], end[1] + y]);
        }
      } else {
        for (let y = 1; y < Math.abs(height); y++) {
          increaseMapValue(map, [start[0], start[1] + y]);
        }
      }
    }
    if (isHorizontal([start, end])) {
      increaseMapValue(map, start);
      increaseMapValue(map, end);

      const width = start[0] - end[0];
      if (width > 0) {
        for (let x = 1; x < width; x++) {
          increaseMapValue(map, [end[0] + x, start[1]]);
        }
      } else {
        for (let x = 1; x < Math.abs(width); x++) {
          increaseMapValue(map, [start[0] + x, start[1]]);
        }
      }
    }
  }

  const overlappingPoints = countOverlappingPoints(map);

  return `${overlappingPoints}`;
};

const executePart2 = (input: ParsedInput): string => {
  const map: Map = {};

  for (const [start, end] of input) {
    const width = start[0] - end[0];
    const height = start[1] - end[1];

    if (isVertical([start, end])) {
      if (height > 0) {
        for (let y = 0; y <= height; y++) {
          increaseMapValue(map, [start[0], start[1] - y]);
        }
      } else {
        for (let y = 0; y <= Math.abs(height); y++) {
          increaseMapValue(map, [start[0], start[1] + y]);
        }
      }
    } else if (isHorizontal([start, end])) {
      if (width > 0) {
        for (let x = 0; x <= width; x++) {
          increaseMapValue(map, [start[0] - x, start[1]]);
        }
      } else {
        for (let x = 0; x <= Math.abs(width); x++) {
          increaseMapValue(map, [start[0] + x, start[1]]);
        }
      }
    } else {
      // Diagonal. Width and height are equal.
      if (width > 0 && height > 0) {
        for (let y = 0; y <= height; y++) {
          increaseMapValue(map, [start[0] - y, start[1] - y]);
        }
      } else if (width > 0 && height <= 0) {
        for (let y = 0; y <= Math.abs(height); y++) {
          increaseMapValue(map, [start[0] - y, start[1] + y]);
        }
      } else if (width <= 0 && height > 0) {
        for (let y = 0; y <= height; y++) {
          increaseMapValue(map, [start[0] + y, start[1] - y]);
        }
      } else {
        // width <= 0 && height <= 0
        for (let y = 0; y <= Math.abs(height); y++) {
          increaseMapValue(map, [start[0] + y, start[1] + y]);
        }
      }
    }
  }

  return `${countOverlappingPoints(map)}`;
};

const day5: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day5;
