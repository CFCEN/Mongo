package Db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Criteria struct {
	sql map[string]interface{}
}

func Where(key string, value interface{}) *Criteria {
	criteria := Criteria{}
	criteria.sql = make(map[string]interface{})
	criteria.sql[key] = value
	return &criteria
}
func (criteria *Criteria) And(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = value
	return *criteria
}

func (criteria *Criteria) Or(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql["$or"] = []bson.M{{key: value}}
	return *criteria
}

func (criteria *Criteria) In(key string, value []interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$in": value}
	return *criteria
}

func (criteria *Criteria) Ne(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$ne": value}
	return *criteria
}

func (criteria *Criteria) Nin(key string, value []interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$nin": value}
	return *criteria
}

func (criteria *Criteria) Lt(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$lt": value}
	return *criteria
}

func (criteria *Criteria) Lte(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$lte": value}
	return *criteria
}

func (criteria *Criteria) Gt(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$gt": value}
	return *criteria
}

func (criteria *Criteria) Gte(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$gte": value}
	return *criteria
}

func (criteria *Criteria) Eq(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = value
	return *criteria
}

func (criteria *Criteria) Regex(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$regex": value}
	return *criteria
}

func (criteria *Criteria) Like(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$regex": value}
	return *criteria
}

func (criteria *Criteria) NotLike(key string, value interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$not": bson.M{"$regex": value}}
	return *criteria
}

func (criteria *Criteria) Between(key string, value []interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$gte": value[0], "$lte": value[1]}
	return *criteria
}

func (criteria *Criteria) NotBetween(key string, value []interface{}) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$not": bson.M{"$gte": value[0], "$lte": value[1]}}
	return *criteria
}

func (criteria *Criteria) IsNull(key string) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = nil
	return *criteria
}

func (criteria *Criteria) IsNotNull(key string) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$ne": nil}
	return *criteria
}

func (criteria *Criteria) IsEmpty(key string) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = ""
	return *criteria
}

func (criteria *Criteria) IsNotEmpty(key string) Criteria {
	if criteria.sql == nil {
		panic("please use Where first")
	}
	criteria.sql[key] = bson.M{"$ne": ""}
	return *criteria
}

func (criteria *Criteria) GetCriteria() bson.M {
	return criteria.sql
}

func ConvertStringToObjectId(id string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return objectId
}
