/*
 * @Descripttion: 模块信息介绍
 * @Company: HIK
 * @version: 1.0.0
 * @Author: TK
 * @Date: 2020-01-09 19:28:41
 * @LastEditors  : TK
 * @LastEditTime : 2020-01-16 14:40:45
 */
package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mattn/go-colorable"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	out := zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat: time.RFC3339}
	//out := consoleEx.ConsoleWriterEx{Out: colorable.NewColorableStdout()}
	zerolog.CallerSkipFrameCount = 2
	out.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%-6s", i))
	}
	out.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("- %s", i)
	}
	out.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	out.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	logger = zerolog.New(out).With().Caller().Timestamp().Logger()
	logger.Info().Str("foo", "bar").Msg("Hello World")
}

var dst chan int

func read1(ctx context.Context) {
	for data := range dst {
		// 打印通道数据
		fmt.Println(data)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}

}
func write1(ctx context.Context) {
	i := 0
	for {
		i++
		dst <- i
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {

	logger.Info().Msg("main函数开始")
	dst = make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go write1(ctx)
	//	go read1(ctx)

	time.Sleep(time.Second * 1)
	//close(dst)
	cancel()

	logger.Info().Msg("main函数结束")
	/* for {

	} */
}

/*
取消任务，可以使用上下文（context），主要用来通知下游业务，上游已经取消，或者超时取消。可以使用通道，但是通道不是协成安全的。所以应该优先使用context
通道优雅的关闭，（主要在多次关闭同一通道问题上）可以使用锁进行管理，对通道进行封装，统一关闭。
协成间通信，可以多使用通道，减少共享内存方式，使用通信进行共享内存。
*/
