const s = "bababx";

const longestPalindrome = (s) => {
  let result = "";

  for (let i = 0; i < s.length; i++) {
    let l = i,
      r = i;

    while (l >= 0 && r < s.length && s[l] === s[r]) {
      const currentPalindrome = s.slice(l, r + 1);
      // check if current palindrome length is larger than the one in result (if any)
      if (currentPalindrome.length > result.length) {
        result = currentPalindrome;
      }
      l--;
      r++;
    }

    //    Check for even lettered palidnrome
    while (l >= 0 && r < s.length && s[l] === s[r]) {
      const currentPalindrome = s.slice(l, r + 1);
      if (currentPalindrome.length > result.length) {
        result = currentPalindrome;
      }

      l--;
      r++;
    }
  }
  return result;
};

const res = longestPalindrome(s);
console.log("result: ", res);
