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
      position1 = (position1 + rolled) % 10 || 10;
      return position1;
    } else {
      position2 = (position2 + rolled) % 10 || 10;
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

const playQuantumDice = (start1: number, start2: number) => {
  const diceMap: Record<number, number> = {};
  // Compute probability
  for (let i = 1; i <= 3; i++) {
    for (let j = 1; j <= 3; j++) {
      for (let k = 1; k <= 3; k++) {
        const sum = i + j + k;
        diceMap[sum] = (diceMap[sum] || 0) + 1;
      }
    }
  }

  let wins1Count = 0;
  let wins2Count = 0;

  const play = (
    player: 1 | 2,
    pos1: number,
    pos2: number,
    score1: number,
    score2: number,
    accUniverses: number
  ) => {
    for (let roll = 3; roll <= 9; roll++) {
      const universes = accUniverses * diceMap[roll];

      if (player === 1) {
        // Save state for next loop
        let currentP = pos1;
        let currentS = score1;
        // Apply
        pos1 = (pos1 + roll) % 10 || 10;
        score1 += pos1;

        if (score1 >= 21) wins1Count += universes;
        else play(2, pos1, pos2, score1, score2, universes);

        // Reset state for next loop
        pos1 = currentP;
        score1 = currentS;
      } else {
        // Same as before but for player2
        let currentP = pos2;
        let currentS = score2;

        pos2 = (pos2 + roll) % 10 || 10;
        score2 += pos2;

        if (score2 >= 21) wins2Count += universes;
        else play(1, pos1, pos2, score1, score2, universes);

        pos2 = currentP;
        score2 = currentS;
      }
    }
  };

  play(1, start1, start2, 0, 0, 1);

  return { wins1Count, wins2Count };
};

const executePart1 = (input: ParsedInput) => {
  const { score1, score2, rollCount } = playDiracDice(input[0], input[1]);

  const loser = score1 < score2 ? score1 : score2;

  return loser * rollCount;
};

const executePart2 = (input: ParsedInput) => {
  const { wins1Count, wins2Count } = playQuantumDice(input[0], input[1]);

  return wins1Count > wins2Count ? wins1Count : wins2Count;
};

const day21: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day21;
