# code
```
package main

import (
    "bytes"
    "fmt"
    "strings"
    "time"
)

var way map[int]string

func benchmarkStringFunction(n int, index int) (d time.Duration) {
    v := "ni shuo wo shi bu shi tai wu liao le a?"
    var s string
    var buf bytes.Buffer

    t0 := time.Now()
    for i := 0; i < n; i++ {
        switch index {
        case 0: // fmt.Sprintf
            s = fmt.Sprintf("%s[%s]", s, v)
        case 1: // string +
            s = s + "[" + v + "]"
        case 2: // strings.Join
            s = strings.Join([]string{s, "[", v, "]"}, "")
        case 3: // stable bytes.Buffer
            buf.WriteString("[")
            buf.WriteString(v)
            buf.WriteString("]")
        }

    }
    d = time.Since(t0)
    if index == 3 {
        s = buf.String()
    }
    fmt.Printf("string len: %d\t", len(s))
    fmt.Printf("time of [%s]=\t %v\n", way[index], d)
    return d
}

func main() {
    way = make(map[int]string, 5)
    way[0] = "fmt.Sprintf"
    way[1] = "+"
    way[2] = "strings.Join"
    way[3] = "bytes.Buffer"

    k := 4
    d := [5]time.Duration{}
    for i := 0; i < k; i++ {
        d[i] = benchmarkStringFunction(10000, i)
    }
}

string len: 410000      time of [fmt.Sprintf]=   426.001476ms
string len: 410000      time of [+]=     307.044147ms                                                              
string len: 410000      time of [strings.Join]=  738.44362ms                                                       
string len: 410000      time of [bytes.Buffer]=  742.248µs   
```
# 结论
- strings.Join 最慢
- fmt.Sprintf 和 string + 差不多
- bytes.Buffer又比上者快约500倍

> https://gocn.io/question/265
