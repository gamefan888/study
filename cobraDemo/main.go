/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

cobra框架：命令行框架

1、框架安装： go get -u github.com/spf13/cobra@latest
2、客户端工具安装（cobra-cli）： cobra-cli用于生成代码基本框架。

	安装命令： go install github.com/spf13/cobra-cli@latest

3、客户端工具的使用：

	A、自动生成代码框架：
	命令如下： cobra-cli init -a 作者名 -l 开源协议名   比如：cobra-cli init -a hjj -l MIT
	license文件： 授权协议，主流开源协议的对比：https://www.runoob.com/w3cnote/open-source-license.html
	main.go: 入口文件
	cmd/root.go ：根命令

	B、添加一个子命令：cobra-cli add 子命令名 ，比如：cobra-cli add version
*/
package main

import "cobraDemo/cmd"

func main() {
	cmd.Execute()
}
