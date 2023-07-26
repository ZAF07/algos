// SAME AS TwoSum BUT THIS ONE EPECTS AN ARRAY OF THE NUMBERS THAT MAKE UP THE TARGET VALUE

// const nums = [5,1,9,5,4,6,5,10,0,5]; // 10: [9,1][5,5] [6,4] [10,0]
// const target = 10;

const nums = [3, 4, 3, 4];
const target = 7;

// const nums = [3,3,3,3]; // [3,3]
// const target = 6;

// const nums = [3,3,3,4,2,6,0,3];
// const target = 6;

// const nums = [10, 0];
// const target = 10;

// const nums = [3,3,3,3];
// const target = 6;

// const nums = [11, -1, 9, 9, 1, 1, 5, 5, 6, 10, 0, 15, -5, 8, 2, 2];
// const target = 10;

// const nums = [];
// const nums = [10];
// const target = 10;

// does not work with same number pair
const findPairsWithSum = (nums, target) => {
  const result = [];
  const missingMap = {};

  for (let i = 0; i < nums.length; i++) {
    const currentPair = nums[i];
    const missingPair = target - currentPair;

    // If current is seen but it's pair is not seen
    if (missingMap[currentPair] && !missingMap[missingPair]) {
      result.push([currentPair, missingPair]);
    }

    missingMap[missingPair] = true;
  }
  return result;
};

// works with same number pair
const findPairsWithSum2 = (nums, target) => {
  result = [];
  missingMap = {};

  for (let i = 0; i < nums.length; i++) {
    const currentPair = nums[i];
    const missingPair = target - currentPair;

    if (
      missingMap[currentPair] !== undefined &&
      missingMap[currentPair] !== -1
    ) {
      // if (missingMap[currentPair] !== -1) {
      // result.push([currentPair, nums[missingMap[currentPair]]])
      result.push([currentPair, missingPair]);
      missingMap[currentPair] = -1;
      // }
      continue;
    }

    // Necessary for 3,4,3,4
    if (!missingMap[missingPair]) {
      missingMap[missingPair] = i;
    }
  }
  return result;
};

// works with both cases above
const findPairsWithSum3 = (nums, target) => {
  if (nums.length < 2) {
    return [];
  }

  const result = [];
  const missingMap = {};

  for (let i = 0; i < nums.length; i++) {
    const currentPair = nums[i];
    const missingPair = target - currentPair;

    // If both pairs were used before ...
    if (missingMap[currentPair] == false && missingMap[missingPair] == false) {
      continue;
    }

    // If the current pair was seen as a missing pair before
    if (missingMap[currentPair] == true) {
      result.push([currentPair, missingPair]);
      missingMap[currentPair] = false;
      missingMap[missingPair] = false;
    } else {
      missingMap[missingPair] = true;
    }
  }
  return result;
};

console.log(
  "Does not work with same number pair: ",
  findPairsWithSum(nums, target)
);
console.log("Works with same number pair: ", findPairsWithSum2(nums, target));
console.log("Should work with both", findPairsWithSum3(nums, target));
