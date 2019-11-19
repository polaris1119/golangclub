/*
 * Copyright (c) 2019. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 * https://golangclub.com
 * Author:polaris	polaris@studygolang.com
 */

package controller

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	new(IndexController).RegisterRoutes(e)
	new(RepoController).RegisterRoutes(e)
}
