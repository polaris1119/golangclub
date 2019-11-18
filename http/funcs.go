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
	"encoding/json"
	"html/template"
	"math"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"github.com/polaris1119/golangclub/global"
)

// funcMap 自定义模板函数
var funcMap = template.FuncMap{
	// 转为前端显示需要的时间格式
	"formatTime": func(i interface{}) string {
		ctime, ok := i.(string)
		if !ok {
			return ""
		}
		t, _ := time.Parse("2006-01-02 15:04:05", ctime)
		return t.Format(time.RFC3339) + "+08:00"
	},
	"format": func(i interface{}, format string) string {
		switch i.(type) {
		case time.Time:
			return (i.(time.Time)).Format(format)
		case int64:
			val := i.(int64)
			return time.Unix(val, 0).Format(format)
		}

		return ""
	},
	"hasPrefix": func(s, prefix string) bool {
		if strings.HasPrefix(s, prefix) {
			return true
		}
		return false
	},
	"add": func(nums ...interface{}) int {
		total := 0
		for _, num := range nums {
			if n, ok := num.(int); ok {
				total += n
			}
		}
		return total
	},
	"mod": func(num1, num2 int) int {
		if num1 == 0 {
			num1 = rand.Intn(500)
		}

		return num1 % num2
	},
	"divide": func(num1, num2 int) int {
		return int(math.Ceil(float64(num1) / float64(num2)))
	},
	"explode": func(s, sep string) []string {
		return strings.Split(s, sep)
	},
	"noescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"timestamp": func(ts ...time.Time) int64 {
		if len(ts) > 0 {
			return ts[0].Unix()
		}
		return time.Now().Unix()
	},
	"parseJSON": func(str string) map[string]interface{} {
		result := make(map[string]interface{})
		json.Unmarshal([]byte(str), &result)
		return result
	},
	"genList": func(n int, steps ...int) []int {
		step := 1
		if len(steps) > 0 {
			step = steps[0]
		}
		num := int(math.Ceil(float64(n) / float64(step)))
		nums := make([]int, num)
		for i := 0; i < num; i++ {
			nums[i] = i + 1
		}

		return nums
	},
}

// tplInclude 支持 include 模板
func tplInclude(file string, dot map[string]interface{}) template.HTML {
	var buffer = &bytes.Buffer{}
	tpl, err := template.New(filepath.Base(file)).Funcs(funcMap).ParseFiles(global.App.TemplateDir + file)
	if err != nil {
		return ""
	}

	err = tpl.Execute(buffer, dot)
	if err != nil {
		return ""
	}

	return template.HTML(buffer.String())
}
