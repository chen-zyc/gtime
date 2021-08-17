# gtime

time package for Go.


## Duration

可生成 JSON 字符串或从 JSON 字符串中解析。

示例：

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/chen-zyc/gtime"
)

type Options struct {
	Timeout gtime.Duration `json:"timeout"`
}

func main() {
	jsonText := `{"timeout":"30s"}`
	opts := &Options{}
	err := json.Unmarshal([]byte(jsonText), opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(opts.Timeout.ToDuration().String()) // 30s

	opts.Timeout = gtime.NewDuration(20 * time.Second)
	jsonData, err := json.Marshal(opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData)) // {"timeout":"20s"}
}
```