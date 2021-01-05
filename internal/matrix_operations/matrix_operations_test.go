package matrix_operations

import (
	"testing"
)

func Test_matrixOperationsImpl_Echo(t *testing.T) {
	type args struct {
		i [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Given a valid matrix, return it",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}},
			want: "1,2,3\n4,5,6\n7,8,9\n",
		},
		{
			name: "Given an invalid matrix containing letters, return an error",
			args: args{i: [][]string{
				{"A", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			wantErr: true,
		},
		{
			name: "Given a not squared matrix, return an error",
			args: args{i: [][]string{
				{"0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &matrixOperationsImpl{}
			got, err := m.Echo(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Echo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixOperationsImpl_Flatten(t *testing.T) {
	type args struct {
		i [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Given a valid matrix, return it flattened",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}},
			want: "1,2,3,4,5,6,7,8,9",
		},
		{
			name: "Given an invalid matrix containing letters, return an error",
			args: args{i: [][]string{
				{"A", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			wantErr: true,
		},
		{
			name: "Given a not squared matrix, return an error",
			args: args{i: [][]string{
				{"0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &matrixOperationsImpl{}
			got, err := m.Flatten(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Flatten() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixOperationsImpl_Invert(t *testing.T) {
	type args struct {
		i [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Given a valid matrix, return it inverted",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}},
			want: "1,4,7\n2,5,8\n3,6,9\n",
		},
		{
			name: "Given an invalid matrix containing letters, return an error",
			args: args{i: [][]string{
				{"A", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			wantErr: true,
		},
		{
			name: "Given a not squared matrix, return an error",
			args: args{i: [][]string{
				{"0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &matrixOperationsImpl{}
			got, err := m.Invert(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Invert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Invert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixOperationsImpl_Multiply(t *testing.T) {
	type args struct {
		i [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Given a valid matrix with only positive numbers, return the correct result",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}},
			want: "362880",
		},
		{
			name: "Given a valid matrix with positive and negative numbers, return the correct result",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "-5", "6"},
				{"7", "8", "9"},
			}},
			want: "-362880",
		},
		{
			name: "Given a valid matrix with only negative numbers, return the correct result",
			args: args{i: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			}},
			want: "-362880",
		},
		{
			name: "Given a valid matrix with only zeroes, return zero",
			args: args{i: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			want: "0",
		},
		{
			name: "Given an invalid matrix containing letters, return an error",
			args: args{i: [][]string{
				{"A", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			want:    "0",
			wantErr: true,
		},
		{
			name: "Given a not squared matrix, return an error",
			args: args{i: [][]string{
				{"0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			want:    "0",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &matrixOperationsImpl{}
			got, err := m.Multiply(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Multiply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrixOperationsImpl_Sum(t *testing.T) {
	type args struct {
		i [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Given a valid matrix with only positive numbers, return the correct sum",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}},
			want: "45",
		},
		{
			name: "Given a valid matrix with positive and negative numbers, return the correct sum",
			args: args{i: [][]string{
				{"1", "2", "3"},
				{"4", "-5", "6"},
				{"7", "8", "9"},
			}},
			want: "35",
		},
		{
			name: "Given a valid matrix with only negative numbers, return the correct sum",
			args: args{i: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			}},
			want: "-45",
		},
		{
			name: "Given a valid matrix with only zeroes, return zero",
			args: args{i: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			want: "0",
		},
		{
			name: "Given an invalid matrix containing letters, return an error",
			args: args{i: [][]string{
				{"A", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			want:    "0",
			wantErr: true,
		},
		{
			name: "Given a not squared matrix, return an error",
			args: args{i: [][]string{
				{"0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			}},
			want:    "0",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &matrixOperationsImpl{}
			got, err := m.Sum(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sum() got = %v, want %v", got, tt.want)
			}
		})
	}
}
