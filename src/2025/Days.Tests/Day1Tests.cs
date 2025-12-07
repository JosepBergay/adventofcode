namespace Days.Tests;

public class Day1Tests
{
    private readonly Day1 day = new();

    [Theory]
    [InlineData(@"L68
L30
R48
L5
R60
L55
L1
L99
R14
L82")]
    public void Day1_P1(string input)
    {
        var res = day.Exec(input);

        Assert.Equal("3", res.Part1);
    }

    [Theory]
    [InlineData(@"L68
L30
R48
L5
R60
L55
L1
L99
R14
L82")]
    public void Day1_P2(string input)
    {
        var res = day.Exec(input);

        Assert.Equal("6", res.Part2);
    }
}
