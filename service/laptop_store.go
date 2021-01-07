package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/lessongo/grpc_ex/pb"
)

// ErrAlreadyExist return exists record
var ErrAlreadyExist = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutext sync.RWMutex
	data   map[string]*pb.Laptop
}

// NewInMemoryLaptopStore return a new NewInMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutext.Lock()
	defer store.mutext.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExist
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[other.Id] = other
	return nil
}

// Find finds to laptop
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutext.RLock()
	defer store.mutext.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	//deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
