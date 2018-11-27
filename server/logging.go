package server

import (
	"context"
	"fmt"
	"encoding/json"
)

func (s *service) StartInputstream(ctx context.Context) (error) {
	eventsInputStream := s.inputStream.Start()
	var passthroughStream chan *Event
	passthroughflag := false

	//todo: check race conditions for this whole block
	for {
		select {
		case msg := <-eventsInputStream:
			fmt.Printf("%s", msg.Message())
			var event *Event
			json.Unmarshal(msg.Message(), event)

			//todo store in ds
			fmt.Printf("e: %+v\n", *event)

			if passthroughflag {
				if passthroughStream != nil {
					passthroughStream <- event
				}
			}
		case eventStream := <-s.PassthroughCh:
			passthroughStream = eventStream
			passthroughflag = true

		case <-s.StopPassthroughCh:
			passthroughflag = false
			close(passthroughStream)
			passthroughStream = nil
		}
	}

	return nil
}

func (s *service) LoadEvents(ctx context.Context, loadEventsRequest *LoadEventsRequest) (*LoadEventsResponse, error) {
	return &LoadEventsResponse{}, nil
}

func (s *service) StreamEvents(streamEventsRequest *StreamEventsRequest, streamEventsServer EventLogger_StreamEventsServer) error {
	eventFilter := streamEventsRequest.GetFilter()
	
	eventsStream := make(chan *Event)
	s.PassthroughCh<- eventsStream

	select {
	case event := <-eventsStream:
		if applyFilter(eventFilter, event) {
			streamEventResp := &StreamEventsResponse{
				Event:event,
			}
			streamEventsServer.Send(streamEventResp)
		}
	}
	return nil
}

func (s *service) StopStreaming(context.Context, *Void) (*Void, error) {
	s.StopPassthroughCh<- true
	return &Void{}, nil
}

func applyFilter(f *Filter, e *Event) bool {
	return true
}