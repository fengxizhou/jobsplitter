logsplitter: main.go
	go pull
	go build .
	cp -f logsplitter ~/bin/
clean:
	rm -rf logsplitter
