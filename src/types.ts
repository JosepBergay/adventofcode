export type AOCDay = () => Promise<AOCDayResponse>;

interface AOCDayResponse {
  level: number;
  part1: string;
  part2: string;
}
