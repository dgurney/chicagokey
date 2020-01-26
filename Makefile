GITVERSION=`git describe --tags --dirty`
LDFLAGS=-ldflags "-w -X main.gitVersion=${GITVERSION}"
PROGRAM=github.com/dgurney/chicagokey
PROGRAMSHORT=chicagokey

install:
	go install ${LDFLAGS} ${PROGRAM}
windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/windows/amd64/${PROGRAMSHORT}.exe ${PROGRAM}
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o build/windows/386/${PROGRAMSHORT}.exe ${PROGRAM}
darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/darwin/amd64/${PROGRAMSHORT} ${PROGRAM}
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o build/darwin/386/${PROGRAMSHORT} ${PROGRAM}
freebsd:
	GOOS=freebsd GOARCH=amd64 go build ${LDFLAGS} -o build/freebsd/amd64/${PROGRAMSHORT} ${PROGRAM}
	GOOS=freebsd GOARCH=386 go build ${LDFLAGS} -o build/freebsd/386/${PROGRAMSHORT} ${PROGRAM}
openbsd:
	GOOS=openbsd GOARCH=amd64 go build ${LDFLAGS} -o build/openbsd/amd64/${PROGRAMSHORT} ${PROGRAM}
	GOOS=openbsd GOARCH=386 go build ${LDFLAGS} -o build/openbsd/386/${PROGRAMSHORT} ${PROGRAM}
linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux/amd64/${PROGRAMSHORT} ${PROGRAM}
	GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o build/linux/arm64/${PROGRAMSHORT} ${PROGRAM}
	GOOS=linux GOARCH=arm GOARM=7 go build ${LDFLAGS} -o build/linux/armv7/${PROGRAMSHORT} ${PROGRAM}
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o build/linux/386/${PROGRAMSHORT} ${PROGRAM}
clean:
	rm -rf build/
cross: windows darwin freebsd linux openbsd
all: install cross
