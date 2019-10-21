BIN = icurl

all: $(BIN)

run: $(BIN)
	./$(BIN)

clean:
	rm -rf $(BIN) *.ico *.o

$(BIN):
	rustc main.rs -o $(BIN) -L ../lib \
	  -C link_args='-Wl,-rpath=../lib'
