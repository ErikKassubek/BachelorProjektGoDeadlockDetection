/*
 * Project: cockroach
 * Issue or PR  : https://github.com/cockroachdb/cockroach/pull/10214
 * Buggy version: 7207111aa3a43df0552509365fdec741a53f873f
 * fix commit-id: 27e863d90ab0660494778f1c35966cc5ddc38e32
 * Flaky: 3/100
 * Description: This deadlock is caused by different order when acquiring
 * coalescedMu.Lock() and raftMu.Lock(). The fix is to refactor sendQueuedHeartbeats()
 * so that cockroachdb can unlock coalescedMu before locking raftMu.
 */
package otherExamples

/* sasha-s
import (
	"unsafe"

	"github.com/sasha-s/go-deadlock"
)

type Store1 struct {
	coalescedMu struct {
		deadlock.Mutex
		heartbeatResponses []int
	}
	mu struct {
		replicas map[int]*Replica1
	}
}

func (s *Store1) sendQueuedHeartbeats() {
	s.coalescedMu.Lock()         // LockA acquire
	defer s.coalescedMu.Unlock() // LockA release
	for i := 0; i < len(s.coalescedMu.heartbeatResponses); i++ {
		s.sendQueuedHeartbeatsToNode() // LockB
	}
}

func (s *Store1) sendQueuedHeartbeatsToNode() {
	for i := 0; i < len(s.mu.replicas); i++ {
		r := s.mu.replicas[i]
		r.reportUnreachable() // LockB
	}
}

type Replica1 struct {
	raftMu deadlock.Mutex
	mu     deadlock.Mutex
	store  *Store1
}

func (r *Replica1) reportUnreachable() {
	r.raftMu.Lock() // LockB acquire
	//+time.Sleep(time.Nanosecond)
	defer r.raftMu.Unlock()
	// LockB release
}

func (r *Replica1) tick() {
	r.raftMu.Lock() // LockB acquire
	defer r.raftMu.Unlock()
	r.tickRaftMuLocked()
	// LockB release
}

func (r *Replica1) tickRaftMuLocked() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.maybeQuiesceLocked() {
		return
	}
}
func (r *Replica1) maybeQuiesceLocked() bool {
	for i := 0; i < 2; i++ {
		if !r.maybeCoalesceHeartbeat() {
			return true
		}
	}
	return false
}
func (r *Replica1) maybeCoalesceHeartbeat() bool {
	msgtype := uintptr(unsafe.Pointer(r)) % 3
	switch msgtype {
	case 0, 1, 2:
		r.store.coalescedMu.Lock() // LockA acquire
	default:
		return false
	}
	r.store.coalescedMu.Unlock() // LockA release
	return true
}

func RunCockroach10214() {
	store := &Store1{}
	responses := &store.coalescedMu.heartbeatResponses
	*responses = append(*responses, 1, 2)
	store.mu.replicas = make(map[int]*Replica1)

	rp1 := &Replica1{
		store: store,
	}
	rp2 := &Replica1{
		store: store,
	}
	store.mu.replicas[0] = rp1
	store.mu.replicas[1] = rp2

	go func() {
		store.sendQueuedHeartbeats()
	}()

	go func() {
		rp1.tick()
	}()
}
*/

/* deadlock-go */
import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
	"unsafe"
)

type Store1 struct {
	coalescedMu struct {
		mu                 deadlock.Mutex
		heartbeatResponses []int
	}
	mu struct {
		replicas map[int]*Replica1
	}
}

func (s *Store1) sendQueuedHeartbeats() {
	s.coalescedMu.mu.Lock()         // LockA acquire
	defer s.coalescedMu.mu.Unlock() // LockA release
	for i := 0; i < len(s.coalescedMu.heartbeatResponses); i++ {
		s.sendQueuedHeartbeatsToNode() // LockB
	}
}

func (s *Store1) sendQueuedHeartbeatsToNode() {
	for i := 0; i < len(s.mu.replicas); i++ {
		r := s.mu.replicas[i]
		r.reportUnreachable() // LockB
	}
}

type Replica1 struct {
	raftMu deadlock.Mutex
	mu     deadlock.Mutex
	store  *Store1
}

func (r *Replica1) reportUnreachable() {
	r.raftMu.Lock() // LockB acquire
	//+time.Sleep(time.Nanosecond)
	defer r.raftMu.Unlock()
	// LockB release
}

func (r *Replica1) tick() {
	r.raftMu.Lock() // LockB acquire
	defer r.raftMu.Unlock()
	r.tickRaftMuLocked()
	// LockB release
}

func (r *Replica1) tickRaftMuLocked() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.maybeQuiesceLocked() {
		return
	}
}
func (r *Replica1) maybeQuiesceLocked() bool {
	for i := 0; i < 2; i++ {
		if !r.maybeCoalesceHeartbeat() {
			return true
		}
	}
	return false
}
func (r *Replica1) maybeCoalesceHeartbeat() bool {
	msgtype := uintptr(unsafe.Pointer(r)) % 3
	switch msgtype {
	case 0, 1, 2:
		r.store.coalescedMu.mu.Lock() // LockA acquire
	default:
		return false
	}
	r.store.coalescedMu.mu.Unlock() // LockA release
	return true
}

func RunCockroach10214() {
	store := &Store1{}
	store.coalescedMu.mu = deadlock.NewLock()
	responses := &store.coalescedMu.heartbeatResponses
	*responses = append(*responses, 1, 2)
	store.mu.replicas = make(map[int]*Replica1)

	rp1 := &Replica1{
		raftMu: deadlock.NewLock(),
		mu:     deadlock.NewLock(),
		store:  store,
	}
	rp2 := &Replica1{
		raftMu: deadlock.NewLock(),
		mu:     deadlock.NewLock(),
		store:  store,
	}
	store.mu.replicas[0] = rp1
	store.mu.replicas[1] = rp2

	go func() {
		store.sendQueuedHeartbeats()
	}()

	go func() {
		rp1.tick()
	}()
}
