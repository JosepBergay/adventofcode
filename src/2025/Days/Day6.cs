public class Day6 : BaseDay<List<List<char>[]>>
{
    public override List<List<char>[]> Parse(string input)
    {
        var lines = input.Split('\n', StringSplitOptions.RemoveEmptyEntries);

        var parsed = new List<List<char>[]>();

        var operation = lines.Select(l => new List<char>()).ToArray();

        for (int i = 0; i < lines[0].Length; i++)
        {
            if (lines.All(l => l[i] == ' '))
            {
                parsed.Add(operation);
                // Reset
                operation = lines.Select(l => new List<char>()).ToArray();
            }
            else
            {
                for (int j = 0; j < lines.Length; j++)
                {
                    operation[j].Add(lines[j][i]);
                }
            }
        }

        parsed.Add(operation);

        return parsed;
    }

    public override string Part1(List<List<char>[]> parsed)
    {
        var ans = 0L;

        for (int i = 0; i < parsed.Count; i++)
        {
            var operation = parsed[i];
            var op = operation.Last();

            var nums = operation.SkipLast(1).Select(chars =>
            {
                var str = chars.Aggregate("", (a, b) => $"{a}{b}");
                return int.Parse(str);
            });

            ans += op[0] == '+' ? nums.Sum() : nums.Aggregate(1L, (a, b) => a * b);
        }

        return ans.ToString();
    }

    public override string Part2(List<List<char>[]> parsed)
    {
        var ans = 0L;

        for (int i = 0; i < parsed.Count; i++)
        {
            var operation = parsed[i];
            var op = operation.Last();
            var lines = operation.SkipLast(1);

            var nums = new int[op.Count];

            for (int j = 0; j < op.Count; j++)
            {
                var str = lines.Select(chars => chars[j]).Aggregate("", (a, b) => $"{a}{b}");
                nums[j] = int.Parse(str);
            }

            ans += op[0] == '+' ? nums.Sum() : nums.Aggregate(1L, (a, b) => a * b);
        }

        return ans.ToString();
    }
}