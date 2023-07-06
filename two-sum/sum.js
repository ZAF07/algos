
/* 
  Find the two numbers that when added together, makes the target value
  Example: Target = 9, Args = [2,4,6,7,8]
  Answer: [0, 1] Indexes of 2 & 7 (2 + 7 = 9)

  NOTES: Hash performance is always faster albeit slightly
*/

// ****** HASH VARIATION ****** 
const main = (arr, target) => {
  
  // Create a map to store the complement:index to the current number that when added together gives the target
  const hash = new Map;

  for (let i = 0; i < arr.length; i++) {
    //  If the current number matches the complement of a previous number, then we know that current number + that number == target, we can return the indexes of both numbers
   if (hash.has(arr[i])) {
    return [hash.get(arr[i]), i]
   } 

  //  Set the complement:index pair of the current number
   hash.set(target - arr[i], i);
  }
  return []
}

const startHash = performance.now()
console.log("Hash variation: ", main([2,4,6,7,8], 9));
console.log("Hash variation: ", main([4,44,65,71,83], 48));
console.log("Hash variation: ", main([1,434,2,9,199], 48));
console.log("Hash variation: ", main([1,434,2,9,199], 201));
console.log("\x1b[33m Hash Performance:  \x1b[0m", performance.measure("Hash : ", startHash).duration);


// ****** ON2 VARIATION ******
const  squareVariation = (arr, target) => {
  // Go through each element in the arr but not including the last element because we dont want to go out of range when comparing with the next element
  for (let i = 0; i < arr.length-1; i++) {
    const curr = arr[i];
    // Start at index 1 because index j is used as the next element to index i. We dont want to compare the same element
    for (let j = 1; j < arr.length; j++) {
      if (curr + arr[j] == target) {
        return [i, j]
      }
    }
  }
  return []
}

const start = performance.now()
console.log("ON2: ", squareVariation([2,4,6,7,8], 9));
console.log("ON2: ", squareVariation([4,44,65,71,83], 48));
console.log("ON2: ", squareVariation([1,434,2,9,199], 48));
console.log("ON2: ", squareVariation([1,434,2,9,199], 201));
console.log("\x1b[33m ON2 Performance:  \x1b[0m", performance.measure("Square: ", start).duration);
