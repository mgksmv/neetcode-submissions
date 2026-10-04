class NumArray:
    def __init__(self, nums: List[int]):
        self.prefixes = []

        total = 0
        for n in nums:
            total += n
            self.prefixes.append(total)

    def sumRange(self, left: int, right: int) -> int:
        right_prefix = self.prefixes[right]
        left_prefix = self.prefixes[left - 1] if left > 0 else 0
        return right_prefix - left_prefix


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)