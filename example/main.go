/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-26 19:10
**/

package main

import (
	"go.mongodb.org/mongo-driver/bson"
	
	"mongo"
)

func main() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"
	
	mgo, _ := mongo.NewMongoClient().Connect(&mongo.Config{Url: url})
	
	err := mgo.RawClient().Ping(nil, mongo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}
	
	mgo.Transaction(func(sessionContext mongo.SessionContext) error {
		
		var err error
		
		_, err = mgo.DB("QGame").C("test1").InsertOneWithSession(sessionContext, bson.M{"hello": "world"})
		
		_, err = mgo.DB("QGame").C("test2").InsertOneWithSession(sessionContext, bson.M{"hello": "world"})
		
		return err
	})
	
}
