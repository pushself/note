package main

import (
	"os"
	"log"
	"github.com/axgle/mahonia"
	"encoding/csv"
)

func main() {
	file,err:=os.Open("shuju.txt")

	if err!=nil{
		log.Fatal(err.Error())
	}
defer file.Close()
	decoder:=mahonia.NewDecoder("gbk")

	reader:=csv.NewReader(decoder.NewReader(file))

	set,err:=reader.ReadAll()

	if  err!=nil{
		log.Fatal(err.Error())
	}

	zz,err:=os.OpenFile("zuizhong.txt",os.O_WRONLY,0777)
	if err!=nil{
		log.Fatal("打开文件失败2")
	}


	for _,data:=range set{
		if data[1]!="123456"{
			zz.WriteString(data[0]+","+data[1]+"\n")
			println(data[1])
		}
	}
}
