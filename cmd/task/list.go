package task

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v9"
	"github.com/spf13/cobra"
)

func NewListCommand(rdb *redis.Client) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Print list of all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := rdb.ZRange(context.Background(), "tasks", 0, -1).Result()
			if err != nil {
				panic(err)
			}

			for ix, t := range tasks {
				fmt.Printf("%d. %s\n", ix+1, t)
			}
		},
	}

	return listCmd
}
