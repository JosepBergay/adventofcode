public abstract class BaseDay<T> : IDay
{
    public abstract T Parse(string input);

    public abstract string Part1(T parsed);

    public abstract string Part2(T parsed);

    public DayResult Exec(string input)
    {
        var parsed = Parse(input);

        return new(Part1(parsed), Part2(parsed));
    }
}