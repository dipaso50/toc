all: compileLinux compileWindows

compileLinux:
	go build -o bin/linux/toc main.go

compileWindows:
	GOOS=windows GOARCH=386 go build -o bin/windows/toc.exe main.go

install:
	cp bin/linux/toc ${GOPATH}/bin/

clean:
	rm -f bin/linux/toc ; rm -f bin/windows/toc.exe
	