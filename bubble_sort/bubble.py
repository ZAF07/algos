# nums = [9, 6, 8, 2, 5, 1]
nums = [9, 5, 2, 6, 1, 8]


def bubble(arr):
    for i in range(len(arr)):
        for j in range(len(arr) - 1 - i):
            if arr[j] > arr[j + 1]:
                tmp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = tmp
    return arr


result = bubble(nums)
print(result)
