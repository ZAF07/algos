/*
Given a string s, find the length of the longest
substring
 without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

//   lr
// a a c b d a b i
const s = "aacbdabi";

const longestSubstringWithoutRepeating = (s) => {
  let result = 0;
  const seenMap = {};
  let l = 0;
  let r = 0;
  //   As long as the left side of the window does not go past the right and the right does not exceed the length of the array we continue looping
  while (r < s.length && l <= r) {
    // If we have seen the letter before, we first check if the location of the seen letter is at or after the left position
    if (seenMap[s[r]] != undefined) {
      const location = seenMap[s[r]];
      //   If the last seen letter is at the current left pointer (means the right and the left pointers has the same value), or somewhere after the left pointer (means that the right pointer shares the same value with one letter in the substring) we shift the left pointer one letter after the last seen.
      //   However, if the last seen letter was seen somewhere BEFORE the left pointer (meaning that the last seen letter IS NOT in the current l-r window)
      if (location >= l) {
        l = location + 1;
      }
    }

    const cur = s.slice(l, r + 1);
    if (cur.length > result) {
      result = cur.length;
    }
    // Keep updating the seenMap to show the latest poosition of the last seen letter
    seenMap[s[r]] = r;
    r++;
  }
  return result;
};

// Set method (No good. Does no achieve the desired outcome. Will redo when time permits)
const setLongestNoRepeating = (s) => {
  const noRepeat = new Set();

  for (let i = 0; i < s.length; i++) {
    noRepeat.add(s[i]);
  }
  console.log("SET: ", noRepeat);

  return noRepeat.size;
};

const res = longestSubstringWithoutRepeating(s);
console.log("result --> ", res);

// const res1 = setLongestNoRepeating(s);
// console.log("set: ", res1);
