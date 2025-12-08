public class Day2 : BaseDay<IEnumerable<long>>
{
    public override IEnumerable<long> Parse(string input) => input.Split(",")
            .SelectMany(str =>
            {
                var split = str.Split("-");
                var low = long.Parse(split[0]);
                var high = long.Parse(split[1]);
                return low.ToRange(high);
            });

    // Sequence of digits repeated twice
    static public bool IsInvalidIdP1(string str)
    {
        for (int i = 0, j = str.Length / 2; i < str.Length / 2; i++, j++)
        {
            if (str[i] != str[j]) return false;
        }

        return true;
    }

    static public bool IsInvalidId(string str)
    {
        var factors = str.Length.Factors().ToArray();
        // System.Console.WriteLine($"{str} -> {String.Join(',', factors)}");
        factors.Sort();

        var res = factors.Skip(1).Any(f =>
        {
            var sliceLen = str.Length / f;

            var indexes = Enumerable.Range(0, f).Select(i => i * sliceLen).ToArray();

            while (indexes[0] < sliceLen)
            {
                var currV = str[indexes[0]];

                for (int i = 0; i < indexes.Length; i++)
                {
                    if (i != 0)
                    {
                        if (currV != str[indexes[i]]) return false;
                    }

                    indexes[i]++;
                }
            }

            return true;
        });

        return res;
    }

    public override string Part1(IEnumerable<long> parsed)
    {
        return parsed
            .Where(i => (Math.Truncate(Math.Log10(i)) + 1) % 2 == 0) // Even number of digits
            .Where(i => IsInvalidIdP1(i.ToString()))
            .Sum()
            .ToString();
    }



    public override string Part2(IEnumerable<long> parsed)
    {
        return parsed
            .Where(i => IsInvalidId(i.ToString()))
            .Sum()
            .ToString();
    }
}