interface AOCDayResponse {
  level: number;
  part1: string;
  part2: string;
}

export type AOCDay = () => Promise<AOCDayResponse>;

export interface AOCModule {
  default: AOCDay;
}
