package main
//文件的学习
import(
	"fmt"
	"os"
	"bufio"
	"log"
)

func ReadFile1(path string) {
		input,_  := os.Open(path)
		reader := bufio.NewReader(input)
		buff := make([]byte,1)
		for {
			n,_ := reader.Read(buff)
			if n == 0{
				break
			}
			fmt.Printf("%s",string(buff[0:n]))
		}
}	

func ReadFile2(path string) {
	currFile,err := os.Open(path)
	defer currFile.Close()
	if(err == nil){
		fmt.Println("打开成功！")
		buf := make([]byte,1)
		for{
			n,_:=currFile.Read(buf)
			if( 0 == n){
				break
			}
			os.Stdout.Write(buf[:n])
		}

	}else{

		fmt.Println("打开失败！")
	}
}


func main()  {
	fmt.Printf("第四天学习")
	
	/*目录*/
	os.Mkdir("filename",0777)
	os.MkdirAll("filname/1/2",0777)
	err :=os.Remove("filename")
	if(err != nil){
		fmt.Println(err)
	}
	err = os.RemoveAll("filname")
	if err != nil{
		fmt.Println(err)
	}

	/*创建文件*/
	/*
	变量必须使用不然会报错
	*/
	//fileName := "testFile"
    currFile,err:=os.Create("testFile1.txt")
	if err != nil{
		fmt.Println(err)
		log.Fatal(err)
	}
	currFile.WriteString("hello world\n")

	currFile.WriteAt([]byte("111"),5)



	/*打开文件*/

	currFile,err = os.Open("testFile1.txt")
	defer currFile.Close()
	if(err == nil){
		fmt.Println("打开成功！")
		buf := make([]byte,1024)
		for{
			n,_:=currFile.Read(buf)
			if( 0 == n){
				break
			}
			os.Stdout.Write(buf[:n])
		}
		fmt.Println(buf)

	}else{

		fmt.Println("打开失败！")
	}

	/*读文件*/
	ReadFile1("testFile.txt")
}