using T = Point2D[];

public class Day9 : BaseDay<T>
{
    public override T Parse(string input)
    {
        var points = new List<Point2D>();
        var reader = new StringReader(input);

        while (true)
        {
            var line = reader.ReadLine();

            if (line == null) break;

            points.Add(Point2D.From(line.Split(',').Select(c => int.Parse(c))));
        }

        return [.. points];
    }

    public override string Part1(T parsed)
    {
        var it = parsed
            .SkipLast(1)
            .SelectMany((p1, y) =>
                parsed
                    .Skip(y + 1)
                    .Select((p2) =>
                    {
                        var diffX = Math.Abs(p1.X - p2.X) + 1L;
                        var diffY = Math.Abs(p1.Y - p2.Y) + 1L;
                        return (p1, p2, Area: diffX * diffY);
                    }))
            .OrderByDescending(it => it.Area)
            .First();

        return it.Area.ToString();
    }

    public override string Part2(T parsed)
    {
        return "";
    }
}
