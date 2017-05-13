package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Activity struct {
		Id bson.ObjectId
	}
	Task struct {
	}
	Process    struct{}
	Vendor     struct{}
	User       struct{}
	Role       struct{}
	Permission struct{}
	Note       struct{}
	Approval   struct{}
	Event      struct{}
)
