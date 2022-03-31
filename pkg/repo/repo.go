package repo

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Repo struct {
	mu sync.Mutex
	file *os.File
}

func (repo *Repo) WriteLine(user uint64, msg string) {
	repo.mu.Lock()
	_, err := fmt.Fprintf(repo.file, "%v: `%s`\n", user, msg)
	if err != nil {
		log.Printf("can't write to file: %v", err)
	}
	repo.mu.Unlock()
}