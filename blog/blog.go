package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type MongoSessin struct {
	Session *mgo.Session
}

var MgSess *MongoSessin

type BlogInfo struct {
	Title   string
	Length  int64
	FileDir string
}

func (m *MongoSessin) Connect(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	m.Session = session
}

func (m *MongoSessin) UploadFile(title, dir string, length int64) error {
	fmt.Println("call  UploadFile")
	table := m.Session.DB("myblog").C("blogs")
	return table.Insert(&BlogInfo{title, length, dir})
}

func (m *MongoSessin) Lists() ([]BlogInfo, error) {
	fmt.Println("call  Lists")
	var blogInfos []BlogInfo
	err := m.Session.DB("myblog").C("blogs").Find(nil).All(&blogInfos)
	return blogInfos, err
}
