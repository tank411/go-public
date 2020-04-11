package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/transport/tcp"
)

func server(port string) {

	// 子连接的流水线配置
	var childPipelineInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			// 最大允许包长128字节，使用\n分割包, 丢弃分隔符
			AddLast(frame.LengthFieldCodec(binary.LittleEndian, 1024, 0, 2, 0, 2)).
			// 解包出来的bytes转换为字符串
			AddLast(format.TextCodec()).
			// 日志处理器, 打印连接建立断开消息，收到的消息
			AddLast(LoggerHandler{}).
			AddLast(EchoHandler{"Server"}).
			// 业务处理器 (将字符串全部大写)
			AddLast(UpperHandler{})
	}

	// 配置服务器
	netty.NewBootstrap().
		// 配置子链接的流水线配置
		ChildInitializer(childPipelineInitializer).
		// 配置传输使用的方式
		Transport(tcp.New()).
		// 配置监听地址
		Listen("0.0.0.0:" + port).
		// 等待退出信号
		Action(netty.WaitSignal(os.Interrupt)).
		// 打印退出消息
		Action(func(bs netty.Bootstrap) {
			fmt.Println("server exited")
		})
}

type LoggerHandler struct{}

func (LoggerHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Println("go-netty:", "->", "active:", ctx.Channel().RemoteAddr())
	// 写入欢迎信息
	ctx.Write("欢迎你的到来！")
}

func (LoggerHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	//fmt.Println("接收信息:", message)
	// 交给下一个处理器处理(按照处理器的注册顺序, 此例下一个处理器应该是UpperHandler)
	ctx.HandleRead(message)
}

func (LoggerHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Println("go-netty:", "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	// 连接断开了，默认处理是关闭连接
	ctx.HandleInactive(ex)
}

type UpperHandler struct{}

func (UpperHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	name := ""
	prompt := &survey.Input{
		Message: "回复内容:",
	}
	survey.AskOne(prompt, &name, survey.WithValidator(survey.Required))
	// 写入返回结果给客户端
	ctx.Write(name)
}
