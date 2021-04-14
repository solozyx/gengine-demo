package repository

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gengine/common"
	"gengine/repository/model"
)

const (
	_mgoRule = "rule"
)

type mgoRuleRepo struct {
	m *mgo.Session
}

func NewMgoRuleRepo(cfg *common.Config) *mgoRuleRepo {
	return &mgoRuleRepo{
		m: GetMgoSession(cfg),
	}
}

func (p *mgoRuleRepo) Create(entity model.MgoRule) error {
	session := p.m.Copy()
	defer session.Close()

	entity.ID = bson.NewObjectId()
	entity.CreatedAt = time.Now().Format(time.RFC3339)

	collection := session.DB(_mgoDB).C(_mgoRule)
	err := collection.Insert(&entity)
	if err != nil {
		return err
	}
	return nil
}
