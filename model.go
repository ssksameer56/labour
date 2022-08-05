package labour

import (
	"context"
	"time"
)

/*
The Worker struct to initialize a worker
Parent Context - context to be passed
NotificationChannel on which data/work is returned
PingTime - In case periodic pings are to be enabled
*/
type worker struct {
	parentContext       context.Context
	notificationChannel chan interface{}
	pingTime            time.Duration
	enableChannelPing   bool
}

//creates a new worker with context, the notif channel, ping duration and if pings are to be enabled
func NewWorker(ctx context.Context, notificationChannel chan interface{}, pingTime time.Duration, enablePing bool) worker {
	return worker{
		parentContext:       ctx,
		notificationChannel: notificationChannel,
		pingTime:            pingTime,
		enableChannelPing:   enablePing,
	}
}
