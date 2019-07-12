[Build Status]: https://travis-ci.org/YoussefKaib/nister
[Build Status Badge]:https://travis-ci.com/YoussefKaib/nister.svg?branch=master

# Nister

[![Build Status][Build Status Badge]][Build Status] [![Go Report Card](https://goreportcard.com/badge/github.com/YoussefKaib/nister)](https://goreportcard.com/report/github.com/YoussefKaib/nister)

Nister is lightweight Go package returns most recent and modified [CVE](https://cve.mitre.org/) Per **Product or Programing Language** from [National Vulnerability Database](https://nvd.nist.gov/vuln/data-feeds)

## Installation

```md
go get github.com/YoussefKaib/nister
```

### CLI Example

![](https://media.giphy.com/media/hvGbporNaP1xozXbxn/giphy.gif)

![nister_cli](https://github.com/YoussefKaib/nister/blob/master/images/nister_cli_example.png)

```go
package main

import (
	"fmt"
	"os"

	"github.com/youssefkaib/nister"
)

func main() {
	products := os.Args

	// Parsed CVE Data
	data := nister.ParseCVEReport()

	report := nister.ProductChecker(data, products)
	
	fmt.Println(report)
}

```

### Contributing

Pull requests, bug fixes, and new features are welcome!

1. Fork the repository
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -a -m 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request on GitHub
