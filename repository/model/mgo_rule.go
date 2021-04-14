package model

import (
	"gopkg.in/mgo.v2/bson"
)

type MgoRule struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Rule      string        `bson:"rule" json:"-"`                // 规则文本
	CreatedAt string        `bson:"created_at" json:"created_at"` //
	UpdatedAt string        `bson:"updated_at" json:"updated_at"` //
}
