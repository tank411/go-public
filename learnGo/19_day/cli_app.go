package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "say a greeting"
	app.Action = func(c *cli.Context) error {
		for {
			name := ""
			prompt := &survey.Input{
				Message: "请输入命令:",
			}
			survey.AskOne(prompt, &name, survey.WithValidator(survey.Required))
			r, e := HandleCmd(strings.TrimSpace(name))
			if r[0] == "quit" {
				fmt.Println("Bye~ Bye!!")
				break
			}
			invokeCmd(r, e)
		}
		println("Greetings")
		return nil
	}

	app.Run(os.Args)
}

func HandleCmd(name string) ([]string, map[string]string) {
	all := strings.Split(name, "-")
	if len(all) > 2 {
		fmt.Println("额外参数隔断 - 只需一个")
	}
	r := strings.Split(all[0], " ")
	eCmd := make(map[string]string)

	if len(all) == 2 {
		for _, v := range strings.Split(all[1], " ") {
			n := strings.Split(v, "=")
			eCmd[n[0]] = n[1]
		}
	}

	return r, eCmd
}

/**
解析命令
*/
func invokeCmd(r []string, e map[string]string) {
	cmd := r[0]
	p := r[1:]
	fmt.Println(p)
	switch cmd {
	case "keys":
		fmt.Println("keys")
	case "get":
		fmt.Println("get")
	default:
		fmt.Println("命令不存在")
	}
}
