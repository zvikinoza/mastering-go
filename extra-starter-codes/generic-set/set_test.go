package set

import (
	"fmt"
	"testing"
)

func populateSet(count int, start int) *ItemSet {
	set := ItemSet{}
	for i := 0; i < count; i++ {
		set.Add(fmt.Sprintf("item%d", start+i))
	}
	return &set
}

func TestAdd(t *testing.T) {
	set := populateSet(3, 0)
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	set.Add("item1") //should not add it, already there
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	set.Add("item4") //should not add it, already there
	if size := set.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}

func TestClear(t *testing.T) {
	set := populateSet(3, 0)
	set.Clear()
	if size := set.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func TestDelete(t *testing.T) {
	set := populateSet(3, 0)
	set.Delete("item2")
	if size := set.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestHas(t *testing.T) {
	set := populateSet(3, 0)
	has := set.Has("item2")
	if !has {
		t.Errorf("expected item2 to be there")
	}
	set.Delete("item2")
	has = set.Has("item2")
	if has {
		t.Errorf("expected item2 to be removed")
	}
	set.Delete("item1")
	has = set.Has("item1")
	if has {
		t.Errorf("expected item1 to be removed")
	}
}

func TestItems(t *testing.T) {
	set := populateSet(3, 0)
	items := set.Items()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	set = populateSet(520, 0)
	items = set.Items()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}

func TestSize(t *testing.T) {
	set := populateSet(3, 0)
	items := set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
	set = populateSet(0, 0)
	items = set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
	set = populateSet(10000, 0)
	items = set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
}

func TestUnion(t *testing.T) {
	s1 := populateSet(3, 0)
	s2 := populateSet(2, 3)
	s3 := s1.Union(s2)

	if s3.Size() != 5 {
		t.Errorf("worng count, expected 5 got %d", s3.Size())
	}

	if len(s1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", s1.Size())
	}

	if len(s2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", s2.Size())
	}
}

func TestIntersection(t *testing.T) {
	s1 := populateSet(3, 0)
	s2 := populateSet(2, 2)
	s3 := s1.Intersection(s2)

	if len(s3.Items()) != 1 {
		t.Errorf("worng count, expected 1 got %d", s3.Size())
	}

	if len(s1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", s1.Size())
	}

	if len(s2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", s2.Size())
	}
}

func TestDifference(t *testing.T) {
	s1 := populateSet(3, 0)
	s2 := populateSet(2, 2)
	s3 := s1.Difference(s2)

	if len(s3.Items()) != 2 {
		t.Errorf("worng count, expected 2 got %d", s3.Size())
	}

	if len(s1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", s1.Size())
	}

	if len(s2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", s2.Size())
	}
}

func TestSubset(t *testing.T) {
	s1 := populateSet(9, 0)
	s2 := populateSet(6, 0)

	if !s2.Subset(s1) {
		t.Errorf("expected true, got false")
	}
	if s1.Subset(s2) {
		t.Errorf("expected false, got true")
	}

	if len(s1.Items()) != 9 {
		t.Errorf("wrong count, expected 3 and got %d", s1.Size())
	}
	if len(s2.Items()) != 6 {
		t.Errorf("wrong count, expected 2 and got %d", s2.Size())
	}
}
