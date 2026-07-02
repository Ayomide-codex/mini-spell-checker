# Mini Spell Checker

A Go program that checks words against a dictionary and suggests the closest matching word when a misspelling is detected.

## Example

**Dictionary:**
```
hello
world
golang
project
```

**Input:**
```
helo wrld
```

**Output:**
```
Did you mean:
hello
world
```

## How It Works

1. The input string is split into individual words.
2. Each word is compared against every word in the dictionary.
3. If a word is similar enough to a dictionary word (based on length difference and matching characters), it gets suggested.
4. All suggestions are printed under "Did you mean:".

## Similarity Logic

Two words are considered similar if:
- Their lengths differ by **2 or less**.
- At least `len(dictionaryWord) - 2` characters match at the same position.

This keeps the logic simple and readable without needing a complex algorithm.

## Project Structure

```
mini-spell-checker/
├── main.go          # Entry point — defines dictionary, input, and prints suggestions
├── checker.go       # findClosest — loops through dictionary to find a match
└── similarity.go    # isSimilar — compares two words by length and character matching
```

All files share `package main` and compile together as one program.

## Concepts Demonstrated

- **Byte-level string comparison** — checking characters position by position.
- **Slice of strings** — used for both the dictionary and collecting suggestions.
- **Helper functions** — logic split cleanly across files by responsibility.
- **Multi-file structure** — entry point, matching logic, and similarity check separated.

## Running the Program

From inside the project folder:

```bash
go run .
```

> Use `go run .` (not `go run main.go`) since the project spans multiple files.

## Edge Cases Handled

- A correctly spelled word, e.g. `"hello"` → still matched and suggested (0 differences).
- A completely unrecognisable word, e.g. `"zzzzz"` → no suggestion printed.
- Empty input `""` → `strings.Fields` returns an empty slice, nothing crashes.
- A word close to multiple dictionary words → each match is suggested independently.

## Author

Ayomide Ajisegiri (Show Manny)