import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 1;

type ParsedInput = number[];

const parser = (input: string): ParsedInput =>
  input.split("\n").map((i) => parseInt(i));

const executePart1 = (input: ParsedInput): string => {
  let result = 0;

  input.reduce((p, c) => {
    if (c > p) ++result;
    return c;
  });

  return `${result}`;
};

const executePart2 = (input: ParsedInput): string => {
  let result = 0;
  let previous;

  for (let i = 0; i + 2 < input.length; i++) {
    const current = input[i] + input[i + 1] + input[i + 2];
    if (previous != undefined && current > previous) {
      result++;
    }
    previous = current;
  }

  return `${result}`;
};

const day1: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day1;
