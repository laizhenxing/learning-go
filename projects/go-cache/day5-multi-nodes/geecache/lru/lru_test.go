package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	v, ok := lru.Get("key1")
	t.Logf("v := %v", v)
	if  !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed!\n")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache missing key2 failed\n")
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	k1,k2,k3 := "k1","k2","k3"
	v1,v2,v3 := "v1","v2","v3"
	cap := len(k1+v1+k2+v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("k1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest k1 failed!")
	}
}

func TestCache_OnEvicted(t *testing.T)  {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := New(int64(10), callback)
	lru.Add("k1", String("1"))
	lru.Add("k2", String("2"))
	lru.Add("k3", String("3"))
	lru.Add("k4", String("4"))

	expectd := []string{"k1"}

	if !reflect.DeepEqual(expectd, keys) {
		t.Fatalf("Call OnEvicted failed, expected keys equal to %s", keys)
	}
}
