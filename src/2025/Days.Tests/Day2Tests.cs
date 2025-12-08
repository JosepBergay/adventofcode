
namespace Days.Tests;

public class Day2Tests
{
    private readonly Day2 day = new();

    public static IEnumerable<object[]> Data() => new List<object[]>()
    {
        new object[]{"11-22", "33", "33"},
        new object[]{"95-115", "99", "210"},
        new object[]{"998-1012", "1010", "2009"},
        new object[]{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124", "1227775554", "4174379265"},
    };

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay2_P1(string input, string expected, string _)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay2_P2(string input, string _, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part2(parsed);

        Assert.Equal(expected, res);
    }

}