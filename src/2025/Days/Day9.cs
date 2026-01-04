using T = TheaterFloor;

public record PointPair(Point2D P1, Point2D P2, long Area);

public record TheaterFloor(
    List<Point2D> Points,
    IOrderedEnumerable<PointPair> Pairs
);

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

            points.Add(Point2D.From(line.Split(',').Select(int.Parse)));
        }

        var pairs = points
            .SkipLast(1)
            .SelectMany((p1, i) =>
                points
                    .Skip(i + 1)
                    .Select(p2 =>
                    {
                        var diffX = Math.Abs(p1.X - p2.X) + 1L;
                        var diffY = Math.Abs(p1.Y - p2.Y) + 1L;
                        return new PointPair(p1, p2, Area: diffX * diffY);
                    }))
            .OrderByDescending(it => it.Area);

        return new(points, pairs);
    }

    public override string Part1(T parsed) => parsed.Pairs.First().Area.ToString();

    public override string Part2(T parsed)
    {
        var p = parsed.Pairs.First(pair =>
        {
            var p3 = new Point2D(pair.P1.X, pair.P2.Y);
            var p4 = new Point2D(pair.P2.X, pair.P1.Y);

            var sides = new Vector2D[4]
            {
                new (pair.P1, p3),
                new (p3, pair.P2),
                new (pair.P2, p4),
                new (p4, pair.P1),
            };

            return parsed.Points
                .Zip(parsed.Points.Skip(1).Concat(parsed.Points), (p1, p2) => new Vector2D(p1, p2))
                .All(v =>
                {
                    var isEqual = pair.P1 == v.From ||
                        pair.P2 == v.From ||
                        pair.P1 == v.To ||
                        pair.P2 == v.To;

                    if (isEqual) return true;

                    foreach (var side in sides)
                    {
                        var intersection = v.IntersectsAt(side);
                        if (intersection != null)
                        {
                            var diagonal = new Vector2D(pair.P1, pair.P2);
                            if ((intersection != v.From && intersection != v.To)
                                || (intersection == v.From && v.To.IsInSquare(diagonal))
                                || (intersection == v.To && v.From.IsInSquare(diagonal)))
                            {
                                return false;
                            }
                        }
                    }

                    return true;
                });
        });

        return p.Area.ToString();
    }
}
