package main

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type Topic struct {
	id          int
	title       string
	content     string
	create_time int64
}
type Post struct {
	id          int
	parent_id   int
	content     string
	create_time int64
}

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[int64(topic.id)] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}
func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		var i = int64(post.parent_id)
		postTmpMap[i] = append(postTmpMap[i], &post)
	}
	postIndexMap = postTmpMap
	return nil
}

//查询

// 单例模式
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

// 结构体内部方法
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}

// 回帖实现
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(id int64) []*Post {
	return postIndexMap[id]
}
