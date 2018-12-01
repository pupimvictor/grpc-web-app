package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NYTimes/gizmo/pubsub"
	"time"
)

func (s *Service) StartInputstream(ctx context.Context, eventsInputStream <-chan pubsub.SubscriberMessage) error {
	var passthroughStream map[int64]chan *StreamEventsResponse

	//todo: check race conditions for this whole block
	for {
		select {
		case msg := <-eventsInputStream:
			var event *Event
			json.Unmarshal(msg.Message(), event)

			//todo store in ds
			fmt.Printf("e: %+v\n", *event)

			if len(passthroughStream) > 0 {
				for id, eventCh := range passthroughStream {
					streamResp := &StreamEventsResponse{
						StreamId: &StreamId{id},
						Event:    event,
					}
					eventCh<- streamResp
				}
			}
		case eventStream := <-s.passthroughCh:
			streamId := time.Now().Unix()
			passthroughStream[streamId] = eventStream

		case streamId := <-s.stopPassthroughCh:
			close(passthroughStream[streamId])
			delete(passthroughStream, streamId)

		case <-ctx.Done():
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
