# wordcount

A small clone of the Unix `wc` command, written in Go. It counts the lines, words, and characters in one or more files.

## Install

```sh
go install github.com/Olamilekan-12/wordcount@latest
```

Or build from source:

```sh
git clone https://github.com/Olamilekan-12/wordcount.git
cd wordcount
go build
```

## Usage

```
wordcount <file> [file...]
```

Output is one line per file: lines, words, characters, then the file name.

```
$ wordcount testdata/unicode.txt main.go
0 2 7 testdata/unicode.txt
69 159 1132 main.go
```

If a file can't be opened, it prints an error to stderr and exits with status 1:

```
$ wordcount nope.txt
wordcount: open nope.txt: no such file or directory
```

## How it counts

The counts match `wc -lwm`:

- **Lines**: the number of newline (`\n`) characters. A last line without a trailing newline isn't counted, the same as `wc`.
- **Words**: runs of non-whitespace characters, where whitespace is anything `unicode.IsSpace` accepts (spaces, tabs, newlines, and so on).
- **Characters**: Unicode characters (runes), not bytes. `héllo 🚀` is 7 characters but 11 bytes.

## Limitations

- No `-l`, `-w`, `-c` flags yet. All three counts are always printed.
- Can't read from stdin yet.
