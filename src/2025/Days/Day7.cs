public class Day7 : BaseDay<(Map2D<char>, Point2D)>
{
    public override (Map2D<char>, Point2D) Parse(string input)
    {
        var map = Map2D<char>.FromString(input);
        var (_, start) = map.Iter().First((it) => it.Item == 'S');
        return (map, start);
    }

    public override string Part1((Map2D<char>, Point2D) parsed)
    {
        var (map, start) = parsed;

        var count = 0;
        var curr = new HashSet<Point2D>() { start };

        while (curr.Count > 0)
        {
            var tmp = new HashSet<Point2D>();
            foreach (var c in curr)
            {
                var moved = c + Direction.N;
                if (map.IsOutOfBounds(moved)) continue;
                if (map.Get(moved) == '^')
                {
                    count++;
                    tmp.Add(moved + Direction.W);
                    tmp.Add(moved + Direction.E);
                }
                else
                {
                    tmp.Add(moved);
                }
            }
            curr = tmp;
        }

        return count.ToString();
    }

    private static long MoveParticleDfs(Point2D curr, Map2D<char> map, Dictionary<Point2D, long> cache)
    {
        var moved = curr + Direction.N;

        if (map.IsOutOfBounds(moved)) return 1;

        if (cache.ContainsKey(moved)) return cache[moved];

        var count = 0L;
        if (map.Get(moved) == '^')
        {
            count += MoveParticleDfs(moved + Direction.W, map, cache);
            count += MoveParticleDfs(moved + Direction.E, map, cache);
        }
        else
        {
            count += MoveParticleDfs(moved, map, cache);
        }

        cache[moved] = count;

        return count;
    }

    public override string Part2((Map2D<char>, Point2D) parsed)
    {
        return MoveParticleDfs(parsed.Item2, parsed.Item1, new()).ToString();
    }
}