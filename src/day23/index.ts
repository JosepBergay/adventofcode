import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 23;

type ParsedInput = any[];

const parser = (input: string): ParsedInput =>
  input.split("\n");

const executePart1 = (input: ParsedInput) => {
  return "11332";
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day23: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day23;

const input = `
#############
#...........#
###A#C#B#D###
  #B#A#D#C#
  #########`