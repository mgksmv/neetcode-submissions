class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_asc = []
        for i, num in enumerate(nums):
            nums_asc.append((num, i))

        nums_asc.sort()

        l = 0
        r = len(nums) - 1

        while l <= r:
            res = nums_asc[l][0] + nums_asc[r][0]

            if res == target:
                return [min(nums_asc[l][1], nums_asc[r][1]), max(nums_asc[l][1], nums_asc[r][1])]

            if res > target:
                r -= 1
            else:
                l += 1

        return []
