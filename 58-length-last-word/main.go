package lengthlastword

import "strings"

/* O(n) space.

Alernatively we could count backwards, starting the count when we find a character that is not " "
Entding the count when we find a character that is " "
this could be done with O(1) space

*/

func lengthOfLastWord(s string) int {
	words := strings.Split(strings.Trim(s, " "), " ")
	return len(words[len(words)-1])
}
