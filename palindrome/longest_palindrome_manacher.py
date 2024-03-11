def longest_palindrome(str: str) -> str:
    
    if len(str) < 2:
        return str

    res = ""

    for i in range(len(str)):
        l,r = i, i 

        # Odd palindrome (ABA)
        while l >= 0 and r < len(str) and str[l] == str[r]:
            # calc the len of the current palindrome
            cur_len = len(str[l:r+1])

            # replace res if curLen is larger
            if cur_len > len(res):
                res = str[l:r+1]

            # Move the pointers outwards (left and right)
            l -= 1
            r += 1

        # reset the pointers to be side by side
        l = i
        r = i + 1

        # Even palindrome (ABBA)
        while l >= 0 and r < len(str) and str[l] == str[r]:
            # calc the len of the current palindrome
            cur_len = len(str[l:r+1])

            # replace res if curLen is larger
            if cur_len > len(res):
                res = str[l:r+1]
            
            # Move pointers outwards
            l -= 1
            r += 1
        
    return res

print('Longest palindrome --> ', longest_palindrome('a'))
print('Longest palindrome --> ', longest_palindrome('aa'))
print('Longest palindrome --> ', longest_palindrome('abbac'))
print('Longest palindrome --> ', longest_palindrome('xkabbaxkaabbaa'))