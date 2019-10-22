BIN = icurl

all: $(BIN)

run: $(BIN)
	./$(BIN)

clean:
	rm -rf $(BIN) *.o curled

$(BIN):
	rustc main.rs -o $(BIN)
