import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import { copyGrid, getAdjacents } from "../gridUtils.js";

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

const flashAndIncrease = (dumboGrid: Dumbo[][]) => {
  let flashCount = 0;
  for (const [y, dumboRow] of dumboGrid.entries()) {
    for (const [x, dumbo] of dumboRow.entries()) {
      if (dumbo.energy > 9 && !dumbo.hasFlashed) {
        dumbo.hasFlashed = true;
        flashCount++;
        const adjacents = getAdjacents([x, y], dumboGrid, true);
        increaseEnergyByOne([adjacents.map((a) => a[1])]);
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

const executeStep = (dumboGrid: ParsedInput) => {
  // First increase energy by 1.
  increaseEnergyByOne(dumboGrid);

  // Then, flash and increase energy
  const flashCount = flashAndIncrease(dumboGrid);

  // Finally, reset those that flashed
  resetFlashed(dumboGrid);

  return flashCount;
};

const executePart1 = (input: ParsedInput): string => {
  const dumboGrid = copyGrid(input, (d) => ({ ...d }));

  let flashCount = 0;

  for (let step = 0; step < 100; step++) {
    flashCount += executeStep(dumboGrid);
  }

  return `${flashCount}`;
};

const executePart2 = (input: ParsedInput): string => {
  const dumboGrid = copyGrid(input, (d) => ({ ...d }));

  let step = 0;
  do {
    executeStep(dumboGrid);
    step++;
  } while (!dumboGrid.every((row) => row.every((d) => !d.energy)));

  return `${step}`;
};

const day11: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day11;
