public static class Extensions
{
    extension(string str)
    {
        public int WordCount() =>
            str.Split([' ', '.', '?'], StringSplitOptions.RemoveEmptyEntries).Length;

        public bool IsPallindrome()
        {
            for (int i = 0; i < (str.Length / 2); i++)
            {
                if (str[i] != str[str.Length - i - 1]) return false;
            }

            return true;
        }
    }

    extension(long l)
    {
        /**
         * Returns an IEnumerable<long> matching the range [l, end]
         */
        public IEnumerable<long> ToRange(long end)
        {
            var curr = l;
            while (curr <= end)
            {
                yield return curr;
                curr++;
            }
        }
    }

    extension(int i)
    {
        public IEnumerable<int> Factors()
        {
            var curr = 1;
            var sqrt = Math.Sqrt(i);
            while (curr <= sqrt)
            {
                if (i % curr == 0)
                {
                    yield return curr;

                    var div = i / curr;
                    if (div != curr)
                    {
                        yield return div;
                    }
                }
                curr++;
            }
        }
    }
}