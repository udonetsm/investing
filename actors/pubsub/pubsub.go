package pubsub

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/udonetsm/investing/actors/stack"
	"github.com/udonetsm/investing/general"
	"github.com/udonetsm/investing/interfaces"
)

const STREAMNAME = "transactions"

const (
	SubTypeRequests = "requests"
	SubTypeReplies  = "replies"
	SubTypeAllRead  = "read"
)

const URL = "nats://:4222"

const (
	SubjectString = "%s.%s"
)

type PubSub struct {
	// Some message to the  target subject of stream of nats-server
	Message     string
	SubjectType string
	// Configurtion of stream. Can be rebuild manually
	StreamConf jetstream.StreamConfig
	// User who works with jetstream
	ConsumerConf jetstream.ConsumerConfig
	// Fill when switch to jetstream mode. Stream builds on it
	Jetstream jetstream.JetStream
	// Stream with StreamConf.
	Stream jetstream.Stream
	// Consumer which works with stream
	Consumer jetstream.Consumer
	// Contains errors from any step
	// Should check always
	Error           error
	All, Show, Read bool
	// Contains some key-value storage to store cache on server
	SeterGeter interfaces.LastSeenSetGeter
}

func (ps *PubSub) Subscribe(uid string) *stack.Stack {
	if ps.Read {
		general.SetLastSeen(ps.SeterGeter, uid, time.Now().UnixNano())
		return &stack.Stack{}
	}
	con := connect()
	defer con.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	defer ctx.Done()
	ps.jetstreamMode(con)
	ps.stream(uid, ctx)
	if ps.Error != nil {
		return &stack.Stack{}
	}
	ps.consumer(uid, ctx)
	if ps.Error != nil {
		return &stack.Stack{}
	}
	messagesContext, err := ps.Consumer.Messages()
	if err != nil {
		ps.Error = err
		return &stack.Stack{}
	}
	defer messagesContext.Stop()
	messagesChan := make(chan jetstream.Msg)
	counter := make(chan struct{})
	stack := stack.NewStack()
	go next(ps, messagesContext, messagesChan, counter)
	for {
		select {
		case msg := <-messagesChan:
			defer stack.Add(msg)
		case <-counter:
			stack.Count++
		case <-time.After(200 * time.Millisecond):
			log.Printf("[%s] got from subject %v of stream [%s]\n", ps.ConsumerConf.Name, ps.StreamConf.Subjects, ps.StreamConf.Name)
			return stack
		}
	}
}

func next(ps *PubSub, msgs jetstream.MessagesContext, messages chan jetstream.Msg, counter chan struct{}) {
	for {
		msg, err := msgs.Next()
		if err != nil {
			return
		}
		if ps.Show {
			messages <- msg
		} else {
			counter <- struct{}{}
		}
		err = msg.Ack()
		if err != nil {
			ps.Error = err
			return
		}
	}
}

func (ps *PubSub) Publish(targetUID string) {
	con := connect()
	defer con.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer ctx.Done()
	defer cancel()
	ps.jetstreamMode(con)
	ps.streamConfig(targetUID)
	ps.stream(targetUID, ctx)
	_, err := ps.Jetstream.Publish(ctx, ps.StreamConf.Subjects[0], []byte(ps.Message))
	if err != nil {
		ps.Error = err
		return
	}
	log.Printf("Published %s to the %v of %s\n", ps.Message, ps.StreamConf.Subjects, ps.StreamConf.Name)
}

func connect() *nats.Conn {
	con, err := nats.Connect(URL)
	if err != nil {
		log.Fatal(err)
	}
	return con
}

func (ps *PubSub) jetstreamMode(con *nats.Conn) {
	js, err := jetstream.New(con)
	if err != nil {
		log.Fatal(err)
	}
	ps.Jetstream = js
}

func (ps *PubSub) streamConfig(uid string) {
	subject := fmt.Sprintf(SubjectString, ps.SubjectType, uid)
	ps.StreamConf = jetstream.StreamConfig{
		Name:     STREAMNAME,
		Subjects: []string{subject},
	}
}

func (ps *PubSub) consumerConfig(uid string) {
	startopt := time.Unix(0, general.GetLastSeen(ps.SeterGeter, uid))
	var conf jetstream.ConsumerConfig
	if !ps.All {
		conf = jetstream.ConsumerConfig{
			DeliverPolicy: jetstream.DeliverByStartTimePolicy,
			OptStartTime:  &startopt,
		}
	} else {
		conf = jetstream.ConsumerConfig{}
	}
	conf.FilterSubjects = ps.StreamConf.Subjects
	ps.ConsumerConf = conf
	ps.ConsumerConf.Name = uid
	ps.ConsumerConf.InactiveThreshold = time.Second

}

func (ps *PubSub) stream(uid string, ctx context.Context) {
	ps.streamConfig(uid)
	ps.Stream, ps.Error = ps.Jetstream.CreateOrUpdateStream(ctx, ps.StreamConf)
}

func (ps *PubSub) consumer(uid string, ctx context.Context) {
	ps.consumerConfig(uid)
	ps.Consumer, ps.Error = ps.Stream.CreateOrUpdateConsumer(ctx, ps.ConsumerConf)
}
