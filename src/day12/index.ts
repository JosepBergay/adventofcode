import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import { logGrid } from "../gridUtils.js";

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

const getNextCave = (
  lastCave: string,
  link: Link,
  isNextCave: (nextCave: string) => boolean
) => {
  const nextCave =
    lastCave === link[0] ? link[1] : lastCave === link[1] ? link[0] : undefined;
  return nextCave && isNextCave(nextCave) ? nextCave : undefined;
};

const notVisitedSmallCave = (cave: string, current: string[]) =>
  isBigCave(cave) || !current.includes(cave);

const noSmallCaveVisitedTwice = (current: string[]) => {
  const smallCaves = current.filter(isSmallCave);
  const repeated = smallCaves.filter(
    (c1) => smallCaves.filter((c2) => c2 === c1).length === 2
  );
  return repeated.length === 0;
};

const notVisitedTooManySmallCaves = (cave: string, current: string[]) =>
  notVisitedSmallCave(cave, current) ||
  (cave !== start && noSmallCaveVisitedTwice(current));

const computePaths = (
  current: string[],
  links: Link[],
  isNextCave: (nextCave: string, currentPath: string[]) => boolean
): string[][] => {
  const lastCave = current[current.length - 1];
  if (lastCave === end) {
    return [current];
  }
  const newPaths: string[][] = [];
  for (const link of links) {
    const nextCave = getNextCave(lastCave, link, (c) => isNextCave(c, current));
    if (nextCave) {
      newPaths.push(...computePaths([...current, nextCave], links, isNextCave));
    }
  }
  return newPaths;
};

const executePart1 = (input: ParsedInput): string => {
  const paths = computePaths([start], input, notVisitedSmallCave);

  return `${paths.length}`;
};

const executePart2 = (input: ParsedInput): string => {
  const paths = computePaths([start], input, notVisitedTooManySmallCaves);

  return `${paths.length}`;
};

const day12: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day12;
