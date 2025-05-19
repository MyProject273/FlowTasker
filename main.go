package main

import (
	_ "FlowTasker/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"FlowTasker/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
