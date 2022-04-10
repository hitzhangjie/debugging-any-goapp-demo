all:
	go build -gcflags 'all=-N -l' -o main 
	./main
