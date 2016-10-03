# ICS
An iCalendar parser writen in Go

An example can be found in main/main.go

## Installing

    go get github.com/zskamljic/ics
    
## Usage

Calendar object can be obtained by calling `ics.NewCalendar(data string)`. It accepts a string containing the ics that you want to convert. The class is ready to be converted to JSON.
