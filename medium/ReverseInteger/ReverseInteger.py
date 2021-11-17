class Solution:
    def reverse(self, x: int) -> int:
        output = 0
        temp = x if x > 0 else x*(-1)

        while temp > 0:
            digit = temp % 10
            temp -= digit
            temp = (temp // 10)
            output = output * 10 + digit

        if x < 0:
            output *= -1
        if abs(output) <= (pow(2, 31) - 1) and abs(output) >= (-1 * pow(2,31)):
            return output
        else:
            return 0
