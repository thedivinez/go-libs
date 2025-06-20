package storage

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	DeleteOne(string, any) error
	DeleteMany(string, any) error
	Count(string, any) (int64, error)
	GenerateID(letters int, digits int) string
	InsertOne(string, any) (string, error)
	InsertMany(collection string, data ...any) (err error)
	Aggregate(collection string, filter any, results any) error
	UpdateOne(collection string, filter any, values any, opts ...*options.UpdateOptions) error
	GetPage(collection string, filter any, page string, limit, sort int64, results any) error
	ReplaceOne(collection string, filter any, values any, opts ...*options.ReplaceOptions) error
	Find(collection string, filter any, results interface{}, opts ...*options.FindOptions) error
	FindOne(collection string, filter any, results interface{}, opts ...*options.FindOneOptions) error
}
