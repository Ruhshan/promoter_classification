package main

import (
	"biogo"
	"fmt"
	"strings"
	"strconv"
)

func makegcskewdata(window int, step int, maxlen int){
	parsed, _ := biogo.ParseFasta("ctv_total.txt")
	gcskews, _ := biogo.GCSkewSlidingWindow(parsed, window, step, maxlen)
	count:= 0
	for _, g := range gcskews {
		results := strings.Join(g.Values, ";")
		fmt.Println("ctv_"+strconv.Itoa(count)+";"+results + ";1")
		count++
	}

	count = 0
	parsed, _ = biogo.ParseFasta("si_total.txt")
	gcskews, _ = biogo.GCSkewSlidingWindow(parsed, window, step, maxlen)

	for _, g := range gcskews {
		results := strings.Join(g.Values, ";")
		fmt.Println("si_"+strconv.Itoa(count)+";"+results + ";0")
		count++
	}
}


func makebasestackingdata(window int, step int, maxlen int){
	parsed, _ := biogo.ParseFasta("ctv_total.txt")
	bs_engs, _ := biogo.BaseStackingSlidingWindow(parsed, window, step, maxlen)
	count:= 0
	for _, g := range bs_engs {
		results := strings.Join(g.Values, ";")
		fmt.Println("ctv_"+strconv.Itoa(count)+";"+results + ";1")
		count++
	}

	count = 0
	parsed, _ = biogo.ParseFasta("si_total.txt")
	bs_engs, _ = biogo.BaseStackingSlidingWindow(parsed, window, step, maxlen)

	for _, g := range bs_engs {
		results := strings.Join(g.Values, ";")
		fmt.Println("si_"+strconv.Itoa(count)+";"+results + ";0")
		count++
	}
}
func main() {
	//makegcskewdata(100, 5, 1000)

	makebasestackingdata(200,20,1000)

}
