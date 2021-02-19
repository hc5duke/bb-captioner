# bb-captioner

caption data extractor and saver. Assumes existance of seasons and episodes.

## Requirements

- go 1.x
- postgres

## Running

### Input:

- SRT files with file name format "s[SEASON]e[EPISODE]" somewhere in the file or dir name

### Output:

- episode caption data saved in db

### Example usage:

```bash
$ go run main.go ./s04e01.txt 
```
