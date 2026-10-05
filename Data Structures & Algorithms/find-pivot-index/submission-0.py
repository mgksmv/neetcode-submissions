class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        prefix_sum = []
        total = 0
        for n in nums:
            total += n
            prefix_sum.append(total)

        for i in range(len(nums)):
            left_sum = prefix_sum[i - 1] if i > 0 else 0
            right_sum = prefix_sum[len(nums) - 1] - prefix_sum[i]
            if left_sum == right_sum:
                return i

        return -1
