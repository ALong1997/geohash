package geohash

import (
	"reflect"
	"testing"
)

func TestNewTrie(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		{
			name: "",
			want: &Trie{root: &node{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrie(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Get(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	t1.Run("TestTrie_Get", func(t1 *testing.T) {
		_, got := t.Get(NewPoint(121.4871639, 31.2388556, "上海和平饭店").Geohash())
		if got != false {
			t1.Errorf("Get() want %v", false)
		}
		got1, got2 := t.Get(p1.Geohash())
		if got2 != true {
			t1.Errorf("Get() got1 = %v, want %v", got1, true)
		}
	})
}

func TestTrie_GetBoxesByPrefix(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	t1.Run("TestTrie_GetBoxesByPrefix", func(t1 *testing.T) {
		want := []*Box{NewBox(p1.Geohash(), map[string]*Point{p1.key(): p1})}
		if got := t.GetByPrefix(string(p1.Geohash())[:7]); !reflect.DeepEqual(got, want) {
			t1.Errorf("GetByPrefix() = %v, want %v", got, want)
		}
		if got := t.GetByPrefix("?"); got != nil {
			t1.Errorf("GetByPrefix() = %v, want %v", got, nil)
		}
	})
}

func TestTrie_Put(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
}

func TestTrie_Delete(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	t1.Run("TestTrie_Delete", func(t1 *testing.T) {
		if got := t.Delete("SQC8B49R"); got != true {
			t1.Errorf("Delete() = %v, want %v", got, true)
		}
		if got := t.Delete("WTW3SZYP"); got != true {
			t1.Errorf("Delete() = %v, want %v", got, true)
		}
		if got := t.Delete("WTW3SZYP"); got != false {
			t1.Errorf("Delete() = %v, want %v", got, false)
		}
	})
}

func TestTrie_GetPointsByCircle(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	t1.Run("TestTrie_GetPointsByCircle 1", func(t1 *testing.T) {
		got, err := t.GetPointsByCircle(p1, 1)
		if err != nil {
			t1.Errorf("GetPointsByCircle() error = %v, wantErr %v", err, nil)
			return
		}
		if !reflect.DeepEqual(got, []*Point{p1}) {
			t1.Errorf("GetPointsByCircle() got = %v, want %v", got, []*Point{p1})
		}
	})
	t1.Run("TestTrie_GetPointsByCircle 2", func(t1 *testing.T) {
		got, err := t.GetPointsByCircle(NewPoint(0, 0, nil), 1000)
		if err != nil {
			t1.Errorf("GetPointsByCircle() error = %v, wantErr %v", err, nil)
			return
		}
		if !reflect.DeepEqual(got, []*Point{}) {
			t1.Errorf("GetPointsByCircle() got = %v, want %v", got, []*Point{})
		}
	})
	t1.Run("TestTrie_GetPointsByCircle 3", func(t1 *testing.T) {
		got, err := t.GetPointsByCircle(NewPoint(121.4871639, 31.2388556, "上海和平饭店"), 10000)
		if err != nil {
			t1.Errorf("GetPointsByCircle() error = %v, wantErr %v", err, nil)
			return
		}
		if !reflect.DeepEqual(got, []*Point{p2}) {
			t1.Errorf("GetPointsByCircle() got = %v, want %v", got, []*Point{p2})
		}
	})
}

func TestTrie_search(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	t1.Run("TestTrie_search", func(t1 *testing.T) {
		got := t.search("BB")
		if got != nil {
			t1.Errorf("Get() got = %v, want %v", got, nil)
		}
		got1 := t.search("SQC8B49R")
		if !reflect.DeepEqual(got1, &node{
			children:  [32]*node{},
			passCount: 0,
			isLeaf:    true,
			Box:       NewBox("SQC8B49R", map[string]*Point{p1.key(): p1}),
		}) {
			t1.Errorf("Get() got1 = %v, want %v", got, p1)
		}
	})
}

func Test_node_dfs(t1 *testing.T) {
	t := NewTrie()
	p1 := NewPoint(13.361389, 38.115556, "Palermo")
	p2 := NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	t1.Run("Test_node_dfs", func(t1 *testing.T) {
		got := t.root.dfs()
		want := []*Box{
			NewBox("SQC8B49R", map[string]*Point{p1.key(): p1}),
			NewBox("WTW3SZYP", map[string]*Point{p2.key(): p2}),
		}
		if !reflect.DeepEqual(got, want) {
			t1.Errorf("Get() got = %v, want %v", got, want)
		}
	})
}
