package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//1. 连接到mongo
	sess, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("failed to dial mongo", err)
		return
	}
	fmt.Println(sess)

	//构造一个collection
	tabName := "person"
	ps := sess.DB("myblog").C(tabName)
	//添加一个perso
	//func (c *Collection) Insert(docs ...interface{}) error
	err = ps.Insert(&Person{"yekai", 30})
	if err != nil {
		panic(err)
	}

	//查询person
	//p1 := Person{}
	//func (c *Collection) Find(query interface{}) *Query
	//ps.Find(bson.M{"name": "yekai"}).One(&p1)
	p2 := make(map[string]interface{})
	ps.Find(bson.M{"name": "yekai"}).One(&p2)
	fmt.Println(p2)

	//修改person
	//func (c *Collection) Update(selector interface{}, update interface{}) error
	err = ps.Update(bson.M{"name": "yekai"}, bson.M{"name": "yekai", "age": 35})
	if err != nil {
		panic(err)
	}

	//删除全部纪录
	//(c *Collection) RemoveAll(selector interface{}) (info *ChangeInfo, err error)
	ps.RemoveAll(bson.M{"name": "yekai"})
}
