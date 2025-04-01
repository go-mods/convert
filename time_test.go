package convert

import (
	"reflect"
	"testing"
	"time"

	"github.com/dromara/carbon/v2"
)

var now = time.Date(2022, 7, 2, 11, 45, 2, 651387237, time.UTC)
var nowStr = carbon.CreateFromStdTime(now, "UTC").ToRfc3339String("UTC")
var nowDate = time.Date(2022, 7, 2, 0, 0, 0, 0, time.Local)
var nowDateTime = time.Date(2022, 7, 2, 11, 45, 2, 0, time.Local)

func TestToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected time.Time
	}{
		{
			name:     "valid time",
			input:    now,
			expected: now,
		},
		{
			name:     "valid time pointer",
			input:    &now,
			expected: now,
		},
		{
			name:     "valid date string",
			input:    "2022-07-02",
			expected: nowDate,
		},
		{
			name:     "valid datetime string",
			input:    "2022-07-02 11:45:02",
			expected: nowDateTime,
		},
		{
			name:     "empty string",
			input:    "",
			expected: time.Time{},
		},
		{
			name:     "invalid date string",
			input:    "not-a-date",
			expected: time.Time{},
		},
		{
			name:     "nil input",
			input:    nil,
			expected: time.Time{},
		},
		{
			name:     "integer input",
			input:    42,
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToTime(tt.input)
			if tt.expected.IsZero() {
				if !got.IsZero() {
					t.Errorf("ToTime() = %v, want zero time for input %v", got, tt.input)
				}
			} else {
				if got.IsZero() {
					t.Errorf("ToTime() returned zero time for input %v, expected non-zero time", tt.input)
				}
				if !got.UTC().Equal(tt.expected.UTC()) {
					t.Errorf("ToTime() = %v, want %v", got.UTC(), tt.expected.UTC())
				}
			}
		})
	}
}

