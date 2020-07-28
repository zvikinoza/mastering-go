package set

import "sync"

type Type interface{}

type Item Type

type ItemSet struct {
	items map[Item]bool
	lock  sync.RWMutex
}

func (s *ItemSet) Add(t Item) *ItemSet {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.items == nil {
		s.items = make(map[Item]bool)
	}

	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

func (s *ItemSet) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = make(map[Item]bool)
}

func (s *ItemSet) Delete(t Item) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.items[t]
	if ok {
		delete(s.items, t)
	}
	return ok
}

func (s *ItemSet) Has(t Item) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.items[t]
	return ok
}

func (s *ItemSet) Items() []Item {
	s.lock.RLock()
	defer s.lock.RUnlock()

	items := []Item{}
	for k := range s.items {
		items = append(items, k)
	}
	return items
}

func (s *ItemSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.items)
}

func (s *ItemSet) Union(o *ItemSet) *ItemSet {
	us := ItemSet{}
	us.items = make(map[Item]bool)

	s.lock.RLock()
	for k := range s.items {
		us.items[k] = true
	}
	s.lock.RUnlock()

	o.lock.RLock()
	for k := range o.items {
		_, ok := us.items[k]
		if !ok {
			us.items[k] = true
		}
	}
	o.lock.RUnlock()

	return &us
}

func (s *ItemSet) Intersection(o *ItemSet) *ItemSet {
	s.lock.RLock()
	o.lock.RLock()
	defer s.lock.RUnlock()
	defer o.lock.RUnlock()

	is := ItemSet{}
	is.items = make(map[Item]bool)
	for k := range s.items {
		_, ok := o.items[k]
		if ok {
			is.items[k] = true
		}
	}

	return &is
}

func (s *ItemSet) Difference(o *ItemSet) *ItemSet {
	s.lock.RLock()
	o.lock.RLock()
	defer s.lock.RUnlock()
	defer o.lock.RUnlock()

	ds := ItemSet{}
	ds.items = make(map[Item]bool)
	for k := range s.items {
		_, ok := o.items[k]
		if !ok {
			ds.items[k] = true
		}
	}
	return &ds
}

func (s *ItemSet) Subset(o *ItemSet) bool {
	s.lock.RLock()
	o.lock.RLock()
	defer s.lock.RUnlock()
	defer o.lock.RUnlock()

	for k := range s.items {
		_, ok := o.items[k]
		if !ok {
			return false
		}
	}
	return true
}
