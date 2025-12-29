using System.Collections;
using System.Numerics;

public class Day8 : BaseDay<List<Point3D>>
{
    public override List<Point3D> Parse(string input)
    {
        var list = new List<Point3D>();

        var reader = new StringReader(input);

        while (true)
        {
            var line = reader.ReadLine();

            if (line is null) break;

            list.Add(Point3D.From(line.Split(",").Select(int.Parse)));
        }

        return list;
    }

    private static (int, int) FindMinValueIdxs(double[][] table)
    {
        var minValue = double.MaxValue;
        (int, int)? idxs = null;

        for (int y = 0; y < table.Length; y++)
        {
            var row = table[y];
            for (int x = 0; x < row.Length; x++)
            {
                var d = row[x];
                if (d < minValue)
                {
                    idxs = (y, x);
                    minValue = d;
                }
            }
        }

        if (idxs is null) throw new SystemException("There should be an available connection");

        return idxs.Value;
    }

    public override string Part1(List<Point3D> parsed)
    {
        var table = parsed
            .SkipLast(1)
            .Select((p1, y) =>
                parsed
                    .Skip(y + 1)
                    .Select(p2 => p1.Distance(p2))
                    .ToArray())
            .ToArray();

        var maxCount = parsed.Count > 20 ? 1000 : 10;
        var adjDict = new Dictionary<Point3D, HashSet<Point3D>>();

        for (int count = 0; count < maxCount; count++)
        {
            var (y, x) = FindMinValueIdxs(table);
            var p1 = parsed[y];
            var p2 = parsed[x + y + 1];

            table[y][x] = double.MaxValue;

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

        var sizes = circuits.Select(c => c.Count).ToArray();
        sizes.Sort();

        return sizes.TakeLast(3).Aggregate((a, b) => a * b).ToString();
    }

    public override string Part2(List<Point3D> parsed)
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