# wordcount

A small clone of the Unix `wc` command, written in Go. It counts the lines, words, and characters in files or in standard input.

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
wordcount [-l] [-w] [-m] [file ...]
```

| Flag | Prints |
|---|---|
| `-l` | lines |
| `-w` | words |
| `-m` | characters |

With no flags, all three are printed. Columns always appear in the order lines, words, characters, followed by the file name.

```
$ wordcount testdata/unicode.txt go.mod
0 2 7 testdata/unicode.txt
3 4 53 go.mod
3 6 60 total

$ wordcount -l -w go.mod
3 4 go.mod
```

When more than one file is given, a `total` line is printed at the end.

With no files, it reads from standard input:

```
$ echo "hello world" | wordcount
1 2 12
```

If a file can't be opened, it prints an error to stderr, carries on with the remaining files, and exits with status 1 at the end.

## How it counts

The counts match `wc -lwm`:

- **Lines**: the number of newline (`\n`) characters. A last line without a trailing newline isn't counted, the same as `wc`.
- **Words**: runs of non-whitespace characters, where whitespace is anything `unicode.IsSpace` accepts (spaces, tabs, newlines, and so on).
- **Characters**: Unicode characters (runes), not bytes. `héllo 🚀` is 7 characters but 11 bytes.

## Tests

```sh
go test ./...
```

## Limitations

- Flags have to be written separately: `-l -w` works but `-lw` doesn't. Go's standard `flag` package doesn't support bundled short flags.
- No `-c` (byte count).
