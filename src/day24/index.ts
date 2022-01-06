import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 24;

type Var = "w" | "x" | "y" | "z";

const isVar = (varOrNum: string): varOrNum is Var =>
  varOrNum === "w" || varOrNum === "x" || varOrNum === "y" || varOrNum === "z";

type IntVars = [w: number, x: number, y: number, z: number];

const createALU = (...intVars: IntVars) => {
  let w = intVars[0];
  let x = intVars[1];
  let y = intVars[2];
  let z = intVars[3];
  let setNum = 0;
  return {
    set: (variable: Var, value: number) => {
      setNum++;
      if (isNaN(value)) throw new Error(`Invalid set: ${variable}; #${setNum}`);
      switch (variable) {
        case "w":
          w = value;
          break;
        case "x":
          x = value;
          break;
        case "y":
          y = value;
          break;
        case "z":
          z = value;
          break;
        default:
          throw new Error(`Invalid variable ${variable}`);
      }
    },
    get: (variable: Var) => {
      switch (variable) {
        case "w":
          return w;
        case "x":
          return x;
        case "y":
          return y;
        case "z":
          return z;
        default:
          throw new Error(`Invalid variable ${variable}`);
      }
    },
  };
};

type ALU = ReturnType<typeof createALU>;

type Instruction = (alu: ALU, getInput?: () => number) => void;

const createInp =
  (a: Var): Instruction =>
  (alu, getInput) =>
    alu.set(a, getInput!());

const createAdd =
  (a: Var, b: string): Instruction =>
  (alu) =>
    alu.set(a, alu.get(a) + (isVar(b) ? alu.get(b) : Number(b)));

const createMul =
  (a: Var, b: string): Instruction =>
  (alu) =>
    alu.set(a, alu.get(a) * (isVar(b) ? alu.get(b) : Number(b)));

const createDiv =
  (a: Var, b: string): Instruction =>
  (alu) => {
    const divisor = isVar(b) ? alu.get(b) : Number(b);
    if (!divisor) throw new Error(`Invalid divisor: ${divisor}`);
    alu.set(a, Math.floor(alu.get(a) / divisor));
  };

const createMod =
  (a: Var, b: string): Instruction =>
  (alu) => {
    const numB = isVar(b) ? alu.get(b) : Number(b);
    const numA = alu.get(a);
    if (numA < 0 || numB <= 0)
      throw new Error(`Invalid MOD op: a = ${numA}; b = ${numB}`);
    alu.set(a, numA % numB);
  };

const createEql =
  (a: Var, b: string): Instruction =>
  (alu) =>
    alu.set(a, alu.get(a) === (isVar(b) ? alu.get(b) : Number(b)) ? 1 : 0);

type ParsedInput = Instruction[][];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l != "")
    .reduce((p, l) => {
      const splitted = l.split(" ");
      const a = splitted[1] as Var;
      let instr: Instruction | undefined;
      switch (splitted[0]) {
        case "inp":
          p.push([createInp(a)]);
          break;
        case "add":
          instr = createAdd(a, splitted[2]);
          break;
        case "mul":
          instr = createMul(a, splitted[2]);
          break;
        case "div":
          instr = createDiv(a, splitted[2]);
          break;
        case "mod":
          instr = createMod(a, splitted[2]);
          break;
        case "eql":
          instr = createEql(a, splitted[2]);
          break;
        default:
          throw new Error(`Invalid instruction ${l}`);
      }
      if (instr) {
        p[p.length - 1].push(instr);
      }
      return p;
    }, [] as Instruction[][]);

const executeSection = (
  input: number,
  intVars: IntVars,
  instructions: Instruction[]
): IntVars => {
  const alu = createALU(...intVars);
  for (const instruction of instructions) {
    instruction(alu, () => input);
  }
  return [alu.get("w"), alu.get("x"), alu.get("y"), alu.get("z")];
};

const cache = new Map<string, IntVars>();

const getOrExecSection = (
  input: number,
  intVars: IntVars,
  instructions: Instruction[],
  currentSection: number
) => {
  // Map key is currentSection + input + intVars
  const key = `${currentSection},${input},${intVars}`;
  if (cache.has(key)) {
    console.log("hit cache!");
    return cache.get(key)!;
  }
  return executeSection(input, intVars, instructions);
};

const executeSectionRec = (
  intVars: IntVars,
  sections: Instruction[][],
  currentSection: number
): number | undefined => {
  for (let i = 9; i > 0; i--) {
    if (currentSection < 7) {
      console.log(i * Math.pow(10, sections.length - 1 - currentSection));
    }
    const vars = executeSection(i, intVars, sections[currentSection]);
    if (currentSection === sections.length - 1) {
      if (vars[3] === 0) return i;
      else return;
    }
    const res = executeSectionRec(vars, sections, currentSection + 1);
    if (res) {
      return i * Math.pow(10, sections.length - 1 - currentSection) + res;
    }
  }
};

const bruteForce = (input: ParsedInput) =>
  executeSectionRec([0, 0, 0, 0], input, 0);

const executePart1 = (input: ParsedInput) => {
  const highestModelNum = bruteForce(input);

  return highestModelNum;
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day24: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day24;
