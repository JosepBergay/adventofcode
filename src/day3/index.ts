import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 3;

type ParsedInput = string[];

const parser = (input: string): ParsedInput => {
  const parsed = input.split("\n");
  parsed.splice(-1);
  return parsed;
};

const getNumberOfBits = (input: ParsedInput, pos: number) => {
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
  return { numberOfZeroes, numberOfOnes };
};

const findMostCommonBit = (input: ParsedInput, pos: number) => {
  const { numberOfOnes, numberOfZeroes } = getNumberOfBits(input, pos);

  return numberOfOnes > numberOfZeroes ? "1" : "0";
};

const filterRowOnNumberOfBits = (
  input: ParsedInput,
  bitPos: number,
  filterFn: (row: ParsedInput[0], ones: number, zeros: number) => boolean
) => {
  const { numberOfZeroes, numberOfOnes } = getNumberOfBits(input, bitPos);

  return input.filter((row) => filterFn(row, numberOfOnes, numberOfZeroes));
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
  let o2inputs = input;
  let co2inputs = input;

  for (let bitPos = 0; bitPos < input[0].length; bitPos++) {
    // Compute Oxygen Generator rating
    if (o2inputs.length > 1) {
      o2inputs = filterRowOnNumberOfBits(
        o2inputs,
        bitPos,
        (row, ones, zeroes) => row[bitPos] == (ones >= zeroes ? "1" : "0")
      );
    }

    // Compute CO2 Scrubber rating
    if (co2inputs.length > 1) {
      co2inputs = filterRowOnNumberOfBits(
        co2inputs,
        bitPos,
        (row, ones, zeroes) => row[bitPos] == (ones >= zeroes ? "0" : "1")
      );
    }

    if (o2inputs.length == 1 && co2inputs.length == 1) break;
  }

  const o2rating = parseInt(o2inputs[0], 2);

  const co2rating = parseInt(co2inputs[0], 2);

  return `${o2rating * co2rating}`;
};

const day3: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day3;
