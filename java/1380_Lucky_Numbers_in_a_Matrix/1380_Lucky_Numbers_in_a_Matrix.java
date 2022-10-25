import java.util.ArrayList;
import java.util.List;

public class P1380 {
	public List<Integer> luckyNumbers(int[][] matrix) {
		List<Integer> result = new ArrayList<Integer>();
		outer: 
		for (int i = 0; i < matrix.length; i++) {
			int min = matrix[i][0];
			int mincol = 0;
			for (int j = 0; j < matrix[0].length; j++) { // row check
				if (matrix[i][j] < min) {
					min = matrix[i][j];
					mincol = j;
				}
			}
			for (int k = 0; k < matrix.length; k++) { // col check
				if (matrix[k][mincol] > min) {
					continue outer;
				}
			}
			result.add(min);
		}
		return result;
	}
}