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
  for (const crabPosition of crabs) {
    fuel = fuel + Math.abs(median - crabPosition);
  }

  return `${fuel}`;
};

const executePart2 = (input: ParsedInput): string => {
  const crabs = input.slice();

  const computeFuelPerCrab = (from: number, to: number) => {
    let n = 0;
    for (let i = 0; i <= Math.abs(from - to); i++) {
      n = n + i;
    }
    return n;
  };

  const computeFuel = (position: number) =>
    crabs.reduce((total, c) => computeFuelPerCrab(c, position) + total, 0);

  let bestPosition = -1;
  let minimumFuel = Infinity;
  for (let pos = 0; pos <= Math.max(...crabs); pos++) {
    if (bestPosition !== pos) {
      const fuel = computeFuel(pos);
      if (fuel < minimumFuel) {
        minimumFuel = fuel;
        bestPosition = pos;
      }
    }
  }

  return `${minimumFuel}`;
};

const day7: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day7;
