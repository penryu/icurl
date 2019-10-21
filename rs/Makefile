BIN = icurl

all: $(BIN)
run: $(BIN)
	./$(BIN)
clean:
	rm -rf $(BIN) *.ico *.o libfetch*

$(BIN):
	rm -f libfetch*
	ln -sf ../fetch/libfetch* .
	rustc main.rs -o $(BIN) -L ../fetch
