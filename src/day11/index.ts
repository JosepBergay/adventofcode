import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import {
  isBotEdge,
  isLeftEdge,
  isRightEdge,
  isTopEdge,
  Point,
} from "../gridUtils.js";

const level = 11;

type Dumbo = {
  hasFlashed: boolean;
  energy: number;
};

type ParsedInput = Dumbo[][];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((s) => s != "")
    .map((row) =>
      row.split("").map((i) => ({ hasFlashed: false, energy: parseInt(i) }))
    );

const increaseEnergyByOne = (dumboGrid: ParsedInput) => {
  for (const dumboRow of dumboGrid) {
    for (const dumbo of dumboRow) {
      dumbo.energy++;
    }
  }
};

const createAdjacentGrid = ([x, y]: Point, dumboGrid: ParsedInput) => {
  const notLeftEdge = !isLeftEdge([x, y]);
  const notRightEdge = !isRightEdge([x, y], dumboGrid[y].length);
  const adjacents: Dumbo[][] = [];

  if (!isTopEdge([x, y])) {
    const topRow: Dumbo[] = [];
    topRow.push(dumboGrid[y - 1][x]);
    if (notLeftEdge) topRow.push(dumboGrid[y - 1][x - 1]);
    if (notRightEdge) topRow.push(dumboGrid[y - 1][x + 1]);
    adjacents.push(topRow);
  }

  const midRow: Dumbo[] = [];
  if (notLeftEdge) midRow.push(dumboGrid[y][x - 1]);
  if (notRightEdge) midRow.push(dumboGrid[y][x + 1]);
  adjacents.push(midRow);

  if (!isBotEdge([x, y], dumboGrid.length)) {
    const botRow: Dumbo[] = [];
    botRow.push(dumboGrid[y + 1][x]);
    if (notLeftEdge) botRow.push(dumboGrid[y + 1][x - 1]);
    if (notRightEdge) botRow.push(dumboGrid[y + 1][x + 1]);
    adjacents.push(botRow);
  }
  return adjacents;
};

const flashAndIncrease = (dumboGrid: Dumbo[][]) => {
  let flashCount = 0;
  for (const [y, dumboRow] of dumboGrid.entries()) {
    for (const [x, dumbo] of dumboRow.entries()) {
      if (dumbo.energy > 9 && !dumbo.hasFlashed) {
        dumbo.hasFlashed = true;
        flashCount++;
        const adjacent = createAdjacentGrid([x, y], dumboGrid);
        increaseEnergyByOne(adjacent);
      }
    }
  }
  if (flashCount) {
    flashCount += flashAndIncrease(dumboGrid);
  }
  return flashCount;
};

const resetFlashed = (dumboGrid: ParsedInput) => {
  for (const dumboRow of dumboGrid) {
    for (const dumbo of dumboRow) {
      if (dumbo.hasFlashed) {
        dumbo.energy = 0;
        dumbo.hasFlashed = false;
      }
    }
  }
};

const executePart1 = (input: ParsedInput): string => {
  const dumboGrid = input.map((row) => row.slice());

  let flashCount = 0;

  for (let step = 0; step < 100; step++) {
    // First increase energy by 1.
    increaseEnergyByOne(dumboGrid);

    // Then, flash and increase energy
    flashCount += flashAndIncrease(dumboGrid);

    // Finally, reset those that flashed
    resetFlashed(dumboGrid);
  }

  return `${flashCount}`;
};

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day11: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day11;
