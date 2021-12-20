import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import { Point } from "../gridUtils";

const level = 13;

type FoldInstructions = { dir: "x" | "y"; value: number };

type ParsedInput = [Point[], FoldInstructions[]];

const parser = (input: string): ParsedInput => {
  const splitted = input.split("\n\n");
  const points = splitted[0]
    .split("\n")
    .map((l) => l.split(",").map((n) => parseInt(n)) as Point);
  const common = "fold along ";
  const instructions = splitted[1]
    .split("\n")
    .filter((l) => l !== "")
    .map((l) => {
      const instr = l.substring(common.length).split("=");
      return { dir: instr[0], value: parseInt(instr[1]) } as FoldInstructions;
    });
  return [points, instructions];
};

const groupBy = (points: Point[], condition: (point: Point) => boolean) =>
  points.reduce(
    ([truthy, falsy], p) =>
      (condition(p) ? [[...truthy, p], falsy] : [truthy, [...falsy, p]]) as [
        Point[],
        Point[]
      ],
    [[], []] as [Point[], Point[]]
  );

const removeRepeated = (points: Point[]) =>
  points.reduce(
    (acc, c) =>
      acc.find((p) => p[0] === c[0] && p[1] === c[1]) ? acc : [...acc, c],
    [] as Point[]
  );

const foldLeft = (points: Point[], y_axis: number) => {
  const [leftSide, rightSide] = groupBy(points, (p) => p[0] < y_axis);

  for (const [x, y] of rightSide) {
    leftSide.push([x - (x - y_axis) * 2, y]);
  }

  return removeRepeated(leftSide);
};

const foldUp = (points: Point[], x_axis: number) => {
  const [top, bot] = groupBy(points, (p) => p[1] < x_axis);

  for (const [x, y] of bot) {
    top.push([x, y - (y - x_axis) * 2]);
  }

  return removeRepeated(top);
};

const executePart1 = (input: ParsedInput) => {
  let [points, instructions] = input;

  for (const instr of [instructions[0]]) {
    if (instr.dir === "x") {
      points = foldLeft(points, instr.value);
    } else {
      points = foldUp(points, instr.value);
    }
  }

  return points.length;
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day13: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day13;
