package otherExamples

/* sasha-s
import (
	"github.com/sasha-s/go-deadlock"
)

type Gossip struct {
	mu     deadlock.Mutex
	closed bool
}

func (g *Gossip) bootstrap() {
	for {
		g.mu.Lock()
		if g.closed {
			/// Missing g.mu.Unlock
			break
		}
		g.mu.Unlock()
		break
	}
}

func (g *Gossip) manage() {
	for {
		g.mu.Lock()
		if g.closed {
			/// Missing g.mu.Unlock
			break
		}
		g.mu.Unlock()
		break
	}
}
func RunCockroach584() {
	g := &Gossip{
		closed: true,
	}
	go func() {
		g.bootstrap()
		g.manage()
	}()
}

*/

/* Deadlock-go */
import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

type Gossip struct {
	mu     deadlock.Mutex
	closed bool
}

func (g *Gossip) bootstrap() {
	for {
		g.mu.Lock()
		if g.closed {
			/// Missing g.mu.Unlock
			break
		}
		g.mu.Unlock()
		break
	}
}

func (g *Gossip) manage() {
	for {
		g.mu.Lock()
		if g.closed {
			/// Missing g.mu.Unlock
			break
		}
		g.mu.Unlock()
		break
	}
}
func RunCockroach584() {
	g := &Gossip{
		mu:     *deadlock.NewLock(),
		closed: true,
	}
	go func() {
		g.bootstrap()
		g.manage()
	}()
}
