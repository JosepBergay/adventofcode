import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 14;

type ParsedInput = {
  templatePairs: string[]; // [ "NN", "NC", "CB" ]
  rules: [string, string][]; // [["CH", "B"], ...]
};

const templateToPairs = (template: string) => {
  const chars = template.split("");
  const pairs: string[] = [];
  for (let i = 1; i < chars.length; i++) {
    pairs.push(`${chars[i - 1]}${chars[i]}`);
  }
  return pairs;
};

const parser = (input: string): ParsedInput => {
  const [first, ...rest] = input.split("\n").filter((l) => l != "");
  const templatePairs = templateToPairs(first);
  const rules = rest.map((l) => l.split(" -> ") as [string, string]);
  return { templatePairs, rules };
};

const pairsToTemplate = (pairs: string[]) => {
  let template = "";
  for (const pair of pairs) {
    template = `${template}${pair[0]}`;
  }
  const lastChar = pairs[pairs.length - 1][1];
  return `${template}${lastChar}`;
};

const executePart1 = (input: ParsedInput) => {
  let pairs = input.templatePairs.slice();

  for (let step = 0; step < 10; step++) {
    const aux: string[] = [];
    for (const pair of pairs) {
      let matched = false;
      for (const rule of input.rules) {
        if (pair === rule[0]) {
          aux.push(`${pair[0]}${rule[1]}`, `${rule[1]}${pair[1]}`);
          matched = true;
        }
      }
      if (!matched) {
        aux.push(pair);
      }
    }
    pairs = aux;
  }

  const template = pairsToTemplate(pairs);

  const quantities = template.split("").reduce((p, c) => {
    if (!p[c]) p[c] = 0;
    p[c]++;
    return p;
  }, {} as Record<string, number>);

  const { most, least } = Object.keys(quantities).reduce(
    (p, key) => {
      const value = quantities[key];
      if (value > p.most.value) {
        p.most = { key, value };
      }
      if (value < p.least.value) {
        p.least = { key, value };
      }
      return p;
    },
    { most: { key: "", value: 0 }, least: { key: "", value: Infinity } }
  );

  return most.value - least.value;
};

const executePart2 = (input: ParsedInput) => {
  const map = new Map<string, number>();
  const charMap = new Map<string, number>();

  // Fill initial maps.
  for (const pair of input.templatePairs) {
    map.set(pair, (map.get(pair) ?? 0) + 1);
    const [first, _] = pair.split("");
    charMap.set(first, (charMap.get(first) ?? 0) + 1);
  }

  // Insert last char of last pair (last char of template).
  const [_, last] = input.templatePairs[input.templatePairs.length - 1];
  charMap.set(last, (charMap.get(last) ?? 0) + 1);

  for (let step = 0; step < 40; step++) {
    const iteratingMap = new Map(map.entries()); // Can't modify the same map we are iterating. Lesson learned!
    for (const [pair, counter] of iteratingMap) {
      const rule = input.rules.find(([p, _]) => p === pair);

      if (rule) {
        const insertedChar = rule[1];
        const newPair1 = `${pair[0]}${insertedChar}`;
        const newPair2 = `${insertedChar}${pair[1]}`;

        map.set(pair, map.get(pair)! - counter);
        map.set(newPair1, (map.get(newPair1) ?? 0) + counter);
        map.set(newPair2, (map.get(newPair2) ?? 0) + counter);
        charMap.set(insertedChar, (charMap.get(insertedChar) ?? 0) + counter);
      }
    }
  }

  let most = { key: "", value: 0 };
  let least = { key: "", value: Infinity };
  for (const [char, counter] of charMap) {
    if (most.value < counter) {
      most = { key: char, value: counter };
    }
    if (least.value > counter) {
      least = { key: char, value: counter };
    }
  }

  return most.value - least.value;
};

const day14: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day14;
