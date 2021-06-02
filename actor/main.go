package main

import (
	"fmt"
	"log"
	"reflect"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type hello struct{ Who string }

// HelloActor
type HelloActor struct {
	Acc int64
}

func NewHelloActor() actor.Actor {
	return &HelloActor{Acc: 160}
}

// Receive
func (HA *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

// Printer 用于打印
func Printer(next actor.ReceiverFunc) actor.ReceiverFunc {

	// env 值要发送的
	return func(ctx actor.ReceiverContext, env *actor.MessageEnvelope) {
		message := env.Message
		if ca, ok := ctx.Actor().(*HelloActor); ok {
			log.Printf("当前的值是:%d,%+v", ca.Acc, ctx.Message())
		}
		log.Printf("%v got %v %+v", ctx.Self(), reflect.TypeOf(message), message)
		next(ctx, env)
		log.Printf("ctx ==> %+v   %+v", ctx.Actor(), ctx.Message())
	}
}

func main() {
	system := actor.NewActorSystem()
	rootContext := system.Root
	props := actor.PropsFromProducer(NewHelloActor).WithReceiverMiddleware(Printer)
	pid := rootContext.Spawn(props)
	rootContext.Send(pid, &hello{Who: "Roger"})
	_, _ = console.ReadLine()
}
