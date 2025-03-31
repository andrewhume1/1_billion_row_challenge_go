# 1 Billion Row Challenge - Go Implementation

![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)

## 🚀 Challenge Overview
Two optimized Go implementations demonstrating progressively advanced techniques for extreme-scale data processing.

## 📂 Project Structure

- 📂 __1\_billion\_row\_challenge\_go__
   - 📄 [README.md](README.md)
   - 📂 __data__
   - 📂 __data\_generator__
     - 📄 [data\_generator.go](data_generator/data_generator.go)
   - 📂 __src__
     - 📂 __v1__
       - 📄 [v1\_basic.go](src/v1/v1_basic.go)
     - 📂 __v2__
       - 📄 [v2\_implementation.go](src/v2/v2_implementation.go)



## 🧠 Design Approach

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