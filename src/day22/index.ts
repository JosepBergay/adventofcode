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

const executePart2 = (input: ParsedInput) => {
//   const grid = create3DGrid<boolean>();

//   const set = new Set();

//   for (const [a, rangeX, rangeY, rangeZ] of input) {
//     for (let x = rangeX[0]; x <= rangeX[1]; x++) {
//       for (let y = rangeY[0]; y <= rangeY[1]; y++) {
//         for (let z = rangeZ[0]; z <= rangeZ[1]; z++) {
//           if (a === "on") {
//             set.add(`${x},${y},${z}`);
//           } else {
//             set.delete(`${x},${y},${z}`);
//           }
//         //   console.log(set.size);
//           //   grid.set(x, y, z, a === "on");
//         }
//       }
//     }
//   }
//   return set.size;
//   return grid.count((isOn) => isOn);
};

const day22: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day22;
