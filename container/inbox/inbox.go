package inbox

import (
	"doko/container/ring"
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
)

var ErrInboxStopped = errors.New("inbox stopped")

const (
	STOPPED int32 = iota
	STARTING
	IDLE
	RUNNING
	STOPPING
)

const (
	DefaultInitialSize = 1024
	DefaultThroughput  = 300
	DefaultBatchSize   = 1024 * 4
)

type (
	Processor[T any] interface {
		Invoke(envelopes []T)
	}
	Inbox[T any] struct {
		buffer    *ring.Ring[T]
		processor Processor[T]
		scheduler scheduler
		status    int32
		wg        sync.WaitGroup
		config
	}
)

func NewInbox[T any](opts ...Option) *Inbox[T] {
	conf := defaultConfig()
	for _, opt := range opts {
		opt(conf)
	}
	return &Inbox[T]{
		buffer:    ring.New[T](conf.initialSize),
		scheduler: scheduler(conf.throughput),
		status:    STOPPED,
		wg:        sync.WaitGroup{},
		config:    conf,
	}
}

func (slf *Inbox[T]) Start(processor Processor[T]) {
	if atomic.CompareAndSwapInt32(&slf.status, STOPPED, STARTING) {
		slf.processor = processor
		atomic.StoreInt32(&slf.status, IDLE)
		slf.schedule()
	}
}

func (slf *Inbox[T]) Stop(graceful ...bool) {
	if len(graceful) > 0 && graceful[0] {
		atomic.StoreInt32(&slf.status, STOPPING)
		slf.wg.Wait()
	} else {
		atomic.StoreInt32(&slf.status, STOPPED)
	}
}

func (slf *Inbox[T]) Send(envelope T) error {
	if atomic.LoadInt32(&slf.status) != STOPPING && atomic.LoadInt32(&slf.status) != STOPPED {
		slf.buffer.Push(envelope)
		slf.schedule()
		return nil
	}
	return ErrInboxStopped
}

func (slf *Inbox[T]) schedule() {
	if atomic.CompareAndSwapInt32(&slf.status, IDLE, RUNNING) {
		slf.wg.Add(1)
		slf.scheduler.Schedule(slf.process)
	}
}

func (slf *Inbox[T]) process() {
	slf.run()
	if atomic.LoadInt32(&slf.status) == STOPPING {
		atomic.StoreInt32(&slf.status, STOPPED)
	} else {
		atomic.StoreInt32(&slf.status, IDLE)
	}
	slf.wg.Done()
}

func (slf *Inbox[T]) run() {
	count, throughput := 0, slf.scheduler.Throughput()
	for atomic.LoadInt32(&slf.status) != STOPPED {
		if count > throughput {
			count = 0
			runtime.Gosched()
		}
		count++
		if envelopes, ok := slf.buffer.PopN(slf.config.batchSize); ok && len(envelopes) > 0 {
			slf.processor.Invoke(envelopes)
		} else {
			return
		}
	}
}

type scheduler int

func (slf scheduler) Schedule(fn func()) {
	go fn()
}

func (slf scheduler) Throughput() int {
	return int(slf)
}

type config struct {
	initialSize int64
	throughput  int64
	batchSize   int64
}

func defaultConfig() config {
	return config{
		initialSize: DefaultInitialSize,
		throughput:  DefaultThroughput,
		batchSize:   DefaultBatchSize,
	}
}

type Option func(config)

// WithInitialSize sets the initial size of the inbox buffer.
func WithInitialSize(initialSize int64) Option {
	return func(conf config) {
		conf.initialSize = initialSize
	}
}

// WithThroughput sets the throughput of the inbox scheduler.
func WithThroughput(throughput int64) Option {
	return func(conf config) {
		conf.throughput = throughput
	}
}

// WithBatchSize sets the batch size of the inbox buffer pop operation.
func WithBatchSize(batchSize int64) Option {
	return func(conf config) {
		conf.batchSize = batchSize
	}
}
