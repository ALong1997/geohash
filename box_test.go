package geohash

import (
	"reflect"
	"testing"
)

func TestNewBox(t *testing.T) {
	type args struct {
		geohash  Geohash
		pointSet map[string]*Point
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		{
			name: "TestNewBox 1",
			args: args{
				geohash:  "",
				pointSet: nil,
			},
			want: nil,
		},
		{
			name: "TestNewBox 2",
			args: args{
				geohash:  "A",
				pointSet: nil,
			},
			want: nil,
		},
		{
			name: "TestNewBox 3",
			args: args{
				geohash: "WTW3SZYP",
				pointSet: map[string]*Point{
					"121.506377_31.245105": &Point{
						Lng: 121.506377,
						Lat: 31.245105,
						Val: "东方明珠",
					}},
			},
			want: &Box{
				Geohash: "WTW3SZYP",
				PointSet: map[string]*Point{
					"121.506377_31.245105": &Point{
						Lng: 121.506377,
						Lat: 31.245105,
						Val: "东方明珠",
					}},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBox(tt.args.geohash, tt.args.pointSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetGeohash(t *testing.T) {
	type fields struct {
		Geohash  Geohash
		PointSet map[string]*Point
	}
	tests := []struct {
		name   string
		fields fields
		want   Geohash
	}{
		{
			name: "TestBox_GetGeohash 1",
			fields: fields{
				Geohash:  "",
				PointSet: nil,
			},
			want: "",
		},
		{
			name: "TestBox_GetGeohash 2",
			fields: fields{
				Geohash: "WTW3SZYP",
				PointSet: map[string]*Point{
					"121.506377_31.245105": &Point{
						Lng: 121.506377,
						Lat: 31.245105,
						Val: "东方明珠",
					}},
			},
			want: "WTW3SZYP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Geohash:  tt.fields.Geohash,
				PointSet: tt.fields.PointSet,
			}
			if got := b.GetGeohash(); got != tt.want {
				t.Errorf("GetGeohash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetPointSet(t *testing.T) {
	type fields struct {
		Geohash  Geohash
		PointSet map[string]*Point
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]*Point
	}{
		{
			name: "TestBox_GetGeohash 1",
			fields: fields{
				Geohash:  "",
				PointSet: nil,
			},
			want: nil,
		},
		{
			name: "TestBox_GetGeohash 2",
			fields: fields{
				Geohash: "WTW3SZYP",
				PointSet: map[string]*Point{
					"121.506377_31.245105": &Point{
						Lng: 121.506377,
						Lat: 31.245105,
						Val: "东方明珠",
					}},
			},
			want: map[string]*Point{
				"121.506377_31.245105": &Point{
					Lng: 121.506377,
					Lat: 31.245105,
					Val: "东方明珠",
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Geohash:  tt.fields.Geohash,
				PointSet: tt.fields.PointSet,
			}
			if got := b.GetPointSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPointSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetAllPoints(t *testing.T) {
	type fields struct {
		Geohash  Geohash
		PointSet map[string]*Point
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Point
	}{
		{
			name: "TestBox_GetAllPoints 1",
			fields: fields{
				Geohash:  "",
				PointSet: nil,
			},
			want: []*Point{},
		},
		{
			name: "TestBox_GetAllPoints 2",
			fields: fields{
				Geohash: "WTW3SZYP",
				PointSet: map[string]*Point{
					"121.506377_31.245105": &Point{
						Lng: 121.506377,
						Lat: 31.245105,
						Val: "东方明珠",
					}},
			},
			want: []*Point{{
				Lng: 121.506377,
				Lat: 31.245105,
				Val: "东方明珠",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Geohash:  tt.fields.Geohash,
				PointSet: tt.fields.PointSet,
			}
			if got := b.GetAllPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_add(t *testing.T) {
	type fields struct {
		Geohash  Geohash
		PointSet map[string]*Point
	}
	type args struct {
		point *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestBox_add 1",
			fields: fields{
				Geohash:  "",
				PointSet: nil,
			},
			args: args{
				point: NewPoint(13.361389, 38.115556, "Palermo"),
			},
		},
		{
			name: "TestBox_add 2",
			fields: fields{
				Geohash: "WTW3SZYP",
				PointSet: map[string]*Point{
					"121.506377_31.245105": &Point{
						Lng: 121.506377,
						Lat: 31.245105,
						Val: "东方明珠",
					}},
			},
			args: args{
				point: NewPoint(13.361389, 38.115556, "Palermo"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Geohash:  tt.fields.Geohash,
				PointSet: tt.fields.PointSet,
			}
			b.add(tt.args.point)
		})
	}
}
