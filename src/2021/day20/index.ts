import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";
import { getAdjacentsInfinity, Grid } from "../gridUtils.js";

const level = 20;

type Pixel = "#" | ".";

type ParsedInput = {
  algorithm: string;
  image: Grid<Pixel>;
};

const parser = (input: string): ParsedInput => {
  const [first, ...rest] = input.split("\n").filter((l) => l !== "");
  const algorithm = first;
  const image = rest.map((line) => line.split("") as Pixel[]);
  return { algorithm, image };
};

const pixelsToInt = (pixels: Pixel[]) =>
  parseInt(pixels.map((p) => (p === "#" ? 1 : 0)).join(""), 2);

const parseImage = (grid: Grid<Pixel>, algo: string, infinityValue: Pixel) => {
  const parsedImage: Grid<Pixel> = [];
  // Extending grid with -1, +1
  for (let y = -1; y < grid.length + 1; y++) {
    const parsedRow: Pixel[] = [];
    for (let x = -1; x < grid[0].length + 1; x++) {
      const adjacents = getAdjacentsInfinity([x, y], grid, infinityValue);
      const adjacentsWithPoint = [
        ...adjacents.slice(0, 4),
        (grid[y] && grid[y][x]) ?? infinityValue,
        ...adjacents.slice(-4),
      ];

      const int = pixelsToInt(adjacentsWithPoint);
      const realPixel = algo[int] as Pixel;

      parsedRow.push(realPixel);
    }
    parsedImage.push(parsedRow);
  }
  return parsedImage;
};

const countLitPixels = (grid: Grid<Pixel>) => {
  let litPixels = 0;
  for (const row of grid) {
    for (const pixel of row) {
      litPixels += pixel === "#" ? 1 : 0;
    }
  }
  return litPixels;
};

const parseImageMultipleTimes = (
  input: ParsedInput,
  infinityValue: Pixel,
  steps: number
) => {
  let parsedImage: Grid<Pixel> = input.image;
  let infValue = infinityValue;
  for (let i = 0; i < steps; i++) {
    parsedImage = parseImage(parsedImage, input.algorithm, infValue);
    // const int = pixelsToInt(new Array(9).fill(infValue));
    // Instead of computing int each time, we know it's either 0 or 511.
    infValue = input.algorithm[infValue === "." ? 0 : 511] as Pixel;
  }
  return parsedImage;
};

const executePart1 = (input: ParsedInput) => {
  // First time infinity value is "."
  const infinityValue = ".";
  const parsed = parseImage(input.image, input.algorithm, ".");

  // Second time, other way around, all 3x3 -> 111111111 -> 511
  const int = pixelsToInt(new Array(9).fill(infinityValue));
  const secondInfValue = input.algorithm[int] as Pixel;

  const parsedTwice = parseImage(parsed, input.algorithm, secondInfValue);

  const count = countLitPixels(parsedTwice);

  return count;
};

const executePart2 = (input: ParsedInput) => {
  const parsed = parseImageMultipleTimes(input, ".", 50);

  const count = countLitPixels(parsed);

  return count;
};

const day20: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day20;