func TestToTimeOrDefault(t *testing.T) {
	defaultTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name         string
		input        interface{}
		defaultValue time.Time
		expected     time.Time
	}{
		{
			name:         "valid time",
			input:        now,
			defaultValue: defaultTime,
			expected:     now,
		},
		{
			name:         "valid time pointer",
			input:        &now,
			defaultValue: defaultTime,
			expected:     now,
		},
		{
			name:         "valid date string",
			input:        "2022-07-02",
			defaultValue: defaultTime,
			expected:     nowDate,
		},
		{
			name:         "valid datetime string",
			input:        "2022-07-02 11:45:02",
			defaultValue: defaultTime,
			expected:     nowDateTime,
		},
		{
			name:         "empty string",
			input:        "",
			defaultValue: defaultTime,
			expected:     defaultTime,
		},
		{
			name:         "invalid date string",
			input:        "not-a-date",
			defaultValue: defaultTime,
			expected:     defaultTime,
		},
		{
			name:         "nil input",
			input:        nil,
			defaultValue: defaultTime,
			expected:     defaultTime,
		},
		{
			name:         "integer input",
			input:        42,
			defaultValue: defaultTime,
			expected:     defaultTime,
		},
		{
			name:         "zero default value",
			input:        "invalid",
			defaultValue: time.Time{},
			expected:     time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToTimeOrDefault(tt.input, tt.defaultValue)
			if tt.expected.IsZero() {
				if !got.IsZero() {
					t.Errorf("ToTimeOrDefault() = %v, want zero time", got)
				}
			} else {
				if !got.UTC().Equal(tt.expected.UTC()) {
					t.Errorf("ToTimeOrDefault() = %v, want %v", got.UTC(), tt.expected.UTC())
				}
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
		{
			name:    "empty string with empty layout",
			args:    args{"", ""},
			wantRes: time.Time{},
			wantErr: true,
		},
		{
			name:    "empty string with valid layout",
			args:    args{carbon.DateLayout, ""},
			wantRes: time.Time{},
			wantErr: true,
		},
		{
			name:    "valid string with empty layout",
			args:    args{"", "2022-07-02"},
			wantRes: time.Time{},
			wantErr: true,
		},
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
			gotRes, err := ToLayoutTimeE(tt.args.layout, tt.args.value)
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

func TestToTimeString(t *testing.T) {

	type args struct {
		t      time.Time
		format []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"FromDate", args{nowDate, []string{carbon.DateLayout}}, "2022-07-02", false},
		{"FromDate", args{nowDate, []string{"Y-m-d"}}, "2022-07-02", false},

		{"FromDate", args{nowDate, []string{"m-d-y"}}, "07-02-22", false},
		{"FromDate", args{nowDate, []string{"m-d-Y"}}, "07-02-2022", false},
		{"FromDate", args{nowDate, []string{"m-j-Y"}}, "07-2-2022", false},
		{"FromDate", args{nowDate, []string{"n-d-Y"}}, "7-02-2022", false},
		{"FromDate", args{nowDate, []string{"n-j-Y"}}, "7-2-2022", false},
		{"FromDate", args{nowDate, []string{"n-j-y"}}, "7-2-22", false},

		{"FromDate", args{nowDate, []string{"d/m/y"}}, "02/07/22", false},
		{"FromDate", args{nowDate, []string{"d/m/Y"}}, "02/07/2022", false},
		{"FromDate", args{nowDate, []string{"j/m/Y"}}, "2/07/2022", false},
		{"FromDate", args{nowDate, []string{"d/n/Y"}}, "02/7/2022", false},
		{"FromDate", args{nowDate, []string{"j/n/Y"}}, "2/7/2022", false},
		{"FromDate", args{nowDate, []string{"j/n/y"}}, "2/7/22", false},

		{"FromDate", args{nowDateTime, []string{carbon.DateTimeLayout}}, "2022-07-02 11:45:02", false},
		{"FromDate", args{nowDateTime, []string{"Y-m-d h:i:s"}}, "2022-07-02 11:45:02", false},

		{"FromDate", args{nowDateTime, []string{"m-d-y h:i:s"}}, "07-02-22 11:45:02", false},
		{"FromDate", args{nowDateTime, []string{"m-d-Y H:i:s"}}, "07-02-2022 11:45:02", false},
		{"FromDate", args{nowDateTime, []string{"m-j-Y h:i:s"}}, "07-2-2022 11:45:02", false},
		{"FromDate", args{nowDateTime, []string{"n-d-Y h:i:s"}}, "7-02-2022 11:45:02", false},
		{"FromDate", args{nowDateTime, []string{"n-j-Y h:i:s"}}, "7-2-2022 11:45:02", false},
		{"FromDate", args{nowDateTime, []string{"n-j-y h:i:s"}}, "7-2-22 11:45:02", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToTimeStringE(tt.args.t, tt.args.format...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTimeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToTimeString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTimeE(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr error
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: ErrEmptyString,
		},
		{
			name:    "valid time",
			input:   now,
			wantErr: nil,
		},
		{
			name:    "valid time pointer",
			input:   &now,
			wantErr: nil,
		},
		{
			name:    "valid date string",
			input:   "2022-07-02",
			wantErr: nil,
		},
		{
			name:    "valid datetime string",
			input:   "2022-07-02 11:45:02",
			wantErr: nil,
		},
		{
			name:    "invalid date string",
			input:   "not-a-date",
			wantErr: carbon.Parse("not-a-date").Error,
		},
		{
			name:    "nil input",
			input:   nil,
			wantErr: carbon.Parse("").Error,
		},
		{
			name:    "integer input",
			input:   42,
			wantErr: carbon.Parse("42").Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToTimeE(tt.input)
			if tt.wantErr == nil && err != nil {
				t.Errorf("ToTimeE() error = %v, wantErr nil", err)
				return
			}
			if tt.wantErr != nil && err == nil {
				t.Errorf("ToTimeE() error = nil, wantErr %v", tt.wantErr)
				return
			}
			if tt.wantErr == nil {
				if got.IsZero() {
					t.Errorf("ToTimeE() returned zero time for input %v", tt.input)
				}

				if inputTime, ok := tt.input.(time.Time); ok {
					if !got.Equal(inputTime) {
						t.Errorf("ToTimeE() = %v, want %v", got, inputTime)
					}
				}
				if inputTimePtr, ok := tt.input.(*time.Time); ok {
					if !got.Equal(*inputTimePtr) {
						t.Errorf("ToTimeE() = %v, want %v", got, *inputTimePtr)
					}
				}
			}
		})
	}
}
