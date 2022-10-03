package convert

import (
	"github.com/golang-module/carbon/v2"
	"reflect"
	"testing"
	"time"
)

var now = time.Date(2022, 7, 2, 11, 45, 2, 651387237, time.UTC)
var nowDate = time.Date(2022, 7, 2, 0, 0, 0, 0, time.UTC)
var nowDateTime = time.Date(2022, 7, 2, 11, 45, 2, 0, time.UTC)

func TestToTime(t *testing.T) {

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes time.Time
		wantErr bool
	}{
		{"FromTime", args{now}, now, false},
		{"FromDateString", args{"2022-07-02"}, nowDate, false},
		{"FromDateTimeString", args{"2022-07-02 11:45:02"}, nowDateTime, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ToTime(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes.Year(), tt.wantRes.Year()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Month(), tt.wantRes.Month()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Day(), tt.wantRes.Day()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Hour(), tt.wantRes.Hour()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Minute(), tt.wantRes.Minute()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Second(), tt.wantRes.Second()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestToLayoutTime(t *testing.T) {

	type args struct {
		layout string
		value  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes time.Time
		wantErr bool
	}{
		{"FromDateString", args{carbon.DateLayout, "2022-07-02"}, nowDate, false},

		{"FromDateString", args{"m-d-y", "07-02-22"}, nowDate, false},
		{"FromDateString", args{"m-d-Y", "07-02-2022"}, nowDate, false},
		{"FromDateString", args{"m-j-Y", "07-2-2022"}, nowDate, false},
		{"FromDateString", args{"n-d-Y", "7-02-2022"}, nowDate, false},
		{"FromDateString", args{"n-j-Y", "7-2-2022"}, nowDate, false},
		{"FromDateString", args{"n-j-y", "7-2-22"}, nowDate, false},

		{"FromDateString", args{"d/m/y", "02/07/22"}, nowDate, false},
		{"FromDateString", args{"d/m/Y", "02/07/2022"}, nowDate, false},
		{"FromDateString", args{"j/m/Y", "2/07/2022"}, nowDate, false},
		{"FromDateString", args{"d/n/Y", "02/7/2022"}, nowDate, false},
		{"FromDateString", args{"j/n/Y", "2/7/2022"}, nowDate, false},
		{"FromDateString", args{"j/n/y", "2/7/22"}, nowDate, false},

		{"FromDateString", args{carbon.DateTimeLayout, "2022-07-02 11:45:02"}, nowDateTime, false},

		{"FromDateString", args{"m-d-y h:i:s", "07-02-22 11:45:02"}, nowDateTime, false},
		{"FromDateString", args{"m-d-Y H:i:s", "07-02-2022 11:45:02"}, nowDateTime, false},
		{"FromDateString", args{"m-j-Y h:i:s", "07-2-2022 11:45:02"}, nowDateTime, false},
		{"FromDateString", args{"n-d-Y h:i:s", "7-02-2022 11:45:02"}, nowDateTime, false},
		{"FromDateString", args{"n-j-Y h:i:s", "7-2-2022 11:45:02"}, nowDateTime, false},
		{"FromDateString", args{"n-j-y h:i:s", "7-2-22 11:45:02"}, nowDateTime, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ToLayoutTime(tt.args.layout, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLayoutTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes.Year(), tt.wantRes.Year()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Month(), tt.wantRes.Month()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Day(), tt.wantRes.Day()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Hour(), tt.wantRes.Hour()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Minute(), tt.wantRes.Minute()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotRes.Second(), tt.wantRes.Second()) {
				t.Errorf("ToTime() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
