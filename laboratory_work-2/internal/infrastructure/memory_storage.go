package infrastructure

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/n1jke/oop-bsuir-2025/lr-2/internal/domain"
)

var ErrNotFound = errors.New("not found")

type CacheStorage[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewCacheStorage[K comparable, V any]() *CacheStorage[K, V] {
	return &CacheStorage[K, V]{
		mu:   sync.RWMutex{},
		data: make(map[K]V),
	}
}

func (cs *CacheStorage[K, V]) Save(k K, v V) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.data[k] = v
}

func (cs *CacheStorage[K, V]) ByID(k K) (*V, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	val, exist := cs.data[k]
	if !exist {
		return nil, ErrNotFound
	}

	return &val, nil
}

func (cs *CacheStorage[K, V]) Delete(k K) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, exists := cs.data[k]; !exists {
		return ErrNotFound
	}

	delete(cs.data, k)
	return nil
}

type AccountCache struct {
	*CacheStorage[uuid.UUID, *domain.Account]
}

func NewAccountCacheStorage() *AccountCache {
	cs := NewCacheStorage[uuid.UUID, *domain.Account]()
	return &AccountCache{cs}
}

func (ac *AccountCache) UpdateStatus(accountID uuid.UUID, status domain.AccountStatus) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	account, exists := ac.data[accountID]
	if !exists {
		return ErrNotFound
	}

	account.ChangeStatus(status)
	return nil
}

type EventCache struct {
	*CacheStorage[uuid.UUID, *domain.Event]
}

func NewEventCacheStorage() *EventCache {
	cs := NewCacheStorage[uuid.UUID, *domain.Event]()
	return &EventCache{cs}
}

func (ec *EventCache) QueryAll() []*domain.Event {
	ec.mu.RLock()
	defer ec.mu.RUnlock()

	events := make([]*domain.Event, len(ec.data), 0)
	for _, v := range ec.data {
		events = append(events, v)
	}

	return events
}
