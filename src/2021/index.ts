import type { AOCModule } from "./types";
import { promises } from "fs";
import consola from "consola";

const main = async () => {
  const dirs = await promises.readdir("./dist", { withFileTypes: true });

  const dayNum = process.argv[2];

  if (dayNum && isNaN(parseInt(dayNum))) {
    consola.error(new Error("Argument is not a number"));
    return;
  }

  const days: AOCModule[] = await Promise.all(
    dirs
      .filter((d) =>
        d.isDirectory() && dayNum
          ? d.name === `day${dayNum}`
          : d.name.startsWith("day")
      )
      .map((d) => import(`./${d.name}/index.js`))
  );

  if (!days.length) {
    consola.warn(dayNum ? `Day ${dayNum} not done yet` : "No exercises found");
    return;
  }

  consola.start(
    `Running ${days.length > 1 ? `${days.length} exercises` : `Day ${dayNum}`}`
  );

  const start = Date.now();

  const responses = await Promise.all(days.map((d) => d.default()));

  consola.success(`Done in ${Date.now() - start} ms`);

  consola.info(`Answer${dayNum ? "" : "s"}:`);
  for (const res of responses) {
    consola.info(`Day #${res.level} part1: ${res.part1}, part2: ${res.part2}`);
  }
};

main();
