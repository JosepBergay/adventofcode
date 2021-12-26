import type { AOCDay } from "../types";
import type { Point } from "../gridUtils";
import { fetchInput } from "../helpers.js";

const level = 17;

type ParsedInput = { minX: number; maxX: number; minY: number; maxY: number };

const parser = (input: string): ParsedInput => {
  const [_, ...rest] = input.split("=");
  const str = rest
    .map((r) => (r.endsWith(", y") ? r.slice(0, -3) : r))
    .map((r) => r.split(".."));
  const ints = str.flatMap((n) => [parseInt(n[0]), parseInt(n[1])]);
  return {
    minX: ints[0],
    maxX: ints[1],
    minY: ints[2],
    maxY: ints[3],
  };
};

type PosVel = [Point, Point];

const move = (pos: Point, vel: Point) => {
  const nextPos = [pos[0] + vel[0], pos[1] + vel[1]];
  const nextVel = [
    vel[0] ? (vel[0] > 0 ? vel[0] - 1 : vel[0] + 1) : vel[0],
    vel[1] - 1,
  ];
  return [nextPos, nextVel] as PosVel;
};

const isInTarget = ([x, y]: Point, area: ParsedInput) =>
  x >= area.minX && x <= area.maxX && y >= area.minY && y <= area.maxY;

const wontReachTarget = ([x, y]: Point, [vx, vy]: Point, area: ParsedInput) =>
  (vx === 0 && (x < area.minX || x > area.maxX)) || y < area.minY;

const computeTrajectory = (start: Point, v0: Point, area: ParsedInput) => {
  let currentPos = start;
  let currentVel: Point = v0;
  const trajectory: Point[] = [];
  while (true) {
    [currentPos, currentVel] = move(currentPos, currentVel);
    trajectory.push(currentPos);
    if (isInTarget(currentPos, area)) {
      return trajectory;
    }
    if (wontReachTarget(currentPos, currentVel, area)) {
      return undefined;
    }
  }
};

const executePart1 = (input: ParsedInput) => {
  const start: Point = [0, 0];

  const validTrajectories: { trajectory: Point[]; v0: Point }[] = [];
  // To do it in style vy should be > 0 and vx should be less than target minX.
  // Heuristic: vy should be less than target minY (a max value for vy is needed).
  for (let vx = 0; vx < input.minX; vx++) {
    for (let vy = 0; vy < Math.abs(input.minY); vy++) {
      const trajectory = computeTrajectory(start, [vx, vy], input);
      if (trajectory) {
        validTrajectories.push({ trajectory, v0: [vx, vy] });
      }
    }
  }

  // Compute highest trajectory
  const maxY = validTrajectories
    .map((t) => ({
      maxY: Math.max(...t.trajectory.map((t) => t[1])),
      v0: t.v0,
    }))
    .reduce((p, c) => (p.maxY > c.maxY ? p : c));

  return maxY.maxY;
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day17: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day17;
