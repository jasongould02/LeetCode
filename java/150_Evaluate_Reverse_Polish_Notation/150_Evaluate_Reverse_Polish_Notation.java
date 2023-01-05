class Solution {
    public int evalRPN(String[] tokens) {
		// Slow solution
        ArrayList<String> ops = new ArrayList<String>();
        ops.add("/");
        ops.add("*");
        ops.add("-");
        ops.add("+");
        Stack<Integer> stack = new Stack<Integer>();
        for (String s : tokens) {
            if (ops.contains(s) && !stack.isEmpty()) {
                int v2 = stack.pop();
                int v1 = stack.pop();
                if (s.equals("/")) {
                    stack.push(v1 / v2);
                } else if (s.equals("*")) {
                    stack.push(v1 * v2);
                } else if (s.equals("-")) {
                    stack.push(v1 - v2);
                } else if (s.equals("+")) {
                    stack.push(v1 + v2);
                }
            } else {
                stack.push(Integer.parseInt(s));
            }
        }
        return stack.pop();
    }
}
