# evilcombo

A Go-powered web tool for generating structured wordlist combinations.\
It creates every possible pairing of two text components (first names,
last name, relatives) and appends numeric tails derived from birth dates
or custom numbers.

------------------------------------------------------------------------

## Features

-   Input multiple:
    -   first names\
    -   relatives\
    -   custom numeric tails\
-   Birth date parsing (`DD.MM.YYYY`) producing several numeric endings:
    -   `DDMM`
    -   `DDMMYY`
    -   `DDMMYYYY`
    -   `YYYY`
    -   `YYYYMMDD`
-   Custom numbers appended exactly as typed (digits only)
-   Resulting combinations always follow the pattern:\
    **text + text + number**
-   Case variations for each combination:
    -   lowercase\
    -   capitalized components\
    -   UPPERCASE\
    -   Capitalized whole string
-   Output downloaded as `combinations.txt`
*i do not know why u need this for*

------------------------------------------------------------------------

## How It Works

1.  Launches a local server on `http://localhost:8080`
2.  The main page displays a form:
    -   First names (comma-separated)
    -   Last name
    -   Birth date (DD.MM.YYYY)
    -   Additional numbers (comma-separated)
    -   Relatives (comma-separated)
3.  On form submission, the generator:
    -   Builds all text pair combinations
    -   Appends all numeric tails
    -   Applies letter-case variations
    -   Writes all unique combinations to a TXT file
4.  Browser downloads the file automatically.

------------------------------------------------------------------------

## Installation

``` bash
git clone <your-repo-url>
cd evilcombo
go mod init evilcombo
go run evilwordlist.go
```

------------------------------------------------------------------------

## Usage

1.  Run:

``` bash
go run main.go
```

2.  Open:

```{=html}
<!-- -->
```
    http://localhost:8080

3.  Fill in all desired fields\
4.  Click **Generate TXT file**\
5.  Receive a `.txt` wordlist

------------------------------------------------------------------------

## Example Output

    kyliejohn1010
    KylieJohn1010
    KYLIEJOHN1010
    Kyliejohn1010
    johnkylie2001
    JohnKylie2001
    ...

------------------------------------------------------------------------

## Notes

-   All non-digit characters in custom numbers are removed
-   If birth date is invalid, only custom numbers are used
-   If no numbers exist at all, the generator returns an error

------------------------------------------------------------------------

## License

MIT

