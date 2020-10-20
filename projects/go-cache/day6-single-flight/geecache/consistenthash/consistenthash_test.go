package consistenthash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	hash := NewMap(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	hash.Add("8")

	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}

func TestMap_Add(t *testing.T) {
	hash := NewMap(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	hash.Add("1","3","5")

	testCase := []int{
		1,3,5,11,13,15,21,23,25,
	}

	actual := hash.keys
	fmt.Println("actual keys: ", actual)

	for k, v := range testCase {
		if actual[k] != v {
			t.Errorf("the node is not correct! %d", v)
		}
	}
}