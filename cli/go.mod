module gomosaics-cli

go 1.21.0

require github.com/urfave/cli/v2 v2.25.7

require gomosaics v0.0.0

replace gomosaics => ../gomosaics

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
)
