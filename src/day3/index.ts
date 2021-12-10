import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 3;

type ParsedInput = string[];

const parser = (input: string): ParsedInput => {
  const parsed = input.split("\n");
  parsed.splice(-1);
  return parsed;
};

const findMostCommonBit = (input: ParsedInput, pos: number) => {
  let numberOfZeroes = 0;
  let numberOfOnes = 0;

  for (let index = 0; index < input.length; index++) {
    const bit = input[index][pos];
    if (bit === "0") {
      numberOfZeroes++;
    } else {
      numberOfOnes++;
    }
  }

  return numberOfOnes > numberOfZeroes ? "1" : "0";
};

const executePart1 = (input: ParsedInput): string => {
  let g = "";
  let e = "";

  for (let i = 0; i < input[0].length; i++) {
    const nextBit = findMostCommonBit(input, i);
    g = `${g}${nextBit}`;
    e = `${e}${nextBit === "0" ? "1" : "0"}`;
  }

  const gamma = parseInt(g, 2);
  const epsilon = parseInt(e, 2);

  // Alternate way to compute epsilon. Negate gamma, push 32 'zeroes' from the left and then remove
  // all the 'ones' before the desired 'length'.
  // (~22 >>> 32 & parseInt('11111', 2)) == 9 ; From 10110 to 01001
  // const epsilong = (~gamma >>> 32 & parseInt('111111111111', 2));

  return `${gamma * epsilon}`;
};

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day3: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day3;
