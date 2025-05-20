package main

import (
	"github.com/MyProject273/FlowTasker/internal/cmd"
	"github.com/MyProject273/FlowTasker/internal/storage/postgres"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	db := postgres.GetDatabaseConnection()
	cmd.Main.Run(gctx.GetInitCtx())
	defer db.Close()
}
