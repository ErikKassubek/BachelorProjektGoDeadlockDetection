package otherExamples

/* sasha-s
import (
	"context"
	"fmt"

	"github.com/sasha-s/go-deadlock"
)

var ErrConnClosed error

type Client struct {
	mu     deadlock.Mutex
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cancel == nil {
		return
	}
	c.cancel()
	c.cancel = nil
	c.mu.Unlock()
	c.mu.Lock() // block here
}

type remoteClient struct {
	client *Client
	mu     deadlock.Mutex
}

func (r *remoteClient) acquire(ctx context.Context) error {
	for {
		r.client.mu.Lock()
		closed := r.client.cancel == nil
		r.mu.Lock()
		r.mu.Unlock()
		if closed {
			return ErrConnClosed // Missing RUnlock before return
		}
		r.client.mu.Unlock()
	}
}

type kv struct {
	rc *remoteClient
}

func (kv *kv) Get(ctx context.Context) error {
	return kv.Do(ctx)
}

func (kv *kv) Do(ctx context.Context) error {
	for {
		err := kv.do(ctx)
		if err == nil {
			return nil
		}
		return err
	}
}

func (kv *kv) do(ctx context.Context) error {
	err := kv.getRemote(ctx)
	return err
}

func (kv *kv) getRemote(ctx context.Context) error {
	return kv.rc.acquire(ctx)
}

type KV interface {
	Get(ctx context.Context) error
	Do(ctx context.Context) error
}

func NewKV(c *Client) KV {
	return &kv{rc: &remoteClient{
		client: c,
	}}
}
func RunEtcd5509() {
	ctx, cancel := context.WithCancel(context.TODO())
	cli := &Client{
		ctx:    ctx,
		cancel: cancel,
	}
	kv := NewKV(cli)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		err := kv.Get(context.TODO())
		if err != nil && err != ErrConnClosed {
			fmt.Println("Expect ErrConnClosed")
		}
	}()

	cli.Close()

	<-donec
}
*/

/* deadlock-go */
import (
	"context"
	"fmt"
	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

var ErrConnClosed error

type Client struct {
	mu     deadlock.Mutex
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cancel == nil {
		return
	}
	c.cancel()
	c.cancel = nil
	c.mu.Unlock()
	c.mu.Lock() // block here
}

type remoteClient struct {
	client *Client
	mu     deadlock.Mutex
}

func (r *remoteClient) acquire(ctx context.Context) error {
	for {
		r.client.mu.Lock()
		closed := r.client.cancel == nil
		r.mu.Lock()
		r.mu.Unlock()
		if closed {
			return ErrConnClosed // Missing RUnlock before return
		}
		r.client.mu.Unlock()
	}
}

type kv struct {
	rc *remoteClient
}

func (kv *kv) Get(ctx context.Context) error {
	return kv.Do(ctx)
}

func (kv *kv) Do(ctx context.Context) error {
	for {
		err := kv.do(ctx)
		if err == nil {
			return nil
		}
		return err
	}
}

func (kv *kv) do(ctx context.Context) error {
	err := kv.getRemote(ctx)
	return err
}

func (kv *kv) getRemote(ctx context.Context) error {
	return kv.rc.acquire(ctx)
}

type KV interface {
	Get(ctx context.Context) error
	Do(ctx context.Context) error
}

func NewKV(c *Client) KV {
	return &kv{rc: &remoteClient{
		client: c,
		mu:     *deadlock.NewLock(),
	}}
}
func RunEtcd5509() {
	ctx, cancel := context.WithCancel(context.TODO())
	cli := &Client{
		mu:     *deadlock.NewLock(),
		ctx:    ctx,
		cancel: cancel,
	}
	kv := NewKV(cli)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		err := kv.Get(context.TODO())
		if err != nil && err != ErrConnClosed {
			fmt.Println("Expect ErrConnClosed")
		}
	}()

	cli.Close()

	<-donec
}
