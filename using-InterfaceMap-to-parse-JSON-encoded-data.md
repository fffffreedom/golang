# 使用interface解析Json数据

## code

```
package main

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func main () {
	f, err := os.Open("./ceph-statics/ceph-df.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

  // 定义map，key类型为string，值类型为interface{}，即可以保存任意类型的数据
	dataJson := make(map[string]interface{})

  // Unmarshal parses the JSON-encoded data
	err = json.Unmarshal([]byte(data), &dataJson)
	if err != nil {
		log.Fatal(err)
	}

  // 打印数据中key为pools的值
	for _, v := range dataJson["pools"].([]interface{}) {
		fmt.Printf("%+v\n", v)
		fmt.Println(v.(map[string]interface{})["name"].(string))
	}
}
```

## JSON-encoded data

```
{
    "stats": {
        "total_bytes": 179945629286400,
        "total_used_bytes": 39811350528,
        "total_avail_bytes": 179905817935872
    },
    "pools": [
        {
            "name": "rbd",
            "id": 0,
            "stats": {
                "kb_used": 14769,
                "bytes_used": 15122606,
                "max_avail": 56955725839252,
                "objects": 25
            }
        },
        {
            "name": "cephfs_data",
            "id": 1,
            "stats": {
                "kb_used": 2,
                "bytes_used": 1247,
                "max_avail": 56955725839252,
                "objects": 1
            }
        },
        {
            "name": "cephfs_metadata",
            "id": 2,
            "stats": {
                "kb_used": 6771,
                "bytes_used": 6933044,
                "max_avail": 56955725839252,
                "objects": 21
            }
        },
        {
            "name": "fiobench",
            "id": 3,
            "stats": {
                "kb_used": 36874893,
                "bytes_used": 37759889424,
                "max_avail": 56955725839252,
                "objects": 9260
            }
        }
    ]
}
```
