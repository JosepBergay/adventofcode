
namespace Days.Tests;

public class Day7Tests
{
    private readonly Day7 day = new();

    public static IEnumerable<object[]> Data() => new List<object[]>()
    {
        new object[]{@".......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
", "21", ""},
    };

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay7_P1(string input, string expected, string _)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay7_P2(string input, string _, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part2(parsed);

        Assert.Equal(expected, res);
    }
}