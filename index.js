const twoSum = (nums, target) => {
  const seenMap = {};

  for (let i = 0; i < nums.length; i++) {
    if (seenMap[nums[i]] !== undefined) {
      return [seenMap[nums[i]], i];
    }

    seenMap[target - nums[i]] = i;
  }
  return [];
};

const nums = [3, 4, 1, 2, 6, 7];
const target = 7;
const res = twoSum(nums, target);
console.log("result : ", res);
