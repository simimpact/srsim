package queue

import (
	"container/heap"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type minHeap []Task

type Task struct {
	Execute    func()
	Priority   info.InsertPriority
	AbortFlags []model.BehaviorFlag
	Source     key.TargetID
	id         int
}

type Handler struct {
	actions *minHeap
	counter int
}

func New() *Handler {
	heap := make(minHeap, 0, 10)
	return &Handler{
		actions: &heap,
		counter: 0,
	}
}

func (s *Handler) Insert(t Task) {
	t.id = s.counter
	heap.Push(s.actions, t)
	s.counter += 1
}

func (s *Handler) Pop() Task {
	return heap.Pop(s.actions).(Task)
}

func (s *Handler) IsEmpty() bool {
	return s.actions.Len() == 0
}

// --- min heap functions

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority || (h[i].Priority == h[j].Priority && h[i].id < h[j].id)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
