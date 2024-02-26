# phishtank

Go wrapper for the phishtank API & offline database

## Setup

- To install this package in your project directory run the command:

```
go get github.com/nemzyxt/phishtank
```

## Usage

You can either query an __offline database__ or the __online API__

### Offline database

Example:
```
package main

import (
	"fmt"

	"github.com/nemzyxt/phishtank/db"
)

func main() {
	db.LoadDatabase("verified_online.csv")
	record := db.CheckURL("https://assiua.site.tb-hosting.com/")
	if record != nil {
		fmt.Println(record)
	} else {
		fmt.Println("No such record")
	}
}
```
Sample result:
```
[8418748 https://assiua.site.tb-hosting.com/ http://www.phishtank.com/phish_detail.php?phish_id=8418748 2024-01-12T08:25:01+00:00 yes 2024-01-12T08:32:39+00:00 yes Other]
```

### Online API

In progress :smiley:
