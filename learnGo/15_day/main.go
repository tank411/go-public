/*
 * @Descripttion: 模块信息介绍
 * @Company: HIK
 * @version: 1.0.0
 * @Author: TK
 * @Date: 2019-12-06 11:34:16
 * @LastEditors: TK
 * @LastEditTime: 2019-12-16 11:37:22
 */ /*
package main

import "fmt"

//"gitee.com/veni0/robotgo"

type Fn func(x, y int) int

func (fn Fn) Chain(f Fn) Fn {
	return func(x, y int) int {
		fmt.Println(fn(x, y))
		return f(x, y)
	}
}

func add(x, y int) int {
	fmt.Println("%d + %d=", x, y)
	return x + y
}

func minus(x, y int) int {
	fmt.Println("%d - %d=", x, y)
	return x - y
}
func mul(x, y int) int {
	fmt.Println("%d * %d=", x, y)
	return x * y
}
func divide(x, y int) int {
	fmt.Println("%d / %d=", x, y)
	return x / y
}
func main() {

	/* ========================= 键盘按下 ======================== */

// 模拟按下1个键：打开开始菜单（win）
//robotgo.KeyTap(`command`)
// 模拟按下2个键：打开资源管理器（win + e）
//robotgo.KeyTap(`e`, `command`)
// 模拟按下3个键：打开任务管理器（Ctrl + Shift + ESC）
//robotgo.KeyTap(`esc`, `control`, `shift`)

/* ========================= 按住不放 ======================== */

// 一直按住 A键不放
//robotgo.KeyToggle(`a`, `down`)
// 解除按住 A键
//robotgo.KeyToggle(`a`, `up`) */
/*
	var result = Fn(add).Chain(Fn(minus)).Chain(Fn(mul)).Chain(Fn(divide))(10, 5)
	fmt.Println(result)
}
*/

package main

import (
	"github.com/keepeye/log4go"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

/***************************************************************************************************
    Global Variables
***************************************************************************************************/
var gCurCookies []*http.Cookie

var gLogger log4go.Logger

/***************************************************************************************************
    Functions
***************************************************************************************************/
//do init before all others
func initAll() {
	gCurCookies = nil
	gLogger = nil

	initLogger()
	initCrifanLib()
}

//de-init for all
func deinitAll() {
	gCurCookies = nil
	if nil == gLogger {
		gLogger.Close()
		gLogger = nil
	}
}

//do some init for crifanLib
func initCrifanLib() {
	gLogger.Debug("init for crifanLib")
	gCurCookies = nil
	return
}

//init for logger
func initLogger() {
	var filenameOnly string
	filenameOnly = GetCurFilename()
	var logFilename string = filenameOnly + ".log"

	//gLogger = log4go.NewLogger()
	gLogger = make(log4go.Logger)
	//for console
	//gLogger.AddFilter("stdout", log4go.INFO, log4go.NewConsoleLogWriter())
	gLogger.AddFilter("stdout", log4go.INFO, log4go.NewConsoleLogWriter())
	//for log file
	if _, err := os.Stat(logFilename); err == nil {
		//fmt.Printf("found old log file %s, now remove it\n", logFilename)
		os.Remove(logFilename)
	}
	//gLogger.AddFilter("logfile", log4go.FINEST, log4go.NewFileLogWriter(logFilename, true))
	var l = log4go.NewFileLogWriter(logFilename, false)
	l.SetFormat("2006-01-02T15:04:05.051+08:00")
	gLogger.AddFilter("logfile", log4go.FINEST, l)

	gLogger.Info("Current time is : %s", time.Now().Format("2006-01-02T15:04:05.051+08:00")) //2006-01-02T15:04:05.051+08:00

	return
}

// GetCurFilename
// Get current file name, without suffix
func GetCurFilename() string {
	_, fulleFilename, _, _ := runtime.Caller(0)
	//fmt.Println(fulleFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fulleFilename)
	//fmt.Println("filenameWithSuffix=", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	//fmt.Println("fileSuffix=", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	//fmt.Println("filenameOnly=", filenameOnly)

	return filenameOnly
}

//get url response html
func GetUrlRespHtml(url string) string {
	gLogger.Debug("GetUrlRespHtml, url=%s", url)
	var respHtml string = ""

	resp, err := http.Get(url)
	if err != nil {
		gLogger.Warn("http get url=%s response errror=%s\n", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//gLogger.Debug("body=%s\n", body)

	gCurCookies = resp.Cookies()

	respHtml = string(body)

	return respHtml
}

func printCurCookies() {
	var cookieNum int = len(gCurCookies)
	gLogger.Info("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		//gLogger.Info("curCk.Raw=%s", curCk.Raw)
		gLogger.Info("------ Cookie [%d]------", i)
		gLogger.Info("Name\t=%s", curCk.Name)
		gLogger.Info("Value\t=%s", curCk.Value)
		gLogger.Info("Path\t=%s", curCk.Path)
		gLogger.Info("Domain\t=%s", curCk.Domain)
		gLogger.Info("Expires\t=%s", curCk.Expires)
		gLogger.Info("RawExpires=%s", curCk.RawExpires)
		gLogger.Info("MaxAge\t=%d", curCk.MaxAge)
		gLogger.Info("Secure\t=%t", curCk.Secure)
		gLogger.Info("HttpOnly=%t", curCk.HttpOnly)
		gLogger.Info("Raw\t=%s", curCk.Raw)
		gLogger.Info("Unparsed=%s", curCk.Unparsed)
	}
}

func main() {
	log4go.Debug("in main")
	initAll()

	gLogger.Info("this is EmulateLoginBaidu.go")
	var baiduMainUrl string
	baiduMainUrl = "http://www.baidu.com/"
	gLogger.Info("baiduMainUrl=%s", baiduMainUrl)
	respHtml := GetUrlRespHtml(baiduMainUrl)
	gLogger.Debug("respHtml=%s", respHtml)
	printCurCookies()

	deinitAll()
}
