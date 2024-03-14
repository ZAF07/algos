/**
 *
 * @param {number} nums
 * @returns {number}
 * @description Takes in an array of integers and returns the sum of the largest continuos subarray
 */
const maxSubarrSum = (nums) => {
  let maxSum = Number.MIN_SAFE_INTEGER;
  let curSum = 0;

  for (let i = 0; i < nums.length; i++) {
    // Capture the current element
    const cur = nums[i];

    // set curSum to the larger between current element or previous curSum + current element
    curSum = Math.max(cur, curSum + cur);

    // set maxSum (result) to the larger between current maxSum or curSum
    maxSum = Math.max(maxSum, curSum);
  }

  return maxSum;
};

console.log("Max Subarr Sum ---> ", maxSubarrSum([-1, 1, 2, 5, -8, 6])); // 8
console.log("Max Subarr Sum ---> ", maxSubarrSum([-1, 1, 2, 5, -8, 9])); // 9
console.log("Max Subarr Sum ---> ", maxSubarrSum([-1, 1, 2, 5 - 5, 6])); // 9
console.log("Max Subarr Sum ---> ", maxSubarrSum([-1, 1, 2, 5, -3, 6])); // 11
