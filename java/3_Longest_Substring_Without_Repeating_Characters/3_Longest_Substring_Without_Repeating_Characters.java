class Solution {
    public int lengthOfLongestSubstring(String s) {
		if (s.length() == 1) {
			return 1;
		}
		int max = 0;
        HashSet<Character> set = new HashSet<Character>();
		for (int i = 0; i < s.length() - 1; i++) {
			set.add(s.charAt(i));
			for (int j = i + 1; j < s.length(); j++) {
				if (set.contains(s.charAt(j))) {
					max = (set.size() > max) ? set.size() : max;
					break;
				} else {
					set.add(s.charAt(j));
					max = (set.size() > max) ? set.size() : max;
				}
			}
            set.clear();
		}
		return max;
	}
}