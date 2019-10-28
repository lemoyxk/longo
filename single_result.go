/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:32
**/

package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type SingleResult struct {
	singleResult *mongo.SingleResult
}

func (sg *SingleResult) Get(result interface{}) error {
	return sg.singleResult.Decode(result)
}