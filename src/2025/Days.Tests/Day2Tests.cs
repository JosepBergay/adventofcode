
namespace Days.Tests;

public class Day2Tests
{
    private readonly Day2 day = new();

    public static IEnumerable<object[]> P1Data() => new List<object[]>()
    {
        new object[]{"11-22", "33"},
        new object[]{"95-115", "99"},
        new object[]{"998-1012", "1010"},
        new object[]{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124", "1227775554"},
    };

    [Theory]
    [MemberData(nameof(P1Data))]
    public void TestDay2_P1(string input, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

}