
namespace Days.Tests;

public class Day3Tests
{
    private readonly Day3 day = new();

    public static IEnumerable<object[]> Data() => new List<object[]>()
    {
        new object[]{@"987654321111111", "98", "987654321111"},
        new object[]{@"811111111111119", "89", "811111111119"},
        new object[]{@"234234234234278", "78", "434234234278"},
        new object[]{@"818181911112111", "92", "888911112111"},
        new object[]{@"987654321111111
811111111111119
234234234234278
818181911112111
", "357", "3121910778619"},
    };

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay3_P1(string input, string expected, string _)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay3_P2(string input, string _, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part2(parsed);

        Assert.Equal(expected, res);
    }
}