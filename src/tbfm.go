package main

import (
	"fmt"

	"./lib/bfm"
)

func main() {
	blockFile := bfm.BlockFile{FileName: "./data.dat", BlockSize: 1024}
	pageManager := PageManager{File: &blockFile, PageSize: 1024, NumOfPagesInMemory: 1024}

	blockFile.CreateFile()
	defer blockFile.Fp.Close()

	fmt.Println("all is cool")
}
