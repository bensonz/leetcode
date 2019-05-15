'''
first we sort the clips by its first element

then for each element in the list, get its start and end time as i and j
if end2 >= T, which means we can cover all of it now with the cases we have seen so far, or if i > end2, which means we havea gap now, we break out ofthe look and return immediately

if i is in between first end and the final end, we add this clip to our result (ineffectual here), add 1 to res, which is the number of clips we have so far, and set end to end2.

then in each loop we check whether j can be updated.
Finally we return res if we can cover all, otherwise we return -1.
'''
def videoStitching(clips, T):
      end, end2, res = -1, 0, 0
      for i, j in sorted(clips):
          # print(i,j)
          if end2 >= T or i > end2:
              # print('Finished: {} | i = {}'.format(end2, i))
              break
          elif end < i <= end2:
              res, end = res + 1, end2
              # print('Res: {}'.format(res))
              # print('Found new: {} | {} | {}'.format(end,i,end2))
          end2 = max(end2, j)
          # print('updated : end {} | end2 {}'.format(end, end2))
      return res if end2 >= T else -1

c = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]
print(videoStitching(c, 10) == 3)
print(videoStitching(c, 5) == 2)
print(videoStitching(c, 11) == -1)
