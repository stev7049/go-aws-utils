GOLIST = $(shell go list ./... | grep -v /vendor/)

ifeq ($(OS),Windows_NT)
	INSTALLPRE = "c:\Faws\bin"
else
	INSTALLPRE = /usr/local/bin
endif

build: bin/awsresources bin/instancelist

bin/awsresources:
	@echo "Building awsresources..."
	go build -o bin/awsresources awsresources/main.go

bin/instancelist:
	@echo "Building instancelist..."
	go build -o bin/instancelist instancelist/main.go

test:
		@test -z "$(gofmt -s -l . | tee /dev/stderr)"
		@test -z "$(golint $(GOLIST) | tee /dev/stderr)"
		@go test -v -race $(GOLIST)
		@go vet $(GOLIST)

clean:
		@echo "Cleaning up..."
ifeq ($(OS),Windows_NT)
	powershell.exe -Command "if(Test-path .\bin ){ rm .\bin -Recurse -Force}"
else
	rm -rf bin
endif

rebuild: clean build

install:
ifeq ($(OS),Windows_NT)
		@echo "Installing..."
		powershell.exe -Command "if(-not(Test-path 'C:\Program Files (x86)\Faws')){ New-Item 'C:\Faws\bin' -ItemType Directory -Force}"
		setx path "%path%;C:\Faws\bin"
		cp bin/awsresources $(INSTALLPRE)
		cp bin/instancelist $(INSTALLPRE)
else
		@echo "Installing..."
		cp bin/awsresources $(INSTALLPRE)
		cp bin/instancelist $(INSTALLPRE)
endif

.PHONY: build test clean
