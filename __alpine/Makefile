all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./nexusutils.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.22.0 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/nxrmutils.exe $(DESTDIR)$(PREFIX)/bin/nexusutils
	strip $(DESTDIR)$(PREFIX)/bin/nexusutils
