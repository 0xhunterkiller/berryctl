package cmd

import (
	"fmt"
	"strings"
)

func parseAccess(entries []string) map[string][]string {
	accessMap := make(map[string][]string)

	for _, entry := range entries {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Skipping invalid access entry: %s (expected format: resource=verb1,verb2)\n", entry)
			continue
		}
		resource := strings.TrimSpace(parts[0])
		verbs := strings.Split(parts[1], ",")
		for i := range verbs {
			verbs[i] = strings.TrimSpace(verbs[i])
		}
		accessMap[resource] = verbs
	}

	return accessMap
}
