using System.Numerics;

public record Point2D(int X, int Y) : IAdditionOperators<Point2D, Point2D, Point2D>, ISubtractionOperators<Point2D, Point2D, Point2D>
{
    public static Point2D operator +(Point2D left, Point2D right)
    {
        return new Point2D(left.X + right.X, left.Y + right.Y);
    }

    public static Point2D operator -(Point2D left, Point2D right)
    {
        return new Point2D(left.X - right.X, left.Y - right.Y);
    }

    public static Point2D From(IEnumerable<int> collection)
    {
        if (collection.Count() != 2) throw new ArgumentException("Must have exactly 2 elements");

        return new(collection.First(), collection.ElementAt(1));
    }

    public double Length() => Math.Sqrt((long)X * X + (long)Y * Y);

    public double Distance(Point2D other) => (this - other).Length();

    public bool IsXInProjectionOf(Vector2D vector)
    {
        var minX = Math.Min(vector.From.X, vector.To.X);
        var maxX = Math.Max(vector.From.X, vector.To.X);
        return minX <= X && X <= maxX;
    }

    public bool IsYInProjectionOf(Vector2D vector)
    {
        var minY = Math.Min(vector.From.Y, vector.To.Y);
        var maxY = Math.Max(vector.From.Y, vector.To.Y);
        return minY <= Y && Y <= maxY;
    }

    // Assuming Vector2D is the diagonal of a square. This method returns true if the point is
    // inside that square.
    public bool IsInSquare(Vector2D vector) =>
        IsXInProjectionOf(vector) && IsYInProjectionOf(vector);
}


public record Vector2D(Point2D From, Point2D To)
{

    public Point2D? IntersectsAt(Vector2D other)
    {
        // From linear equation: y = m*x + n -> y = vy/vx * x + n
        // n = y - vy/vx * x
        // To avoid division we use: vy * x - vx * y + vy * n = 0
        // (y1-y2) * x + (x2-x1) * y + (x1*y2 - x2*y1) = 0
        Point2D intersection;

        var v = From - To;
        var otherV = other.From - other.To;
        if (v.X == 0 && otherV.X == 0)
        {
            // Parallel vertical lines.
            // Assume that overlapping lines (exactly same lines) do not intersect.
            // They would intersect in all the points of the smallest of the two vectors.
            return null;
        }
        else if (v.X == 0)
        {
            // v is vertical
            if (From.IsXInProjectionOf(other))
            {
                var m = (double)otherV.Y / otherV.X;
                var n = other.From.Y - m * other.From.X;
                var y = m * From.X + n;
                intersection = new(From.X, (int)y);
            }
            else
            {
                return null;
            }
        }
        else if (otherV.X == 0)
        {
            // otherV is vertical
            if (other.From.IsXInProjectionOf(this))
            {
                var m = (double)v.Y / v.X;
                var n = From.Y - m * From.X;
                var y = m * other.From.X + n;
                intersection = new(other.From.X, (int)y);
            }
            else
            {
                return null;
            }
        }
        else if (v.Y == 0 && otherV.Y == 0)
        {
            // Both horizontal, so no intersection
            return null;
        }
        else
        {
            var m1 = (double)v.Y / v.X;
            var m2 = (double)otherV.Y / otherV.X;

            if (m1 == m2) return null; // Parallel

            // Find n: n = y - vy/vx * x
            var n1 = From.Y - m1 * From.X;
            var n2 = other.From.Y - m2 * other.From.X;

            // Solve -> x = (n2 - n1) / (m1 - m2)
            var x = (n2 - n1) / (m1 - m2);

            intersection = new Point2D((int)x, (int)(m1 * x + n1));
        }

        // We have an intersection but vectors are not infinite, so check if intersection is valid.
        if (intersection.IsInSquare(this) && intersection.IsInSquare(other)) return intersection;

        return null;
    }

    public bool Intersects(Vector2D other)
    {
        return IntersectsAt(other) is not null;
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
