class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        res = ''

        l = 0
        r = 0

        while l < len(word1) or r < len(word2):
            if l < len(word1):
                res += word1[l]
            if r < len(word2):
                res += word2[r]

            l += 1
            r += 1

        return res
