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
