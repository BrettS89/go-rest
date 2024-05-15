package db

import (
	"context"

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

func (m *MongoDB) Table(s string) DbInterface {
	m.table = s

	return m
}

func (m *MongoDB) Get(id string) ([]byte, error) {
	return []byte{}, nil
}

func (m *MongoDB) Create(d any) ([]byte, error) {
	return []byte{}, nil
}

var Mongo = MongoDB{}
