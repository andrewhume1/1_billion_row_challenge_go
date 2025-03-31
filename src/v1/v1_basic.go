package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StationStats struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	// Check command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run v1_basic.go <data_file>")
		os.Exit(1)
	}

	dataFile := os.Args[1]
	fmt.Printf("Processing file: %s\n", dataFile)
	startTime := time.Now()

	// Open file
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a map to store statistics for each station
	stationStats := make(map[string]*StationStats)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	lineCount := int64(0)

	// Read and process each line
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// Parse station and temperature
		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			fmt.Printf("Warning: Invalid line format at line %d: %s\n", lineCount, line)
			continue
		}

		station := parts[0]
		temperature, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Printf("Warning: Invalid temperature at line %d: %s\n", lineCount, parts[1])
			continue
		}

		// Update statistics for this station
		stats, exists := stationStats[station]
		if !exists {
			stationStats[station] = &StationStats{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			// Update existing stats
			if temperature < stats.Min {
				stats.Min = temperature
			}
			if temperature > stats.Max {
				stats.Max = temperature
			}
			stats.Sum += temperature
			stats.Count++
		}

		// Print progress every 10 million lines
		if lineCount%10_000_000 == 0 {
			elapsedTime := time.Since(startTime).Seconds()
			linesPerSecond := float64(lineCount) / elapsedTime
			fmt.Printf("Processed %d lines (%.2f lines/s)\n", lineCount, linesPerSecond)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Sort stations for display
	var stations []string
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	// Calculate and display results
	fmt.Println("\nResults:")
	fmt.Printf("%-15s %8s %8s %8s\n", "Station", "Min", "Mean", "Max")
	fmt.Println(strings.Repeat("-", 41))

	for _, station := range stations {
		stats := stationStats[station]
		mean := stats.Sum / float64(stats.Count)
		fmt.Printf("%-15s %8.1f %8.1f %8.1f\n", station, stats.Min, mean, stats.Max)
	}

	// Calculate total processed rows
	var totalRows int64
	for _, stats := range stationStats {
		totalRows += stats.Count
	}

	elapsedTime := time.Since(startTime).Seconds()
	fmt.Printf("\nProcessed %d rows in %.2f seconds (%.2f rows/s)\n", totalRows, elapsedTime, float64(totalRows)/elapsedTime)
}