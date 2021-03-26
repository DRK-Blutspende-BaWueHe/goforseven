# goforseven

> HL7 Tool for [Go](https://golang.org/). 

## Features

* Marshal and unmarshal HL7 Messages

## Installation

To install the library and command line program, use the following:

```bash
go get -v github.com/DRK-Blutspende-BaWueHe/goforseven/...
```

## Usage
```
import "github.com/DRK-Blutspende-BaWueHe/goforseven"
```

To read a message 
```go
  hl7 := goforseven.Init25()
  
  str, _err := ioutil.Read('test.hl7')
  if err != nil {
    log.Fatal(err)
  }  
  
  orm, err := hl7.UnmarshalORM_R01(str)
  if err != nil {
    log.Fatal(err)
  }
```
