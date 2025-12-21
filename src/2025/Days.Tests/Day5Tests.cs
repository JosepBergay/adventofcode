
namespace Days.Tests;

public class Day5Tests
{
    private readonly Day5 day = new();

    public static IEnumerable<object[]> Data() => new List<object[]>()
    {
        new object[]{@"3-5
10-14
16-20
12-18

1
5
8
11
17
32
", "3", "14"},
    };

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay5_P1(string input, string expected, string _)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay5_P2(string input, string _, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part2(parsed);

        Assert.Equal(expected, res);
    }
}