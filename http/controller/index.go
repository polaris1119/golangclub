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

type IndexController struct{}

// RegisterRoutes 注册路由
func (i IndexController) RegisterRoutes(e *echo.Echo) {
	e.GET("/", i.index)
	e.GET("/solutions", i.solution)
	e.GET("/learn", i.learn)
	e.GET("/gopher", i.gopher)
	e.GET("/about", i.about)
	e.GET("/solutions/clis/", i.solutionCLIs)
	e.GET("/solutions/webdev/", i.webdev)
	e.GET("/solutions/:typ", i.solutionProxy)
}

// index 首页
func (i IndexController) index(ctx echo.Context) error {
	return Render(ctx, "index.html", nil)
}

// solution 解决方案
func (i IndexController) solution(ctx echo.Context) error {
	return Render(ctx, "solution.html", map[string]interface{}{"solution_active": "Header-menuItem--active"})
}

// learn 学习资源
func (i IndexController) learn(ctx echo.Context) error {
	return Render(ctx, "learn.html", map[string]interface{}{"learn_active": "Header-menuItem--active"})
}

// gopher 名人
func (i IndexController) gopher(ctx echo.Context) error {
	return Render(ctx, "gopher.html", map[string]interface{}{"gopher_active": "Header-menuItem--active"})
}

// about 关于
func (i IndexController) about(ctx echo.Context) error {
	return Render(ctx, "about.html", map[string]interface{}{"about_active": "Header-menuItem--active"})
}

// Designed for today's needs : CLIs (Go for Command-line Interfaces)
func (i IndexController) solutionCLIs(ctx echo.Context) error {
	return Render(ctx, "solutions/clis.html", nil)
}

// solutions Go for Web Development
func (i IndexController) webdev(ctx echo.Context) error {
	return Render(ctx, "solutions/webdev.html", map[string]interface{}{"solution_active": "Header-menuItem--active"})
}

// solutionProxy  代理所有 solutions 子路由
func (i IndexController) solutionProxy(ctx echo.Context) error {
	typ := ctx.Param("typ")
	return Render(ctx, "solutions/"+typ+".html", map[string]interface{}{"solution_active": "Header-menuItem--active"})
}
