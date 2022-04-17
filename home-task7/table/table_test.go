package table

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortStrings(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				s: []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"},
			},
			want: []string{"", "%*&^*&^&", "***", "Hello", "bar", "f00", "foo", "foo"},
		},
		{
			name: "test2",
			args: args{
				s: []string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.args.s)
			sortedStrings := tt.args.s
			if got := sortedStrings; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
