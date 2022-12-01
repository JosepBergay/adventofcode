import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 22;

type Range = [number, number];

type Action = "on" | "off";

type Procedure = [Action, Range, Range, Range];

type ParsedInput = Procedure[];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l !== "")
    .map((line) => {
      const [first, rest] = line.split(" ");
      const action = first as Action;
      const [x1, x2] = rest
        .split("x=")[1]
        .split("..")
        .map((s, i) => (i === 0 ? s : s.split(",")[0]));
      const [y1, y2] = rest
        .split("y=")[1]
        .split("..")
        .map((s, i) => (i === 0 ? s : s.split(",")[0]));
      const [z1, z2] = rest
        .split("z=")[1]
        .split("..")
        .map((s, i) => (i === 0 ? s : s.split(",")[0]));
      return [
        action,
        [Number(x1), Number(x2)],
        [Number(y1), Number(y2)],
        [Number(z1), Number(z2)],
      ];
    });

const create3DGrid = <T>() => {
  const map = new Map<string, T>();
  const getKey = (x: number, y: number, z: number) => `${x},${y},${z}`;
  const getPoint = (k: string) =>
    k.split(",").map(Number) as [number, number, number];

  return {
    set: (x: number, y: number, z: number, v: T) => {
      map.set(getKey(x, y, z), v);
    },
    get: (x: number, y: number, z: number) => map.get(getKey(x, y, z)),
    map,
    count: (filterFn?: (t: T) => boolean) => {
      if (!filterFn) {
        return map.size;
      }
      let count = 0;
      for (const value of map.values()) {
        if (filterFn(value)) {
          ++count;
        }
      }
      return count;
    },
  };
};

const executePart1 = (input: ParsedInput) => {
  const isInitRange = (r: Range) => r[0] >= -50 && r[1] <= 50;
  const filtered = input.filter(
    (p) => isInitRange(p[1]) && isInitRange(p[2]) && isInitRange(p[3])
  );

  const grid = create3DGrid<boolean>();

  for (const [a, rangeX, rangeY, rangeZ] of filtered) {
    for (let x = rangeX[0]; x <= rangeX[1]; x++) {
      for (let y = rangeY[0]; y <= rangeY[1]; y++) {
        for (let z = rangeZ[0]; z <= rangeZ[1]; z++) {
          grid.set(x, y, z, a === "on");
        }
      }
    }
  }

  return grid.count((isOn) => isOn);
};

type Cube = {
  x1: number;
  x2: number;
  y1: number;
  y2: number;
  z1: number;
  z2: number;
};

type Cuboid = {
  cube: Cube;
  isOn: boolean;
  holes: Cuboid[];
};

type Proc = Omit<Cuboid, "holes">;

const intersect = (a: Cube, b: Cube) => {
  const x1 = Math.max(a.x1, b.x1);
  const x2 = Math.min(a.x2, b.x2);
  const y1 = Math.max(a.y1, b.y1);
  const y2 = Math.min(a.y2, b.y2);
  const z1 = Math.max(a.z1, b.z1);
  const z2 = Math.min(a.z2, b.z2);

  if (x2 >= x1 && y2 >= y1 && z2 >= z1) {
    return { x1, x2, y1, y2, z1, z2 } as Cube;
  }
  return false;
};

const getHole = (a: Cuboid, b: Cube): Cuboid | null => {
  const intersection = intersect(a.cube, b);
  if (!intersection) {
    return null;
  }
  // If intersecting, create a hole of the opposite state to not count it twice.
  return {
    cube: intersection,
    isOn: !a.isOn,
    holes: a.holes.map((h) => getHole(h, b)).filter((h) => h) as Cuboid[],
  };
};

const countCubes = ({ cube, holes }: Cuboid) => {
  let total =
    (cube.x2 - cube.x1 + 1) * (cube.y2 - cube.y1 + 1) * (cube.z2 - cube.z1 + 1);
  total -= holes.reduce((p, c) => p + countCubes(c), 0);
  return total;
};

const executePart2 = (input: ParsedInput) => {
  const procs = input.map(
    ([a, [x1, x2], [y1, y2], [z1, z2]]) =>
      ({ cube: { x1, x2, y1, y2, z1, z2 }, isOn: a === "on" } as Proc)
  );

  const parsed: Cuboid[] = [];
  for (const proc of procs) {
    for (const p of parsed) {
      const hole = getHole(p, proc.cube);
      if (hole) {
        p.holes.push(hole);
      }
    }

    if (proc.isOn) {
      parsed.push({ ...proc, holes: [] });
    }
  }
  return parsed.reduce((p, c) => p + countCubes(c), 0);
};

const day22: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day22;
