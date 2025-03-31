package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
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
		fmt.Println("Usage: go run v2_improved.go <data_file>")
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
	stationStats := make(map[string]*StationStats, 50) // Pre-allocate for typical number of stations

	// Use a larger buffer size for faster reading
	const bufferSize = 256 * 1024 // 256KB buffer
	reader := bufio.NewReaderSize(file, bufferSize)

	lineCount := int64(0)
	lastReportTime := startTime

	// Use a buffer for building strings to reduce allocations
	var lineBuf strings.Builder
	lineBuf.Grow(64) // Pre-allocate typical line size

	// Read and process each line in larger chunks
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

		if line != "" {
			// Remove trailing newline if present
			if len(line) > 0 && line[len(line)-1] == '\n' {
				line = line[:len(line)-1]
			}
			if len(line) > 0 && line[len(line)-1] == '\r' {
				line = line[:len(line)-1]
			}

			// Parse station and temperature
			semicolonIndex := strings.IndexByte(line, ';')
			if semicolonIndex == -1 {
				continue // Skip invalid lines
			}

			station := line[:semicolonIndex]
			temperatureStr := line[semicolonIndex+1:]
			temperature, err := strconv.ParseFloat(temperatureStr, 64)
			if err != nil {
				continue // Skip invalid temperatures
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

			lineCount++

			// Print progress every 10 million lines or every 5 seconds
			if lineCount%10_000_000 == 0 || time.Since(lastReportTime).Seconds() >= 5 {
				elapsedTime := time.Since(startTime).Seconds()
				linesPerSecond := float64(lineCount) / elapsedTime
				fmt.Printf("Processed %d lines (%.2f lines/s)\n", lineCount, linesPerSecond)
				lastReportTime = time.Now()

				// Force garbage collection to prevent memory build-up
				runtime.GC()
			}
		}

		if err == io.EOF {
			break
		}
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