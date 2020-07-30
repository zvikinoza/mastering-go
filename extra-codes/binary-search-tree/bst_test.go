package bst

import (
	"fmt"
	"testing"
)

func stringComparator(a, b Item) int {
	s1 := a.(string)
	s2 := b.(string)
	if s1 < s2 {
		return -1
	}
	if s1 == s2 {
		return 0
	}
	return 1
}

func populateSet(count int, start int) *BST {
	set := BST{comparator: stringComparator}
	for i := 0; i < count/2; i++ {
		set.Add(fmt.Sprintf("item%d", start+i*2))
	}
	for i := 0; i < (count+1)/2; i++ {
		set.Add(fmt.Sprintf("item%d", start+i*2+1))
	}
	return &set
}

func TestAdd(t *testing.T) {
	set := populateSet(3, 0)

	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
		t.Errorf("inoreder tree view=%v", set.Items())
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set.Add("item1") //should not add it, already there
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set.Add("item4") //should not add it, already there
	if size := set.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
		t.Errorf("inoreder tree view=%v", set.Items())
	}
}

func TestClear(t *testing.T) {
	set := populateSet(3, 0)
	set.Clear()
	if size := set.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
		t.Errorf("inoreder tree view=%v", set.Items())
	}
}

func TestDelete(t *testing.T) {
	set := populateSet(3, 0)
	set.Delete("item0")
	if size := set.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
		t.Errorf("inoreder tree view=%v", set.Items())
	}
}

func TestHas(t *testing.T) {
	set := populateSet(3, 0)
	has := set.Has("item0")
	if !has {
		t.Errorf("expected item0 to be there")
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set.Delete("item0")

	has = set.Has("item0")
	if has {
		t.Errorf("expected item0 to be removed")
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set.Delete("item1")
	has = set.Has("item1")
	if has {
		t.Errorf("expected item1 to be removed")
		t.Errorf("inoreder tree view=%v", set.Items())
	}
}

func TestItems(t *testing.T) {
	set := populateSet(3, 0)
	items := set.Items()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set = populateSet(520, 0)
	items = set.Items()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
		t.Errorf("inoreder tree view=%v", set.Items())
	}
}

func TestSize(t *testing.T) {
	set := populateSet(3, 0)
	items := set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set = populateSet(0, 0)
	items = set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
		t.Errorf("inoreder tree view=%v", set.Items())
	}
	set = populateSet(10000, 0)
	items = set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
		t.Errorf("inoreder tree view=%v", set.Items())
	}
}
