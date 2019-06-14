package mongo

import "testing"

/*
BenchmarkGetForumById_normal-8         	    3000	    512039 ns/op
BenchmarkGetForumById_view-8           	    2000	    776554 ns/op
BenchmarkGetAllForum_view_sortById-8   	      10	 146616570 ns/op
BenchmarkGetAllForum_view-8            	     100	  11843170 ns/op
BenchmarkGetAllForum_normal-8          	    2000	    635638 ns/op
*/

func BenchmarkGetForumById_normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetForumById_normal(1)
	}
}

func BenchmarkGetForumById_view(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetForumById_view(1)
	}
}

func BenchmarkGetAllForum_view_sortById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAllForum_view_sortById(1)
	}
}

func BenchmarkGetAllForum_view(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAllForum_view(1)
	}
}

func BenchmarkGetAllForum_normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAllForum_normal(1)
	}
}

