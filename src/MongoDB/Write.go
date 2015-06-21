package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ben", 10})
	if err != nil {
		panic(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name" : "Ben"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Age:", result.Age)
}
