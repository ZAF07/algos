from typing import List
def maxSubarraySum(nums: List[int]) -> int:
    maxSum = float('-inf')
    curSum = 0

    for n in nums:
        curSum = max(n, curSum+n)
        maxSum = max(curSum, maxSum)
    
    return maxSum

print('Kadanes algo --> ', maxSubarraySum([-1,2,4,-1,-1,5,5])) # 14
print('Kadanes algo --> ', maxSubarraySum([2,5,-1,-1,-1,-1,-1,-1,-1,-1,2,3])) # 7
print('Kadanes algo --> ', maxSubarraySum([2,2,-1,-1,-1,-1,-1,-1,-1,-1,2,3])) # 5
