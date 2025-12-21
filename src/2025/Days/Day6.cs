using Microsoft.VisualBasic;

public class Day6 : BaseDay<List<List<int>>>
{
    public override List<List<int>> Parse(string input)
    {
        var table = new List<List<int>>();

        var line = new List<int>();

        var str = "";

        for (int i = 0; i < input.Length; i++)
        {
            var c = input[i];

            if (c == ' ' || c == '\n' || i == input.Length - 1)
            {
                if (str.Length > 0)
                {
                    if (int.TryParse(str, out var num))
                    {
                        line.Add(num);
                        str = "";
                    }
                }
            }
            if (c == '\n' || i == input.Length - 1)
            {
                table.Add(line);
                line = new(line.Count);
            }
            else if (c == '+') line.Add(1);
            else if (c == '*') line.Add(2);
            else str += c;
        }

        return table;
    }

    public override string Part1(List<List<int>> parsed)
    {
        var ans = new List<long>();

        var operators = parsed.Last();
        for (int i = 0; i < operators.Count; i++)
        {
            var op = operators[i];

            var nums = parsed.SkipLast(1).Select(line => line[i]);

            ans.Add(op == 1 ? nums.Sum() : nums.Aggregate(1L, (a, b) => a * b));
        }

        return ans.Sum().ToString();
    }

    public override string Part2(List<List<int>> parsed)
    {
        return "";
    }
}
// 16815738482 is too low