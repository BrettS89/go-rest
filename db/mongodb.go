package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	table  string
}

func (m *MongoDB) setClient(c *mongo.Client) {
	m.client = c
}

func (m *MongoDB) Connect(c string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c))

	if err != nil {
		return err
	}

	m.setClient(client)

	return nil
}

func (m *MongoDB) Unmarshal(b []byte, i interface{}) error {
	return bson.Unmarshal(b, i)
}

func (m *MongoDB) Table(s string) DbInterface {
	m.table = s

	return m
}

func (m *MongoDB) Get(id string) ([]byte, error) {
	coll := m.client.Database("bank").Collection(m.table)

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return []byte{}, err
	}

	b, err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objId}}).Raw()

	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func (m *MongoDB) Create(d interface{}) ([]byte, error) {
	coll := m.client.Database("bank").Collection(m.table)

	result, err := coll.InsertOne(context.TODO(), d)

	if err != nil {
		return []byte{}, err
	}

	b, err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: result.InsertedID}}).Raw()

	if err != nil {
		return b, err
	}

	return b, nil
}

func (m *MongoDB) Patch(id string, d interface{}) ([]byte, error) {
	coll := m.client.Database("bank").Collection(m.table)

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	update := bson.D{{Key: "$set", Value: d}}

	result, err := coll.UpdateByID(context.TODO(), objId, update)

	if err != nil {
		return nil, err
	}

	b, err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: result.UpsertedID}}).Raw()

	if err != nil {
		return nil, err
	}

	return b, nil
}

var Mongo = MongoDB{}
