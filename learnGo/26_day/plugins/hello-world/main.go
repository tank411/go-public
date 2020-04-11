package main

import "github.com/dullgiulio/pingo"

// 创建要导出的对象
type MyPlugin struct{}

// 导出的方法，带有rpc签名
func (p *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "Hello, " + name
	return nil
}

func main() {
	plugin := &MyPlugin{}

	// 注册要导出的对象
	pingo.Register(plugin)
	// 运行主程序
	pingo.Run()
}
