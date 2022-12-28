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

func (m *mockCollection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	c := &mongo.SingleResult{}
	return c
}

func (m *mockCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	c := &mongo.Cursor{}
	return c, nil
}

func (m *mockCollection) ReplaceOne(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	c := &mongo.UpdateResult{}
	return c, nil
}
func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	return c, nil
}

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}
	res, err := insertData(mockCol, Event{"5", "you", "walking", "27-12-2022", "12:36", "12:37", "12:37", "harsha", "harsha"})
	res2 := findDataById(mockCol, Event{Id: "5"})
	res3, err2 := getData(mockCol, Event{})
	res4, err3 := updateData(mockCol, Event{Id: "5"})
	res5, err4 := deleteData(mockCol, Event{Id: "5"})
	assert.IsType(t, &mongo.DeleteResult{}, res5)
	assert.Nil(t, err4)
	assert.IsType(t, &mongo.UpdateResult{}, res4)
	assert.Nil(t, err3)
	assert.IsType(t, &mongo.Cursor{}, res3)
	assert.Nil(t, err2)
	assert.IsType(t, &mongo.SingleResult{}, res2)
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
