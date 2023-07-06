def merge_sort(arr):
    print("This is the arr in merge_sort: ", arr)
    if len(arr) < 2:
        print("DONE: ", arr)
        return arr

    # get the midpoint
    midpoint = len(arr) // 2

    print("Calling merge now: ", arr[:midpoint], arr[midpoint:])
    return merge(
        # recursively call merge_sort()
        merge_sort(arr[:midpoint]),
        merge_sort(arr[midpoint:]),
    )


def merge(left, right):
    if len(left) == 0:
        return left
    if len(right) == 0:
        return right
    print("merging: ", left, right)
    result = []
    index_left = 0
    index_right = 0

    while len(result) < len(left) + len(right):
        if left[index_left] <= right[index_right]:
            result.append(left[index_left])
            index_left += 1
        else:
            result.append(right[index_right])
            index_right += 1

        # If we reached here, it means that we have reached the end of the right array. Meaning that whatever is left in the right array is the biggest digit because we are always appending the smaller digit before the larger one in the previous step.
        if index_right == len(right):
            # So here we can safely append whatever is remaining in the right array to the result and break the
            result += left[index_left:]
            break

        if index_left == len(left):
            result += right[index_right:]
            break

    print("returning: ", result)
    return result


nums = [4, 3, 2, 1]
print(merge_sort(nums))
