package queue

import (
	"testing"
)

func TestAdd(t *testing.T) {
	q := newQueue(3)

	if err := q.Add(0, 10); err != nil {
		t.Fatal(err)
	}

	if q.data[0] != 10 {
		t.Fatal("invalid value")
	}
}

func TestAddDuplicate(t *testing.T) {
	q := newQueue(3)

	if err := q.Add(0, 10); err != nil {
		t.Fatal(err)
	}

	// duplicate key
	if err := q.Add(0, 10); err != ErrKeyExists {
		t.Fatal("key already exists")
	}
}

func TestAddFull(t *testing.T) {
	q := newQueue(3)

	go func() {
		// fill the bucket and
		// add 2 extra items
		for i := 0; i < q.size+2; i++ {
			key := int64(i)
			q.Add(key, i*10)
		}
	}()

	for i := 0; i < 2; i++ {
		if val := <-q.Out(); val != i*10 {
			t.Fatal("invalid value")
		}
	}
}
