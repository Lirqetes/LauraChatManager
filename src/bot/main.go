package main

import (
	"context"
	"time"

	"lirqetes.ru/thelanc3/laura/queue"
)

func main() {
	ctx := context.Background()
	actions := make(chan queue.BotAction, 200)
	go actionsWorker(ctx, actions, 34*time.Millisecond)
}

func actionsWorker(
	ctx context.Context,
	actions <-chan queue.BotAction,
	interval time.Duration,
) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		case action := <-actions:
			{
				<-ticker.C
				action(ctx)
			}
		}
	}
}
