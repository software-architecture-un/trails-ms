package models

import "gopkg.in/mgo.v2/bson"

type Trail struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Usertrail     int           `bson:"usertrail" json:"usertrail"`
	Nametrail     string        `bson:"nametrail" json:"nametrail"`
	Origintrail   int       `bson:"origintrail" json:"origintrail"`
	Destinytrail  int       `bson:"destinytrail" json:"destinytrail"`
}