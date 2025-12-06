public class DayRunner()
{
    private HttpClient? client;

    private static HttpClient InitializeHttpClient(string session)
    {
        var c = new HttpClient();
        c.DefaultRequestHeaders.Add("Cookie", $"session={session}");
        c.BaseAddress = new Uri("https://adventofcode.com/2025/day/");
        return c;
    }

    public async Task<string> FetchDayInput(int dayNum)
    {
        var session = Environment.GetEnvironmentVariable("SESSION_COOKIE")
            ?? throw new ArgumentException(
                "Session ID must be provided as environment variable at SESSION_COOKIE");

        client ??= InitializeHttpClient(session);

        using var res = await client.GetAsync($"{dayNum}/input");

        res.EnsureSuccessStatusCode();

        return await res.Content.ReadAsStringAsync();
    }


    private async Task<string> GetInputAsync(int dayNum)
    {
        var path = $"src/2025/days/day{dayNum}.txt";
        string input;

        if (!File.Exists(path))
        {
            input = await FetchDayInput(dayNum);

            File.WriteAllText(path, input);
        }
        else
        {
            input = File.ReadAllText(path);
        }

        return input;
    }

    public async Task RunDayAsync(int num, IDay day)
    {
        var input = await GetInputAsync(num);

        var stopWatch = System.Diagnostics.Stopwatch.StartNew();

        var result = day.Exec(input);

        var elapsed = stopWatch.Elapsed;

        Console.WriteLine($"Day {num}: [Part1]: {result.part1} [Part2]: {result.part2} ({elapsed})");
    }
}