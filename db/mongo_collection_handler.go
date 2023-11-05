package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient(connectionURL string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionURL))
	if err != nil {
		return nil, err
	}
	return client, client.Ping(context.Background(), readpref.Primary())
}

type MongoCollection[DataType interface{}] struct {
	collection *mongo.Collection
}

func NewMongoCollection[DataType interface{}](client *mongo.Client, dbName, collectionName, indexName string) (*MongoCollection[DataType], error) {
	mc := MongoCollection[DataType]{
		collection: client.Database(dbName).Collection(collectionName),
	}
	var err error
	if indexName != "" {
		indexModel := mongo.IndexModel{
			Keys: bson.D{{Key: indexName, Value: 1}},
		}
		_, err = mc.collection.Indexes().CreateOne(context.Background(), indexModel)
	}

	return &mc, err
}

func (mc *MongoCollection[DataType]) GetElementBy(key string, value interface{}) (DataType, error) {
	return mc.FindOneWithCustomFilter(bson.D{bson.E{Key: key, Value: value}})
}

func (mc *MongoCollection[DataType]) GetElementsBy(key string, value interface{}) ([]DataType, error) {
	filter := bson.D{bson.E{Key: key, Value: value}}
	return mc.FindManyWithCustomFilter(filter)
}

func (mc *MongoCollection[DataType]) GetElementsFromTo(fromKey, toKey string, from, to interface{}) ([]DataType, error) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: fromKey, Value: bson.D{{Key: "$gt", Value: from}}}},
				bson.D{{Key: toKey, Value: bson.D{{Key: "$lte", Value: to}}}},
			},
		},
	}
	return mc.FindManyWithCustomFilter(filter)
}

func (mc *MongoCollection[DataType]) UpdateBy(key string, value interface{}, updated DataType) (interface{}, error) {
	return mc.collection.UpdateOne(context.Background(), bson.D{{Key: key, Value: value}}, updated)
}

func (mc *MongoCollection[DataType]) InsertElement(element *DataType) (interface{}, error) {
	if element == nil {
		return nil, fmt.Errorf("Can't insert empty element in collection.")
	}
	res, err := mc.collection.InsertOne(context.Background(), *element)
	return res.InsertedID, err
}

func (mc *MongoCollection[DataType]) InsertElements(elements []DataType) (interface{}, error) {
	var docs []interface{}
	for _, element := range elements {
		serializedDoc, err := bson.Marshal(&element)
		if err != nil {
			return nil, err
		}
		docs = append(docs, serializedDoc)
	}

	return mc.collection.InsertMany(context.Background(), docs)
}

func (mc *MongoCollection[DataType]) RemoveElement(key, value string) (interface{}, error) {
	return mc.collection.DeleteOne(context.Background(), bson.D{bson.E{Key: key, Value: value}})
}

func (mc *MongoCollection[DataType]) FindOneWithCustomFilter(filter bson.D) (DataType, error) {
	var result DataType
	return result, mc.collection.FindOne(context.Background(), filter).Decode(&result)
}

func (mc *MongoCollection[DataType]) FindManyWithCustomFilter(filter bson.D) ([]DataType, error) {
	cursor, err := mc.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var results []DataType
	return results, cursor.All(context.Background(), &results)
}
