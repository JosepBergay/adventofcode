import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 18;

type SnailFishNum = [number | SnailFishNum, number | SnailFishNum];

// We can model a snailfish number as a Binary Tree where values are stored on leaves nodes.
type Tree = {
  left: Tree | number;
  right: Tree | number;
};

type ParsedInput = Tree[];

const buildTree = (sfNum: SnailFishNum): Tree => {
  let left;
  let right;

  if (Array.isArray(sfNum[0])) {
    left = buildTree(sfNum[0]);
  } else {
    left = sfNum[0];
  }
  if (Array.isArray(sfNum[1])) {
    right = buildTree(sfNum[1]);
  } else {
    right = sfNum[1];
  }
  return { left, right };
};

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l != "")
    .map((l) => JSON.parse(l))
    .map(buildTree);

const getMaxDepth = (arr: number | SnailFishNum): number =>
  typeof arr === "number" ? 0 : 1 + Math.max(...arr.map(getMaxDepth));

const treeToArray = (tree: Tree): SnailFishNum => {
  const leftStr = isNum(tree.left) ? tree.left : treeToArray(tree.left);
  const rightStr = isNum(tree.right) ? tree.right : treeToArray(tree.right);
  return [leftStr, rightStr];
};

const logTree = (tree: Tree) => console.log(JSON.stringify(treeToArray(tree)));

const isNum = (n: any): n is number => typeof n === "number";

type Direction = "left" | "right";

const findExploded = (
  tree: Tree,
  depth: number
): [Direction[], [number, number]?] => {
  if (depth === 4) {
    // Exploding pairs will always consist of two regular numbers.
    // Redundant check to keep TS happy.
    if (isNum(tree.left) && isNum(tree.right))
      return [[], [tree.left, tree.right]];
  }

  if (!isNum(tree.left)) {
    const [path, expl] = findExploded(tree.left, depth + 1);
    if (expl) {
      return [["left", ...path], expl];
    }
  }
  if (!isNum(tree.right)) {
    const [path, expl] = findExploded(tree.right, depth + 1);
    if (expl) {
      return [["right", ...path], expl];
    }
  }
  return [[], undefined];
};

const getBranch = (tree: Tree, direction: Direction) =>
  (direction === "left" ? tree.left : tree.right) as Tree;

const explodePair = (tree: Tree, path: Direction[]) => {
  const dir = path.shift();
  if (path.length === 0) {
    // Explode it!
    if (dir === "left") tree.left = 0;
    else tree.right = 0;
  } else explodePair(getBranch(tree, dir!), path.slice());
};

const assignWithPath = (tree: Tree, path: Direction[], v: number) => {
  let current = tree;
  for (const dir of path) {
    if (dir === "left" && isNum(current.left)) {
      current.left = current.left + v;
      break;
    } else if (dir === "right" && isNum(current.right)) {
      current.right = current.right + v;
      break;
    } else {
      // Keep going
      current = getBranch(current, dir);
    }
  }
};

const assignLeft = (tree: Tree, path: Direction[], v: number) => {
  if (path.every((d) => d === "left")) return; // Exploded is already leftmost

  // When last went right go left then right all the way
  const lastRight = path.lastIndexOf("right");
  const dirs = [
    ...path.slice(0, lastRight),
    "left",
    ...path.slice(lastRight + 1).map((_) => "right"),
    // Exploded pair is at len 4, adding extra step in case left is at len 5
    "right",
  ] as Direction[];

  assignWithPath(tree, dirs, v);
};

const assignRight = (tree: Tree, path: Direction[], v: number) => {
  if (path.every((d) => d === "right")) return; // Exploded is already rightmost

  // When last went left go right then left all the way
  const lastLeft = path.lastIndexOf("left");
  const dirs = [
    ...path.slice(0, lastLeft),
    "right",
    ...path.slice(lastLeft + 1).map((_) => "left"),
    // Exploded pair is at len 4, adding extra step in case right is at len 5
    "left",
  ] as Direction[];

  assignWithPath(tree, dirs, v);
};

const explodeTree = (tree: Tree) => {
  const [path, toExplode] = findExploded(tree, 0);
  if (!toExplode) return false;

  explodePair(tree, path.slice());
  assignLeft(tree, path, toExplode[0]);
  assignRight(tree, path, toExplode[1]);
  return true;
};

const splitTree = (tree: Tree): boolean => {
  if (isNum(tree.left) && tree.left >= 10) {
    tree.left = {
      left: Math.floor(tree.left / 2),
      right: Math.round(tree.left / 2),
    };
    return true;
  }
  let splitted = false;
  if (!isNum(tree.left)) {
    splitted = splitTree(tree.left);
  }
  if (!splitted) {
    if (isNum(tree.right) && tree.right >= 10) {
      tree.right = {
        left: Math.floor(tree.right / 2),
        right: Math.round(tree.right / 2),
      };
      return true;
    }
    if (!isNum(tree.right) && !splitted) {
      splitted = splitTree(tree.right);
    }
  }
  return splitted;
};

const add = (first: Tree, second: Tree) => {
  let added: Tree = { left: first, right: second };

  while (explodeTree(added) || splitTree(added)) {}

  return added;
};

const computeMagnitude = (tree: Tree) => {
  let left: number;
  if (isNum(tree.left)) {
    left = tree.left * 3;
  } else {
    left = computeMagnitude(tree.left) * 3;
  }
  let right: number;
  if (isNum(tree.right)) {
    right = tree.right * 2;
  } else {
    right = computeMagnitude(tree.right) * 2;
  }
  return left + right;
};

const executePart1 = (input: ParsedInput) => {
  const addition = input.reduce(add);

  // console.log("addition");
  // logTree(addition);

  return computeMagnitude(addition);
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day18: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day18;
