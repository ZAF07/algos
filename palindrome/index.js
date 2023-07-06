const nums = [1,2,1];
const num = 22;

/**
 This takes O(n log n) time (better than O(n))
 O(nlogn) because i am only iterating up until the middle of the array only. Halving the input length
 */
const checkPalindromeArray = (arr) => {
  let result = true;
  for (let i = 0; i < arr.length; i++) {
    
    let rightIndex = arr.length - 1; 
    const left = arr[i];
    const right = arr[rightIndex - i]; // get the right side equilvalant position of the given array

    // make sure that i does not go past the middle element, converging with rightIndex (which has already been compared to with i)
    // If this statement is true, that means that we have reached the middle of the array. We can safely return true here because we have checked that left and right are equal in the previous passes
    if (i == rightIndex - i) {
      console.log(`We have reached the middle of the array:\n i: ${i} || rightIndex - i: ${rightIndex-i}`);
      return result
    }
    
    if (left !== right) { // Check if left and right elements are equal
      console.log('Not a palindrome');
      result = false
      return result
    }
  }
  return result
}

const checkPalindromeInt = (int) => {
  // if given int is 0 or negative number or is a number that is a decimal of 2 (example: 10,20,30,100,1000 etc...) the int cannot be a palindrome so return false
  if (int < 0 || (int % 10 == 0 && int != 0)) return false

  let right = 0;
  let left = int;
  while (left > right) {
    right = Math.floor(right * 10 + left % 10);
    left = Math.floor(left / 10);
  }
  return left == Math.floor(right/10) || right == left
}

console.log(`Array version: ${nums}:\n`, checkPalindromeArray(nums));
console.log(`Integer version: ${num}:\n`, checkPalindromeInt(num));