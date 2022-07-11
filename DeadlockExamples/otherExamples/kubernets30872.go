package otherExamples

/* sasha-s
import (
	"github.com/sasha-s/go-deadlock"
)

type PopProcessFunc func()

type ProcessFunc func()

func Util1(f func(), stopCh <-chan struct{}) {
	JitterUntil(f, stopCh)
}

func JitterUntil(f func(), stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
		}
		func() {
			f()
		}()
	}
}

type Queue interface {
	HasSynced()
	Pop(PopProcessFunc)
}

type Config struct {
	Queue
	Process ProcessFunc
}

type Controller struct {
	config Config
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	Util1(c.processLoop, stopCh)
}

func (c *Controller) HasSynced() {
	c.config.Queue.HasSynced()
}

func (c *Controller) processLoop() {
	c.config.Queue.Pop(PopProcessFunc(c.config.Process))
}

type ControllerInterface interface {
	Run(<-chan struct{})
	HasSynced()
}

type ResourceEventHandler interface {
	OnAdd()
}

type ResourceEventHandlerFuncs struct {
	AddFunc func()
}

func (r ResourceEventHandlerFuncs) OnAdd() {
	if r.AddFunc != nil {
		r.AddFunc()
	}
}

type informer struct {
	controller ControllerInterface

	stopChan chan struct{}
}

type federatedInformerImpl struct {
	deadlock.Mutex
	clusterInformer informer
}

func (f *federatedInformerImpl) ClustersSynced() {
	f.Lock()
	defer f.Unlock()
	f.clusterInformer.controller.HasSynced()
}

func (f *federatedInformerImpl) addCluster() {
	f.Lock()
	defer f.Unlock()
}

func (f *federatedInformerImpl) Start() {
	f.Lock()
	defer f.Unlock()

	f.clusterInformer.stopChan = make(chan struct{})
	go f.clusterInformer.controller.Run(f.clusterInformer.stopChan)
}

func (f *federatedInformerImpl) Stop() {
	f.Lock()
	defer f.Unlock()
	close(f.clusterInformer.stopChan)
}

type DelayingDeliverer struct{}

func (d *DelayingDeliverer) StartWithHandler(handler func()) {
	go func() {
		handler()
	}()
}

type FederationView interface {
	ClustersSynced()
}

type FederatedInformer interface {
	FederationView
	Start()
	Stop()
}

type NamespaceController struct {
	namespaceDeliverer         *DelayingDeliverer
	namespaceFederatedInformer FederatedInformer
}

func (nc *NamespaceController) isSynced() {
	nc.namespaceFederatedInformer.ClustersSynced()
}

func (nc *NamespaceController) reconcileNamespace() {
	nc.isSynced()
}

func (nc *NamespaceController) Run(stopChan <-chan struct{}) {
	nc.namespaceFederatedInformer.Start()
	go func() {
		<-stopChan
		nc.namespaceFederatedInformer.Stop()
	}()
	nc.namespaceDeliverer.StartWithHandler(func() {
		nc.reconcileNamespace()
	})
}

type DeltaFIFO struct {
	lock deadlock.Mutex
}

func (f *DeltaFIFO) HasSynced() {
	f.lock.Lock()
	defer f.lock.Unlock()
}

func (f *DeltaFIFO) Pop(process PopProcessFunc) {
	f.lock.Lock()
	defer f.lock.Unlock()
	process()
}

func NewFederatedInformer() FederatedInformer {
	federatedInformer := &federatedInformerImpl{}
	federatedInformer.clusterInformer.controller = NewInformer(
		ResourceEventHandlerFuncs{
			AddFunc: func() {
				federatedInformer.addCluster()
			},
		})
	return federatedInformer
}

func NewInformer(h ResourceEventHandler) *Controller {
	fifo := &DeltaFIFO{}
	cfg := &Config{
		Queue: fifo,
		Process: func() {
			h.OnAdd()
		},
	}
	return &Controller{config: *cfg}
}

func NewNamespaceController() *NamespaceController {
	nc := &NamespaceController{}
	nc.namespaceDeliverer = &DelayingDeliverer{}
	nc.namespaceFederatedInformer = NewFederatedInformer()
	return nc
}

func RunKubernetes30872() {
	namespaceController := NewNamespaceController()
	stop := make(chan struct{})
	namespaceController.Run(stop)
	close(stop)
}
*/

