package history

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type MsgHistory struct {
	store    map[string]*list.Element
	list     *list.List // used for GC
	notifier map[string]chan *discordgo.Message
	sync.RWMutex
}

func NewMsgHistory() *MsgHistory {
	mh := &MsgHistory{}
	mh.list = list.New()
	mh.store = make(map[string]*list.Element, 0)
	mh.notifier = make(map[string]chan *discordgo.Message, 0)
	return mh
}

func (mh *MsgHistory) AddMsg(msg *discordgo.Message) {
	mh.Lock()
	defer mh.Unlock()
	if elem, exists := mh.store[msg.ID]; exists {
		mh.list.MoveToFront(elem)
		return
	} else {
		mh.list.PushFront(msg)
		mh.NotifyAll(msg)
	}
}

func (mh *MsgHistory) GC(maxLifeTime time.Duration) {
	time.AfterFunc(
		maxLifeTime,
		func() {
			mh.Lock()
			defer mh.Unlock()
			elem := mh.list.Back()
			if elem == nil {
				mh.GC(maxLifeTime)
				return
			}
			msg := elem.Value.(*discordgo.Message)
			var lstModTime *time.Time
			if editTime := msg.EditedTimestamp; editTime != nil {
				lstModTime = editTime
			} else {
				lstModTime = &msg.Timestamp
			}
			if time.Now().After(lstModTime.Add(maxLifeTime)) {
				mh.list.Remove(elem)
				delete(mh.store, msg.ID)
				fmt.Printf("A data has been cleared from the 'MsgHistory' cache, id:%s, content:%s\n", msg.ID, msg.Content)
			}
			mh.GC(maxLifeTime)
		},
	)
}

func (mh *MsgHistory) NotifyAll(m *discordgo.Message) {
	go func() {
		mh.RLock()
		defer mh.RUnlock()
		for _, ch := range mh.notifier {
			ch <- m
		}
	}()
}

func (mh *MsgHistory) NewFilter(
	name string,
	criteriaFunc func(m *discordgo.Message) bool,
	max int, // When this amount of data is collected, it will stop. Use -1 to ignore this setting.
	timeout time.Duration, // Set -1 if you don't want to use the timeout.
	loop bool,
	callbackFunc func([]*discordgo.Message),
) {
	for {
		ch := make(chan *discordgo.Message)

		mh.Lock()
		if _, exists := mh.notifier[name]; exists {
			panic("has existed")
		}
		mh.notifier[name] = ch
		mh.Unlock()
		fmt.Printf("Filter: %q Start\n", name)
		collect := make([]*discordgo.Message, 0)

		if timeout != -1 {
			time.AfterFunc(timeout, func() {
				fmt.Printf("timeout: %q\n", name)
				close(ch) // Do not receive messages anymore after the timeout.
			})
		}

		for {
			msg, isOpen := <-ch
			isDone := false
			if !isOpen {
				isDone = true
			} else {
				if criteriaFunc(msg) {
					collect = append(collect, msg)
				}
				if max != -1 && len(collect) >= max {
					isDone = true
				}
			}
			if isDone {
				mh.Lock()
				delete(mh.notifier, name)
				mh.Unlock()
				callbackFunc(collect)
				if loop {
					break
				} else {
					return
				}
			}
		}
	}
}
