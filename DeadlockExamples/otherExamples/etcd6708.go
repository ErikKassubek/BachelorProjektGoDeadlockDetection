package otherExamples

/* sasha-s
import (
	"context"

	"github.com/sasha-s/go-deadlock"
)

type EndpointSelectionMode int

const (
	EndpointSelectionRandom EndpointSelectionMode = iota
	EndpointSelectionPrioritizeLeader
)

type MembersAPI interface {
	Leader(ctx context.Context)
}

type Client1 interface {
	Sync(ctx context.Context)
	SetEndpoints()
	httpClient
}

type httpClient interface {
	Do(context.Context)
}

type httpClusterClient struct {
	deadlock.Mutex
	selectionMode EndpointSelectionMode
}

func (c *httpClusterClient) getLeaderEndpoint() {
	mAPI := NewMembersAPI(c)
	mAPI.Leader(context.Background())
}

func (c *httpClusterClient) SetEndpoints() {
	switch c.selectionMode {
	case EndpointSelectionRandom:
	case EndpointSelectionPrioritizeLeader:
		c.getLeaderEndpoint()
	}
}

func (c *httpClusterClient) Do(ctx context.Context) {
	c.Lock() // block here
	c.Unlock()
}

func (c *httpClusterClient) Sync(ctx context.Context) {
	c.Lock()
	defer c.Unlock()

	c.SetEndpoints()
}

type httpMembersAPI struct {
	client httpClient
}

func (m *httpMembersAPI) Leader(ctx context.Context) {
	m.client.Do(ctx)
}

func NewMembersAPI(c Client1) MembersAPI {
	return &httpMembersAPI{
		client: c,
	}
}
func RunEtcd6708() {
	hc := &httpClusterClient{
		selectionMode: EndpointSelectionPrioritizeLeader,
	}
	hc.Sync(context.Background())
}

/* deadlock-go */
import (
	"context"
	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

type EndpointSelectionMode int

const (
	EndpointSelectionRandom EndpointSelectionMode = iota
	EndpointSelectionPrioritizeLeader
)

type MembersAPI interface {
	Leader(ctx context.Context)
}

type Client1 interface {
	Sync(ctx context.Context)
	SetEndpoints()
	httpClient
}

type httpClient interface {
	Do(context.Context)
}

type httpClusterClient struct {
	mu            deadlock.Mutex
	selectionMode EndpointSelectionMode
}

func (c *httpClusterClient) getLeaderEndpoint() {
	mAPI := NewMembersAPI(c)
	mAPI.Leader(context.Background())
}

func (c *httpClusterClient) SetEndpoints() {
	switch c.selectionMode {
	case EndpointSelectionRandom:
	case EndpointSelectionPrioritizeLeader:
		c.getLeaderEndpoint()
	}
}

func (c *httpClusterClient) Do(ctx context.Context) {
	c.mu.Lock() // block here
	c.mu.Unlock()
}

func (c *httpClusterClient) Sync(ctx context.Context) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.SetEndpoints()
}

type httpMembersAPI struct {
	client httpClient
}

func (m *httpMembersAPI) Leader(ctx context.Context) {
	m.client.Do(ctx)
}

func NewMembersAPI(c Client1) MembersAPI {
	return &httpMembersAPI{
		client: c,
	}
}
func RunEtcd6708() {
	hc := &httpClusterClient{
		mu:            *deadlock.NewLock(),
		selectionMode: EndpointSelectionPrioritizeLeader,
	}
	hc.Sync(context.Background())
}
