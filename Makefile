huecat:
	go build

install: huecat
	cp huecat /usr/local/bin/

clean:
	go clean
