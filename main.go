package main

import (
	"fmt"
	libs "project/libs"
)

func main() {
	user := libs.GetUsername()
    path := libs.GeneratePaths(user)
    filename := []string{"History","places.sqlite","History.db"}
    err := libs.FindAndCopyFile(path, filename)
    if err != nil {
        fmt.Println(err)
    }
}
