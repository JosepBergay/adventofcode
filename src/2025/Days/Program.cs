
var allDays = new Dictionary<int, IDay>{
    { 1, new Day1() },
    { 2, new Day2() },
    { 3, new Day3() },
};

DayRunner runner = new();

List<(int, IDay)> days = new();

foreach (var arg in args)
{
    if (int.TryParse(arg, out var d) && 0 < d && d <= 25 && allDays.TryGetValue(d, out var day))
    {
        days.Add((d, day));
    }
}

if (!days.Any())
{
    days.AddRange(allDays.Select((k, v) => (k.Key, k.Value)));
}

Console.WriteLine($"Running days {String.Join(", ", days.Select(d => d.Item1))}");

var tasks = days
    .Select(d => runner.RunDayAsync(d.Item1, d.Item2));

await Task.WhenAll(tasks);