public class Day7 : BaseDay<Map2D<char>>
{
    public override Map2D<char> Parse(string input)
    {
        return Map2D<char>.FromString(input);
    }

    public override string Part1(Map2D<char> parsed)
    {
        var (_, s) = parsed.Iter().First((it) => it.Item == 'S');

        var count = 0;

        var curr = new HashSet<Point2D>() { s };

        while (curr.Count > 0)
        {
            var tmp = new HashSet<Point2D>();
            foreach (var c in curr)
            {
                var moved = c + Direction.N;
                if (parsed.IsOutOfBounds(moved)) continue;
                if (parsed.Get(moved) == '^')
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
            // curr.Select(c => c + Direction.S).SelectMany(c => parsed.Get(c) == '^' ? new Point2D[2]{
            //     c + Direction.SE,
            //     c + Direction.SW
            // } : c);

        }

        return count.ToString();
    }

    public override string Part2(Map2D<char> parsed)
    {
        return "";
    }
}