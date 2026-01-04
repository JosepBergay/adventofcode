
namespace Days.Tests;

public class Day9Tests
{
    private readonly Day9 day = new();

    public static IEnumerable<object[]> Data() => new List<object[]>()
    {
        new object[]{@"7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
", "50", "24"},
    };

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay9_P1(string input, string expected, string _)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay9_P2(string input, string _, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part2(parsed);

        Assert.Equal(expected, res);
    }
}