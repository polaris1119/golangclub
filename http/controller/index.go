/*
 * Copyright (c) 2019. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 * https://golangclub.com
 * Author:polaris	polaris@studygolang.com
 */

package controller

import (
	"github.com/labstack/echo/v4"

	. "github.com/polaris1119/golangclub/http"
)

type IndexController struct {}

func (i IndexController) RegisterRoutes(e *echo.Echo) {
	e.GET("/", i.index)
}

func (i IndexController) index(ctx echo.Context) error {
	return Render(ctx, "index.html", nil)
}