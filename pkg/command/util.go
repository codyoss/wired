package command

import (
	"strconv"
	"strings"
	"sync"
)

// parseID reads the command lin argument that should be a id to look up an entity.
func parseID(args []string) (int, bool) {
	if len(args) != 1 {
		return 0, false
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, false
	}

	return id, true
}

// extractIDs takes the ids out of Referer link strings.
func extractIDs(urls []string) []int {
	var ids []int
	for _, v := range urls {
		splitURL := strings.Split(v, "/")
		if len(splitURL) <= 2 {
			continue
		}
		sID := splitURL[len(splitURL)-2]
		id, err := strconv.Atoi(sID)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

// collectNames fans out work to collect names for future templating.
func collectNames(ids []int, block func(id int) string) []string {
	var wg sync.WaitGroup
	ch := make(chan string, 10)
	done := make(chan bool)
	var names []string

	for _, v := range ids {
		wg.Add(1)
		go func(id int) {
			name := block(id)
			if name == "" {
				return
			}
			ch <- name
			wg.Done()
		}(v)
	}

	go func() {
		for name := range ch {
			names = append(names, name)
		}
		close(done)
	}()

	wg.Wait()
	close(ch)
	<-done

	return names
}
