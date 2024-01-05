package cli

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/urfave/cli/v2"

	helpers "github.com/safciplak/capila/src/helpers/environment"
)

const scanCount = 50

// registerCacheCmd registers cache store commands
func (capila CapilaCLI) registerCacheCmd() {
	cmd := &cli.Command{
		Name:  "cache",
		Usage: "Cache store commands",
		Subcommands: []*cli.Command{
			cacheInfoCmd(),
			cacheFlushCmd(),
		},
	}

	capila.cli.Commands = append(capila.cli.Commands, cmd)
}

func cacheInfoCmd() *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "Shows information about the cache command",
		Action: func(context *cli.Context) error {
			return nil
		},
	}
}

func cacheFlushCmd() *cli.Command {
	return &cli.Command{
		Name:  "flush",
		Usage: "Flushes cache start with the cache prefix.",
		Action: func(ctx *cli.Context) error {
			cachePrefix := ctx.Args().First()

			if len(cachePrefix) < 1 {
				fmt.Printf("%s", "Please provide a cache prefix \n")

				return nil
			}

			deletedKeysCount, err := removeCacheKeysWithPrefix(ctx, cachePrefix)
			if err != nil {
				return err
			}

			fmt.Printf("Total deleted keys: %d\n", deletedKeysCount)

			return nil
		},
	}
}

func removeCacheKeysWithPrefix(ctx *cli.Context, prefix string) (int, error) {
	var (
		cursor           uint64
		deletedKeysCount int
		match            = fmt.Sprintf("%s*", prefix)
		envHelper        = helpers.NewEnvironmentHelper()
	)

	rClient := getRedisClient(envHelper.Get("REDIS_HOST"), envHelper.Get("REDIS_PASSWORD"), 0)

	for {
		var (
			keys []string
			err  error
		)

		keys, cursor, err = rClient.Scan(ctx.Context, cursor, match, scanCount).Result()
		if err != nil {
			return 0, fmt.Errorf("redis scan error: %w", err)
		}

		for _, key := range keys {
			err = rClient.Del(ctx.Context, key).Err()
			if err != nil {
				return 0, fmt.Errorf("redis delete key error: %w", err)
			}
			deletedKeysCount++
		}

		if cursor == 0 {
			break
		}
	}

	return deletedKeysCount, nil
}

func getRedisClient(host, password string, db int) *redis.Client {
	opts := &redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	}

	return redis.NewClient(opts)
}
