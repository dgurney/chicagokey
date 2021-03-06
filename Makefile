GITVERSION=`git describe --tags --dirty`
LDFLAGS=-ldflags "-w -X main.gitVersion=${GITVERSION}"
PROGRAMROOT=github.com/dgurney/chicagokey
COMMAND=${PROGRAMROOT}/cmd/chicagokey
PROGRAMSHORT=chicagokey

install:
	go install ${LDFLAGS} ${COMMAND}
windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/windows/amd64/${PROGRAMSHORT}.exe ${COMMAND}
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o build/windows/386/${PROGRAMSHORT}.exe ${COMMAND}
	GOOS=windows GOARM=7 GOARCH=arm go build ${LDFLAGS} -o build/windows/arm/${PROGRAMSHORT}.exe ${COMMAND}
darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/darwin/amd64/${PROGRAMSHORT} ${COMMAND}
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o build/darwin/386/${PROGRAMSHORT} ${COMMAND}
freebsd:
	GOOS=freebsd GOARCH=amd64 go build ${LDFLAGS} -o build/freebsd/amd64/${PROGRAMSHORT} ${COMMAND}
	GOOS=freebsd GOARCH=386 go build ${LDFLAGS} -o build/freebsd/386/${PROGRAMSHORT} ${COMMAND}
openbsd:
	GOOS=openbsd GOARCH=amd64 go build ${LDFLAGS} -o build/openbsd/amd64/${PROGRAMSHORT} ${COMMAND}
	GOOS=openbsd GOARCH=386 go build ${LDFLAGS} -o build/openbsd/386/${PROGRAMSHORT} ${COMMAND}
linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux/amd64/${PROGRAMSHORT} ${COMMAND}
	GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o build/linux/arm64/${PROGRAMSHORT} ${COMMAND}
	GOOS=linux GOARCH=arm GOARM=7 go build ${LDFLAGS} -o build/linux/armv7/${PROGRAMSHORT} ${COMMAND}
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o build/linux/386/${PROGRAMSHORT} ${COMMAND}
test:
	rm -rf coverage
	mkdir coverage
	go test ${PROGRAMROOT}/pkg/generator -coverprofile=coverage/generator
cross: windows darwin freebsd linux openbsd
all: install cross
