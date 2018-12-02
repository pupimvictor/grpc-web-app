package server

import (
	"context"
	"encoding/json"
	"github.com/NYTimes/gizmo/pubsub"
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

	for _, test := range tests {
		eventsInputStream := make(chan pubsub.SubscriberMessage)

		ctx, _ := context.WithCancel(context.Background())
		s := Service{}
		s.passthroughCh = make(chan chan *StreamEventsResponse)
		s.stopPassthroughCh = make(chan int64)

		go s.StartPipeline(ctx, eventsInputStream)

		t.Run(test.name, func(ts *testing.T) {
			for i, e := range test.inputEvents {
				if streams := shouldStart(test.passthroughs, i); len(streams) > 0 {
					for _, j := range streams {
						s.NewStreamId = test.passthroughs[j].testStreamId
						s.passthroughCh <- test.passthroughs[j].passthroughStream
					}
				}
				if streams := shouldCancel(test.passthroughs, i); len(streams) > 0 {
					for _, j := range streams {
						s.stopPassthroughCh <- test.passthroughs[j].testStreamId()
					}
				}

				bResp, _ := json.Marshal(e)
				msg := &TestMsg{msg: bResp}
				eventsInputStream <- msg

				for _, p := range test.passthroughs {
					if p.start <= i && p.cancel > i {
						result := <-p.passthroughStream

						if result.StreamId.Id != p.expectedOutput[i].StreamId.Id {
							ts.Errorf("expected streamId %d got %d", p.expectedOutput[i].StreamId, result.StreamId)
						}

						if applyFilter(p.filter, result.Event) {
							if result.Event.Msg != p.expectedOutput[i].Event.Msg {
								ts.Errorf("expected %+v\n got %+v", p.expectedOutput[i].Event, result.Event)
							}
						}
					}
				}
			}
		})
		//close(eventsInputStream)
		//done()
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
