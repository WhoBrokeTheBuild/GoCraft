
SRC = $(wildcard *.go)
OUT = GoCraft

all: $(OUT)

$(OUT): $(SRC)
	go build .

run: all
	./$(OUT)
