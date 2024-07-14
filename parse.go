package con

import (
	`bufio`
	`os`
	`path/filepath`
	`strings`
	`sync`
)

type parsedItem struct {
	Key     string
	Value   string
	Factors map[string]string
}

func (c *Con) extract(line string, factors map[string]string) (*parsedItem, error) {
	item := parsedItem{
		Factors: factors,
	}
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return nil, nil
	}

	if matchesAssignment(line) {
		assignment := strings.Split(line, assignmentSeparator)
		item.Key = strings.TrimSpace(assignment[0])
		item.Value = strings.TrimSpace(assignment[1])
		return &item, nil
	}

	return nil, nil
}

func (c *Con) parseDir() error {
	if len(c.dir) == 0 {
		c.dir = defaultConDir
	}

	fileCh := make(chan *string)
	resultCh := make(chan *parsedItem)
	var wg sync.WaitGroup

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileCh <- &path
		}

		return nil
	}

	go func() {
		err := filepath.Walk(c.dir, walkFn)
		if err != nil {
			close(fileCh)
			panic(err)
		}

		close(fileCh)
	}()

	go func() {
		for path := range fileCh {
			go func(path *string) {
				wg.Add(1)
				defer wg.Done()

				file, err := os.Open(*path)
				if err != nil {
					close(fileCh)
					close(resultCh)
					panic(err)
				}
				defer file.Close()

				factors := make(map[string]string)
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					item, err := c.extract(scanner.Text(), factors)
					if err != nil {
						close(fileCh)
						close(resultCh)
						panic(err)
					}

					factors = item.Factors

					resultCh <- item
				}
			}(path)
		}
	}()

	go func() {
		for result := range resultCh {
			c.addContext(result.Key, result.Value)
		}
	}()

	wg.Wait()
	close(resultCh)

	return nil
}
