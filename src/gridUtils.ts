export type Point = [number, number];

export type Grid<T> = T[][];

export const isSamePoint = (p1: Point, p2: Point) =>
  p1[0] === p2[0] && p1[1] === p2[1];

export const isTopEdge = ([x, y]: Point) => y === 0;

export const isBotEdge = ([x, y]: Point, mapHeight: number) =>
  y === mapHeight - 1;

export const isLeftEdge = ([x, y]: Point) => x === 0;

export const isRightEdge = ([x, y]: Point, mapWidth: number) =>
  x === mapWidth - 1;

export const logGrid = <T>(
  grid: Grid<T>,
  prefix?: string,
  valueMapper?: (v: T) => any
) =>
  console.log(
    prefix || "",
    grid.map((row) => (valueMapper ? row.map(valueMapper).join() : row.join()))
  );

export const copyGrid = <T>(grid: Grid<T>, mapper?: (v: T) => T) =>
  grid.map((row) => (mapper ? row.map(mapper) : row.slice()));

export const buildGrid = <T>(points: Point[], value: T, emptyValue: T) => {
  const max_x = points.reduce((max, [x, _]) => (x > max ? x : max), 0);
  const max_y = points.reduce((max, [_, y]) => (y > max ? y : max), 0);

  const grid: Grid<T> = [];
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

export const getAdjacentsInfinity = <T>(
  [x, y]: Point,
  grid: Grid<T>,
  defaultValue: T
) => {
  const adjacents: Point[] = [];
  // Top Row
  adjacents.push([x - 1, y - 1], [x, y - 1], [x + 1, y - 1]);
  // Mid Row
  adjacents.push([x - 1, y], [x + 1, y]);
  // Bot Row
  adjacents.push([x - 1, y + 1], [x, y + 1], [x + 1, y + 1]);

  return adjacents.map(([x, y]) => (grid[y] && grid[y][x]) ?? defaultValue);
};

export const getAdjacents = <T>(
  [x, y]: Point,
  grid: Grid<T>,
  getDiagonals: boolean
) => {
  const notLeftEdge = !isLeftEdge([x, y]);
  const notRightEdge = !isRightEdge([x, y], grid[y].length);
  const adjacents: Point[] = [];

  if (!isTopEdge([x, y])) {
    adjacents.push([x, y - 1]);
    if (notLeftEdge && getDiagonals) adjacents.push([x - 1, y - 1]);
    if (notRightEdge && getDiagonals) adjacents.push([x + 1, y - 1]);
  }

  if (notLeftEdge) adjacents.push([x - 1, y]);
  if (notRightEdge) adjacents.push([x + 1, y]);

  if (!isBotEdge([x, y], grid.length)) {
    adjacents.push([x, y + 1]);
    if (notLeftEdge && getDiagonals) adjacents.push([x - 1, y + 1]);
    if (notRightEdge && getDiagonals) adjacents.push([x + 1, y + 1]);
  }

  return adjacents.map(([x, y]) => [[x, y], grid[y][x]] as [Point, T]);
};

export const gridToPoints = <T>(grid: Grid<T>) =>
  grid.flatMap((row, y) => row.map((_, x) => [x, y] as Point));

export const pointToStr = (p: Point) => `${p[0]},${p[1]}`;
