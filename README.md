# timeutil

> Utilities around time with respect of `TZ` environment variable

## Install

```shell
go get github.com/avakarev/go-timeutil
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/avakarev/go-timeutil"
)

func main() {
	// `TZ` env var is set to "Europe/Berlin"
	t, _ := time.Parse(time.RFC3339, "2022-06-03T16:26:15Z")
	fmt.Println(t.Format(time.RFC3339))                 // => 2022-06-03T16:26:15Z
	fmt.Println(timeutil.Local(t).Format(time.RFC3339)) // => 2022-06-03T18:26:15+02:00
}
```


## License

`go-logutil` is licensed under MIT license. (see [LICENSE](./LICENSE))
