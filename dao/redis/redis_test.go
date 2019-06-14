package redis

import (
	"fmt"
	"testing"
)

func TestGetForumById(t *testing.T) {
	res := GetForumById("1")
	fmt.Println(string(res))
}
