package con

import (
	`bufio`
	`fmt`
	`os`
	`path/filepath`
	`strings`
	`sync`
)

type parsedItem struct {
	Key                 string
	Value               string
	Factors             map[string]string
	NewFactorDiscovered bool
}

func (c *Con) extract(line string, factors map[string]string) (*parsedItem, error) {
	item := parsedItem{
		Factors:             factors,
		NewFactorDiscovered: false,
	}
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return &item, nil
	}

	if matchesAssignmentPattern(line) {
		assignment := strings.Split(line, assignmentSeparator)
		item.Key = strings.TrimSpace(assignment[0])
		item.Value = strings.TrimSpace(assignment[1])
		return &item, nil
	}

	if matchesFactorsPattern(line) {
		line = line[1 : len(line)-1]
		definition := strings.Split(line, fmt.Sprintf("%s %s", factorsSuffix, factorsPrefix))

		for _, def := range definition {
			pair := strings.Split(def, ":")
			item.Factors[strings.ToUpper(strings.TrimSpace(pair[0]))] = strings.ToUpper(strings.TrimSpace(pair[1]))
		}
		item.NewFactorDiscovered = true
		return &item, nil
	}

	return nil, nil
}

func (c *Con) parseDir() error {

	if len(c.dir) == 0 {
		c.dir = defaultConDir
	}

	var files []string
	resultCh := make(chan *parsedItem)
	var wg sync.WaitGroup

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".con") {
			files = append(files, path)
		}

		return nil
	}

	err := filepath.Walk(c.dir, walkFn)
	if err != nil {
		return err
	}

	for _, path := range files {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			file, err := os.Open(path)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			factors := make(map[string]string)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				item, err := c.extract(scanner.Text(), factors)
				if err != nil {
					fmt.Println("Error extracting item:", err)
					return
				}

				if item != nil {
					factors = item.Factors
				}

				resultCh <- item
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error scanning file:", err)
			}
		}(path)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	weight := 0.0
	calculate := false
	for result := range resultCh {
		if result != nil && result.Key != "" && result.Value != "" {
			_, weight = c.addContext(result.Key, result.Value, result.Factors, weight, calculate)
		} else if result != nil {
			calculate = result.NewFactorDiscovered
		}
	}

	fmt.Println("Processing complete.")
	return nil
}
