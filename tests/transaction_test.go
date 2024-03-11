package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/udonetsm/investing/actors/pubsub"
	"github.com/udonetsm/investing/actors/storage/myredis"
	"github.com/udonetsm/investing/general"
	"github.com/udonetsm/investing/models"
	"github.com/udonetsm/investing/tests/testmodels"
)

func TestPub(t *testing.T) {
	transaction := testmodels.Transaction()
	ps := &pubsub.PubSub{
		SeterGeter:  myredis.NewRedis(myredis.REQUESTS),
		Message:     transaction.Tid,
		SubjectType: pubsub.SubTypeRequests,
	}
	recieverUID := ""
	if transaction.Type == models.INVESTING {
		recieverUID = transaction.Payer.Uid
	}
	if transaction.Type == models.RETURNING {
		recieverUID = transaction.Reciever.Uid
	}
	general.Publish(ps, recieverUID)
	if ps.Error != nil {
		fmt.Println(ps.Error)
	}
}

// func TestChecReqAndDontReadIt(t *testing.T) {
// transaction := testmodels.Transaction()
// ps := &pubsub.PubSub{SaveGeter: myredis.NewRedis(myredis.REQUESTS)}
// _, i := general.CheckRequests(ps, transaction.Reciever.(*startuper.Startuper).User.User_id)
// if ps.Error != nil {
// fmt.Println("Too often. Wait 3 second before new checking!")
// return
// }
// if i > 0 {
// fmt.Println(i, "new requests")
// return
// }
// fmt.Println("No new requests")
// }

// func TestChecReqAndReadIt(t *testing.T) {
// transaction := testmodels.Transaction()
// ps := &pubsub.PubSub{SaveGeter: myredis.NewRedis(myredis.REQUESTS), Show: true}
// messages, _ := general.CheckRequsets(ps, transaction)
// if ps.Error != nil {
// fmt.Println("Too often. Wait 3 second before new checking!")
// return
// }
// if len(messages) > 0 {
// for _, msg := range messages {
// fmt.Println(string(msg.Data()))
// }
// } else {
// fmt.Println("No new messages")
// }
// }

func TestPubAnswer(t *testing.T) {
	transaction := testmodels.Transaction()
	recieverUID := ""
	if transaction.Type == models.INVESTING {
		recieverUID = transaction.Reciever.Uid
	}
	if transaction.Type == models.RETURNING {
		recieverUID = transaction.Payer.Uid
	}
	transaction.Accepted = true
	ps := &pubsub.PubSub{
		Message:     transaction.Tid,
		SeterGeter:  myredis.NewRedis(myredis.REQUESTS),
		SubjectType: pubsub.SubTypeReplies,
	}
	general.Publish(ps, recieverUID)
	if ps.Error != nil {
		log.Println(ps.Error)
		return
	}
}

// func TestCheckAnsAndDontRead(t *testing.T) {
// transaction := testmodels.Transaction()
// ps := &pubsub.PubSub{SaveGeter: myredis.NewRedis(myredis.REQUESTS)}
// _, i := general.CheckAnswer(ps, transaction)
// if ps.Error != nil {
// fmt.Println("Too often. Wait 3 second before new checking!")
// return
// }
// if i < 1 {
// fmt.Println("No new answers")
// return
// }
// fmt.Println(i, " new answers")
// }

// func TestCheckAnsAndRead(t *testing.T) {
// transaction := testmodels.Transaction()
// ps := &pubsub.PubSub{Show: true, All: true, SaveGeter: myredis.NewRedis(myredis.REQUESTS)}
// messages, _ := general.CheckAnswer(ps, transaction)
// if ps.Error != nil {
// log.Println(ps.Error.Error())
// return
// }
// if len(messages) < 1 {
// fmt.Println("No new answers")
// return
// }
// for _, msg := range messages {
// fmt.Println(msg.Subject(), string(msg.Data()))
// }
// }
