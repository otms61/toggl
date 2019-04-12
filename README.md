# Toggl API in Go

[![Build Status](https://travis-ci.com/otms61/toggl.svg?branch=master)](https://travis-ci.com/otms61/toggl)

This library partially supports the `Toggl API v8` REST calls.

## Installing

```console
go get -u github.com/otms61/toggl
```

## Examples

### Getting Running TimeEntry

```golang
package main

import (
	"context"
	"fmt"

	"github.com/otms61/toggl"
)

func main() {
	api := toggl.New("API KEY")

	r, err := api.GetRunningTimeEntry(context.Background())
	fmt.Println(r)
	if err != nil {
		fmt.Printf("Err %s\n", err)
		return
	}
}

```

## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.

## License

MIT License
