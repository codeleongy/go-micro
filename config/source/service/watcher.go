package service

import (
	"github.com/micro/go-micro/config/source"
	proto "github.com/micro/go-micro/config/source/service/proto"
)

type watcher struct {
	stream proto.Service_WatchService
}

func newWatcher(stream proto.Service_WatchService) (source.Watcher, error) {
	return &watcher{stream: stream}, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	var rsp proto.WatchResponse
	err := w.stream.RecvMsg(&rsp)
	if err != nil {
		return nil, err
	}
	return toChangeSet(rsp.ChangeSet), nil
}

func (w *watcher) Stop() error {
	return w.stream.Close()
}
