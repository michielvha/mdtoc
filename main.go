package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	// Get directory from args or use current directory
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	outputFile := "INDEX.md"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	if err := generateTOC(dir, outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated %s\n", outputFile)
}

func generateTOC(rootDir, outputFile string) error {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var lines []string
	lines = append(lines, "## Table of Content")

	// Get all subdirectories
	var folders []string
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			folders = append(folders, entry.Name())
		}
	}
	sort.Strings(folders)

	// Process each folder
	for _, folder := range folders {
		folderPath := filepath.Join(rootDir, folder)
		mdFiles, err := filepath.Glob(filepath.Join(folderPath, "*.md"))
		if err != nil {
			continue
		}

		if len(mdFiles) == 0 {
			continue
		}

		sort.Strings(mdFiles)

		// Add folder header
		lines = append(lines, fmt.Sprintf("### [%s](%s)", folder, folder))

		// Add each markdown file
		for _, mdFile := range mdFiles {
			title := extractTitle(mdFile)
			relativePath := filepath.Join(folder, filepath.Base(mdFile))
			lines = append(lines, fmt.Sprintf("- [%s](%s)", title, relativePath))
		}

		lines = append(lines, "") // blank line between sections
	}

	// Write to output file
	outputPath := filepath.Join(rootDir, outputFile)
	content := strings.Join(lines, "\n")
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

func extractTitle(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return filenameToTitle(filepath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "#"))
		}
	}

	// Fallback to filename
	return filenameToTitle(filepath)
}

func filenameToTitle(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}
