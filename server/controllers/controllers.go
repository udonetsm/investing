package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/udonetsm/investing/actors/pubsub"
	"github.com/udonetsm/investing/actors/storage/myredis"
	"github.com/udonetsm/investing/actors/storage/postgresql"
	"github.com/udonetsm/investing/general"
	"github.com/udonetsm/investing/models"
)

const (
	SubTypeRequests = pubsub.SubTypeRequests
	SubTypeAnswers  = pubsub.SubTypeReplies
	SubTypeRead     = pubsub.SubTypeAllRead
)

const (
	All         = "all"
	Show        = "show"
	Read        = "read"
	Transaction = "transaction"
	User        = "user"
)

const (
	ErrTooFast   = "Wait for 1 sec before try else"
	ErrNoContent = "No new messages"
)

type repl models.Reply

func SubOnSomething(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	requests_replies := vars["requests-replies"]
	all_new := vars["all-new"]
	count_show_read := vars["count-show-read"]
	encoder := json.NewEncoder(w)
	ps := &pubsub.PubSub{
		SeterGeter:  myredis.NewRedis(myredis.REQUESTS),
		SubjectType: requests_replies,
		All:         all_new == All,
		Show:        count_show_read == Show,
		Read:        count_show_read == Read,
	}
	stack := general.Subcribe(ps, uid)
	if ps.Error != nil {
		encoder.Encode(&repl{Error: "Try else"})
		return
	}
	if len(stack.Mess) > 0 {
		encoder.Encode(stack.Mess)
		return
	}
	encoder.Encode(&repl{Content: stack.Count})
}

func CaptureSomething(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	something := vars["something"]
	encoder := json.NewEncoder(w)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		encoder.Encode(&repl{Error: err.Error()})
		return
	}
	object := selectSomething(something, uid, data)
	if _, ok := object.(*models.Stub); ok {
		encoder.Encode(&repl{Error: "Unsupported target object"})
		return
	}
	db := db()
	reply := general.Save(db, object)
	if db.Err != nil {
		encoder.Encode(&repl{Error: "Database error"})
		return
	}
	encoder.Encode(reply)
}

func selectSomething(something, id string, data []byte) any {
	if something == User {
		return &models.Userentry{Uid: id, Baseuser: string(data)}
	}
	if something == Transaction {
		return &models.Transactionentry{Tid: id, Transaction: string(data)}
	}
	return &models.Stub{}
}

func Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	about := vars["about"]
	id := vars["id"]
	encoder := json.NewEncoder(w)
	model := selectAbout(about, id)
	if _, ok := model.(*models.Stub); ok {
		encoder.Encode(&repl{Error: "Unsupported target"})
		return
	}
	db := db()
	reply := general.Get(db, model)
	if db.Err != nil {
		encoder.Encode(&repl{Error: "Database error"})
		return
	}
	encoder.Encode(reply)
}

func selectAbout(about, id string) any {
	if about == User {
		return &models.Userentry{Uid: id}
	}
	if about == Transaction {
		return &models.Transactionentry{Tid: id}
	}
	return &models.Stub{}
}

func db() *postgresql.Database {
	return &postgresql.Database{
		Host:    "127.0.0.1",
		User:    os.Getenv("DBUSER"),
		Pass:    os.Getenv("DBPASS"),
		Dbname:  os.Getenv("DBNAME"),
		Port:    5432,
		SSLMode: "disable",
	}
}

func PubMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid := vars["tid"]
	subtype := vars["requests-replies"]
	recieverUID := vars["uid"]
	// check recieverUID exists
	ps := &pubsub.PubSub{
		SeterGeter:  myredis.NewRedis(myredis.REQUESTS),
		Message:     tid,
		SubjectType: subtype,
	}
	encoder := json.NewEncoder(w)
	general.Publish(ps, recieverUID)
	if ps.Error != nil {
		encoder.Encode(&repl{Error: ps.Error.Error()})
		return
	}
	encoder.Encode(&repl{Content: "OK"})
}
