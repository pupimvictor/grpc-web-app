package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NYTimes/gizmo/pubsub"
	"sync"
	"time"
)

func (s *Service) StartPipeline(ctx context.Context, eventsInputStream <-chan pubsub.SubscriberMessage) error {
	passthroughStream := make(map[int64]chan *StreamEventsResponse)
	var mux sync.Mutex

	for {
		select {
		case msg := <-eventsInputStream:
			var event Event
			json.Unmarshal(msg.Message(), &event)

			//todo store in ds
			fmt.Printf("e: %+v\n", event.Msg)

			mux.Lock()
			if len(passthroughStream) > 0 {
				for id, eventCh := range passthroughStream {
					streamResp := &StreamEventsResponse{
						StreamId: &StreamId{id},
						Event:    &event,
					}
					fmt.Printf("send msg %d to stream %d\n", streamResp.Event.Id, id)
					eventCh<- streamResp
				}
			}
			mux.Unlock()
		case eventStream := <-s.passthroughCh:
			streamId := s.NewStreamId()
			fmt.Printf("receive passthrough ch for stream %d\n", streamId)
			mux.Lock()
			passthroughStream[streamId] = eventStream
			mux.Unlock()

		case streamId := <-s.stopPassthroughCh:
			fmt.Printf("receive cancel for stream %d\n", streamId)
			mux.Lock()
			close(passthroughStream[streamId])
			delete(passthroughStream, streamId)
			mux.Unlock()

		case <-ctx.Done():
			fmt.Printf("input stream closed\n")
			return fmt.Errorf("input stream closed")
		}
	}
	return nil
}

func (s *Service) LoadEvents(ctx context.Context, loadEventsRequest *LoadEventsRequest) (*LoadEventsResponse, error) {
	return &LoadEventsResponse{}, nil
}

func (s *Service) StreamEvents(streamEventsRequest *StreamEventsRequest, streamEventsServer EventLogger_StreamEventsServer) error {
	eventFilter := streamEventsRequest.GetFilter()

	eventsStream := make(chan *StreamEventsResponse)
	s.passthroughCh <- eventsStream

	for eventResp := range eventsStream {
		if applyFilter(eventFilter, eventResp.Event) {
			streamEventsServer.Send(eventResp)
		}
	}
	return nil
}

func (s *Service) StopStreaming(context.Context, *Void) (*Void, error) {
	s.stopPassthroughCh <- 1
	return &Void{}, nil
}

func applyFilter(f *Filter, e *Event) bool {
	return true
}

func NewStreamId() int64{
	return time.Now().Unix()
}
