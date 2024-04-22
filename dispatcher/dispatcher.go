package dispatcher

import (
	"math/rand"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/request"
)

type Dispatcher struct {
	// 工作池大小
	workerPoolSize int
	// 任务队列长度
	taskQueueSize int
	// 任务列表
	taskQueue []chan iface.Request
}

func NewDispatcher(workerPoolSize int, taskQueueSize int) iface.Dispatcher {
	return &Dispatcher{
		workerPoolSize: workerPoolSize,
		taskQueueSize:  taskQueueSize,
		taskQueue:      make([]chan iface.Request, workerPoolSize),
	}
}

func (d *Dispatcher) Dispatch(conn iface.Connection, msg []byte) {
	// 随机分配一个worker
	workerId := rand.Intn(d.workerPoolSize)

	// 写入工作队列
	d.taskQueue[workerId] <- request.NewRequest(conn, msg)
}

func (d *Dispatcher) StartWorkerPool() {
	for i := 0; i < d.workerPoolSize; i++ {
		d.taskQueue[i] = make(chan iface.Request, d.taskQueueSize)

		go d.startWorker(i, d.taskQueue[i])
	}
}

func (d *Dispatcher) startWorker(workerId int, taskQueue chan iface.Request) {
	for {
		select {
		case msg := <-taskQueue:
			d.doHandler(workerId, msg)
		}
	}
}

func (d *Dispatcher) doHandler(workerId int, req iface.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("workerId: %d handle msg panic: %v", workerId, err)
		}
	}()

	req.GetConn().Pipeline().Handle(req.GetMsg())
}
