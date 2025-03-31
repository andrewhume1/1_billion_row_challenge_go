package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
    // Check command line arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run generator.go <number_of_rows>")
        os.Exit(1)
    }

    // Parse number of rows
    numRows, err := strconv.Atoi(os.Args[1])
    if err != nil || numRows <= 0 {
        fmt.Println("Error: Invalid number of rows. Please provide a positive integer.")
        os.Exit(1)
    }

    // Define weather station names
    stations := []string{
        "Hamburg", "Berlin", "Munich", "Cologne", "Frankfurt",
        "Stuttgart", "Düsseldorf", "Leipzig", "Dortmund", "Essen",
        "Bremen", "Dresden", "Hanover", "Nuremberg", "Duisburg",
        "Bochum", "Wuppertal", "Bielefeld", "Bonn", "Münster",
    }

    // Create data directory if it doesn't exist
    os.MkdirAll("data", os.ModePerm)

    // Generate output file name with timestamp
    timestamp := time.Now().Format("20060102_150405")
    outputFile := filepath.Join("data", fmt.Sprintf("measurements_%d_%s.csv", numRows, timestamp))

    // Open file for writing
    file, err := os.Create(outputFile)
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    fmt.Printf("Generating %d rows of weather data...\n", numRows)
    startTime := time.Now()

    // Initialize random number generator
    rand.Seed(time.Now().UnixNano())

    // Generate and write data
    for i := 0; i < numRows; i++ {
        station := stations[rand.Intn(len(stations))]
        // Generate temperature between -30.0 and 45.0 degrees Celsius
        temperature := -30.0 + rand.Float64()*75.0
        // Round to 1 decimal place
        temperature = float64(int(temperature*10+0.5)) / 10

        // Write to file
        _, err := fmt.Fprintf(file, "%s;%.1f\n", station, temperature)
        if err != nil {
            fmt.Printf("Error writing to file: %v\n", err)
            os.Exit(1)
        }

        // Print progress every 10 million rows
        if (i+1)%10_000_000 == 0 {
            progress := float64(i+1) / float64(numRows) * 100
            elapsedTime := time.Since(startTime).Seconds()
            rowsPerSecond := float64(i+1) / elapsedTime
            fmt.Printf("Progress: %.2f%% (%d rows, %.2f rows/s)\n", progress, i+1, rowsPerSecond)
        }
    }

    elapsedTime := time.Since(startTime).Seconds()
    fmt.Printf("Done! Generated %d rows in %.2f seconds.\n", numRows, elapsedTime)
    fmt.Printf("Output file: %s\n", outputFile)
}