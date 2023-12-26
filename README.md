# **Parallel-Word-Counter**

## **Introduction**
Implemented project that uses parallel processing techniques, using goroutines, to count occurrences of user-specified words
concurrently in a text file. The program provides efficient word counting with exclusion of repetitions, providing optimised
solution for analyzing large text files.
## Installation
Clone this repository
  ```bash
    git clone https://github.com/NRKA/Parallel-Word-Counter.git
```
## How to Run
```bash
  go run cmd/main.go
```
## **After running the project:**
  ``` 
  Please enter words separated by spaces: that there i then they
  ```
## **Result of counting:**
```
there: 1365
i: 72345
then: 0
they: 0
total: 77805
```
## **Run unit tests**
```
  go test ./...
```

## **Run unit tests with coverage**
```
go test ./... -cover
```
