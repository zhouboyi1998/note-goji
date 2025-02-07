package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Command 命令结构体
type Command struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Command string             `bson:"command" json:"command"`
	Usage   string             `bson:"usage" json:"usage"`
	Params  []struct {
		Param       string `bson:"param" json:"param"`
		Description string `bson:"description" json:"description"`
	} `bson:"params" json:"params"`
}
