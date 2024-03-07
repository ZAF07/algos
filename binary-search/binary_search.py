from typing import List

def first_and_last_pos(nums: List[int], target: int) -> List[int]:
    res = [-1,-1]
    
    # Init start and end pointers for binary search
    l, r = 0, len(nums)-1

    while l <= r:
        # Get the mid point
        mid = (l + r) // 2

        cur = nums[mid]

        if cur == target:
            tmp = mid
            # find the start. Keep going left as long as tmp > 0 and left value == target
            while tmp > 0 and nums[tmp-1] == target:
                tmp -= 1
            
            # set start point
            res[0] = tmp

            # reset mid point
            tmp = mid

            # find the end. keep going right as long as tmp <= len(nums)-2 so we dont go out of bounds at +1 step and that right value == target
            while  tmp <= len(nums) -2 and nums[tmp+1] == target:
                tmp += 1
            
            res[1] = tmp
            return res

        if cur < target:
            l = mid + 1
        
        if cur > target:
            r = mid -1
    return res

nums = [3,3,3,3,3]
target = 3
res = first_and_last_pos(nums, target)
print(f'Given: {nums}, With target: {target}. Result: {res}')