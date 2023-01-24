package task

import (
	"strconv"

	redis "github.com/go-redis/redis/v9"
	"github.com/spf13/cobra"
)

func NewDoCommand(rdb *redis.Client) *cobra.Command {
	doCmd := &cobra.Command{
		Use:   "do",
		Short: "Complete task",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				panic(err)
			}

			rdb := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			})

			argInt, err := strconv.Atoi(args[0])
			argInt-- // fix coz redis accept indexes from 0
			if err != nil {
				panic(err)
			}

			rdb.ZRemRangeByRank(cmd.Context(), "tasks", int64(argInt), int64(argInt))
		},
	}

	return doCmd
}
