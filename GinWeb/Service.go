package main

import "sync"

type PageInfo struct {
	Topic    *Topic
	PostList []*Post
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	// 参数检验
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

	}()
	go func() {

	}()
	wg.Wait()
	return nil
}
