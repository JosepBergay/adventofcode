import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 7;

type ParsedInput = number[];

const parser = (input: string): ParsedInput =>
  input.split(",").map((s) => parseInt(s));

const executePart1 = (input: ParsedInput): string => {
  const crabs = input.slice();

  crabs.sort((a, b) => a - b);
  const median = crabs[crabs.length / 2];

  let fuel = 0;
  for (const crabPosition of input) {
    fuel = fuel + Math.abs(median - crabPosition);
  }

  return `${fuel}`;
};

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day7: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day7;
