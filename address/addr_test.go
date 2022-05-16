package address

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddress_Checksum(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, 11592},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, 58659},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, 28813},
		{"4", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, 24233},
		{"4", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 1, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, 54133},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.Checksum(); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_Data(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_Dump(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, "human-readable address: EQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nC1I isBounceable: false, isTestnetOnly: false, data.len: 32"},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, "human-readable address: EQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULOUj isBounceable: false, isTestnetOnly: false, data.len: 32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.Dump(); got != tt.want {
				t.Errorf("Dump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_FlagsToByte(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name      string
		fields    fields
		wantFlags byte
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, byte(0b00010001)},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, byte(0b00010001)},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, byte(0b01010001)},
		{"4", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, byte(0b01010001)},
		{"5", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, byte(0b10010001)},
		{"6", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, byte(0b10010001)},
		{"7", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, byte(0b11010001)},
		{"8", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, byte(0b11010001)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if gotFlags := a.FlagsToByte(); gotFlags != tt.wantFlags {
				t.Errorf("FlagsToByte() = %v, want %v", gotFlags, tt.wantFlags)
			}
		})
	}
}

func TestAddress_IsBounceable(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, false},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, false},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, true},
		{"4", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, true},
		{"5", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, true},
		{"6", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 2, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, true},
		{"5", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, true},
		{"6", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 2, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.IsBounceable(); got != tt.want {
				t.Errorf("IsBounceable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_IsTestnetOnly(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, false},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, false},
		{"3", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, true},
		{"4", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, true},
		{"5", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, true},
		{"6", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 2, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, true},
		{"5", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, true},
		{"6", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 2, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.IsTestnetOnly(); got != tt.want {
				t.Errorf("IsTestnetOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetBounce(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	type args struct {
		bouncable bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, args{bouncable: true}, true},
		{"2", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, args{bouncable: false}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			a.SetBounce(tt.args.bouncable)
			if got := a.IsBounceable(); got != tt.want {
				t.Errorf("IsBounceable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetTestnetOnly(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	type args struct {
		testnetOnly bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, args{testnetOnly: true}, true},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, args{testnetOnly: false}, false},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, args{testnetOnly: true}, true},
		{"4", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, args{testnetOnly: false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			a.SetTestnetOnly(tt.args.testnetOnly)
			if got := a.IsTestnetOnly(); got != tt.want {
				t.Errorf("IsTestnetOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_String(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, "EQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nC1I"},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, "EQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULOUj"},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, "UQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nHCN"},
		{"4", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, "UQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULLjm"},
		{"5", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, "kQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nJbC"},
		{"6", fields{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, "kQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULF6p"},
		{"7", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, "0QC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nMsH"},
		{"8", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, "0QCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULANs"},
		{"9", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, "0QG6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nEbb"},
		{"10", fields{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 1, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, "0QGTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULI6w"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_Workchain(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, byte(0)},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, byte(0)},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, byte(1)},
		{"4", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 1, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, byte(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.Workchain(); got != tt.want {
				t.Errorf("Workchain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_prepareChecksumData(t *testing.T) {
	type fields struct {
		flags     AddressFlags
		workchain byte
		data      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"1", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, []byte{17, 0, 186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}},
		{"2", fields{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, []byte{17, 0, 147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}},
		{"3", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, []byte{81, 1, 186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}},
		{"4", fields{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 1, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, []byte{81, 1, 147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				flags:     tt.fields.flags,
				workchain: tt.fields.workchain,
				data:      tt.fields.data,
			}
			if got := a.prepareChecksumData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareChecksumData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParseAddr(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *Address
	}{
		{"1", "EQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nC1I", &Address{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}},
		{"2", "EQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULOUj", &Address{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}},
		{"3", "UQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nHCN", &Address{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}},
		{"4", "UQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULLjm", &Address{flags: AddressFlags{bounceable: true, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}},
		{"5", "kQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nJbC", &Address{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}},
		{"6", "kQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULF6p", &Address{flags: AddressFlags{bounceable: false, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}},
		{"7", "0QC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nMsH", &Address{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}},
		{"8", "0QCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULANs", &Address{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}},
		{"9", "0QG6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nEbb", &Address{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 1, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}},
		{"10", "0QGTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULI6w", &Address{flags: AddressFlags{bounceable: true, testnet: true}, workchain: 1, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustParseAddr(tt.args); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("test: %#v", got.flags)
				t.Errorf("MustParseAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddressFromBytes(t *testing.T) {
	type args struct {
		flags     byte
		workchain byte
		data      []byte
	}
	tests := []struct {
		name string
		args args
		want *Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddress(tt.args.flags, tt.args.workchain, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddressFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAddr(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Address
		wantErr bool
	}{
		{"1", "EQC6KV4zs8TJtSZapOrRFmqSkxzpq-oSCoxekQRKElf4nC1I", &Address{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{186, 41, 94, 51, 179, 196, 201, 181, 38, 90, 164, 234, 209, 22, 106, 146, 147, 28, 233, 171, 234, 18, 10, 140, 94, 145, 4, 74, 18, 87, 248, 156}}, false},
		{"2", "EQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULOUj", &Address{flags: AddressFlags{bounceable: false, testnet: false}, workchain: 0, data: []byte{147, 13, 85, 51, 152, 10, 186, 17, 252, 216, 24, 69, 169, 84, 235, 245, 235, 42, 62, 31, 149, 112, 220, 29, 43, 146, 215, 34, 119, 63, 212, 44}}, false},
		{"err 1", "AQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULOUj", nil, true},
		{"err 2", "EQCTDVUzmAq6EfzYGEWpVOv16yo-H5Vw3B0rktcidz_ULOUB", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAddr(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAddr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFlags(t *testing.T) {
	type args struct {
		data byte
	}
	tests := []struct {
		name string
		args args
		want AddressFlags
	}{
		{"1", args{data: byte(0b00010001)}, AddressFlags{bounceable: false, testnet: false}},
		{"2", args{data: byte(0b01010001)}, AddressFlags{bounceable: true, testnet: false}},
		{"3", args{data: byte(0b10010001)}, AddressFlags{bounceable: false, testnet: true}},
		{"4", args{data: byte(0b11010001)}, AddressFlags{bounceable: true, testnet: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFlags(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}