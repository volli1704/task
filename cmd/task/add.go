package task

import (
	"time"

	redis "github.com/go-redis/redis/v9"
	"github.com/spf13/cobra"
)

func NewAddCmd(rdb *redis.Client) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add new task",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				panic(err.Error())
			}

			tasks := make([]redis.Z, len(args))

			for ix, t := range args {
				tasks[ix] = redis.Z{
					Score:  float64(time.Now().UnixNano()),
					Member: t,
				}
			}

			rdb.ZAdd(cmd.Context(), "tasks", tasks...)
		},
	}

	return addCmd
}
