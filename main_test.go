package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestDBClient struct {
	success bool
}

func (tс *TestDBClient) Get(interface{}, string, ...interface{}) error {
	if tс.success {
		return nil
	}
	return fmt.Errorf("This is a test error")
}

func TestCheckEmployee(t *testing.T) {
	type args struct {
		db BaseDBClient
	}
	tests := []struct {
		name       string
		args       args
		wantErr    error
		wantExists bool
	}{
		{
			name: "Employee exists",
			args: args{
				db: &TestDBClient{success: true},
			},
			wantErr:    nil,
			wantExists: true,
		}, {
			name: "Employee don't exists",
			args: args{
				db: &TestDBClient{success: false},
			},
			wantErr:    fmt.Errorf("This is a test error"),
			wantExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, gotExists := CheckEmployee(tt.args.db, "some phone")
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("CheckEmployee() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if gotExists != tt.wantExists {
				t.Errorf("CheckEmployee() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}

}
