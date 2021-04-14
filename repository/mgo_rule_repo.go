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
	entity.UpdatedAt = entity.CreatedAt

	collection := session.DB(_mgoDB).C(_mgoRule)
	err := collection.Insert(&entity)
	if err != nil {
		return err
	}
	return nil
}

func (p *mgoRuleRepo) GetLatestFoodRule() (*model.MgoRule, error) {
	session := p.m.Copy()
	defer session.Close()
	collection := session.DB(_mgoDB).C(_mgoRule)

	list := make([]model.MgoRule, 0)

	err := collection.Find(bson.M{"rule_type": model.RuleTypeFood}).
		Sort("-_id").Limit(1).All(&list)
	if err != nil {
		return nil, err
	}

	return &list[0], nil
}
