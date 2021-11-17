class Solution:
    def maxArea(self, height: List[int]) -> int:
        length = len(height)
        volume = 0
        i = 0
        j = length-1
        while i != j:
            volume = max(volume, min(height[i], height[j]) * (j-i))
            if height[i] > height[j]:
                j -= 1
            else:
                i += 1
        return volume
