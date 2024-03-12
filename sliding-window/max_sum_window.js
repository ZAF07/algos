const maxSumWindow = (nums, k) => {
  let maxSum = 0;

  // Calcute the value of the first window
  for (let i = 0; i < k; i++) {
    maxSum += nums[i];
  }

  let curSum = maxSum;
  for (let j = k; j < nums.length; j++) {
    curSum += nums[j] - nums[j - k];
    maxSum = Math.max(curSum, maxSum);
  }
  return maxSum;
};

console.log("Max value of window --> ", maxSumWindow([2, 2, 3, 1, 1, 5], 4));
