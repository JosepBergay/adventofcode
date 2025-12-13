using Microsoft.VisualBasic;

public class Day3 : BaseDay<string>
{
    private readonly int zero = '0';

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
        long total = 0;

        var digits = new int[12];
        long currBest = 0;
        foreach (var c in parsed)
        {
            if (c == '\n')
            {
                total += currBest;
                digits = new int[12];
                currBest = 0;
                continue;
            }

            var replaceIdx = digits.Length - 1;
            for (int i = 0; i < digits.Length - 1; i++)
            {
                if (digits[i] < digits[i + 1])
                {
                    replaceIdx = i;
                    break;
                }
            }


            long candidate = 0;
            for (int i = 0; i < digits.Length; i++)
            {
                if (i == replaceIdx) continue;

                var d = digits[i];
                if (i < replaceIdx) candidate += (long)(d * Math.Pow(10, 11 - i));
                else
                    // i > minIdx
                    candidate += (long)(d * Math.Pow(10, 12 - i)); // Shift digit left
            }

            var n = c - zero;
            candidate += n; // Always add current char at the end

            if (candidate > currBest)
            {
                for (int i = 0; i < digits.Length; i++)
                {
                    if (i >= replaceIdx)
                    {
                        if (i == digits.Length - 1)
                        {
                            digits[i] = n;
                        }
                        else
                        {
                            digits[i] = digits[i + 1];
                        }
                    }
                }

                currBest = candidate;
            }
        }

        // In case input does not end with '\n'.
        total += currBest;

        return total.ToString();
    }
}
