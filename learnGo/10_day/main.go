package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

type test struct {
	V interface{}
	M string
}

type User struct {
	Name string
}

func main() {
	a := f(1)
	a()

	c := test{
		V: &struct {
			M string
		}{
			M: "12",
		},
		M: "23",
	}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	any := User{
		Name: "fidding",
	}
	interfaceTest(any)

	//res, ok := anyInterface.(someType)
	fmt.Println("////////////////当前时间//////////////")
	fmt.Println(time.Now())
	tn := time.Now().Unix() - int64(3600)*int64(1)
	timeBody := time.Unix(tn, 0)
	fmt.Println(timeBody)

	timeStr := time.Now().Format("2006-01-02T15")
	etime, err := time.Parse("2006-01-02T15", timeStr)
	if err == nil {
		fmt.Println(etime) //打印结果：2017-04-11 13:33:37 +0000 UTC
	}
	fmt.Println("////////////////结束时间//////////////")
	fmt.Println(etime)
	bt := etime.Unix() - int64(3600)
	btime := time.Unix(bt, 0)
	fmt.Println("////////////////开始时间//////////////")
	fmt.Println(btime)

	t := time.Now()
	bTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	fmt.Println(bTime)
	zero := bTime.Unix() - int64(3600)
	endTime := time.Unix(zero, 0)

	cc := time.Date(t.Year(), t.Month(), t.Day()+1, t.Hour(), 0, 0, 0, t.Location())
	fmt.Println(cc)
	cc = time.Now()
	fmt.Println(cc)

	fmt.Println("////////////////ToTime0//////////////")
	fmt.Println(endTime)
	var aa interface{}
	aa = 1573015069
	test1 := ToTime(aa)
	fmt.Println(test1.UnixNano())
	var bb interface{}
	bb = "2019-11-06T12:37:49.438+08:00"
	test2 := ToTime(bb)
	fmt.Println(test2.UnixNano())
	fmt.Println("////////////////ToTime1//////////////")

	groupId := "1234654%"
	arr := strings.Split(groupId, "%")
	fmt.Println(arr)
	fmt.Println(len(arr))

	var mp map[string]int
	mp = make(map[string]int)
	mp["2019-11-07T05:00:00.000+08:00"] = 1
	mp["2019-11-07T01:00:00.000+08:00"] = 45
	mp["2019-11-06T16:00:00.000+08:00"] = 3
	mp["2019-11-07T08:00:00.000+08:00"] = 22

	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		fmt.Println("根据key排序后的新集合》》   key:", v, "    value:", mp[v])
	}
	fmt.Println(mp)
	now := time.Now()
	temp := now.UnixNano() / 1e6
	fmt.Println(temp)
	fmt.Println(IntToTimeStr(0, int64(temp), "2006-01-02T15:04:05.000-07:00"))

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		fmt.Println("--------1------------------")
		fmt.Println(val, "->", val)
		m[key] = &val
		fmt.Println("-----------2---------------")
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
		fmt.Println("-----------3---------------")
		fmt.Println(k, "->", *v)
		fmt.Println("-----------4---------------")
	}

	et := time.Now()
	fmt.Println(et)
	et = et.Add(time.Minute * -2)
	fmt.Println(et)

	{
		d := struct {
			Name string
		}{
			Name: "123",
		}
		var cc interface{} = d
		n, ok := cc.(struct {
			Name string
		})
		if ok {
			fmt.Println(n.Name)
		}
	}
	fmt.Println(000)
	{
		s1 := []string{"1", "2", "3", "4", "5"}
		var linetype string
		linetype = fmt.Sprint(s1[:3])
		fmt.Println(linetype)
		//str := string(s1[:3])
		//fmt.Println(str)
	}

	{

		RuleDes := "检查规则：1、xface/xbody/xvehicle资源查询接口调用是否正常（前端触发：给出最新数据）2、展示最近24小时人脸抓拍数据。\n检测周期：30分钟"
		fmt.Println(RuleDes)
	}
}

func IntToTimeStr(k int8, t int64, format string) string {
	if k == 0 {
		t = t / 1000
	}
	formatTimeStr := time.Unix(int64(t), 0).Format(format)
	return formatTimeStr
}

func interfaceTest(value interface{}) {
	res, ok := value.(User)
	if ok {
		fmt.Println(res.Name)
	}
	switch v := value.(type) {
	case string:
		fmt.Println(v)
	case int32, int64:
		fmt.Println(v)
	case User:
		// 可以看到op即为将interface转为User struct类型，并使用其Name对象
		op, ok := value.(User)
		fmt.Println(op.Name, ok)
	default:
		fmt.Println("unknown")
	}
}

func ToTime(t interface{}) time.Time {
	var tm int64
	switch t.(type) {
	case int64:
		{
			newInt, _ := t.(int64)
			tm = time.Unix(newInt, 0).Unix()
		}
	case int32:
		{
			newInt, _ := t.(int32)
			tm = time.Unix(int64(newInt), 0).Unix()
			fmt.Println(newInt)
		}
	case int:
		{
			newInt, _ := t.(int)
			tm = time.Unix(int64(newInt), 0).Unix()
			fmt.Println(newInt)
		}
	case string:
		{
			newString, _ := t.(string)
			temp, _ := time.Parse("2006-01-02T15:04:05.000-07:00", newString)
			tm = temp.Unix()
		}
	}
	fmt.Printf("ToTime:%d", tm)
	return time.Unix(int64(tm), 0)
}
