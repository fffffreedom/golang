package main

import (
	"regexp"
	"fmt"
	"strings"
	"strconv"
)

func main () {
	str := `size 1024 MB in 2560 objects`
	reg := regexp.MustCompile(`size (\d+) ([kKmMgG][bB])`)

	matched := reg.FindStringSubmatch(str)
	fmt.Println(len(matched), matched)

	diff := `Offset       Length  Type
0            2162688 data
4194304      3178496 data
20971520     4194304 data
25165824     32768   data
37748736     3178496 data
41943040     3178496 data
54525952     4194304 data
58720256     32768   data
67108864     4194304 data
71303168     1081344 data
75497472     3178496 data
`
	diff = strings.TrimSuffix(diff, "\n")
	diffSlice := strings.Split(diff, "\n")

	reg = regexp.MustCompile(`(\d+)(\s+)(\d+)(\s+)data`)

	var sum uint64
	sum = 0

	for _, str := range diffSlice {
		matched = reg.FindStringSubmatch(str)
		fmt.Println(len(matched))
		if len(matched) == 0 {
			continue
		}
		fmt.Println(matched[3])
		size, err := strconv.Atoi(matched[3])
		if err != nil {
			fmt.Errorf("==> %v", err)
			continue
		}

		sum += uint64(size)
	}

	fmt.Println(sum)
}
