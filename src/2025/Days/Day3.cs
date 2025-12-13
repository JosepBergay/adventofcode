using Microsoft.VisualBasic;

public class Day3 : BaseDay<string>
{
    public override string Parse(string input)
    {
        return input;
    }

    public override string Part1(string parsed)
    {
        var reader = new StringReader(parsed);
        var total = 0u;
        (uint, uint) curr = ('0', '0');

        while (true)
        {
            var c = reader.Read();

            if (c == '\n' || c == -1)
            {
                total += (curr.Item1 - '0') * 10 + curr.Item2 - '0';
                curr = ('0', '0');

                if (c == -1) break;
                continue;
            }

            var num = Math.Max(curr.Item1, curr.Item2);

            var candidate = num * 10 + c;
            var currInt = curr.Item1 * 10 + curr.Item2;

            if (candidate > currInt)
            {
                curr = (num, (uint)c);
            }
        }

        return total.ToString();
    }

    public override string Part2(string parsed)
    {
        return "";
    }
}
