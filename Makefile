
OUT = GoCraft

all: $(OUT)

$(OUT):
	go build .

run: all
	./$(OUT)
