package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NYTimes/gizmo/pubsub"
	"sync"
	"testing"
	"time"
)

type passthrough struct {
	id     int64
	start  int
	cancel int
	filter *Filter
	//streamIdFn func() int64
	passthroughStream chan *StreamEventsResponse
	expectedOutput    []*StreamEventsResponse
}

func TestStartPipeline(t *testing.T) {
	tests := []struct {
		name         string
		inputEvents  []*Event
		passthroughs []passthrough
	}{
		{
			name: "test1",
			inputEvents: []*Event{
				{Msg: "test 1", Id: 1, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 2", Id: 2, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 3", Id: 3, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 4", Id: 4, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
			},
			passthroughs: []passthrough{
				{
					id:                1,
					start:             1,
					cancel:            3,
					filter:            nil,
					passthroughStream: make(chan *StreamEventsResponse),
					expectedOutput: []*StreamEventsResponse{
						{
							StreamId: &StreamId{},
							Event:    &Event{},
						}, {
							StreamId: &StreamId{Id: 1},
							Event:    &Event{Msg: "test 2", Id: 2, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
						}, {
							StreamId: &StreamId{Id: 1},
							Event:    &Event{Msg: "test 3", Id: 3, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
						}, {
							StreamId: &StreamId{},
							Event:    &Event{},
						},
					},
				},
			},
		}, {
			name: "test2",
			inputEvents: []*Event{
				{Msg: "test 1", Id: 11, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 2", Id: 21, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 3", Id: 31, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 4", Id: 41, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 5", Id: 51, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 6", Id: 61, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
				{Msg: "test 7", Id: 71, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
			},
			passthroughs: []passthrough{
				{
					id:                1,
					start:             1,
					cancel:            3,
					filter:            nil,
					passthroughStream: make(chan *StreamEventsResponse),
					expectedOutput: []*StreamEventsResponse{
						{
							StreamId: &StreamId{},
							Event:    &Event{},
						}, {
							StreamId: &StreamId{Id: 1},
							Event:    &Event{Msg: "test 2", Id: 2, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
						}, {
							StreamId: &StreamId{Id: 1},
							Event:    &Event{Msg: "test 3", Id: 3, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
						}, {
							StreamId: &StreamId{},
							Event:    &Event{},
						},
					},
				},
				{
					id:                2,
					start:             2,
					cancel:            5,
					filter:            nil,
					passthroughStream: make(chan *StreamEventsResponse),
					expectedOutput: []*StreamEventsResponse{
						{
							StreamId: &StreamId{},
							Event:    &Event{},
						}, {
							StreamId: &StreamId{},
							Event:    &Event{},
						}, {
							StreamId: &StreamId{Id: 2},
							Event:    &Event{Msg: "test 3", Id: 3, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
						}, {
							StreamId: &StreamId{2},
							Event:    &Event{Msg: "test 4", Id: 4, Severity: 0, System: "test", Timestamp: time.Now().Unix()},
						}, {
							StreamId: &StreamId{2},
							Event:    &Event{Msg: "test 5", Id: 5, Severity: 01, System: "test", Timestamp: time.Now().Unix()},
						},{
							StreamId: &StreamId{},
							Event:    &Event{},
						},{
							StreamId: &StreamId{},
							Event:    &Event{},
						},
					},
				},
			},
		},
	}
	ctx, done := context.WithCancel(context.Background())
	defer done()

	for _, test := range tests {
		eventsInputStream := make(chan pubsub.SubscriberMessage)

		s := Service{}
		s.passthroughCh = make(chan chan *StreamEventsResponse)
		s.stopPassthroughCh = make(chan int64)

		go s.StartPipeline(ctx, eventsInputStream)

		t.Run(test.name, func(ts *testing.T) {
			for i, e := range test.inputEvents {
				if streams := shouldStart(test.passthroughs, i); len(streams) > 0 {
					go func(streams []int){
						for _, j := range streams {
							s.NewStreamId = test.passthroughs[j].testStreamId
							fmt.Printf("sending passthrough ch %d\n", test.passthroughs[j].id)
							var mux sync.Mutex
							mux.Lock()
							s.passthroughCh <- test.passthroughs[j].passthroughStream
							mux.Unlock()
						}
					}(streams)
				}

				if streams := shouldCancel(test.passthroughs, i); len(streams) > 0 {
					go func(streams []int) {
						for _, j := range streams {
							fmt.Printf("sending cancel  %d\n", test.passthroughs[j].id)
							var mux sync.Mutex
							mux.Lock()
							s.stopPassthroughCh <- test.passthroughs[j].testStreamId()
							mux.Unlock()
						}
					}(streams)
				}

				time.Sleep(time.Millisecond * 1)
				bResp, _ := json.Marshal(e)
				msg := &TestMsg{msg: bResp}
				eventsInputStream <- msg
				for _, p := range test.passthroughs {

					if p.start <= i && p.cancel > i {
						fmt.Printf("passt picked: %d\n", p.id)
						go func(pt passthrough, i int) {
							fmt.Printf("listen to passthrough: %d\n", pt.id)
							result := <-pt.passthroughStream
							if result != nil{
								fmt.Printf("receiving passthrough %v in %d\n", result, pt.id)

								if result.StreamId.Id != pt.expectedOutput[i].StreamId.Id {
									ts.Errorf("expected streamId %d got %d", pt.expectedOutput[i].StreamId, result.StreamId)
								}

								if applyFilter(pt.filter, result.Event) {
									if result.Event.Msg != pt.expectedOutput[i].Event.Msg {
										ts.Errorf("expected %+v\n got %+v", pt.expectedOutput[i].Event, result.Event)
									}
								}
							}

						}(p, i)
					}
				}
			}

		})
	}
}

func (p *passthrough) testStreamId() int64 {
	return p.id
}

type TestMsg struct {
	msg []byte
}

func (t *TestMsg) Message() []byte {
	return t.msg
}
func (t *TestMsg) ExtendDoneDeadline(time.Duration) error {
	return nil
}
func (t *TestMsg) Done() error {
	return nil
}

func shouldStart(passtrhoughs []passthrough, i int) []int {
	shouldStart := make([]int, 0)
	for id, p := range passtrhoughs {
		if p.start == i {
			shouldStart = append(shouldStart, id)
		}
	}
	return shouldStart
}

func shouldCancel(passtrhoughs []passthrough, i int) []int {
	shouldCancel := make([]int, 0)
	for id, p := range passtrhoughs {
		if p.cancel == i {
			shouldCancel = append(shouldCancel, id)
		}
	}
	return shouldCancel
}
