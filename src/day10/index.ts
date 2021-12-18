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

const syntaxCheckerScore: Record<ChunkEnd, number> = {
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

  const score = illegalChars.reduce((p, c) => p + syntaxCheckerScore[c], 0);

  return `${score}`;
};

const autoCompleteScore: Record<ChunkEnd, number> = {
  ")": 1,
  "]": 2,
  "}": 3,
  ">": 4,
};

const findIncompleteStarts = (line: string) => {
  const chunkStarts: ChunkStart[] = [];
  for (const char of line) {
    if (isChunkStart(char)) {
      chunkStarts.push(char);
    } else if (isChunkEnd(char)) {
      chunkStarts.pop();
    }
  }
  return chunkStarts;
};

const computeLineScore = (chunkStarts: ChunkStart[]) =>
  chunkStarts
    .reverse()
    .reduce((p, s) => p * 5 + autoCompleteScore[matching[s]], 0);

const executePart2 = (input: ParsedInput): string => {
  const incompleteStarts: ChunkStart[][] = [];
  for (const line of input) {
    if (!findIllegalChar(line)) {
      // Line is incomplete
      const incomplete = findIncompleteStarts(line);
      if (incomplete.length) {
        incompleteStarts.push(incomplete);
      }
    }
  }

  const lineScores = incompleteStarts
    .map(computeLineScore)
    .sort((a, b) => a - b);
  const middleScore = lineScores[Math.round((lineScores.length - 1) / 2)];

  return `${middleScore}`;
};

const day10: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day10;
