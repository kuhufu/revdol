package redis

import (
	"fmt"
	"testing"
)

func TestGetForumById(t *testing.T) {
	res := GetForumById("17115")
	fmt.Println(res)
}

//BenchmarkGetForumById-8   	   10000	    114833 ns/op
func BenchmarkGetForumById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetForumById("17115")
	}
}
