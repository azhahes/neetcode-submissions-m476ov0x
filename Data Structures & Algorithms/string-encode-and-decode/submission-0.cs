public class Solution {

    public string Encode(IList<string> strs) {
        var sb = new StringBuilder();
		foreach (var item in strs)
		{
			string lenString = item.Length.ToString() + "#";
			sb.Append(lenString);
			sb.Append(item);
		}

		return sb.ToString();
	}
    

    public List<string> Decode(string s) {
var list = new List<string>();
		var sb = new StringBuilder();
		int i = 0;
		while (i < s.Length)
		{
			if (s[i] == '#')
			{
				var count = int.Parse(sb.ToString());
				sb.Clear();
				for (int j = i + 1; j < i + count + 1; j++)
				{
					sb.Append(s[j]);
				}

				list.Add(sb.ToString());
				i += count + 1;
				sb.Clear();
			}
			else
			{
				sb.Append(s[i]);
				i++;
			}
		}

		return list;
   }
}
