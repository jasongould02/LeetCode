class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<Character>();
        for (char c : s.toCharArray()) {
            if (!stack.empty()) {
                if (c == '}' && stack.peek() == '{') {
                    stack.pop();
                    continue;
                } else if (c == ']' && stack.peek() == '[') {
                    stack.pop();
                    continue;
                } else if (c == ')' && stack.peek() == '(') {
                    stack.pop();
                    continue;
                }
            }
            stack.push(c);
        }
        return stack.empty();
    }
}
