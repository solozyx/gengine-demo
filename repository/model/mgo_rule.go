package model

import (
	"gopkg.in/mgo.v2/bson"
)

type MgoRule struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title"`           // 规则标题
	Rule      string        `bson:"rule" json:"-"`                // 规则文本
	RuleType  RuleType      `bson:"rule_type" json:"rule_type"`   // 规则类型
	CreatedAt string        `bson:"created_at" json:"created_at"` //
	UpdatedAt string        `bson:"updated_at" json:"updated_at"` //
}

type RuleType uint

const (
	RuleTypeFood RuleType = iota + 1 // 用餐规则
)
