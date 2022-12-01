import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import { copyGrid, Grid, logGrid, Point } from "../gridUtils.js";

const level = 25;

type ParsedInput = Grid<"." | ">" | "v">;

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l !== "")
    .map((l) => l.split("") as ("." | ">" | "v")[]);

const getNextPos = ([x, y]: Point, isEast: boolean, grid: ParsedInput) => {
  let nextPos: Point;
  if (isEast) {
    nextPos = x === grid[0].length - 1 ? [0, y] : [x + 1, y];
  } else {
    nextPos = y === grid.length - 1 ? [x, 0] : [x, y + 1];
  }
  return grid[nextPos[1]][nextPos[0]] === "." ? nextPos : null;
};

const moveEast = (grid: ParsedInput) => {
  const mustMove: [from: Point, to: Point][] = [];
  for (const [y, row] of grid.entries()) {
    for (const [x, c] of row.entries()) {
      if (c === ">") {
        const nextPos = getNextPos([x, y], true, grid);
        if (nextPos) {
          mustMove.push([[x, y], nextPos]);
        }
      }
    }
  }
  for (const [from, to] of mustMove) {
    grid[from[1]][from[0]] = ".";
    grid[to[1]][to[0]] = ">";
  }
  return !!mustMove.length;
};

const moveSouth = (grid: ParsedInput) => {
  const mustMove: [from: Point, to: Point][] = [];
  for (const [y, row] of grid.entries()) {
    for (const [x, c] of row.entries()) {
      if (c === "v") {
        const nextPos = getNextPos([x, y], false, grid);
        if (nextPos) {
          mustMove.push([[x, y], nextPos]);
        }
      }
    }
  }
  for (const [from, to] of mustMove) {
    grid[from[1]][from[0]] = ".";
    grid[to[1]][to[0]] = "v";
  }
  return !!mustMove.length;
};

const moveCucumbers = (grid: ParsedInput) => {
  const movedEast = moveEast(grid);
  const movedSouth = moveSouth(grid);
  return movedEast || movedSouth;
};

const executePart1 = (input: ParsedInput) => {
  const grid = copyGrid(input);
  let step = 0;
  let moved = true;
  while (moved) {
    moved = moveCucumbers(grid);
    step++;
  }
  return step;
};

const executePart2 = (input: ParsedInput) => {
  return "🎉";
};

const day25: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day25;
