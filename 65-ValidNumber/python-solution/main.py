class Solution:
  # @param {string} s
  # @return {boolean}
  def isNumber(self, s):
    '''
    I guess i kinda cheated on this? LOL
    '''
    try:
      float(s.strip())
      return True
    except:
      return False
