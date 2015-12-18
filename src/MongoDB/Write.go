package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	session, err := mgo.Dial("192.168.1.14:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("people")

	err = c.Insert(&Person{"Bob", 15, })
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ben"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Age:", result.Age)
}
