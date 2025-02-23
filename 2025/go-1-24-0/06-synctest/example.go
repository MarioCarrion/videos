package example

import "time"

type Pusher interface {
	Push(int64) error
}

func Retry(retries, value int64, pusher Pusher) error {
	var err error

	for range retries {
		err = pusher.Push(value)
		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	return err
}
