package main

import (
	"walx/player"
  "github.com/rivo/tview"
)

type Queue struct {
  *tview.List
	items        []*player.AudioFile
	loop         bool
}

func (q *Queue) next() {
  currentIndex := q.GetCurrentItem()
  i := currentIndex + 1
  if currentIndex == q.GetItemCount() - 1 {
    i = 0
  }

  q.SetCurrentItem(i)
}

func (q *Queue) prev() {
  currentIndex := q.GetCurrentItem()
  if currentIndex != 0 {
    q.SetCurrentItem(currentIndex - 1)
  }
}

func makeQueue() *Queue {
  list := tview.NewList()
  queue := &Queue {
    List: list,
  }

  return queue
}
