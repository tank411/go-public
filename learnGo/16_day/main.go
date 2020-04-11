package main

import (
	/*"log"*/

	"github.com/dwdcth/consoleEx"
	"github.com/mattn/go-colorable"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//"gitee.com/veni0/robotgo"

func main() {
	test_log2()
} /*
func test_log1() {
	//onte that the original content will be covered!
	fp, err := os.Create("test.log")
	if err != nil {
		log.Fatal("create file faild")
	}
	defer fp.Close()
	//recommand to use flag 18
	aLoger := log.New(fp, "PreFix :", 0)
	for i := 0; i < 33; i++ {
		aLoger.SetFlags(i)
		aLoger.Println("It is style ", i)
	}
}

//create a new file name by time
func create_file_bytime() (*os.File, error) {
	ntime := time.Now()
	filename := ntime.Format("logs_demo 0102#150405.log")
	file, err := os.Create(filename)
	return file, err
}
*/
//record log in different file to avoid cover the old logs file
func test_log2() {
	/* osfile, err := create_file_bytime()
		if err != nil {
			log.Fatal(err)
		}
		loger := log.New(osfile, "", 18)
		loger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
		loger.SetFlags(log.Lshortfile)
		loger.Println("Create log Scuess !!!....")


		 // unix时间戳，如无此行，则是普通日期格式
	    // zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	    logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	    // 打印代码行
	    // logger = logger.With().Caller().Logger()
	    // 带颜色输出
	    // logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		logger.Info().Str("foo", "bar").Msg("hello world") */

	log.Logger = log.With().Caller().Logger()
	log.Info().Msg("hello world")

	out := consoleEx.ConsoleWriterEx{Out: colorable.NewColorableStdout()}
	zerolog.CallerSkipFrameCount = 2 //这里根据实际，另外获取的是Msg调用处的文件路径和行号
	logger := zerolog.New(out).With().Caller().Timestamp().Logger()

	logger.Info().Msg("info")
	logger.Debug().Msg("debug")

}
