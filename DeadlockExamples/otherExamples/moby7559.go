/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/7559
 * Buggy version: 64579f51fcb439c36377c0068ccc9a007b368b5a
 * fix commit-id: 6cbb8e070d6c3a66bf48fbe5cbf689557eee23db
 * Flaky: 100/100
 */
package otherExamples

/* sasha-s
import (
	"net"

	"github.com/sasha-s/go-deadlock"
)

type UDPProxy struct {
	connTrackLock deadlock.Mutex
}

func (proxy *UDPProxy) Run() {
	for i := 0; i < 2; i++ {
		proxy.connTrackLock.Lock()
		_, err := net.DialUDP("udp", nil, nil)
		if err != nil {
			/// Missing unlock here
			continue
		}
		if i == 0 {
			break
		}
	}
	proxy.connTrackLock.Unlock()
}
func RunMoby7559() {
	proxy := &UDPProxy{}
	go proxy.Run()
}
*/

/* deadlock-go */
import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
	"net"
)

type UDPProxy struct {
	connTrackLock deadlock.Mutex
}

func (proxy *UDPProxy) Run() {
	for i := 0; i < 2; i++ {
		proxy.connTrackLock.Lock()
		_, err := net.DialUDP("udp", nil, nil)
		if err != nil {
			/// Missing unlock here
			continue
		}
		if i == 0 {
			break
		}
	}
	proxy.connTrackLock.Unlock()
}
func RunMoby7559() {
	proxy := &UDPProxy{connTrackLock: *deadlock.NewLock()}
	go proxy.Run()
}
