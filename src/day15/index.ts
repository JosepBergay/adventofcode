import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import {
  getAdjacents as getAdjacentsPointValue,
  isSamePoint,
  Point,
} from "../gridUtils.js";

const level = 15;

type ParsedInput = number[][];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l !== "")
    .map((l) => l.split("").map((i) => parseInt(i)));

// Dijkstra’s Algorithm implementation https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
const findShortestPath = (from: Point, to: Point, graph: number[][]) => {
  // Store in a stack the vertices not yet it in shortest path or not finalized.
  const stack = [[from, 0]] as [Point, number][];

  // Keep track of all visited nodes, going backwards would make no sense.
  const visited = [[from, 0]] as [Point, number][];

  const getAdjacents = (p: Point) => getAdjacentsPointValue(p, graph, false);

  while (stack.length) {
    // Extract vertex with minimum distance.
    const vertex = stack.pop()!;

    const current = visited.find((v) => isSamePoint(v[0], vertex[0]))!;

    for (const [adjacentPoint, value] of getAdjacents(vertex[0])) {
      const previous = visited.find(([p, v]) => isSamePoint(adjacentPoint, p));

      const newValue = current[1] + value;

      // Not visited yet or found a lower value
      if (!previous || previous[1] > newValue) {
        visited.push([adjacentPoint, newValue]);
        stack.push([adjacentPoint, newValue]);
        stack.sort((a, b) => b[1] - a[1]); // Keeping stack sorted so we can pop!
      }
    }
  }
  const destination = visited.find((v) => isSamePoint(to, v[0]));
  return destination && destination[1];
};

const executePart1 = (input: ParsedInput) => {
  const start: Point = [0, 0];
  const end: Point = [input[0].length - 1, input.length - 1];

  const lowestRisk = findShortestPath(start, end, input);
  return lowestRisk;
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day15: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day15;
