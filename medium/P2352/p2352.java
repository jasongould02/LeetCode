import java.util.Arrays;

public class P2352 {
	public int equalPairs(int[][] grid) {
		int n = grid.length;
		int count = 0;
		int[] t = new int[n];
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < n; j++) {
				t[j] = grid[j][i];
			}
			for (int k = 0; k < n; k++) {
				if (Arrays.equals(t, grid[k])) {
					count++;
				}
			}
		}
		return count;
	}
}
