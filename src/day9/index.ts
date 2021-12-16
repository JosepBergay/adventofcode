import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 9;

type ParsedInput = number[][];

type Point = [number, number];

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

  return [up, down, left, right].filter(n => typeof n === "number").every(n => value < n);
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

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day9: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day9;
