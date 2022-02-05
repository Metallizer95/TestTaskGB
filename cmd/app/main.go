package main

import (
	"fmt"
	"github.com/Metallizer95/TestTaskGB/internal/usecases"
)

func main() {
	ucs := usecases.New()
	res, _ := ucs.FindMaxBalanceWalletForLastBlocks(20)
	fmt.Println(res)
}
