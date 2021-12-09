import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 1;

const executePart1 = (input: number[]) => {
  let increased = 0;

  input.reduce((p, c) => {
    if (c > p) ++increased;
    return c;
  });

  return `${increased}`;
};

const executePart2 = (input: number[]) => {
  let increased = 0;
  let previous;

  for (let i = 0; i + 2 < input.length; i++) {
    const current = input[i] + input[i + 1] + input[i + 2];
    if (previous != undefined && current > previous) {
      increased++;
    }
    previous = current;
  }

  return `${increased}`;
};

export const day1: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = input.split("\n").map((i) => parseInt(i));

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};
