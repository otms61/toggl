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
