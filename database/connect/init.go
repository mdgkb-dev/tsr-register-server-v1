package connect

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun/extra/bundebug"

	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"mdgkb/tsr-tegister-server-v1/config"
)

func InitDB(conf *config.Config) *bun.DB {
	fmt.Println(conf.DbName)
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", conf.DbDb, conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	fmt.Println(dsn)
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	_, _ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	_, _ = db.Exec(`CREATE EXTENSION IF NOT EXISTS tablefunc;`)

	return db
}

var client *redis.Client

func InitRedis(conf *config.Config) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort), //redis port
	})
	_, _ = client.Ping().Result()
	//if err != nil {
	//	panic(err)
	//}
	return client
}
