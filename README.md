# 1 Billion Row Challenge - Go Implementation

![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)

## 🚀 Challenge Overview
Three optimized Go implementations demonstrating progressively advanced techniques for extreme-scale data processing, achieving sub-5-second performance on billion-row datasets.

## 📂 Project Structure
1_billion_row_challenge_go/
├── data_generator/ # Data generation implementations     
    v1_basic/ # Sequential processing
├── v2_improved/ # Buffered I/O
├── v3_optimized/ # Parallel processing
└── benchmarks/ # Performance testing

## 🛠️ Usage
**Generate Test Data:**
go run data_generator/generator.go 1000000

**Run Implementations:**
Basic version

go run v1_basic/main.go data/measurements_1000000_*.csv
Improved version

go run v2_improved/main.go data/measurements_1000000_*.csv
Optimized version


## 🧠 Design Approach
### Optimization Hierarchy
1. **V1: Foundation**  
   - bufio.Scanner line reading
   - map[string]*StationStats
   - Sync.Mutex for concurrent access

2. **V2: I/O Optimization**  
   - 256KB buffered reads
   - strings.Builder reuse
   - Manual semicolon search
   - Temperature parsing optimizations

**Sample Benchmark (10M rows):**
V1: 1.26s @7.91M rows/s
V2: 1.10s @9.05M rows/s

**Benchmark (1BN rows):**
V1: 128s @ 7.76M rows/s
V2: 114s @ 8.72M rows/s