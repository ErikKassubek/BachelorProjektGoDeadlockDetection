/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/36114
 * Buggy version: 6d4d3c52ae7c3f910bfc7552a2a673a8338e5b9f
 * fix commit-id: a44fcd3d27c06aaa60d8d1cbce169f0d982e74b1
 * Flaky: 100/100
 * Description:
 *   This is a double lock bug. The the lock for the
 * struct svm has already been locked when calling
 * svm.hotRemoveVHDsAtStart()
 */
package otherExamples

/* sasha-s
import (
	"github.com/sasha-s/go-deadlock"
)

type serviceVM struct {
	deadlock.Mutex
}

func (svm *serviceVM) hotAddVHDsAtStart() {
	svm.Lock()
	defer svm.Unlock()
	svm.hotRemoveVHDsAtStart()
}

func (svm *serviceVM) hotRemoveVHDsAtStart() {
	svm.Lock() // Double lock here
	defer svm.Unlock()
}

func RunMoby36114() {
	s := &serviceVM{}
	go s.hotAddVHDsAtStart()
}
*/

/* deadlock-go */

import deadlock "github.com/ErikKassubek/Deadlock-Go"

type serviceVM struct {
	mu deadlock.Mutex
}

func (svm *serviceVM) hotAddVHDsAtStart() {
	svm.mu.Lock()
	defer svm.mu.Unlock()
	svm.hotRemoveVHDsAtStart()
}

func (svm *serviceVM) hotRemoveVHDsAtStart() {
	svm.mu.Lock() // Double lock here
	defer svm.mu.Unlock()
}

func RunMoby36114() {
	s := &serviceVM{mu: deadlock.NewLock()}
	go s.hotAddVHDsAtStart()
}
