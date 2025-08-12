// Copyright 2025 Bytedance Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package heapsort

import "github.com/bytedance/gg/internal/constraints"

func siftDown[T constraints.Ordered](v []T, less func(i, j int) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(first+child, first+child+1) {
			child++
		}
		if !less(first+root, first+child) {
			return
		}
		v[first+root], v[first+child] = v[first+child], v[first+root]
		root = child
	}
}

func buildHeap[T constraints.Ordered](v []T, less func(i, j int) bool, a, b int) {
	first := a
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(v, less, i, hi, first)
	}
}

func heapSort[T constraints.Ordered](v []T, less func(i, j int) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(v, less, i, hi, first)
	}
	for i := hi - 1; i >= 0; i-- {
		v[first], v[first+i] = v[first+i], v[first]
		siftDown(v, less, lo, i, first)
	}
}

func partialSort[T constraints.Ordered](v []T, less func(i, j int) bool, k int) {
	n := len(v)
	if k <= 0 || n <= 1 {
		return
	}
	if k >= n {
		heapSort(v, less, 0, n)
		return
	}
	buildHeap(v, less, 0, k)
	for i := k; i < n; i++ {
		if less(i, 0) {
			v[0], v[i] = v[i], v[0]
			siftDown(v, less, 0, k, 0)
		}
	}
	heapSort(v, less, 0, k)
}

func Sort[T constraints.Ordered](v []T) {
	if len(v) <= 1 {
		return
	}
	heapSort(v, func(i, j int) bool { return v[i] < v[j] }, 0, len(v))
}

func PartialSort[T constraints.Ordered](k int, v []T) {
	PartialSortBy(func(a, b T) bool { return a < b }, k, v)
}

func PartialSortBy[T constraints.Ordered](less func(a, b T) bool, k int, v []T) {
	n := len(v)
	if k <= 0 || n <= 1 {
		return
	}
	if k >= n {
		heapSort(v, func(i, j int) bool { return less(v[i], v[j]) }, 0, n)
		return
	}
	partialSort(v, func(i, j int) bool { return less(v[i], v[j]) }, k)
}
