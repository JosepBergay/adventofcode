import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import {
  copyGrid,
  getAdjacents as getAdjacentsPointValue,
  Grid,
  isSamePoint,
  Point,
  pointToStr,
} from "../gridUtils.js";

const level = 15;

type ParsedInput = Grid<number>;

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l !== "")
    .map((l) => l.split("").map((i) => parseInt(i)));

// Dijkstra’s Algorithm implementation https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
const findShortestPath = (from: Point, to: Point, graph: Grid<number>) => {
  // Store in a stack the vertices not yet it in shortest path or not finalized.
  const stack = [[from, 0]] as [Point, number][];

  // Keep track of all visited nodes, going backwards would make no sense.
  // const visited = [[from, 0]] as [Point, number][];
  // Optmization: Using Map for (much) better perf.
  const visited = new Map<string, number>([[pointToStr(from), 0]]);

  const getAdjacents = (p: Point) => getAdjacentsPointValue(p, graph, false);

  // Optmization: Condition to break early.
  const reachedEnd = (p: Point) => isSamePoint(p, to);

  while (stack.length) {
    // Extract vertex with minimum distance.
    const vertex = stack.pop()!;

    const current = visited.get(pointToStr(vertex[0]))!;

    if (reachedEnd(vertex[0])) break;

    for (const [adjacentPoint, value] of getAdjacents(vertex[0])) {
      const previous = visited.get(pointToStr(adjacentPoint));

      const newValue = current + value;

      // Not visited yet or found a lower value
      if (!previous || previous > newValue) {
        visited.set(pointToStr(adjacentPoint), newValue);
        stack.push([adjacentPoint, newValue]);
        stack.sort((a, b) => b[1] - a[1]); // Keeping stack sorted so we can pop!
      }
    }
  }
  const destination = visited.get(pointToStr(to));
  return destination;
};

const executePart1 = (input: ParsedInput) => {
  const start: Point = [0, 0];
  const end: Point = [input[0].length - 1, input.length - 1];

  const lowestRisk = findShortestPath(start, end, input);
  return lowestRisk;
};

const executePart2 = (input: ParsedInput) => {
  // Extend map with 8 different diagonals
  const extensions: Grid<number>[] = [copyGrid(input)];
  for (let i = 0; i < 8; i++) {
    const lastExtended = extensions[extensions.length - 1];
    const newExtension = copyGrid(lastExtended, (v) => (v === 9 ? 1 : v + 1));
    extensions.push(newExtension);
  }

  const chooseExtension = ([x, y]: Point) => {
    const xTile = Math.floor(x / input[0].length);
    const yTile = Math.floor(y / input.length);

    // The sum of xTyle + yTile gives us the index of the extension.
    //   [[0, 0]], // Diagonal 0
    //   [[1, 0], [0, 1]], // Diagonal 1
    //   [[2, 0], [1, 1], [0, 2]], // Diagonal 2
    //   [[3, 0], [2, 1], [1, 2], [0,3]], // Diagonal 3
    //   ...
    //   [[4, 4]] // Diagonal 8

    return extensions[xTile + yTile];
  };

  const extendedGrid: ParsedInput = [];
  for (let y = 0; y < input.length * 5; y++) {
    const row: number[] = [];
    for (let x = 0; x < input[0].length * 5; x++) {
      const extension = chooseExtension([x, y]);
      const extY = y % input.length;
      const extX = x % input[0].length;
      row.push(extension[extY][extX]);
    }
    extendedGrid.push(row);
  }

  const start: Point = [0, 0];
  const end: Point = [extendedGrid[0].length - 1, extendedGrid.length - 1];
  const lowestRisk = findShortestPath(start, end, extendedGrid);
  return lowestRisk;
};

const day15: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day15;
