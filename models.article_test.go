package main

import (
	"reflect"
	"testing"
)

func Test_getAllArticles(t *testing.T) {
	tests := []struct {
		name string
		want []article
	}{
		{
			name: "get all articles",
			want: articleList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllArticles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
