const parseInput = (input: string) => input.split("\n").filter((l) => l != "");

export const analyzeInput = (input: string) => {
  const parsed = parseInput(input);

  const grouped = new Map<number, Set<string>>();
  let i = 0;
  for (const instruction of parsed) {
    if (instruction.startsWith("inp")) i = 0;
    const key = i;
    const set = grouped.get(key) ?? new Set();
    set.add(instruction);
    grouped.set(key, set);
    i++;
  }

  const diff = [];
  for (const [i, set] of grouped.entries()) {
    if (set.size > 1) {
      for (const instr of set.values()) {
        diff.push(`index: ${i}; op: ${instr.slice(0, 5)}`);
        break;
      }
    }
  }

  // Only instructions that change between sections are:
  // 'index: 4; op: div z',
  // 'index: 5; op: add x' and
  // 'index: 15; op: add y'
};

/**
 * inp w
 * mul x 0
 * add x z
 * mod x 26
 * div z 1   <- this changes
 * add x 12  <- this changes
 * eql x w
 * eql x 0
 * mul y 0
 * add y 25
 * mul y x
 * add y 1
 * mul z y
 * mul y 0
 * add y w
 * add y 4   <- this changes
 * mul y x
 * add z y
 */
/**
 * W stays at given value. X and Y always start as 0. We need to carry Z value through sections
 */
export const executeDirect = (
  inp: number,
  z: number,
  zDiv: number,
  xAdd: number,
  yAdd: number
) => {
  let w = inp;
  let x = z;
  x %= 26;
  z = Math.floor(z / zDiv);
  x += xAdd;
  /**
   * This:
   * x = x === w ? 1 : 0;
   * x = x === 0 ? 1 : 0;
   * is reduced to:
   */
  x = x === w ? 0 : 1;
  let y = 25;
  y *= x;
  y += 1;
  z *= y;
  y = w + yAdd;
  y *= x;
  z += y;
  return z;
};

type Section = [zDiv: number, xAdd: number, yAdd: number];

const getSections = (parsed: string[]) => {
  const sections: Section[] = [];
  // Each section has 18 instructions
  for (let i = 0; i < parsed.length; i = i + 18) {
    const zDiv = Number(parsed[i + 4].slice(6));
    const xAdd = Number(parsed[i + 5].slice(6));
    const yAdd = Number(parsed[i + 15].slice(6));
    sections.push([zDiv, xAdd, yAdd]);
  }
  return sections;
};

let failedCache = new Map<number, Set<number>>();

// ~2secs for 100M iterations.
const executeLargestRec = (
  iteration: number,
  lastZ: number,
  sections: Section[]
): number | undefined => {
  if (failedCache.has(iteration)) {
    if (failedCache.get(iteration)!.has(lastZ)) return;
  }

  for (let i = 9; i > 0; i--) {
    const newZ = executeDirect(i, lastZ, ...sections[iteration]);
    if (iteration === 13) {
      if (newZ === 0) {
        return i;
      }
    } else {
      const res = executeLargestRec(iteration + 1, newZ, sections);
      if (res) {
        return i * Math.pow(10, sections.length - 1 - iteration) + res;
      }
    }
  }

  // If got here, number wasn't found. Caching
  const fails = failedCache.get(iteration) ?? new Set<number>();
  failedCache.set(iteration, fails.add(lastZ));
};

export const executeSmart = (input: string) => {
  const parsed = parseInput(input);
  const sections = getSections(parsed);

  return executeLargestRec(0, 0, sections);
};

const executeSmallestRec = (
  iteration: number,
  lastZ: number,
  sections: Section[]
): number | undefined => {
  if (failedCache.has(iteration)) {
    if (failedCache.get(iteration)!.has(lastZ)) return;
  }

  for (let i = 1; i < 10; i++) {
    const newZ = executeDirect(i, lastZ, ...sections[iteration]);
    if (iteration === 13) {
      if (newZ === 0) {
        return i;
      }
    } else {
      const res = executeSmallestRec(iteration + 1, newZ, sections);
      if (res) {
        return i * Math.pow(10, sections.length - 1 - iteration) + res;
      }
    }
  }

  // If got here, number wasn't found. Caching
  const fails = failedCache.get(iteration) ?? new Set<number>();
  failedCache.set(iteration, fails.add(lastZ));
};

export const executeSmart2 = (input: string) => {
  const parsed = parseInput(input);
  const sections = getSections(parsed);

  return executeSmallestRec(0, 0, sections);
};
