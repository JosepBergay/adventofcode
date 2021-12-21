export type Point = [number, number];

export const isTopEdge = ([x, y]: Point) => y === 0;

export const isBotEdge = ([x, y]: Point, mapHeight: number) =>
  y === mapHeight - 1;

export const isLeftEdge = ([x, y]: Point) => x === 0;

export const isRightEdge = ([x, y]: Point, mapWidth: number) =>
  x === mapWidth - 1;

export const logGrid = <T>(
  grid: T[][],
  prefix: string,
  valueMapper?: (v: T) => any
) =>
  console.log(
    prefix,
    grid.map((row) => (valueMapper ? row.map(valueMapper).join() : row.join()))
  );

export const copyGrid = <T>(grid: T[][], mapper?: (v: T) => T) =>
  grid.map((row) => (mapper ? row.map(mapper) : row.slice()));

export const buildGrid = <T>(points: Point[], value: T, emptyValue: T) => {
  const max_x = points.reduce((max, [x, _]) => (x > max ? x : max), 0);
  const max_y = points.reduce((max, [_, y]) => (y > max ? y : max), 0);

  const grid: T[][] = [];
  for (let y = 0; y < max_y + 1; y++) {
    const row: T[] = [];
    for (let x = 0; x < max_x + 1; x++) {
      row.push(emptyValue);
    }
    grid.push(row);
  }

  for (const [x, y] of points) {
    grid[y][x] = value;
  }
  return grid;
};
