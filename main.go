package main

import (
	"fmt"
	"getip/pkg"
	"os"
)

func main() {
	ip, err := pkg.GetIp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf(ip)
}
