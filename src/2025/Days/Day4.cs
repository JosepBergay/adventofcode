public class Day4 : BaseDay<Map2D<char>>
{
    public override Map2D<char> Parse(string input)
    {
        return Map2D<char>.FromString(input);
    }

    public override string Part1(Map2D<char> map)
    {
        return map
            .Iter()
            .Count((it) =>
                it.Item == '@'
                && map
                    .GetAdjacents(it.Pos, true)
                    .Count(adj => map.Get(adj) == '@') < 4
            )
            .ToString();
    }

    public override string Part2(Map2D<char> map)
    {
        var total = 0;
        while (true)
        {
            var removed = 0;
            foreach (var (it, p) in map.Iter())
            {
                if (it != '@') continue;

                if (map.GetAdjacents(p, true).Count(adj => map.Get(adj) == '@') < 4)
                {
                    map.Set(p, '.'); // Remove paper roll
                    removed++;
                    total++;
                }
            }

            if (removed == 0) break;
        }

        return total.ToString();
    }
}

public static class Direction
{
    public static readonly Point2D N = new(0, 1);
    public static readonly Point2D NE = new(1, 1);
    public static readonly Point2D E = new(1, 0);
    public static readonly Point2D SE = new(1, -1);
    public static readonly Point2D S = new(0, -1);
    public static readonly Point2D SW = new(-1, -1);
    public static readonly Point2D W = new(-1, 0);
    public static readonly Point2D NW = new(-1, 1);
}