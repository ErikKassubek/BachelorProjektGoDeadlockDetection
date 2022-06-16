package otherExamples

/* sasha-s
import (
	"sync"

	"github.com/sasha-s/go-deadlock"
)

type Address1 int

type Mapping struct {
	mut deadlock.Mutex

	extAddresses map[string]Address1
}

func (m *Mapping) clearAddresses() {
	m.mut.Lock() // First locking
	var removed []Address1
	for id, addr := range m.extAddresses {
		removed = append(removed, addr)
		delete(m.extAddresses, id)
	}
	if len(removed) > 0 {
		m.notify(nil, removed)
	}
	m.mut.Unlock()
}

func (m *Mapping) notify(added, remove []Address1) {
	m.mut.Lock() // Second locking
	m.mut.Unlock()
}

type Service struct {
	mut sync.RWMutex

	mappings []*Mapping
}

func (s *Service) NewMapping() *Mapping {
	mapping := &Mapping{
		extAddresses: make(map[string]Address1),
	}
	s.mut.Lock()
	s.mappings = append(s.mappings, mapping)
	s.mut.Unlock()
	return mapping
}

func (s *Service) RemoveMapping(mapping *Mapping) {
	s.mut.Lock()
	defer s.mut.Unlock()
	for _, existing := range s.mappings {
		if existing == mapping {
			mapping.clearAddresses()
		}
	}
}

func NewService() *Service {
	return &Service{}
}

func RunSyncthing4829() {
	natSvc := NewService()
	m := natSvc.NewMapping()
	m.extAddresses["test"] = 0

	natSvc.RemoveMapping(m)
}
*/

/* deadlock-go */
import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
	"sync"
)

type Address1 int

type Mapping struct {
	mut deadlock.Mutex

	extAddresses map[string]Address1
}

func (m *Mapping) clearAddresses() {
	m.mut.Lock() // First locking
	var removed []Address1
	for id, addr := range m.extAddresses {
		removed = append(removed, addr)
		delete(m.extAddresses, id)
	}
	if len(removed) > 0 {
		m.notify(nil, removed)
	}
	m.mut.Unlock()
}

func (m *Mapping) notify(added, remove []Address1) {
	m.mut.Lock() // Second locking
	m.mut.Unlock()
}

type Service struct {
	mut sync.RWMutex

	mappings []*Mapping
}

func (s *Service) NewMapping() *Mapping {
	mapping := &Mapping{
		mut:          deadlock.NewLock(),
		extAddresses: make(map[string]Address1),
	}
	s.mut.Lock()
	s.mappings = append(s.mappings, mapping)
	s.mut.Unlock()
	return mapping
}

func (s *Service) RemoveMapping(mapping *Mapping) {
	s.mut.Lock()
	defer s.mut.Unlock()
	for _, existing := range s.mappings {
		if existing == mapping {
			mapping.clearAddresses()
		}
	}
}

func NewService() *Service {
	return &Service{}
}

func RunSyncthing4829() {
	natSvc := NewService()
	m := natSvc.NewMapping()
	m.extAddresses["test"] = 0

	natSvc.RemoveMapping(m)
}
