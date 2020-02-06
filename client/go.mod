module github.com/shreve/go-roku/client

go 1.13

require (
	github.com/shreve/go-roku/roku v0.0.0-20200201214839-b0e798196ebe
	github.com/shreve/tui v0.0.0-20191227044428-35e3ce9c440b
)

replace github.com/shreve/go-roku/roku => ../roku

replace github.com/shreve/tui => ../../tui
