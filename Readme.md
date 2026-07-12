# csvtool

A small CLI tool for poking around CSV files from the terminal. Built as a Go learning project before starting [JacksonDB](#), a columnar storage engine.

## Why this exists

I hadn't touched Go before this. Rather than learning the language in a vacuum, I built something small and useful: a command-line tool that loads a CSV and answers basic questions about it. No frameworks, no external dependencies beyond the standard library.

## Commands

```bash
csvtool stats <file>
```
Prints min/max/mean for numeric columns and a count for everything else.

```bash
csvtool head <file> --n 20
```
Prints the first N rows. Defaults to 10 if `--n` isn't set.

```bash
csvtool col <file> --name category
```
Prints every value in a given column.

## Example

```bash
$ csvtool stats markets.csv

category      count=1822
yes_ask       min=0.02  max=0.98  mean=0.51
volume        min=0     max=48210 mean=1204.6

$ csvtool head markets.csv --n 3

ticker              category   yes_ask   volume
KXBTCD-26JUL0517     crypto     0.87      412
KXETHD-26JUL0517     crypto     0.62      198
KXPOL-26JUL0517       politics  0.31      3021

$ csvtool col markets.csv --name category

crypto
crypto
politics
...
```

## What it's not

No writing or editing CSVs. No filtering, joins, or multi-file support. Numeric column detection is automatic, not configurable. This is intentionally scoped small — it's a warmup, not a product.

## Built with

Go standard library only (`encoding/csv`, `os`, `bufio`, `strconv`).