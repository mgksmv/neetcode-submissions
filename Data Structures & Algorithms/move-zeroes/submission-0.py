class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        for i in range(len(nums) - 1):
            j = i + 1

            if nums[i] == 0:
                while j < len(nums) - 1 and nums[j] == 0:
                    j += 1

                nums[i], nums[j] = nums[j], nums[i]
