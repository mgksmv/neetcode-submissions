class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_len = 0
        visited = set()
        l = 0

        for r in range(len(s)):
            while s[r] in visited:
                visited.remove(s[l])
                l += 1
            visited.add(s[r])
            max_len = max(max_len, r - l + 1)

        return max_len
