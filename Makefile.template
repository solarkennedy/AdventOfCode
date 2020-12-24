run: main
	./main

main: input *.go
	go build -o main .

input:
	xdg-open https://adventofcode.com/2020/day/$(shell basename `pwd`)

clean:
	rm main

test:
	go test -v .

benchmark:
	go test -bench=.
