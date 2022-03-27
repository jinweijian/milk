package main

import (
	_ "mlik/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"mlik/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
