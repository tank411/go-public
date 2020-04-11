package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/transport/tcp"
	"github.com/go-netty/go-netty/utils"
)

func client(ip, port string) {

	// 配置客户端
	var bootstrap = netty.NewBootstrap()

	// 客户端流水线配置
	bootstrap.ClientInitializer(func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.LengthFieldCodec(binary.LittleEndian, 1024, 0, 2, 0, 2)).
			AddLast(format.TextCodec()).
			AddLast(EchoHandler{"Client"})
	})

	// connect to the server after 1 second
	time.AfterFunc(time.Second, func() {
		_, err := bootstrap.Connect(ip+":"+port, nil)
		fmt.Println(err)
		utils.Assert(err)
	})

	// setup bootstrap & startup server.
	bootstrap.
		Transport(tcp.New()).
		// 等待退出信号
		Action(netty.WaitSignal(os.Interrupt))
}

type EchoHandler struct {
	role string
}

func (l EchoHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Println(l.role, "->", "active:", ctx.Channel().RemoteAddr())
	//ctx.Write("Hello I'm " + l.role)
}

func (l EchoHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fmt.Println("接收信息:", message)
	name := ""
	prompt := &survey.Input{
		Message: "回复内容:",
	}
	survey.AskOne(prompt, &name, survey.WithValidator(survey.Required))
	// 写入返回结果给服务端
	ctx.Write(name)
}

func (l EchoHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Println(l.role, "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	ctx.HandleInactive(ex)
}
