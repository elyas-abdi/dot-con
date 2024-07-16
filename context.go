package config

import (
	"os"
	"strings"
)

type detail struct {
	weight float64
	value  string
}

func (c *Config) addContext(key string, value string, factors map[string]string, weight float64, calculateWeight bool) (bool, float64) {
	if calculateWeight {
		weight = c.calculateWeight(factors)
	}

	if len(factors) > 0 && weight == 0 {
		return false, weight
	}

	det, ok := c.context[key]
	if !ok {
		c.context[key] = detail{
			value:  value,
			weight: weight,
		}
		return true, weight
	}

	if weight >= det.weight {
		c.context[key] = detail{
			value:  value,
			weight: weight,
		}
		return true, weight
	}

	return false, weight
}

func (c *Config) calculateWeight(factors map[string]string) float64 {
	matched := 0
	for factor, def := range factors {
		_, ok := c.args[factor]
		if !ok {
			argFromEnv, ok := os.LookupEnv(factor)
			if ok {
				c.args[factor] = strings.ToUpper(argFromEnv)
			} else {
				c.args[factor] = "*"
			}
		}

		if matchesEnumPattern(def) {
			items := strings.Split(def, " | ")
			for _, item := range items {
				if c.args[factor] == item {
					matched++
					break
				}
			}
		} else {
			if c.args[factor] != def && def != "*" {
				return 0
			}

			if c.args[factor] == def || def == "*" {
				matched++
			}
		}

	}

	if matched == 0 {
		return 0
	}

	return float64(matched / len(factors))
}
