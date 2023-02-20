class Solution {
    public int strStr(String haystack, String needle) {
        int sourceLength = haystack.length();
        int targetLength = needle.length();

        if (targetLength > sourceLength) {
            return -1;
        }

        if (sourceLength == 0) {
            return -1;
        } else if (targetLength == 0) {
            return 0;
        }

        for (int i = 0; i < sourceLength; i++) {
            if (i + targetLength > sourceLength) {
                break;
            }
            if (haystack.charAt(i) == needle.charAt(0)) {
                for (int j = 0; j < targetLength; j++) {
                    if (needle.charAt(j) != haystack.charAt(i+j)) {
                        break;
                    } else if (j == targetLength - 1) {
                        return i;
                    }
                }
            }
        }
        return -1;
    }
}
