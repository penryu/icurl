BIN = icurl

all: $(BIN)

run: $(BIN)
	./$(BIN)
clean:
	rm -f $(BIN) curled go

$(BIN): $(BIN).go
	go build $(BIN).go
