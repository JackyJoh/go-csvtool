# csvtool

A small CLI tool for poking around CSV files from the terminal. Built as a Go learning project before starting [JacksonDB](#), a columnar storage engine.

## Install

```bash
go install github.com/JackyJoh/go-csvtool/csvtool@latest
```

Or clone and build locally:
```bash
git clone https://github.com/JackyJoh/go-csvtool.git
cd go-csvtool/csvtool
go build -o csvtool
```

## Commands

Prints min/max/mean for numeric columns and a count for everything else.
```bash
csvtool stats <file>
```

Prints the first N rows. Defaults to 10 if `--n` isn't set.
```bash
csvtool head <file> --n 20
```

Prints every value in a given column.
```bash
csvtool col <file> --name <column_name>
```

## Example

```bash
$ csvtool stats fifa_players.csv

short_name{ count: 500 }
overall{ count: 500, min: 54, max: 90, avg: 71.03}
value_eur{ count: 500, min: 70000, max: 99500000, avg: 6999880.00}

$ csvtool head fifa_players.csv --n 3
(showing 4 of 18 columns for readability — head prints every column)

short_name, club_name, overall, value_eur  ...
L. Messi, Inter Miami, 90, 41000000  ...
A. Griezmann, Atlético Madrid, 88, 74000000  ...
Martin Ødegaard, Arsenal, 87, 99500000  ...

$ csvtool col fifa_players.csv --name club_name

Inter Miami
Atlético Madrid
Arsenal
...
```

## Dataset

The included `fifa_players.csv` is a small truncated sample (500 rows) for testing. If you want the full dataset, I published it on Kaggle: [FIFA and IRL Soccer Player Data](https://www.kaggle.com/datasets/jacksonjohannessen/fifa-and-irl-soccer-player-data).