package main

import (
	"log"
	"os"
	"runtime/coverage"
)

func writeCoverageFiles(coverDir string) error {
	if err := os.RemoveAll(coverDir); err != nil {
		log.Printf("Delete coverDir failed: %v", err)
		return err
	}

	if err := os.MkdirAll(coverDir, 0755); err != nil {
		log.Printf("create coverDir failed: %v", err)
		return err
	}

	if err := coverage.WriteMetaDir(coverDir); err != nil {
		log.Printf("Cover writing failed: %v", err)
		return err
	}

	if err := coverage.WriteCountersDir(coverDir); err != nil {
		log.Printf("Cover writing failed: %v", err)
		return err
	}
	return nil
}
