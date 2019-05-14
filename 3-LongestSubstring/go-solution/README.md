so my solution is super slow, but it's the most basic idea.

Here's a super fast code:
```
func lengthOfLongestSubstring(s string) int {
    max := 0
    begin := 0
    hash := make(map[rune]int)
    for i, c := range s {
        index, ok := hash[c]
        length := i - begin + 1
        if ok && index >= begin {
            length = i - index
            begin = index + 1
        }
        hash[c] = i
        if length > max {
            max = length
        }
    }

    return max
}
```

Imma just reason through this code.

First he has a max, a begin set to 0, which I assume would be the same as my cur and longestSubstring and a hash (map), which is also what I used.

Then he loops through the string as I did, with i as index and c as the current value. Then he checks if the character is seen in this hash already, then he set the length to (i - begin + 1)

Then an if statement that says if ok, which means already seen, and index >= begin, which means the value of c in hash is greater than begin, he'd set length to i - index, along with a set for begin to index + 1.

Then he updates the character's value in hash.
