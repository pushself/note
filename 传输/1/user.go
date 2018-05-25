package main

import (
	"os"
	"log"
	"encoding/csv"
	"github.com/axgle/mahonia"
	"regexp"
	"io"
)

func main() {
	file,err:=os.Open("phone.txt")
	if err!=nil{
		log.Fatal("文件打开失败")
	}

	defer file.Close()

	decoder:=mahonia.NewDecoder("gbk")


	reader:=csv.NewReader(decoder.NewReader(file))
	str,err:=reader.ReadAll()

	if err!=nil{
		log.Fatal(err.Error())
	}

	file2,err:=os.OpenFile("shuju.txt",os.O_WRONLY,0777)
	if err!=nil{
		log.Fatal("文件打开失败2")
	}
	defer file2.Close()



	r,_:=regexp.Compile("码为：(.*?)。")
	var phone,pwd string

	for _,data:= range str{

		 pwd=r.FindString(data[2])
		 if pwd!=""{
			 phone=data[3]
			 //println(pwd)
			 //println(phone)

			 _,err:=io.WriteString(file2,phone+","+pwd+"\n")

			 if err!=nil{
			 	log.Fatal("写入错误",err.Error())
			 }
		 }

	}
}
