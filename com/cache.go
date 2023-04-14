package com

import (
	"errors"
	"fmt"
	"github.com/fanjindong/go-cache"
	"time"
)

// Cache 缓存工具
type Cache struct {
	memCache cache.ICache
}

// Create 创建缓存
func (c *Cache) Create() {
	c.memCache = cache.NewMemCache(cache.WithClearInterval(1 * time.Second))
}

// Set 设置
func (c *Cache) Set(key string, value string, exTime time.Duration) error {
	if c.memCache == nil {
		return errors.New("未创建缓存")
	}
	c.memCache.Set(key, value, cache.WithEx(exTime))
	return nil
}

// Get 获取
func (c *Cache) Get(key string) (string, error) {
	if c.memCache == nil {
		return "", errors.New("未创建缓存")
	}
	get, bool := c.memCache.Get(key)
	if bool {
		return fmt.Sprintf("%v", get), nil
	}
	return "", nil
}

// Delete 删除
func (c *Cache) Delete(keys ...string) error {
	if c.memCache == nil {
		return errors.New("未创建缓存")
	}
	for _, key := range keys {
		_ = c.memCache.Del(key)
	}
	return nil
}
