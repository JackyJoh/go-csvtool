# csvtool

A small CLI tool for poking around CSV files from the terminal. Built as a Go learning project before starting [JacksonDB](#), a columnar storage engine.

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
$ csvtool stats fifa_players.csv

short_name{ count: 500 }
overall{ count: 500, min: 54, max: 90, avg: 71.03}
value_eur{ count: 500, min: 70000, max: 99500000, avg: 6999880.00}

$ csvtool head fifa_players.csv --n 3

short_name, club_name, overall, value_eur
L. Messi, Inter Miami, 90, 41000000
A. Griezmann, Atlético Madrid, 88, 74000000
Martin Ødegaard, Arsenal, 87, 99500000

$ csvtool col fifa_players.csv --name club_name

Inter Miami
Atlético Madrid
Arsenal
...
```