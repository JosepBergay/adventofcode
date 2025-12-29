using System.Numerics;

public class Day8 : BaseDay<(Point3D, Point3D, double)[]>
{
    public override (Point3D, Point3D, double)[] Parse(string input)
    {
        var list = new List<Point3D>();

        var reader = new StringReader(input);

        while (true)
        {
            var line = reader.ReadLine();

            if (line is null) break;

            list.Add(Point3D.From(line.Split(",").Select(int.Parse)));
        }

        return list
            .SkipLast(1)
            .SelectMany((p1, y) =>
                list
                    .Skip(y + 1)
                    .Select(p2 => (p1, p2, distance: p1.Distance(p2))))
            .OrderBy(it => it.distance)
            .ToArray();
    }

    public override string Part1((Point3D, Point3D, double)[] parsed)
    {
        var maxCount = parsed.Length > 20 ? 1000 : 10;
        var adjDict = new Dictionary<Point3D, HashSet<Point3D>>();

        foreach (var (p1, p2, _) in parsed.Take(maxCount))
        {
            if (!adjDict.TryGetValue(p1, out HashSet<Point3D>? neighbours1))
            {
                neighbours1 = [];
                adjDict[p1] = neighbours1;
            }
            neighbours1.Add(p2);

            if (!adjDict.TryGetValue(p2, out HashSet<Point3D>? neighbours2))
            {
                neighbours2 = [];
                adjDict[p2] = neighbours2;
            }
            neighbours2.Add(p1);
        }

        var circuits = new List<HashSet<Point3D>>();

        while (adjDict.Count != 0)
        {
            var q = new Queue<Point3D>();
            q.Enqueue(adjDict.Keys.First());
            var circuit = new HashSet<Point3D>();

            while (q.TryDequeue(out var p))
            {
                if (!adjDict.TryGetValue(p, out var neighbours)) continue;

                foreach (var n in neighbours)
                {
                    q.Enqueue(n);
                }
                circuit.Add(p);
                adjDict.Remove(p);
            }

            circuits.Add(circuit);
        }

        return circuits
            .Select(c => c.Count)
            .OrderDescending()
            .Take(3)
            .Aggregate((a, b) => a * b)
            .ToString();
    }

    public override string Part2((Point3D, Point3D, double)[] parsed)
    {
        return "";
    }
}

public record Point3D(int X, int Y, int Z) : ISubtractionOperators<Point3D, Point3D, Point3D>
{
    public static Point3D From(IEnumerable<int> collection)
    {
        if (collection.Count() != 3) throw new ArgumentException("Must have exactly 3 elements");

        return new(collection.First(), collection.ElementAt(1), collection.ElementAt(2));
    }

    public static Point3D operator -(Point3D left, Point3D right)
    {
        return new Point3D(left.X - right.X, left.Y - right.Y, left.Z - right.Z);
    }

    public double Length() => Math.Sqrt((long)X * (long)X + (long)Y * (long)Y + (long)Z * (long)Z);

    public double Distance(Point3D other) => (this - other).Length();

}