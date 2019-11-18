/*
 * Copyright (c) 2019. The StudyGolang Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 * https://golangclub.com
 * Author:polaris	polaris@studygolang.com
 */

package global

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var once = new(sync.Once)

var (
	config = flag.String("config", "config", "配置文件名称，默认 config")
)

func Init() {
	once.Do(func() {
		if !flag.Parsed() {
			flag.Parse()
		}

		// 随机数种子
		rand.Seed(time.Now().UnixNano())

		// 配置文件名称
		viper.SetConfigName(*config)
		// 配置文件查找路径
		viper.AddConfigPath("/etc/golangclub/")
		viper.AddConfigPath("$HOME/.golangclub")
		viper.AddConfigPath(App.RootDir + "/config")
		// 读取配置文件
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// 填充 global.App 需要的数据
		App.fillOtherField()
	})
}
