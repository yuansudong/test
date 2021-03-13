package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Worker 用于描述一个工人
type Worker struct {
	_WID    uint64
	_TaskCh chan int
}

// NewWorker 新建立一个工人
func NewWorker(iWID uint64) *Worker {
	return &Worker{
		_WID:    iWID,
		_TaskCh: make(chan int),
	}
}

// Do 用于工人执行任务
func (w *Worker) Do(ctx context.Context) {
	tIdleDuration := 500 * time.Millisecond
	for {
		select {
		case <-ctx.Done():
			log.Println("进程退出")
			goto end
		case <-time.After(tIdleDuration):
			log.Println("Worker驻留时间已到准备退出!", w._WID)
			goto end
		case task := <-w._TaskCh:
			time.Sleep(time.Second) // 模拟任务处理时长
			log.Println("Worker:", w._WID, " 收到任务", task)
		}
	}
end:
	close(w._TaskCh)
	return
}

// Leader 用于管理工人
type Leader struct {
	sync.WaitGroup
	ctx    context.Context
	idx    uint64
	cancel context.CancelFunc
	works  map[uint64]*Worker
}

// NewLeader 实例化一个leader
func NewLeader() *Leader {
	mCtx, mCancel := context.WithCancel(context.Background())
	return &Leader{
		ctx:    mCtx,
		idx:    0,
		cancel: mCancel,
		works:  make(map[uint64]*Worker),
	}
}

// ClearWorker 用于清理工人
func (l *Leader) ClearWorker(w *Worker) {
	delete(l.works, w._WID)
}

// DispatchTask 用于派发任务
func (l *Leader) DispatchTask(iTask int) {
	bSuccess := false
	for _, work := range l.works {
		if l.NotifyWorker(work, iTask) {
			bSuccess = true
			break
		}
	}
	if !bSuccess {
		l.WorkerStartup(iTask)
	}
}

// NotifyWorker 用于通知
func (l *Leader) NotifyWorker(w *Worker, iTask int) (status bool) {
	defer func() {
		if r := recover(); r != nil {
			// 进入这里代表工人已经退出了
			l.ClearWorker(w)
			log.Println("清理了工人:", w._WID)
			return
		}
	}()
	select {
	case w._TaskCh <- iTask:
		status = true
		return
	default:
		return
	}
}

// WorkerStartup 启动一个工人
func (l *Leader) WorkerStartup(iTask int) {
	w := NewWorker(atomic.AddUint64(&l.idx, 1))
	go func() {
		l.Add(1)
		defer l.Done()
		w.Do(l.ctx)
	}()
	w._TaskCh <- iTask
	l.works[w._WID] = w
}

// Exit 用于退出操作
func (l *Leader) Exit() {
	// 向所有协程通知退出信号
	l.cancel()
	// 等待所有协程退出
	l.Wait()
	log.Println("开启协程的最大编号是:", l.idx)
}

func main() {
	mLeader := NewLeader()
	for i := 0; i < 100; i++ {
		mLeader.DispatchTask(i)
		// 模拟任务的不确定性
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
	time.Sleep(15 * time.Second)
	mLeader.Exit()
}

/*

 1. channel的关闭,所有从该channel读取的消息,都会收到信号通知

 2. channel分为有缓冲和无缓冲,有缓冲channel为1,无缓冲.基于无缓冲channel,可以做协程见的同步
 3. channel分为可读和可写.
    利用可读可写,可以严格控制channel的安全性,毕竟向已经关闭的channel写数据是会发生异常的.
 4. 利用 写channel会发生异常的特性,可以简化代码.
 5. 倘若channel传递的数据没有意义,只是作为信号传递,可以用struct{}作为管道元素.
 6. 巧用select避免执行流阻塞.

    比如在socket读协程中,需要监听,是否收到了监听的信号,倘若接收到了通知,就继续执行
	比如在任务分发的时候,可以利用select default的特性.
	向一个管道中写任务,当任务管道已经满的时候,可以继续轮询.
7. 对于进程内channel的容量,要么是0,要么是1.
8. select 对于管道的选择,是伪随机
9. 基于select,可以做异步任务的超时处理.








*/
