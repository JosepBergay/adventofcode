public class Day5 : BaseDay<(List<(long, long)>, List<long>)>
{
    public override (List<(long, long)>, List<long>) Parse(string input)
    {
        var ranges = new List<(long, long)>();
        var ingredients = new List<long>();

        var reader = new StringReader(input);
        while (true)
        {
            var line = reader.ReadLine();

            if (line is null) break;
            if (line == "") continue;

            var split = line.Split("-");
            if (split.Length == 2)
            {
                ranges.Add((long.Parse(split[0]), long.Parse(split[1])));
            }
            else if (split.Length == 1)
            {
                ingredients.Add(long.Parse(split[0]));
            }
        }

        return (ranges, ingredients);
    }

    public override string Part1((List<(long, long)>, List<long>) parsed)
    {
        var (ranges, ingredients) = parsed;

        long freshCount = 0;
        foreach (var ingredient in ingredients)
        {
            foreach (var (low, high) in ranges)
            {
                if (low <= ingredient && ingredient <= high)
                {
                    freshCount++;
                    break;
                }
            }
        }

        return freshCount.ToString();
    }

    public override string Part2((List<(long, long)>, List<long>) parsed)
    {
        var (ranges, _) = parsed;

        var curr = new List<(long, long)>();

        for (int i = 0; i < ranges.Count; i++)
        {
            var newRange = ranges[i];

            var newRanges = new List<(long, long)>(curr.Count);
            (long, long)? lowerRange = null;
            (long, long)? upperRange = null;
            var skipAdd = false;
            foreach (var oldRange in curr)
            {
                if (oldRange.Item1 <= newRange.Item1 && newRange.Item2 <= oldRange.Item2)
                {
                    // newRange is completely included in oldRange, so items are already included
                    skipAdd = true;
                    break;
                }
                else if (oldRange.Item1 <= newRange.Item1 && newRange.Item1 <= oldRange.Item2)
                {
                    // newRange starts inside oldRange
                    lowerRange = oldRange;
                }
                else if (oldRange.Item1 <= newRange.Item2 && newRange.Item2 <= oldRange.Item2)
                {
                    // newRange ends inside oldRange
                    upperRange = oldRange;
                }
                else if (newRange.Item2 < oldRange.Item1 || oldRange.Item2 < newRange.Item1)
                {
                    // No overlap
                    newRanges.Add(oldRange);
                }
            }

            if (skipAdd)
            {
                continue;
            }

            var lowerN = lowerRange.HasValue ? lowerRange.Value.Item1 : newRange.Item1;
            var upperN = upperRange.HasValue ? upperRange.Value.Item2 : newRange.Item2;

            newRanges.Add((lowerN, upperN));

            curr = newRanges;
        }

        return curr.Sum((r) => r.Item2 - r.Item1 + 1).ToString();
    }
}