import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 4;

type Row = [number, number, number, number, number];
type Board = [Row, Row, Row, Row, Row];

type ParsedInput = {
  draws: number[];
  boards: Board[];
};

const parser = (input: string): ParsedInput => {
  const [drawsStr, ...boardsStr] = input.split("\n").filter((i) => i != "");

  const draws = drawsStr.split(",").map(n => parseInt(n));

  const boards = boardsStr.reduce((p, c, i) => {
    const boardIndex = i % 5 == 0 ? p.length : p.length - 1;
    if (i % 5 == 0) {
      p[boardIndex] = [];
    }
    p[boardIndex][i%5] = c.split(" ").filter(n => n != "").map(n => parseInt(n));
    return p;
  }, [] as number[][][]) as unknown as Board[];

  return { draws, boards };
};

const executePart1 = (input: ParsedInput): string => {
  if (input.boards.length == 0 || input.draws.length == 0)
    return "No boards or draws :<";

  const mark = -1;

  const rowWinCondition = (row: Row) => row.every((n) => n === mark);

  const transposeMatrix = (board: Board) =>
    board.map((_, i) => board.map((r) => r[i])) as Board;

  const winCondition = (board: Board) =>
    board.some(rowWinCondition) || transposeMatrix(board).some(rowWinCondition);

  const markBoard = (draw: number) => (board: Board) => {
    for (let i = 0; i < board.length; i++) {
      const row = board[i];
      for (let j = 0; j < row.length; j++) {
        if (row[j] === draw) row[j] = mark;
      }
    }
  };

  const markBoards = (draw: number) => input.boards.forEach(markBoard(draw));

  // Marking first 4 draws. We can't win until we have drawn 5 times.
  input.draws.slice(0, 3).forEach(markBoards);

  let winningBoard: Board | undefined;
  let winningDraw: number | undefined;

  // Marking and checking for the rest of draws.
  input.draws.slice(4).some((draw) => {
    // Mark boards
    markBoards(draw);
    // Check winning board
    input.boards.some((board, i) => {
      const boardWins = winCondition(board);
      if (boardWins) {
        winningBoard = board;
        winningDraw = draw;
      }
      return boardWins;
    });

    return !!winningBoard;
  });

  if (!winningBoard || !winningDraw) return "There's no winning board :<";

  // Find the sum of all unmarked numbers on the winning board
  const sum = winningBoard.reduce(
    (sum, row) => sum + row.reduce((p, c) => (c != mark ? p + c : p), 0),
    0
  );

  return `${sum * winningDraw}`;
};

const executePart2 = (input: ParsedInput): string => {
  return "";
};

const day4: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = executePart1(parsed);

  const part2 = executePart2(parsed);

  return { level, part1, part2 };
};

export default day4;
