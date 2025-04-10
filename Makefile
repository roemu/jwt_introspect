jwt-introspect:
	go build -o ./build/jwt-introspect

install: jwt-introspect
	install -m 755 ./build/jwt-introspect /usr/local/bin
