package config

import (
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InitCasbin(mongoClient *mongo.Client) *casbin.Enforcer {
	adapter, err := mongodbadapter.NewAdapterByDB(mongoClient, &mongodbadapter.AdapterConfig{})
	if err != nil {
		panic(err)
	}
	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		log.Fatalf("Casbin Enforcer Init Error: %v", err)
	}
	enforcer.EnableAutoSave(true)
	return enforcer
}
