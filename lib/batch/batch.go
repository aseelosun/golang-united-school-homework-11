package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var result []user
	for i := 0; i < int(n); i++ {
		go func() {
			u := getOne(int64(i))
			result = append(result, u)
		}()
	}
	return result
}
