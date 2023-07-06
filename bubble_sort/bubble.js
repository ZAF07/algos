/*
  Sort an array in ascending order
  Example: Given an array [4,5,3,1,2], the array should be sorted like so: [1,2,3,4,5]
*/

const main = (arr) => {
  // I have to loop (n) times because i have to compare all elements side by side
  for (let i = 0; i < arr.length; i++) {
    //  I cut the length of the arr in each iteration because I know that the last element will always be the biggest element after each iteration
    for (let j = 0; j < arr.length -1 -i; j++) {
      if (arr[j] > arr[j + 1]) {
        const tmp = arr[j];
        arr[j] = arr[j+1]
        arr[j+1] = tmp
      }
    }
  }
  return arr
}


console.log(main([2,8,5,7,4]));
console.log(main([9,8,7,6,5]));
console.log(main([10,90,188,94,5]));