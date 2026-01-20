// pkg/queue/asynq.go
package queue

import (
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/UnoraApp/be/internal/config"
)

// NewAsynqClientFromConfig creates a new Asynq client from config
func NewAsynqClientFromConfig(cfg *config.Config) *asynq.Client {
	redisAddr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)

	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}

// NewAsynqServer creates a new Asynq server for processing tasks
func NewAsynqServer(cfg *config.Config) *asynq.Server {
	redisAddr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)

	return asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
}
