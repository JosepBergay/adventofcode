using System.Numerics;

public record Point2D(int X, int Y) : IAdditionOperators<Point2D, Point2D, Point2D>
{
    public static Point2D operator +(Point2D left, Point2D right)
    {
        return new Point2D(left.X + right.X, left.Y + right.Y);
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
