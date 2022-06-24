package main

import "context"

func main() {
	rootCtx := context.Background()
	childCtx1 := context.WithValue(rootCtx, "msgId", "someMsgId")
	childCtx2, cancelFunc := context.WithCancel(childCtx1)
	childCtx3 := context.WithValue(rootCtx, "user_id", "some_user_id")
	childCtx4 := context.WithValue(childCtx1, "current_time", "some_time")
}
