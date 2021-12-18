export type Point = [number, number];

export const isTopEdge = ([x, y]: Point) => y === 0;

export const isBotEdge = ([x, y]: Point, mapHeight: number) =>
  y === mapHeight - 1;

export const isLeftEdge = ([x, y]: Point) => x === 0;

export const isRightEdge = ([x, y]: Point, mapWidth: number) =>
  x === mapWidth - 1;
