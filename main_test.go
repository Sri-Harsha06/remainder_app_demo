package main

import (
	"context"
	"testing"

	"github.com/tj/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct {
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}
	res, err := insertData(mockCol, Event{"4", "you", "walking", "27-12-2022", "12:36", "12:37", "12:37", "harsha", "harsha"})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
