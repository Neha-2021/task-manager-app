package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskList struct {
	ID		primitive.ObjectID	`json:"_id,omiempty" bson:"_id,omitempty"`	
	Task	string				`json:"task.omitempty"`
	Status	bool				`json:"status.omitempty"`
}