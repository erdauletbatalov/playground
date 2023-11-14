package test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// Бенчмарк для time.Now()
func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

// Бенчмарк для time.Now().UnixNano()
func BenchmarkTimeNowUnixNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now().UnixNano()
	}
}

func BenchmarkSample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

func BenchmarkSample2(b *testing.B) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

type testStruct struct {
	X int
	Y string
}

func (t *testStruct) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

func (t *testStruct) toJSON() ([]byte, error) {
	return []byte(`{"X": ` + strconv.Itoa(t.X) + `, "Y": "` + t.Y + `"}`), nil
}

func BenchmarkToJSON1(b *testing.B) {
	tmp := &testStruct{X: 1, Y: "string"}
	js, err := tmp.ToJSON()
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(js)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := tmp.ToJSON(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkToJSON2(b *testing.B) {
	tmp := &testStruct{X: 1, Y: "string"}
	js, err := tmp.toJSON()
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(js)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := tmp.ToJSON(); err != nil {
			b.Fatal(err)
		}
	}
}

type Set struct {
	set map[interface{}]struct{}
	mu  sync.Mutex
}

func (s *Set) Add(x interface{}) {
	s.mu.Lock()
	s.set[x] = struct{}{}
	s.mu.Unlock()
}

func (s *Set) Delete(x interface{}) {
	s.mu.Lock()
	delete(s.set, x)
	s.mu.Unlock()
}

func BenchmarkSetDelete(b *testing.B) {
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		set := Set{set: make(map[interface{}]struct{})}
		for _, elem := range testSet {
			set.Add(elem)
		}
		b.StartTimer()
		for _, elem := range testSet {
			set.Delete(elem)
		}
	}
}
