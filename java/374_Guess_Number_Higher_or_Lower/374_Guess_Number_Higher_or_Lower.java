/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * int guess(int num);
 */

public class Solution extends GuessGame {
    public int guessNumber(int n) {
        int high = n;
        int low = 1;
        while(true) {
            int middle = ((high - low) / 2) + low;
            if (guess(middle) == 1) {
                low = middle + 1;
            } else if (guess(middle) == -1) {
                high = middle - 1;
            } else if (guess(middle) == 0) {
                return middle;
            }
        }
    }
}
