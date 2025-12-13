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
        var path = $"src/2025/Days/day{dayNum}.txt";
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

    private async Task<(DayResult, TimeSpan)> RunDayAsync(string input, IDay day)
    {
        var stopWatch = System.Diagnostics.Stopwatch.StartNew();

        var result = day.Exec(input);

        var elapsed = stopWatch.Elapsed;

        return (result, elapsed);
    }

    public async Task RunDayAsync(int num, IDay day)
    {
        var input = await GetInputAsync(num);

        var (result, elapsed) = await RunDayAsync(input, day);

        Console.WriteLine($"Day {num}: [Part1]: {result.Part1} [Part2]: {result.Part2} ({elapsed})");
    }

    public async Task RunDayWithDiagnosticsAsync(int num, IDay day)
    {
        var input = await GetInputAsync(num);

        var tasks = Enumerable.Range(0, 100)
            .Select(_ => RunDayAsync(input, day));

        var results = await Task.WhenAll(tasks);

        var (result, _) = results.First();

        Console.WriteLine($"Day {num}: [Part1]: {result.Part1} [Part2]: {result.Part2}");

        var sorted = results.Select(it => it.Item2).ToList();
        sorted.Sort();

        var mean = sorted.Aggregate((a, b) => a.Add(b)) / sorted.Count;

        Console.WriteLine($"#1   -> {sorted.First()}");
        Console.WriteLine($"25%  -> {sorted[sorted.Count / 4]}");
        Console.WriteLine($"50%  -> {sorted[sorted.Count / 2]}");
        Console.WriteLine($"75%  -> {sorted[sorted.Count * 3 / 4]}");
        Console.WriteLine($"last -> {sorted.Last()}");
        Console.WriteLine($"~avg -> {mean}");
    }
}