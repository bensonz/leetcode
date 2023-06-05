

def binarySearch(array, target):
    l = len(array)
    left = 0
    right = len(array) - 1
    while left <= right:
        mid = left + (right - left) / 2
        if array[mid] < target:
            left = mid + 1
        elif array[mid] > target:
            right = mid - 1
        else:
            return mid
    return -1


def firstIdx(array, target):
    occr = binarySearch(array, target)
    print("occr", occr)
    while occr >= 0:
        if array[occr-1] == target:
            occr -= 1
        else:
            break
    return occr
