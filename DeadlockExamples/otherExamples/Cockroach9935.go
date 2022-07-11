/*
 * Project: cockroach
 * Issue or PR  : https://github.com/cockroachdb/cockroach/pull/9935
 * Buggy version: 4df302cc3f03328395dc3fefbfba58b7718e4f2f
 * fix commit-id: ed6a100ba38dd51b0888b9a3d3ac6bdbb26c528c
 * Flaky: 100/100
 * Description: This bug is caused by acquiring l.mu.Lock() twice. The fix is
 * to release l.mu.Lock() before acquiring l.mu.Lock for the second time.
 */
package otherExamples

/* sasha-s
import (
	"errors"
	"math/rand"

	"github.com/sasha-s/go-deadlock"
)

type loggingT struct {
	mu deadlock.Mutex
}

func (l *loggingT) outputLogEntry() {
	l.mu.Lock()
	if err := l.createFile(); err != nil {
		l.exit(err)
	}
	l.mu.Unlock()
}
func (l *loggingT) createFile() error {
	if rand.Intn(8)%4 > 0 {
		return errors.New("")
	}
	return nil
}
func (l *loggingT) exit(err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
}
func RunCockroach9935() {
	l := &loggingT{}
	go l.outputLogEntry()
}
*/

/* Deadlock-go */
import (
	"errors"
	deadlock "github.com/ErikKassubek/Deadlock-Go"
	"math/rand"
)

type loggingT struct {
	mu deadlock.Mutex
}

func (l *loggingT) outputLogEntry() {
	l.mu.Lock()
	if err := l.createFile(); err != nil {
		l.exit(err)
	}
	l.mu.Unlock()
}
func (l *loggingT) createFile() error {
	if rand.Intn(8)%4 > 0 {
		return errors.New("")
	}
	return nil
}
func (l *loggingT) exit(err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
}
func RunCockroach9935() {
	l := &loggingT{}
	l.mu = *deadlock.NewLock()
	go l.outputLogEntry()
}
