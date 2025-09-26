package provider

import (
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func MongoClient() *mongo.Client {
	BaseUri := fmt.Sprintf("%s://%s:%s@", config.Get.Mongo.Protocol, config.Get.Mongo.Username, config.Get.Mongo.Password)
	for _, host := range config.Get.Mongo.Hosts {
		BaseUri += host + ":" + config.Get.Mongo.Port + ","
	}
	Uri := fmt.Sprintf(strings.TrimRight(BaseUri, ",")+"/%s?authSource="+config.Get.Mongo.AuthSource, config.Get.Mongo.DbName)
	return config.InitMongoClient(Uri)
}
func Database(client *mongo.Client) *mongo.Database {
	return client.Database(config.Get.Mongo.DbName)
}
