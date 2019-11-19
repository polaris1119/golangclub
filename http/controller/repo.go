package controller

import (
	"bytes"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/polaris1119/golangclub/global"
)

type RepoController struct{}

func (r RepoController) RegisterRoutes(e *echo.Echo) {
	e.POST("/repo/pull", r.pull)
}

// pull 自动拉去仓库最新代码
func (r RepoController) pull(ctx echo.Context) error {
	secret := "L072uFhwQ6"
	_ = secret

	strCmd := "cd " + global.App.RootDir + "; git pull"
	cmd := exec.Command("sh", "-c", strCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	ctx.Logger().Infoj(log.JSON{"pull_result": out.String()})

	return err
}
