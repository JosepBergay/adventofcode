class Day1() : BaseDay<List<(char, int)>>
{
    public override List<(char, int)> Parse(string input)
    {
        var parsed = new List<(char, int)>();

        var reader = new StringReader(input);
        while (true)
        {
            var line = reader.ReadLine();
            if (line is null || line.Length == 0) break;

            var dir = line[0];
            var n = line[1..];
            parsed.Add((dir, int.Parse(n)));
        }

        return parsed;
    }

    public override string Part1(List<(char, int)> parsed)
    {
        var count = 0;
        var max = 100;

        var curr = 50;
        foreach (var (c, i) in parsed)
        {
            if (c == 'L')
            {
                var sub = curr - i;
                curr = (sub + max) % max;
            }
            else
            {
                var added = curr + i;
                curr = added % max;
            }

            if (curr == 0)
            {
                count++;
            }
        }

        return count.ToString();
    }

    public override string Part2(List<(char, int)> parsed)
    {
        return "";
    }
}

