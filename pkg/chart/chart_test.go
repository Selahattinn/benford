package chart

import "testing"

func TestGenerateChart(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Generate Chart", args: args{data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateChart(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("GenerateChart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
