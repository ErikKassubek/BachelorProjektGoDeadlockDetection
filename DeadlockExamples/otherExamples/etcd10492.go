package otherExamples

/* sasha-s
import (
	"context"
	"time"

	"github.com/sasha-s/go-deadlock"
)

type Checkpointer func(ctx context.Context)

type lessor struct {
	mu                 deadlock.Mutex
	cp                 Checkpointer
	checkpointInterval time.Duration
}

func (le *lessor) Checkpoint() {
	le.mu.Lock() // block here
	defer le.mu.Unlock()
}

func (le *lessor) SetCheckpointer(cp Checkpointer) {
	le.mu.Lock()
	defer le.mu.Unlock()

	le.cp = cp
}

func (le *lessor) Renew() {
	le.mu.Lock()
	unlock := func() { le.mu.Unlock() }
	defer func() { unlock() }()

	if le.cp != nil {
		le.cp(context.Background())
	}
}
func RunEtcd10492() {
	le := &lessor{
		checkpointInterval: 0,
	}
	fakerCheckerpointer := func(ctx context.Context) {
		le.Checkpoint()
	}
	le.SetCheckpointer(fakerCheckerpointer)
	le.mu.Lock()
	le.mu.Unlock()
	le.Renew()
}
*/

/* deadlock-go */
import (
	"context"
	deadlock "github.com/ErikKassubek/Deadlock-Go"
	"time"
)

type Checkpointer func(ctx context.Context)

type lessor struct {
	mu                 deadlock.Mutex
	cp                 Checkpointer
	checkpointInterval time.Duration
}

func (le *lessor) Checkpoint() {
	le.mu.Lock() // block here
	defer le.mu.Unlock()
}

func (le *lessor) SetCheckpointer(cp Checkpointer) {
	le.mu.Lock()
	defer le.mu.Unlock()

	le.cp = cp
}

func (le *lessor) Renew() {
	le.mu.Lock()
	unlock := func() { le.mu.Unlock() }
	defer func() { unlock() }()

	if le.cp != nil {
		le.cp(context.Background())
	}
}
func RunEtcd10492() {
	le := &lessor{
		mu:                 *deadlock.NewLock(),
		checkpointInterval: 0,
	}
	fakerCheckerpointer := func(ctx context.Context) {
		le.Checkpoint()
	}
	le.SetCheckpointer(fakerCheckerpointer)
	le.mu.Lock()
	le.mu.Unlock()
	le.Renew()
}
