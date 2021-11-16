class Solution:
    def addBinary(self, a: str, b: str) -> str:
        i = len(a) - 1
        j = len(b) - 1
        output = ""
        carry_bit = 0
        while i >= 0 or j >= 0:
            ai = int(a[i]) if i >= 0 else 0
            bj = int(b[j]) if j >= 0 else 0
            sum = ai + bj + carry_bit
            if sum == 2:
                carry_bit = 1
                sum = 0
            elif sum == 3:
                carry_bit = 1
                sum = 1
            if not ai and not bj and carry_bit == 1:
                carry_bit = 0
            output = str(sum) + output
            i -= 1
            j -= 1
        if carry_bit:
            output = str(carry_bit) + output

        return output
