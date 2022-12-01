import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 6;

type ParsedInput = number[];

const parser = (input: string): ParsedInput =>
  input.split(",").map((s) => parseInt(s));

const processFish = (fish: number) => {
  let mustAddFish = false;
  if (fish === 0) {
    fish = 6;
    mustAddFish = true;
  } else {
    fish--;
  }
  return [fish, mustAddFish] as const;
};

const executePart1 = (input: ParsedInput): string => {
  const fishes = input.slice();

  for (let day = 0; day < 80; day++) {
    const fishesBorn = [];
    for (let [index, fish] of fishes.entries()) {
      const [newFishTimer, mustAdd] = processFish(fish);
      fishes[index] = newFishTimer;
      mustAdd && fishesBorn.push(true);
    }
    fishesBorn.forEach((_) => fishes.push(8));
  }

  return `${fishes.length}`;
};

const executePart2 = (input: ParsedInput): string => {
  const fishes = input.slice();

  // Key is fish timer and value is number of fishes
  let mapOfFishes = new Map<number, number>();

  // Init map
  for (const [index, _] of new Array(9).entries()) {
    mapOfFishes.set(index, 0);
  }

  // Fill map with initial state
  for (const fish of fishes) {
    mapOfFishes.set(fish, mapOfFishes.get(fish)! + 1);
  }

  for (let day = 0; day < 256; day++) {
    const newMap = new Map<number, number>();
    mapOfFishes.forEach((v, k) => {
      if (k === 0) {
        newMap.set(6, v + (newMap.get(6) ?? 0));
        newMap.set(8, mapOfFishes.get(0) ?? 0);
      }
      newMap.set(k, (newMap.get(k) ?? 0) + (mapOfFishes.get(k + 1) ?? 0));
    });
    mapOfFishes = newMap;
  }

  let numOfFishes = 0;
  mapOfFishes.forEach((v) => (numOfFishes = numOfFishes + v));

  return `${numOfFishes}`;
};

const day6: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day6;
