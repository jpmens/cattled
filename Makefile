
PROG = cattled

all: files/$(PROG)-freebsd files/$(PROG)-linux files/$(PROG)-darwin


files/$(PROG)-freebsd: $(PROG).go Makefile
	env GOOS=freebsd GOOARCH=amd64 go build -o $@ $(PROG).go

files/$(PROG)-linux: $(PROG).go Makefile
	env GOOS=linux GOOARCH=amd64 go build -o $@ $(PROG).go

files/$(PROG)-darwin: $(PROG).go Makefile
	go build -o $@ $(PROG).go
