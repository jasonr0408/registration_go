package main

import (
	"net/http"
	delivery "registration/registration/delivery/http"
	"registration/registration/usecase"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 載入router root
	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9487")
}

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:9528",
			"https://127.0.0.1:9528",
			"http://localhost:9528",
			"https://localhost:9528",
			"http://192.168.43.16:9528",
			"https://192.168.43.16:9528",
		},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	redisConn := newRedis()
	usecase := usecase.NewRedisUsecase(redisConn)

	// 載入 Administrator router
	delivery.NewHandler(r, usecase)

	return r
}

// 建立redis
func newRedis() *redis.Pool {
	Conn := &redis.Pool{
		Wait:        true,
		MaxIdle:     20,
		MaxActive:   2000,
		IdleTimeout: 10 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6378")
			if err != nil {
				return nil, err
			}
			// _, err = c.Do("AUTH", "123")
			// if err != nil {
			// 	return nil, err
			// }

			c.Do("SELECT", 4)

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return Conn
}
