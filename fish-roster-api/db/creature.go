package db

import (
	"context"

	"github.com/ataboo/fish-roster/fish-roster-api/graph/model"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreatureRepo struct {
	db *mongo.Collection
}

func NewCreatureRepo(client *mongo.Client) (*CreatureRepo, error) {

	return &CreatureRepo{
		db: client.Database(DBName).Collection("creatures"),
	}, nil
}

func (c *CreatureRepo) FindAll() ([]*model.Creature, error) {
	cursor, err := c.db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	res := []*model.Creature{}
	err = cursor.All(context.TODO(), &res)
	return res, err
}

func (c *CreatureRepo) Find(id string) (*model.Creature, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var res model.Creature
	err = c.db.FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(&res)
	return &res, err
}

func (c *CreatureRepo) Create(name string) (*model.Creature, error) {
	newCreature := model.Creature{Name: name}
	res, err := c.db.InsertOne(context.TODO(), newCreature)
	if err != nil {
		return nil, err
	}

	newCreature.ID = res.InsertedID.(primitive.ObjectID).Hex()
	// return Find(res.InsertedID.(string))

	return &newCreature, err
}

func (c *CreatureRepo) Delete(id string) (bool, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	res, err := c.db.DeleteOne(context.TODO(), bson.D{{"_id", oid}})
	return res.DeletedCount > 0, err
}
