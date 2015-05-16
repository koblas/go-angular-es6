default: build

#clean:
#	rm -f cmd/server/server
#	rm -f cmd/todo/todo

build:
	cd app; \
	go build
	cd cmd; \
	go build

deps:
	cd cmd; \
	go get

#migrate:
#	./cmd/server/server --config config.yaml migratedb

#test:
#	./cmd/server/server --config config.yaml server & \
#	pid=$$!; \
#	go test; \
#	kill $$pid
