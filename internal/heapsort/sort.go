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

import (
	"github.com/bytedance/gg/internal/constraints"
)

// siftDown implements the heap property on v[lo:hi].
func siftDown[T constraints.Ordered](v []T, node int) {
	for {
		child := 2*node + 1
		if child >= len(v) {
			break
		}
		if child+1 < len(v) && v[child] < v[child+1] {
			child++
		}
		if v[node] >= v[child] {
			return
		}
		v[node], v[child] = v[child], v[node]
		node = child
	}
}

func siftDownBy[T any](v []T, node int, less func(i, j int) bool) {
	for {
		child := 2*node + 1
		if child >= len(v) {
			break
		}
		if child+1 < len(v) && less(child+1, child) {
			child++
		}
		if !less(child, node) {
			return
		}
		v[node], v[child] = v[child], v[node]
		node = child
	}
}

func Sort[T constraints.Ordered](v []T) {
	// Build heap with the greatest element at the top.
	for i := (len(v) - 1) / 2; i >= 0; i-- {
		siftDown(v, i)
	}

	// Pop elements into end of v.
	for i := len(v) - 1; i >= 1; i-- {
		v[0], v[i] = v[i], v[0]
		siftDown(v[:i], 0)
	}
}

func sortBy[T any](v []T, less func(T, T) bool) {
	// Build max-heap using the reverse of less.
	heapLess := func(i, j int) bool {
		// For max-heap, compare as greater-than.
		return less(v[j], v[i])
	}

	// Heapify phase: build max-heap
	for i := (len(v) - 1) / 2; i >= 0; i-- {
		siftDownBy(v, i, heapLess)
	}

	// Sort-down phase: move max to the end
	for i := len(v) - 1; i >= 1; i-- {
		v[0], v[i] = v[i], v[0]
		siftDownBy(v[:i], 0, heapLess)
	}
}

func PartialSort[T constraints.Ordered](k int, v []T) {
	n := len(v)

	if k <= 0 {
		return
	}

	if k >= n {
		Sort(v)
		return
	}

	// Build a max-heap from the first k elements
	for j := (k - 1) / 2; j >= 0; j-- {
		siftDown(v[:k], j)
	}

	// Iterate through the rest of the slice
	for j := k; j < n; j++ {
		if v[j] < v[0] {
			v[0], v[j] = v[j], v[0]
			siftDown(v[:k], 0)
		}
	}

	// Sort the heap to get the final k smallest elements in order
	Sort(v[:k])
}

func PartialSortBy[T any](less func(T, T) bool, k int, v []T) {
	n := len(v)

	if k <= 0 {
		return
	}

	if k >= n {
		sortBy(v, less)
		return
	}

	// Build max-heap of first k elements using reverse comparator
	heapLess := func(i, j int) bool {
		return less(v[j], v[i]) // max-heap: "greater" goes up
	}

	for j := (k - 1) / 2; j >= 0; j-- {
		siftDownBy(v[:k], j, heapLess)
	}

	for j := k; j < n; j++ {
		if less(v[j], v[0]) {
			v[0], v[j] = v[j], v[0]
			siftDownBy(v[:k], 0, heapLess)
		}
	}

	// Sort top k using the original comparator
	sortBy(v[:k], less)
}
