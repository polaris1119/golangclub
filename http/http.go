/*
 * Copyright (c) 2019. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 * https://golangclub.com
 * Author:polaris	polaris@studygolang.com
 */

package http

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/polaris1119/golangclub/global"
)

const (
	LayoutTpl = "common/layout.html"
)

// Render html 输出
func Render(ctx echo.Context, contentTpl string, data map[string]interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}

	contentTpl = LayoutTpl + "," + contentTpl
	htmlFiles := strings.Split(contentTpl, ",")
	for i, contentTpl := range htmlFiles {
		htmlFiles[i] = global.App.TemplateDir + contentTpl
	}
	tpl, err := template.New("layout.html").Funcs(funcMap).
		Funcs(template.FuncMap{"include": tplInclude}).ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	data["path"] = ctx.Path()

	return executeTpl(ctx, tpl, data)
}

func executeTpl(ctx echo.Context, tpl *template.Template, data map[string]interface{}) error {
	// css 和 js 可以每个页面保留一些自己特有的

	// 如果没有定义 css 和 js 模板，则定义之
	if jsTpl := tpl.Lookup("js"); jsTpl == nil {
		tpl.Parse(`{{define "js"}}{{end}}`)
	}
	if cssTpl := tpl.Lookup("css"); cssTpl == nil {
		tpl.Parse(`{{define "css"}}{{end}}`)
	}

	// 如果没有 seo 模板，则定义之
	if seoTpl := tpl.Lookup("seo"); seoTpl == nil {
		tpl.Parse(`{{define "seo"}}
			<meta name="keywords" content="` + global.App.SEO["keywords"] + `">
			<meta name="description" content="` + global.App.SEO["description"] + `">
		{{end}}`)
	}

	global.App.SetUptime()
	// global.App.SetCopyright()

	data["app"] = global.App

	// 记录处理时间
	// data["resp_time"] = time.Since(ctx.Get("req_start_time").(time.Time))

	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, data)
	if err != nil {
		return err
	}

	return ctx.HTML(http.StatusOK, buf.String())
}
