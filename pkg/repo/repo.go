package repo

import (
	"fmt"
	"os"
	"sync"
)

type Repo struct {
	mu   *sync.RWMutex // Use RLock to read
	file *os.File
}

func NewRepo(filename string) (*Repo, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}

	return &Repo{
		mu:   &sync.RWMutex{},
		file: file,
	}, nil
}

func (repo *Repo) WriteLine(id int, msg string) error {
	repo.mu.Lock()
	_, err := fmt.Fprintf(repo.file, "%v: `%s`\n", id, msg)
	repo.mu.Unlock()

	if err != nil {
		return fmt.Errorf("can't write to file: %v", err)
	}

	return nil
}
