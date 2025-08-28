
# Tydi - File Organiser CLI

Tydi is a simple CLI tool to organise files in a directory based on grouping patterns.  
Currently supported grouping patterns:  
- **extension** → group files by file extension  
- **prefix** → group files by filename prefix  

---

## Installation

```bash
git clone https://github.com/qlfzn/tydi.git
cd tydi
go build -o tydi
````

---

## Usage

```bash
./tydi -path <directory> -groupby <pattern>
```

### Options

* `-path` : Directory path to organise (default: `.`)
* `-groupby` : Grouping pattern. Options: `extension`, `prefix` (default: `extension`)

---

## Example

```bash
./tydi -path ./downloads -groupby extension
```

This will:

1. Scan `./downloads`.
2. Group files by extension.
3. Show a summary and proposed destination paths.
4. Ask for confirmation before moving files.