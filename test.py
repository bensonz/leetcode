

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
#
# a = [1,2,3,4,5]
# print firstIdx(a, 1)
# print firstIdx(a, 2)
b = [1,1,1,2,2,2,4]
print firstIdx(b, 2)
