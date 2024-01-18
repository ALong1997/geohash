package geohash

import (
	"reflect"
	"testing"
)

func TestGeohash_valid(t *testing.T) {
	tests := []struct {
		name string
		g    Geohash
		want bool
	}{
		{
			name: "TestGeohash_valid 1",
			g:    "",
			want: false,
		},
		{
			name: "TestGeohash_valid 2",
			g:    "C",
			want: false,
		},
		{
			name: "TestGeohash_valid 3",
			g:    "ABCDEFGH",
			want: false,
		},
		{
			name: "TestGeohash_valid 4",
			g:    "12345678",
			want: true,
		},
		{
			name: "TestGeohash_valid 5",
			g:    "WTW3SZYP",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.valid(); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPoint(t *testing.T) {
	type args struct {
		lng float64
		lat float64
		val any
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "TestNewPoint 1",
			args: args{
				lng: 0,
				lat: 0,
				val: nil,
			},
			want: &Point{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
		},
		{
			name: "TestNewPoint 1",
			args: args{
				lng: 121.506377,
				lat: 31.245105,
				val: "东方明珠",
			},
			want: &Point{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint(tt.args.lng, tt.args.lat, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_GetLng(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "TestPoint_GetLng 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			want: 0,
		},
		{
			name: "TestPoint_GetLng 2",
			fields: fields{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
			want: 121.506377,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if got := p.GetLng(); got != tt.want {
				t.Errorf("GetLng() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_GetLat(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{

		{
			name: "TestPoint_GetLat 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			want: 0,
		},
		{
			name: "TestPoint_GetLat 2",
			fields: fields{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
			want: 31.245105,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if got := p.GetLat(); got != tt.want {
				t.Errorf("GetLat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_GetAny(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	tests := []struct {
		name    string
		fields  fields
		wantVal any
	}{
		{
			name: "TestPoint_GetVal 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			wantVal: nil,
		},
		{
			name: "TestPoint_GetVal 2",
			fields: fields{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
			wantVal: "东方明珠",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if gotVal := p.GetAny(); !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("GetAny() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestPoint_Distance(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	type args struct {
		target *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
		{
			name: "TestPoint_Distance 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			args: args{target: NewPoint(1, 1, nil)},
			want: 157249,
		},
		{
			name: "TestPoint_Distance 2",
			fields: fields{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
			args: args{target: NewPoint(121.4871639, 31.2388556, "上海和平饭店")},
			want: 1954,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if got := p.Distance(tt.args.target); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Geohash(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	tests := []struct {
		name   string
		fields fields
		want   Geohash
	}{
		{
			name: "TestPoint_Geohash 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			want: "S0000000",
		},
		{
			name: "TestPoint_Geohash 2",
			fields: fields{
				Lng: 13.361389,
				Lat: 38.115556,
				Val: "Palermo",
			},
			want: "SQC8B49R",
		},
		{
			name: "TestPoint_Geohash 3",
			fields: fields{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
			want: "WTW3SZYP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if got := p.Geohash(); got != tt.want {
				t.Errorf("Geohash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_key(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestPoint_key 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			want: "0_0",
		},
		{
			name: "TestPoint_key 2",
			fields: fields{
				Lng: 13.361389,
				Lat: 38.115556,
				Val: "Palermo",
			},
			want: "13.361389_38.115556",
		},
		{
			name: "TestPoint_key 3",
			fields: fields{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			},
			want: "121.506377_31.245105",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if got := p.key(); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_circumscribedSquarePointsByCircle(t *testing.T) {
	type fields struct {
		Lng float64
		Lat float64
		Val any
	}
	type args struct {
		radius uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [9]*Point
	}{
		{
			name: "TestPoint_circumscribedSquarePointsByCircle 1",
			fields: fields{
				Lng: 0,
				Lat: 0,
				Val: nil,
			},
			args: args{radius: 111},
			want: [9]*Point{
				&Point{-0.001, 0.001, nil},
				&Point{0, 0.001, nil},
				&Point{0.001, 0.001, nil},
				&Point{-0.001, 0, nil},
				&Point{0, 0, nil},
				&Point{0.001, 0, nil},
				&Point{-0.001, -0.001, nil},
				&Point{0, -0.001, nil},
				&Point{0.001, -0.001, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				Lng: tt.fields.Lng,
				Lat: tt.fields.Lat,
				Val: tt.fields.Val,
			}
			if got := p.circumscribedSquarePointsByCircle(tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circumscribedSquarePointsByCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encode(t *testing.T) {
	type args struct {
		coordinate float64
		start      float64
		end        float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_encode 1",
			args: args{
				coordinate: 0,
				start:      minLng,
				end:        maxLng,
			},
			want: "10000000000000000000",
		},
		{
			name: "Test_encode 2",
			args: args{
				coordinate: 0,
				start:      minLat,
				end:        maxLat,
			},
			want: "10000000000000000000",
		},
		{
			name: "Test_encode 3",
			args: args{
				coordinate: 121.506377,
				start:      minLng,
				end:        maxLng,
			},
			want: "11010110011001111000",
		},
		{
			name: "Test_encode 4",
			args: args{
				coordinate: 31.245105,
				start:      minLat,
				end:        maxLat,
			},
			want: "10101100011011111111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.coordinate, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		bit byte
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "Test_decode 1",
			args: args{
				bit: '0',
			},
			want: 0,
		},
		{
			name: "Test_decode 2",
			args: args{
				bit: '1',
			},
			want: 1,
		},
		{
			name: "Test_decode 3",
			args: args{
				bit: 'B',
			},
			want: 10,
		},
		{
			name: "Test_decode 4",
			args: args{
				bit: 'a',
			},
			want: invalidCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.bit); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGeohashLenByDiameter(t *testing.T) {
	type args struct {
		diameter uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "Test_getGeohashLenByDiameter 1",
			args: args{
				diameter: 10,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "Test_getGeohashLenByDiameter 2",
			args: args{
				diameter: 100,
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "Test_getGeohashLenByDiameter 3",
			args: args{
				diameter: 10000,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "Test_getGeohashLenByDiameter 4",
			args: args{
				diameter: 10000000,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGeohashLenByDiameter(tt.args.diameter)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGeohashLenByDiameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getGeohashLenByDiameter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
