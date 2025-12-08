public class Day2 : BaseDay<string>
{
    public override string Parse(string input)
    {
        return input;
    }

    // Sequence of digits repeated twice
    static public bool IsInvalidId(string str)
    {
        for (int i = 0, j = str.Length / 2; i < str.Length / 2; i++, j++)
        {
            if (str[i] != str[j]) return false;
        }

        return true;
    }


    public override string Part1(string parsed)
    {
        return parsed.Split(",")
            .SelectMany(str =>
            {
                var split = str.Split("-");
                var low = long.Parse(split[0]);
                var high = long.Parse(split[1]);
                return low.ToRange(high);
            })
            .Where(i => (Math.Truncate(Math.Log10(i)) + 1) % 2 == 0) // Even number of digits
            .Where(i => IsInvalidId(i.ToString()))
            .Aggregate(0L, (a, b) => a + b)
            .ToString();
    }



    public override string Part2(string parsed)
    {
        return "";
    }
}