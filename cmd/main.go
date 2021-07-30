package main

import (
	"fmt"
	"gbvm"
)

func main() {
    g := &gbvm.GBVM{
        IsRunning: true,
        Program: []int{
             0x1064, 0x11C8, 0x2201, 0x0000,
        },
    }
    g.Run()
}