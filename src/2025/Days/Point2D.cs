using System.Numerics;

public record Point2D(int x, int y) : IAdditionOperators<Point2D, Point2D, Point2D>
{
    public static Point2D operator +(Point2D left, Point2D right)
    {
        return new Point2D(left.x + right.x, left.y + right.y);
    }
}
