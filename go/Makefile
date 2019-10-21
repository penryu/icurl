BIN = icurl

all: $(BIN)

run: $(BIN)
	./$(BIN)
clean:
	rm -f $(BIN) *.ico go *.dylib *.o

$(BIN): $(BIN).go
	rm -f libfetch*
	ln -sf ../fetch/libfetch* .
	go build $(BIN).go
