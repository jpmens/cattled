
PROG = cattled

all: files/$(PROG)-freebsd files/$(PROG)-linux files/$(PROG)-darwin


files/$(PROG)-freebsd: src/$(PROG).go Makefile
	env GOOS=freebsd GOOARCH=amd64 go build -o $@ src/$(PROG).go

files/$(PROG)-linux: src/$(PROG).go Makefile
	env GOOS=linux GOOARCH=amd64 go build -o $@ src/$(PROG).go

files/$(PROG)-darwin: src/$(PROG).go Makefile
	go build -o $@ src/$(PROG).go
