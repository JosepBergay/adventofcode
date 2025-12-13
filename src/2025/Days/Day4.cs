
using System.Collections;
using System.Numerics;

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

public class Map2D<T>
{
    private readonly List<List<T>> map = [];

    public Map2D(List<List<T>> list)
    {
        map = list;
    }

    public IEnumerable<ItemAndPosition<T>> Iter()
    {
        for (int y = 0; y < map.Count; y++)
        {
            for (int x = 0; x < map.Count; x++)
            {
                var item = map[y][x];
                yield return new(item, new Point2D(x, y));
            }
        }
    }

    public T? Get(Point2D p)
    {
        return map[p.y][p.x];

    }
    public T? Get(int x, int y)
    {
        return map[y][x];
    }

    /**
     * Ensure map has enough size or bad things will happen.
     */
    public void Set(Point2D p, T item)
    {
        map[p.y][p.x] = item;
    }

    public IEnumerable<Point2D> GetAdjacents(Point2D p, bool diagonals = false)
    {
        var maxYIdx = map.Count - 1;
        var maxXIdx = maxYIdx > 0 ? map[0].Count - 1 : 0;

        for (int dy = -1; dy <= 1; dy++)
        {
            for (int dx = -1; dx <= 1; dx++)
            {
                if (dx == 0 && dy == 0) continue;

                if (!diagonals && dx != 0 && dy != 0) continue;

                var x = dx + p.x;
                var y = dy + p.y;

                if (y > maxYIdx || x > maxXIdx || y < 0 || x < 0) continue; // Out of bounds!

                yield return new(x, y);
            }
        }
    }

    public static Map2D<char> FromString(string str)
    {
        List<List<char>> list = [];
        var size = (int)Math.Sqrt(str.Length);
        List<char> line = new(size);

        for (int i = 0; i < str.Length; i++)
        {
            var c = str[i];

            if (c == '\n' || i == str.Length - 1)
            {
                list.Add(line);
                line = new(size);
                continue;
            }

            line.Add(c);
        }

        return new(list);
    }
}

public record Point2D(int x, int y) : IAdditionOperators<Point2D, Point2D, Point2D>
{
    public static Point2D operator +(Point2D left, Point2D right)
    {
        return new Point2D(left.x + right.x, left.y + right.y);
    }
}

public record ItemAndPosition<T>(T Item, Point2D Pos);

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