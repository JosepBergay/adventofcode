import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 8;

type Segment = "a" | "b" | "c" | "d" | "e" | "f" | "g";

type SegValue = Segment[];

type Signal = [
  SegValue,
  SegValue,
  SegValue,
  SegValue,
  SegValue,
  SegValue,
  SegValue,
  SegValue,
  SegValue,
  SegValue
];

type OutputValue = [SegValue, SegValue, SegValue, SegValue];

type SignalOutput = [Signal, OutputValue];

type ParsedInput = SignalOutput[];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l != "")
    .map(
      (line) =>
        line.split("|").map((s) =>
          s
            .split(" ")
            .filter((c) => c != "")
            .map((c) => c.split(""))
        ) as SignalOutput
    );

const isOne = (seg: SegValue) => seg.length === 2;

const isFour = (seg: SegValue) => seg.length === 4;

const isSeven = (seg: SegValue) => seg.length === 3;

const isEight = (seg: SegValue) => seg.length === 7;

const executePart1 = (input: ParsedInput): string => {
  const outputValues = input.map(([_, v]) => v);

  const instances = outputValues
    .flat()
    .filter((v) => isOne(v) || isFour(v) || isSeven(v) || isEight(v));

  return `${instances.length}`;
};

/** Display configuration:
 *
 * One: len === 2.
 * Four: len === 4.
 * Seven: len === 3.
 * Eight: len === 7.
 * Three: len === 5 && contains(One).
 * Nine: len === 6 && contains(Three).
 * Zero: len === 6 && contains(One) && (!contains(Nine) || !contains(Three))
 * Six: len === 6 && !isNine && !isZero
 * Two: len === 5 && (contains exactly 2 of four)
 * Five: len === 5 && (contains exactly 3 of four) && !isThree
 */

const eight: SegValue = ["a", "b", "c", "d", "e", "f", "g"];

const isThree = (seg: SegValue, one: SegValue) =>
  seg.length === 5 && one.every((e) => seg.includes(e));

const isNine = (seg: SegValue, three: SegValue) =>
  seg.length === 6 && three.every((e) => seg.includes(e));

const isZero = (seg: SegValue, one: SegValue, nineOrThree: SegValue) =>
  seg.length === 6 &&
  one.every((e) => seg.includes(e)) &&
  !nineOrThree.every((e) => seg.includes(e));

const isSix = (
  seg: SegValue,
  one: SegValue,
  three: SegValue,
  nineOrThree: SegValue
) => seg.length === 6 && !isNine(seg, three) && !isZero(seg, one, nineOrThree);

const isFive = (seg: SegValue, four: SegValue, one: SegValue) =>
  seg.length === 5 &&
  seg.filter((e) => four.includes(e)).length === 3 &&
  !isThree(seg, one);

const isTwo = (seg: SegValue, four: SegValue) =>
  seg.length === 5 && seg.filter((e) => four.includes(e)).length === 2;

const findNums = (seg: SegValue[]) => {
  const map: Record<string, SegValue> = {
    eight: eight,
  };

  const one = seg.find(isOne);
  one && (map["one"] = one);

  const four = seg.find(isFour);
  four && (map["four"] = four);

  const seven = seg.find(isSeven);
  seven && (map["seven"] = seven);

  const three = one && seg.find((s) => isThree(s, one));
  three && (map["three"] = three);

  const nine = three && seg.find((s) => isNine(s, three));
  nine && (map["nine"] = nine);

  const zero = three && seg.find((s) => isZero(s, one, nine || three));
  zero && (map["zero"] = zero);

  const six = three && seg.find((s) => isSix(s, one, three, nine || three));
  six && (map["six"] = six);

  const five = four && one && seg.find((s) => isFive(s, four, one));
  five && (map["five"] = five);

  const two = four && seg.find((s) => isTwo(s, four));
  two && (map["two"] = two);

  return map;
};

const convertKeyToNum = (key: string) => {
  switch (key) {
    case "zero":
      return "0";
    case "one":
      return "1";
    case "two":
      return "2";
    case "three":
      return "3";
    case "four":
      return "4";
    case "five":
      return "5";
    case "six":
      return "6";
    case "seven":
      return "7";
    case "eight":
      return "8";
    case "nine":
      return "9";
    default:
      return "-2";
  }
};

const convertToNumStr = (seg: SegValue, map: Record<string, SegValue>) => {
  for (const key in map) {
    const element = map[key];
    if (
      element.length === seg.length &&
      seg.every((e) => element.includes(e))
    ) {
      return convertKeyToNum(key);
    }
  }
  return "-1";
};

const executePart2 = (input: ParsedInput): string => {
  let sum = 0;
  for (const [signal, output] of input) {
    const map = findNums([...signal, ...output]);

    const numStr = output.reduce(
      (p, o) => `${p}${convertToNumStr(o, map)}`,
      ""
    );

    sum = sum + parseInt(numStr);
  }

  return `${sum}`;
};

const day8: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day8;
