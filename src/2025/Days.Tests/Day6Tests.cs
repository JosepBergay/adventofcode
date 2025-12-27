
namespace Days.Tests;

public class Day6Tests
{
    private readonly Day6 day = new();

    public static IEnumerable<object[]> Data() => new List<object[]>()
    {
        new object[]{@"12 12
 1 1 
+  + 
", "26", "35"},
        new object[]{@"123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
", "4277556", "3263827"},
    };

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay6_P1(string input, string expected, string _)
    {
        var parsed = day.Parse(input);
        var res = day.Part1(parsed);

        Assert.Equal(expected, res);
    }

    [Theory]
    [MemberData(nameof(Data))]
    public void TestDay6_P2(string input, string _, string expected)
    {
        var parsed = day.Parse(input);
        var res = day.Part2(parsed);

        Assert.Equal(expected, res);
    }
}