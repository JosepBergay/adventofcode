import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 10;

type ParsedInput = string[];

const parser = (input: string): ParsedInput =>
  input.split("\n").filter((row) => row != "");

type ChunkStart = "(" | "[" | "{" | "<";
type ChunkEnd = ")" | "]" | "}" | ">";

const matching: Record<ChunkStart, ChunkEnd> = {
  "(": ")",
  "<": ">",
  "[": "]",
  "{": "}",
};

const chunkEndScore: Record<ChunkEnd, number> = {
  ")": 3,
  ">": 25137,
  "]": 57,
  "}": 1197,
};

const isChunkStart = (s: string): s is ChunkStart =>
  s === "(" || s === "[" || s === "{" || s === "<";

const isChunkEnd = (s: string): s is ChunkEnd =>
  s === ")" || s === "]" || s === "}" || s === ">";

const findIllegalChar = (line: string) => {
  const chunkStarts: ChunkStart[] = [];
  for (const char of line) {
    if (isChunkStart(char)) {
      chunkStarts.push(char);
    } else if (isChunkEnd(char)) {
      const lastStart = chunkStarts.pop();
      if (!lastStart || matching[lastStart] !== char) {
        return char;
      }
    }
  }
};

const executePart1 = (input: ParsedInput): string => {
  const illegalChars: ChunkEnd[] = [];

  for (const line of input) {
    const illegalChar = findIllegalChar(line);
    if (illegalChar) {
      illegalChars.push(illegalChar);
    }
  }

  const score = illegalChars.reduce((p, c) => p + chunkEndScore[c], 0);

  return `${score}`;
};

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day10: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day10;
