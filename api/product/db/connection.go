package db

import (
	"fmt"

	"github.com/cbellee/goShop-productService/config"
)

func connectionString(conf config.Config) string {
	if conf.IsLocal {
		// mongodb://cosmosdb-8a9c:`password`@cosmosdb-8a9c.documents.azure.com:10255/?ssl=true&replicaSet=globaldb
		return fmt.Sprintf("mongodb://%s:%d", conf.DbHostName, conf.DbPort)
	}
	return fmt.Sprintf("mongodb://%s:%s@%s.%s:%d/?ssl=true&replicaSet=globaldb", conf.DbHostName, conf.DbPassword, conf.DbHostName, conf.DbHostNameSuffix, conf.DbPort)
}
