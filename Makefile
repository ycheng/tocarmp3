all: tocarmp3

tocarmp3: tocarmp3.go
	go build tocarmp3.go

clean:
	rm tocarmp3
