const getSession = () => {
  if (!process.env.SESSION_COOKIE) throw "NO COOKIE :<";
  return process.env.SESSION_COOKIE!;
};

export const fetchInput = async (day: number) => {
  const session = getSession();

  const res = await fetch(`https://adventofcode.com/2021/day/${day}/input`, {
    headers: {
      cookie: `session=${session}`,
    },
  });

  const body = await res.text();

  return body;
};
