class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        counts = {}

        for i in range(len(nums)):
            counts[nums[i]] = 1 + counts.get(nums[i], 0)

        return max(counts, key=counts.get)
