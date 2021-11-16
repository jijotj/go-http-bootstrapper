package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vars struct {
	missingConfigs   []string
	malformedConfigs []string
}

func (v *Vars) Mandatory(key string) string {
	val := os.Getenv(key)

	if val == "" {
		v.missingConfigs = append(v.missingConfigs, key)
		return ""
	}

	return val
}

func (v *Vars) Optional(key, fallbackValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallbackValue
	}

	return val
}

func (v *Vars) OptionalInt(key string, fallbackValue int) int {
	val := os.Getenv(key)

	if val == "" {
		return fallbackValue
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		v.malformedConfigs = append(v.malformedConfigs, key)
		return 0
	}

	return intVal
}

func (v *Vars) Error() error {
	if len(v.missingConfigs) > 0 {
		return fmt.Errorf("missing mandatory config(s): %s", strings.Join(v.missingConfigs, ", "))
	}

	if len(v.malformedConfigs) > 0 {
		return fmt.Errorf("malformed config(s): %s", strings.Join(v.malformedConfigs, ", "))
	}

	return nil
}
