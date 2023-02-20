import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	var compressed string
	lcount := 0
	currletter := string(s[0])

	for _, letter := range s {
		fmt.Println(string(letter))
		if string(letter) == currletter {
			lcount += 1
		} else {
			compressed += (currletter + strconv.Itoa(lcount))
			lcount = 1
			currletter = string(letter)
		}
	}

	return compressed
}

func main() {
	fmt.Println(compressString("aabcccccaaa"))
}
