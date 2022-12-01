import type { AOCDay } from "../types";
import {
  isBotEdge,
  isLeftEdge,
  isRightEdge,
  isTopEdge,
  Point,
} from "../gridUtils.js";
import { fetchInput } from "../helpers.js";

const level = 9;

type ParsedInput = number[][];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l != "")
    .map((l) => l.split("").map((s) => parseInt(s)));

const isLowPoint = (
  [x, y]: Point,
  value: number,
  map: ParsedInput
): boolean => {
  if (value === 9) {
    return false;
  }
  const up = y > 0 && map[y - 1][x];
  const down = y < map.length - 1 && map[y + 1][x];
  const left = x > 0 && map[y][x - 1];
  const right = x < map[0].length - 1 && map[y][x + 1];

  return [up, down, left, right]
    .filter((n) => typeof n === "number")
    .every((n) => value < n);
};

const executePart1 = (input: ParsedInput): string => {
  let riskLvl = 0;
  for (const [y, row] of input.entries()) {
    for (const [x, value] of row.entries()) {
      if (isLowPoint([x, y], value, input)) riskLvl = riskLvl + value + 1;
    }
  }
  return `${riskLvl}`;
};

const findBasinSize = (
  [x, y]: Point,
  map: ParsedInput,
  visitedMap: boolean[][]
) => {
  let basinSize = 0;

  if (visitedMap[y][x]) return basinSize;

  visitedMap[y][x] = true;

  if (map[y][x] !== 9) {
    // Go Top
    if (!isTopEdge([x, y]))
      basinSize += findBasinSize([x, y - 1], map, visitedMap);

    // Go Left
    if (!isLeftEdge([x, y]))
      basinSize += findBasinSize([x - 1, y], map, visitedMap);

    // Go Down
    if (!isBotEdge([x, y], map.length))
      basinSize += findBasinSize([x, y + 1], map, visitedMap);

    // Go Right
    if (!isRightEdge([x, y], map[y].length))
      basinSize += findBasinSize([x + 1, y], map, visitedMap);

    ++basinSize;
  }

  return basinSize;
};

const executePart2 = (input: ParsedInput): string => {
  const visitedMap = input.map((rows) => rows.map((_) => false));

  const basinSizes: number[] = [];
  for (const [y, row] of input.entries()) {
    for (const [x, _] of row.entries()) {
      if (!visitedMap[y][x])
        basinSizes.push(findBasinSize([x, y], input, visitedMap));
    }
  }

  basinSizes.sort((a, b) => b - a);

  return `${basinSizes[0] * basinSizes[1] * basinSizes[2]}`;
};

const day9: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day9;
