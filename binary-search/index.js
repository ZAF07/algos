/**
 Logarithm Time O(logn)
  When the size of the input data decreases in each step by a certain factor, an algorithm will have logarithmic time complexity. This means as the input size grows, the number of operations that need to be executed grows comparatively much slower.

  💡To better understand 0(logn), let’s think of finding a word in a dictionary. If you want to find a word with the letter “p”, you can go through every alphabet and try finding the word, which is linear time O(n). Another route you can take is to open the book to the exact center page. If the word on the center page comes before “p”, you look for the word in the right half. Otherwise, you look in the left half. In this example, you are reducing your input size by half every time, so the number of operations you will need to perform significantly reduces compared to going through every letter. Thus you will have a time complexity of O(log (n))
  
  Binary search algorithm is faster than linear searching because we do not visit EVERY SINGLE element in the array.
  Binary search works at its best only when your array is sorted. 
 */
const nums = [1,2,3,4,5,6,7,8,9,10];
const binarySearch = (arr, target) => {
  // get the left most index which is arr[0]
  let left = 0;
  // Get the right most index which is arr[arr.length -1] (minus 1 because array indexes start at 0 not 1)
  let right = Math.floor((left + arr.length -1))
  // make sure that the left/lower bound is either lesser than or equals to the right/higher bound. This is so that we dont go out of bounds of the array
  while (left <= right) {

    //  Get the middle element. This is the middle element bewteen left and right index
    let mid = Math.floor((left + right) / 2);

    // store the current guess we land on
    let guess = arr[mid];

    // simply return the index position and the value of the target
    if (guess == target) {
      return `Index: ${mid} || value: ${guess}`
    }

    // if our guess is larger than the target, we simply make the right/higher bound index equal to mid -1. Mid minus one because we know that the current mid value is still larger than the target so we don't have to include that element in the next pass
    if (guess > target) {
      right = mid -1
      continue
    }

    // if our guess is lower than the target, we simply make the left/lower bound index equal to mid +1. Mid plus one because we already know that the current mid value is still smaller than the target so we don't have to include that eklement in the next pass
    if (guess < target) {
      left = mid +1;
      continue
    }

    // If nothing was found
    return `Target: ${target} Not found`
  }
}

console.log(binarySearch(nums, 7));