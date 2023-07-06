/* Beautiful String Problem
  A string is said to be beautiful if b occurs in it no more times than a; c occurs in it no more times than b; etc
 */
function solution(inputString) {
  console.log('Input string is : ', inputString);
  
  const count = {};
  const unicodeOfA = 'a'.charCodeAt(0);

  // Sort the current string in order
  const sortedString = inputString.split('').sort().join(''); // O(n)

  // Store the current letter's unicode & previous letter
  let currUnicode = 0;
  let prevLetter = sortedString[0].charCodeAt(0);

  for (let j = 0; j < inputString.length; j ++) { // O(n)
    // Convert current letter to its unicode value
    const currLetter = sortedString[j].charCodeAt(0);

    // Because the string is already sorted, we first check if the first letter is the letter a(97). If it is not, we can safely return false
    if (j == 0) {
      if (currLetter != unicodeOfA) {
        console.log('First letter is not "a". Returning false');
        return false;
      }
      count[currLetter] = 1;
      continue;
    }

    // We add 1 to the count map if this instance of the letter is not yet captured in the count map
    if (!count[currLetter]) {
      // Check if previous letter comes right before current letter in the alphabet list
      if (!count[currLetter-1]) {
        console.log(`Current letter "${sortedString[j]}" has no direct preceeding letter. This is not a beautiful string ❌`);
        return false;
      }
      count[currLetter] = 1;
      currUnicode = currLetter;
      continue;
    } 
    
    count[currLetter] += 1;
    // Next we check if the currUnicode is equal to the currLetter (meaning that this current letter is a subsequent occurance of the same letter)
    if (currUnicode == currLetter) {

      //  We can safely continue on to the next iteration if the current letter is still a(97) as this is the start of the alphabet list
      if (currLetter == unicodeOfA) {
        continue;
      }
      // console.log('current letter : ', currLetter, count[currLetter] );
      // console.log('previous letter : ', prevLetter, count[prevLetter] );
      // console.log('check : ', count[currLetter] > count[prevLetter]);
      // console.log('count: ', count);

      // We check if this occurance of the letter is already larger than the previous letter in the count map (This current letter is NOT 'a' so we have to run these checks)
      if (count[currLetter] > count[prevLetter]) {
        console.log(`Occurance of current letter "${sortedString[j]}" is already larger than preceeding letter. This is not a beautiful string ❌`);
        return false;
      }
    }
  }
  console.log('This is a beautiful string!! 🔥');
  return true;  // String is beautiful
}

// Example usage:
var inputString1 = process.argv[2];
console.log('CLI VERSION: ', solution(inputString1));

var inputString1 = "aabbccdde";
console.log(solution(inputString1));  // Output: true

var inputString2 = "aaabbbcccdddeeegfhijkl";
console.log(solution(inputString2));  // Output: false

var inputString3 = "aabbb";
console.log(solution(inputString3));  // Output: false

var inputString4 = "aaabbbcccl";
console.log(solution(inputString4));  // Output: false
