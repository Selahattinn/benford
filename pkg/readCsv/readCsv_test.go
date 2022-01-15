package readcsv

import (
	"reflect"
	"testing"
)

func TestReadCsv(t *testing.T) {
	type args struct {
		filePath string
		colIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{args: args{filePath: "../../test.csv", colIndex: 5}, name: "Read csv file", want: []int{43760427, 42918685, 40741509, 39364990, 37916059, 34937789, 34513827, 33741129}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCsv(tt.args.filePath, tt.args.colIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}
