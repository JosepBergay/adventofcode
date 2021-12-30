import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 19;

type V3 = [number, number, number];

type ParsedInput = V3[][];

const parser = (input: string): ParsedInput => {
  const scanners: ParsedInput = [];
  input
    .split("\n")
    .filter((l) => l !== "")
    .forEach((l) =>
      l.startsWith("---")
        ? scanners.push([])
        : scanners[scanners.length - 1].push(
            l.split(",").map((n) => parseInt(n)) as V3
          )
    );
  return scanners;
};

// Turn each scanner beacon into a 3D vector pointing to the rest of beacons.
const getDistancesBetweenBeacons = (scanner: V3[]) => {
  const vectors = [];
  for (const [i, start] of scanner.entries()) {
    const beaconVectors = { beacon: start, relatives: [] as V3[] };
    for (const [j, end] of scanner.entries()) {
      if (i === j) continue; // skip self
      beaconVectors.relatives.push([
        Math.abs(end[0] - start[0]),
        Math.abs(end[1] - start[1]),
        Math.abs(end[2] - start[2]),
      ]);
    }
    beaconVectors.beacon.length && vectors.push(beaconVectors);
  }
  return vectors;
};

// We don't care about orientation in vector distances, using absolute value.
const compareVector = ([x1, y1, z1]: V3, [x2, y2, z2]: V3) =>
  Math.abs(x1) === Math.abs(x2) &&
  Math.abs(y1) === Math.abs(y2) &&
  Math.abs(z1) === Math.abs(z2);

/**
 * x, y, z,
 * x, z, y,
 * y, x, z,
 * y, z, x,
 * z, x, y,
 * z, y, x,
 */
const ROTATIONS: ((v: V3) => V3)[] = [
  ([x, y, z]) => [x, y, z],
  ([x, y, z]) => [x, z, y], // X
  ([x, y, z]) => [y, x, z], // Z
  ([x, y, z]) => [y, z, x],
  ([x, y, z]) => [z, x, y],
  ([x, y, z]) => [z, y, x], // Y
];

// Compare beacons position for each pair of scanners. Each axis can be rotated 4 times.
const findOverlapping = (scannerA: V3[], scannerB: V3[]) => {
  const beaconsA = getDistancesBetweenBeacons(scannerA);
  const beaconsB = getDistancesBetweenBeacons(scannerB);
  for (const [i, rotate] of ROTATIONS.entries()) {
    const matches: [V3, V3][] = [];
    for (const beaconA of beaconsA) {
      // If this beacon has 11 beacon distances equal to another scannerB beacon, they overlap.
      const candidates = beaconA.relatives.flatMap((dA) =>
        beaconsB.flatMap((bB) =>
          bB.relatives
            .map(rotate)
            .filter((dB) => compareVector(dA, dB))
            .map((_) => bB.beacon)
        )
      );

      const mostCommon = candidates.filter(
        (c) => candidates.filter((c2) => compareVector(c, c2)).length >= 11
      );

      if (mostCommon.length) {
        matches.push([beaconA.beacon, mostCommon[0]]);
        if (matches.length === 12) return [matches, i] as const;
      }
    }
  }
  return [null, -1] as const;
};

const sub = ([x1, y1, z1]: V3, [x2, y2, z2]: V3): V3 => [
  x1 - x2,
  y1 - y2,
  z1 - z2,
];

const add = ([x1, y1, z1]: V3, [x2, y2, z2]: V3): V3 => [
  x1 + x2,
  y1 + y2,
  z1 + z2,
];

const getFacing = (distanceA: V3, distanceB: V3) => {
  const face = [
    distanceA[0] === distanceB[0] ? 1 : -1,
    distanceA[1] === distanceB[1] ? 1 : -1,
    distanceA[2] === distanceB[2] ? 1 : -1,
  ];
  return (v: V3): V3 => [v[0] * face[0], v[1] * face[1], v[2] * face[2]];
};

const findOffsetAndTransform = (matches: [V3, V3][], rotation: number) => {
  const rotate = ROTATIONS[rotation];
  let pairs = matches.slice(0, 2).map((vs) => [vs[0], rotate(vs[1])] as const);

  // Compare distances between same beacons from different references.
  const [[vA, vB], [v2A, v2B]] = pairs;
  const distanceA = sub(v2A, vA);
  const distanceB = sub(v2B, vB);

  const face = getFacing(distanceA, distanceB);

  const offset = sub(vA, face(vB));

  const transform = (v: V3) => add(face(rotate(v)), offset);

  return [offset, transform] as const;
};

type Scanner = {
  name: number;
  transformedBeacons: V3[];
  overlapsWith: number[];
  position: V3;
  transformTo0: (v: V3) => V3;
};

const logScanner = (scanner: Scanner) => {
  console.log(`Scanner ${scanner.name}`);
  console.log(`Overlaps with: ${scanner.overlapsWith}`);
  console.log(`Position: ${scanner.position}`);
};

const analyzeScanners = (scanners: ParsedInput) => {
  // Use first scanner as reference.
  const store: Record<number, Scanner> = {
    0: {
      name: 0,
      overlapsWith: [],
      position: [0, 0, 0],
      transformedBeacons: scanners[0],
      transformTo0: (v) => v,
    },
  };

  const analysisQueue = [0];
  while (analysisQueue.length) {
    const current = analysisQueue.pop()!;

    const beaconsA = store[current].transformedBeacons;

    for (const [i, beaconsB] of scanners.entries()) {
      if (i === current || store[i]) continue; // skip self and already analyzed.
      console.log(current, i);
      const [matches, rotationIdx] = findOverlapping(beaconsA, beaconsB);

      if (matches) {
        const [offset, transform] = findOffsetAndTransform(
          matches,
          rotationIdx
        );

        // Update state
        store[i] = {
          name: i,
          overlapsWith: [current],
          position: offset,
          transformTo0: transform,
          transformedBeacons: beaconsB.map(transform),
        };
        store[current].overlapsWith.push(i);
        analysisQueue.push(i);
      }
    }
  }
  return store;
};

const findUniqueBeacons = (scanners: ParsedInput) => {
  // Using Set, it won't allow repeated items.
  const unique = new Set<string>();
  const scannerStore = analyzeScanners(scanners);
  for (const key in scannerStore) {
    const scanner = scannerStore[key];
    scanner.transformedBeacons.forEach((b) => unique.add(b.toString()));
    logScanner(scanner);
  }
  return unique;
};

const executePart1 = (input: ParsedInput) => {
  const unique = findUniqueBeacons(input);

  return unique.size;
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day19: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  console.log("Day19 part1 is going to take a while");
  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day19;
