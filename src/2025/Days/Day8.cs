using System.Numerics;

using T = ((Point3D P1, Point3D P2, double Distance)[] Connections, int Count);

public class Day8 : BaseDay<T>
{
    public override T Parse(string input)
    {
        var list = new List<Point3D>();

        var reader = new StringReader(input);

        while (true)
        {
            var line = reader.ReadLine();

            if (line is null) break;

            list.Add(Point3D.From(line.Split(",").Select(int.Parse)));
        }

        return ([.. list
            .SkipLast(1)
            .SelectMany((p1, y) =>
                list
                    .Skip(y + 1)
                    .Select(p2 => (p1, p2, distance: p1.Distance(p2))))
            .OrderBy(it => it.distance)], list.Count);
    }

    public override string Part1(T parsed)
    {
        var maxCount = parsed.Count > 20 ? 1000 : 10;
        var adjDict = new Dictionary<Point3D, HashSet<Point3D>>();

        foreach (var (p1, p2, _) in parsed.Connections.Take(maxCount))
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

    public override string Part2(T parsed)
    {
        var circuits = new List<HashSet<Point3D>>();
        var i = 0;

        while (circuits.Count != 1 || circuits.First().Count != parsed.Count)
        {
            var (P1, P2, Distance) = parsed.Connections[i++];

            var c1 = circuits.Find(c => c.Contains(P1));
            var c2 = circuits.Find(c => c.Contains(P2));

            if (c1 is null && c2 is null)
            {
                // New circuit
                circuits.Add([P1, P2]);
            }
            else if (c1 == c2)
            {
                // Already added in same circuit
            }
            else if (c1 is not null && c2 is not null)
            {
                // Merge circuits
                c1.UnionWith(c2);
                circuits.Remove(c2);
            }
            else
            {
                c1?.UnionWith([P1, P2]);
                c2?.UnionWith([P1, P2]);
            }
        }

        var connection = parsed.Connections[i - 1];

        return ((long)connection.P1.X * connection.P2.X).ToString();
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