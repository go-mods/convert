package convert

import "testing"

type args struct {
	value interface{}
}

func TestToInt(t *testing.T) {

	type TestInt struct {
		name    string
		args    args
		wantRes int
		wantErr bool
	}

	tests := []TestInt{
		{name: "ToIntFromEmptyString", args: args{""}, wantRes: 0, wantErr: false},
		{name: "ToIntFromString", args: args{"1.2"}, wantRes: 1, wantErr: false},
		{name: "ToIntFromInt", args: args{int(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromInt8", args: args{int8(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromInt16", args: args{int16(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromInt32", args: args{int32(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromInt64", args: args{int64(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUint", args: args{uint(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUint8", args: args{uint8(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUint16", args: args{uint16(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUint32", args: args{uint32(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUint64", args: args{uint64(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUint64", args: args{uint64(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromUintPtr", args: args{uintptr(1.)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromFloat32", args: args{float32(1.2)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromFloat64", args: args{float64(1.2)}, wantRes: 1, wantErr: false},
		{name: "ToIntFromTrue", args: args{true}, wantRes: 1, wantErr: false},
		{name: "ToIntFromFalse", args: args{false}, wantRes: 0, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ToInt(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("ToInt() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
