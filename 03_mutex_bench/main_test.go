package main

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Mutex structure

type Set struct {
	sync.Mutex
	mm map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[string]struct{}{},
	}
}

func (s *Set) Add(i float32) {
	s.Lock()
	str := fmt.Sprintf("%.6f", i)
	s.mm[str] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i float32) bool {
	str := fmt.Sprintf("%.6f", i)
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[str]
	return ok
}

// sync.RWMutex structure

type SetRW struct {
	sync.RWMutex
	mm map[string]struct{}
}

func NewSetRW() *SetRW {
	return &SetRW{
		mm: map[string]struct{}{},
	}
}

func (s *SetRW) Add(i float32) {
	str := fmt.Sprintf("%.6f", i)
	s.Lock()
	s.mm[str] = struct{}{}
	s.Unlock()
}

func (s *SetRW) Has(i float32) bool {
	str := fmt.Sprintf("%.6f", i)
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[str]
	return ok
}

// benchmarks
// 10% write 90% reads sync.Mutex benchmark
func Benchmark10w90rMutex(b *testing.B) {
	var set = NewSet()
	b.Run("10% write, 90% read Mutex", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(123.456)

				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
				set.Has(123.456)
			}
		})
	})
}

// 10% write 90% reads sync.RWMutex benchmark
func Benchmark10w90rRWMut(b *testing.B) {
	var setRW = NewSetRW()
	b.Run("10% write, 90% read RWMutex", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setRW.Add(123.456)

				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
				setRW.Has(123.456)
			}
		})
	})
}

// 50% write 50% reads sync.Mutex benchmark
func Benchmark50w50rMutex(b *testing.B) {
	var set = NewSet()
	b.Run("50% write 50% reads Mutex", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(123.456)
				set.Add(123.4561)
				set.Add(123.4562)
				set.Add(123.4563)
				set.Add(123.4564)

				set.Has(123.456)
				set.Has(123.4561)
				set.Has(123.4562)
				set.Has(123.4563)
				set.Has(123.4564)
			}
		})
	})
}

// 50% write 50% reads sync.RWMutex benchmark
func Benchmark50w50rRWMut(b *testing.B) {
	var setRW = NewSetRW()
	b.Run("50% write 50% reads RWMutex", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setRW.Add(123.456)
				setRW.Add(123.4561)
				setRW.Add(123.4562)
				setRW.Add(123.4563)
				setRW.Add(123.4564)

				setRW.Has(123.456)
				setRW.Has(123.4561)
				setRW.Has(123.4562)
				setRW.Has(123.4563)
				setRW.Has(123.4564)
			}
		})
	})
}

// 90% write 10% reads sync.Mutex benchmark
func Benchmark90w10rMutex(b *testing.B) {
	var set = NewSet()
	b.Run("90% write 10% reads Mutex", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(123.456)
				set.Add(123.4561)
				set.Add(123.4562)
				set.Add(123.4563)
				set.Add(123.4564)
				set.Add(123.4565)
				set.Add(123.4566)
				set.Add(123.4567)
				set.Add(123.4568)

				set.Has(123.456)
			}
		})
	})
}

// 90% write 10% reads sync.RWMutex benchmark
func Benchmark90w10rRWMut(b *testing.B) {
	var setRW = NewSetRW()
	b.Run("90% write 10% reads RWMutex", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setRW.Add(123.456)
				setRW.Add(123.4561)
				setRW.Add(123.4562)
				setRW.Add(123.4563)
				setRW.Add(123.4564)
				setRW.Add(123.4565)
				setRW.Add(123.4566)
				setRW.Add(123.4567)
				setRW.Add(123.4568)

				setRW.Has(123.456)
			}
		})
	})
}
