package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Resp struct {
	Errno  string
	Errmsg string
	Data   interface{}
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func UploadFile2(w http.ResponseWriter, r *http.Request) {
	var resp Resp

	f, h, err := r.FormFile("content")
	if err != nil {
		panic(err)
	}
	dirname := "../file/" + h.Filename
	title := r.FormValue("title")
	file, err := os.Create(dirname)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(file, f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//fmt.Println(h)
	//w.Write([]byte("upload success"))
	//写到 数据库 中
	//fmt.Println(h.Filename, dirname, h.Size, title)
	MgSess.UploadFile(title, h.Filename, h.Size)
	resp.Errno = "0"
	resp.Errmsg = "OK"
	//fmt.Println(resp)
	wdata, err := json.Marshal(&resp)
	//fmt.Println(wdata)
	w.Write(wdata)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	var resp Resp
	title := r.FormValue("title")
	ctx := r.FormValue("content")
	hash := sha256.Sum256([]byte(title))
	fileHash := fmt.Sprintf("%x", hash)
	fmt.Println(ctx, title, fileHash)

	dirname := "static/blogfile/" + fileHash //h.Filename

	file, err := os.Create(dirname)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	//fmt.Println(h)
	//w.Write([]byte("upload success"))
	//写到 数据库 中
	//fmt.Println(h.Filename, dirname, h.Size, title)
	MgSess.UploadFile(title, "blogfile/"+fileHash, 0)
	file.Write([]byte(ctx))
	resp.Errno = "0"
	resp.Errmsg = "OK"
	//fmt.Println(resp)
	wdata, err := json.Marshal(&resp)
	//fmt.Println(wdata)
	w.Write(wdata)
}

func Lists(w http.ResponseWriter, r *http.Request) {
	var resp Resp
	s, err := MgSess.Lists()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	resp.Errno = "0"
	resp.Errmsg = "OK"
	resp.Data = s
	data, err := json.Marshal(resp)
	fmt.Println(string(data))
	w.Write(data)
}
