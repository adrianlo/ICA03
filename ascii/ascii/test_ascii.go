package asciiTest

import "fmt"

var istrue bool = false
var s1 string = "hei"

func testAscii(s1 string) {
	for i := 0; i < len(s1); i++ {
		//fmt.Printf(" %+q \n", ascii[i])
		//s2 := isASCII()

		if isASCII(s1) == true {
			fmt.Print("har bare ascii?", istrue)
		}
	}
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			istrue = false
			return false
		}
	}
	istrue = true
	return true
}
