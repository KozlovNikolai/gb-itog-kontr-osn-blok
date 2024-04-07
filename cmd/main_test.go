package main

import (
	"reflect"
	"testing"
)

func Test_arrayTo3CharStringsArray(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test-0",
			args: args{a: []string{"Hello", "2", "world", ":-"}},
			want: []string{"2", ":-"},
		},
		{
			name: "test-1",
			args: args{a: []string{"1234", "1567", "-2", "computer science"}},
			want: []string{"-2"},
		},
		{
			name: "test-2",
			args: args{a: []string{"Russia", "Denmark", "Kazan"}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayTo3CharStringsArray(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayTo3CharStringsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
