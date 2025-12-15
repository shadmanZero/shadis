package store

import (
	"sync"

	"github.com/shadman/shadis/internal/logger"
	"go.uber.org/zap"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Store {
	logger.Debug("Creating new store")
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if ok {
		logger.Debug("Store GET hit", zap.String("key", key))
	} else {
		logger.Debug("Store GET miss", zap.String("key", key))
	}
	return val, ok
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	logger.Debug("Store SET", zap.String("key", key), zap.Int("value_len", len(value)))
}

func (s *Store) Del(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, existed := s.data[key]
	delete(s.data, key)
	
	if existed {
		logger.Debug("Store DEL", zap.String("key", key), zap.Bool("deleted", true))
	} else {
		logger.Debug("Store DEL", zap.String("key", key), zap.Bool("deleted", false))
	}
	return existed
}

func (s *Store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[key]
	logger.Debug("Store EXISTS", zap.String("key", key), zap.Bool("exists", ok))
	return ok
}

func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	logger.Debug("Store KEYS", zap.Int("count", len(keys)))
	return keys
}

func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}