/* deadlock-go */
import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

type PopProcessFunc func()

type ProcessFunc func()

func Util1(f func(), stopCh <-chan struct{}) {
	JitterUntil(f, stopCh)
}

func JitterUntil(f func(), stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
		}
		func() {
			f()
		}()
	}
}

type Queue interface {
	HasSynced()
	Pop(PopProcessFunc)
}

type Config struct {
	Queue
	Process ProcessFunc
}

type Controller struct {
	config Config
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	Util1(c.processLoop, stopCh)
}

func (c *Controller) HasSynced() {
	c.config.Queue.HasSynced()
}

func (c *Controller) processLoop() {
	c.config.Queue.Pop(PopProcessFunc(c.config.Process))
}

type ControllerInterface interface {
	Run(<-chan struct{})
	HasSynced()
}

type ResourceEventHandler interface {
	OnAdd()
}

type ResourceEventHandlerFuncs struct {
	AddFunc func()
}

func (r ResourceEventHandlerFuncs) OnAdd() {
	if r.AddFunc != nil {
		r.AddFunc()
	}
}

type informer struct {
	controller ControllerInterface

	stopChan chan struct{}
}

type federatedInformerImpl struct {
	mu              deadlock.Mutex
	clusterInformer informer
}

func (f *federatedInformerImpl) ClustersSynced() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.clusterInformer.controller.HasSynced()
}

func (f *federatedInformerImpl) addCluster() {
	f.mu.Lock()
	defer f.mu.Unlock()
}

func (f *federatedInformerImpl) Start() {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.clusterInformer.stopChan = make(chan struct{})
	go f.clusterInformer.controller.Run(f.clusterInformer.stopChan)
}

func (f *federatedInformerImpl) Stop() {
	f.mu.Lock()
	defer f.mu.Unlock()
	close(f.clusterInformer.stopChan)
}

type DelayingDeliverer struct{}

func (d *DelayingDeliverer) StartWithHandler(handler func()) {
	go func() {
		handler()
	}()
}

type FederationView interface {
	ClustersSynced()
}

type FederatedInformer interface {
	FederationView
	Start()
	Stop()
}

type NamespaceController struct {
	namespaceDeliverer         *DelayingDeliverer
	namespaceFederatedInformer FederatedInformer
}

func (nc *NamespaceController) isSynced() {
	nc.namespaceFederatedInformer.ClustersSynced()
}

func (nc *NamespaceController) reconcileNamespace() {
	nc.isSynced()
}

func (nc *NamespaceController) Run(stopChan <-chan struct{}) {
	nc.namespaceFederatedInformer.Start()
	go func() {
		<-stopChan
		nc.namespaceFederatedInformer.Stop()
	}()
	nc.namespaceDeliverer.StartWithHandler(func() {
		nc.reconcileNamespace()
	})
}

type DeltaFIFO struct {
	lock deadlock.Mutex
}

func (f *DeltaFIFO) HasSynced() {
	f.lock.Lock()
	defer f.lock.Unlock()
}

func (f *DeltaFIFO) Pop(process PopProcessFunc) {
	f.lock.Lock()
	defer f.lock.Unlock()
	process()
}

func NewFederatedInformer() FederatedInformer {
	federatedInformer := &federatedInformerImpl{
		mu: *deadlock.NewLock(),
	}
	federatedInformer.clusterInformer.controller = NewInformer(
		ResourceEventHandlerFuncs{
			AddFunc: func() {
				federatedInformer.addCluster()
			},
		})
	return federatedInformer
}

func NewInformer(h ResourceEventHandler) *Controller {
	fifo := &DeltaFIFO{
		lock: *deadlock.NewLock(),
	}
	cfg := &Config{
		Queue: fifo,
		Process: func() {
			h.OnAdd()
		},
	}
	return &Controller{config: *cfg}
}

func NewNamespaceController() *NamespaceController {
	nc := &NamespaceController{}
	nc.namespaceDeliverer = &DelayingDeliverer{}
	nc.namespaceFederatedInformer = NewFederatedInformer()
	return nc
}

func RunKubernetes30872() {
	namespaceController := NewNamespaceController()
	stop := make(chan struct{})
	namespaceController.Run(stop)
	close(stop)
}
