package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func (s *TonApiServer) ServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error
	var m interface{}
	ch := make(chan interface{}, 0)

	go func() {
		s.apiLock.Lock()
		m, err = handler(ctx, req)
		s.apiLock.Unlock()
		if err != nil {
			log.Printf("RPC failed with error %v", err)
			s.apiLock.Lock()
			s.api.UpdateTonConnection()
			s.apiLock.Unlock()
			return
		}
		ch <- m
	}()

	select {
	case <-ctx.Done():
		log.Println("time to return")
		s.apiLock.Lock()
		s.api.UpdateTonConnection()
		s.apiLock.Unlock()
		return nil, fmt.Errorf(ctx.Err().Error())
	case <-ch:
		return m, err
	}
}
