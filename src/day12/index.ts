import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 12;

type Link = [string, string];

type ParsedInput = Link[];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((r) => r != "")
    .map((r) => r.split("-") as Link);

const isSmallCave = (cave: string) => cave.toLowerCase() === cave;

const isBigCave = (cave: string) => !isSmallCave(cave);

const start = "start";
const end = "end";

const getNextCave = (lastCave: string, link: Link) =>
  lastCave === link[0] ? link[1] : lastCave === link[1] ? link[0] : undefined;

const visitedSmallCave = (cave: string, current: string[]) =>
  isSmallCave(cave) && current.includes(cave);

const computePaths = (current: string[], links: Link[]): string[][] => {
  const lastCave = current[current.length - 1];
  if (lastCave === end) {
    return [current];
  }
  const newPaths: string[][] = [];
  for (const link of links) {
    const nextCave = getNextCave(lastCave, link);
    if (nextCave && !visitedSmallCave(nextCave, current)) {
      newPaths.push(...computePaths([...current, nextCave], links));
    }
  }
  return newPaths;
};

const executePart1 = (input: ParsedInput): string => {
  const paths = computePaths([start], input);

  return `${paths.length}`;
};

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day12: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day12;
