###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := build

lint:
	~/go/bin/golint app.go

build: clean
	CC=clang go build -o macosapp app.go

clean:
	-rm -f macosapp

run: build
	./macosapp

test:
	-rm -f macostest
	CC=clang go build -o macostest test.go
	./macostest
