import type { AOCDay } from "./types";
import { day1 } from "./day1/index.js";

const days: AOCDay[] = [day1];

const main = async () => {
  const responses = await Promise.all(days.map((d) => d()));

  for (const res of responses) {
    console.log(`Day #${res.level} part1: ${res.part1}, part2: ${res.part2}`);
  }
};

main();
