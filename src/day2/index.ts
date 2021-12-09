import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 2;

type ParsedInput = {
  dir: "forward" | "down" | "up";
  units: number;
}[];

const parser = (input: string): ParsedInput =>
  input.split("\n").map((i) => {
    const command = i.split(" ");
    return {
      dir: command[0] as ParsedInput[0]["dir"],
      units: parseInt(command[1]),
    };
  });

const executePart1 = (input: ParsedInput): string => {
  const result = input.reduce(
    (p, c) => {
      switch (c.dir) {
        case "forward":
          return { ...p, horizontal: p.horizontal + c.units };
        case "down":
          return { ...p, depth: p.depth + c.units };
        case "up":
          return { ...p, depth: p.depth - c.units };

        default:
          return p;
      }
    },
    { horizontal: 0, depth: 0 }
  );

  return `${result.horizontal * result.depth}`;
};

const executePart2 = (input: ParsedInput): string => {
  const result = input.reduce(
    (p, c) => {
      if (c.dir === "down") {
        return { ...p, aim: p.aim + c.units };
      } else if (c.dir === "up") {
        return { ...p, aim: p.aim - c.units };
      } else if (c.dir === "forward") {
        return {
          ...p,
          horizontal: p.horizontal + c.units,
          depth: p.aim * c.units + p.depth,
        };
      } else {
        return p;
      }
    },
    { horizontal: 0, depth: 0, aim: 0 }
  );

  return `${result.horizontal * result.depth}`;
};

const day2: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day2;
