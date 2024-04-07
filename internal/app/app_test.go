package app

import (
	"reflect"
	"testing"
)

func TestGetDataFromFile(t *testing.T) {
	path := "/home/user/Yandex.Disk/synchro/home/geekbrain/osnov-blok/listing"
	type args struct {
		dataFile *string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test-in",
			args: args{dataFile: &path},
			want: []string{
				"Moscow",
				"Ufa",
				"24",
				"Russia",
				"Обь",
				"Ока",
				"Шуя",
				"Дно Ростовское",
				"123",
				"2022",
				"765 12"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDataFromFile(tt.args.dataFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
