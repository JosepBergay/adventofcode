import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 21;

type ParsedInput = number[];

const parser = (input: string): ParsedInput =>
  input
    .split("\n")
    .filter((l) => l !== "")
    .map((str) => parseInt(str.slice(-1)));

const playDiracDice = (start1: number, start2: number) => {
  let dice = 1;
  let rollCount = 0;
  let score1 = 0;
  let score2 = 0;
  let position1 = start1;
  let position2 = start2;

  const rollDice = () => {
    if (dice > 100) dice = 1;
    ++rollCount;
    return dice++;
  };

  const movePawn = (player: 1 | 2, rolled: number) => {
    if (player === 1) {
      const aux = (position1 + rolled) % 10;
      position1 = aux ? aux : 10;
      return position1;
    } else {
      const aux = (position2 + rolled) % 10;
      position2 = aux ? aux : 10;
      return position2;
    }
  };

  const play = (player: 1 | 2) => {
    // Player rolls 3 times
    let rolled = rollDice();
    rolled += rollDice();
    rolled += rollDice();

    return movePawn(player, rolled);
  };

  const winCondition = (score: number) => score >= 1000;

  while (true) {
    score1 += play(1);
    if (winCondition(score1)) break;
    score2 += play(2);
    if (winCondition(score2)) break;
  }
  return { score1, score2, rollCount };
};

const executePart1 = (input: ParsedInput) => {
  const { score1, score2, rollCount } = playDiracDice(input[0], input[1]);

  const loser = score1 < score2 ? score1 : score2;

  return loser * rollCount;
};

const executePart2 = (input: ParsedInput) => {
  return "";
};

const day21: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day21;
