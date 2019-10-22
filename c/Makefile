BIN = icurl
CFLAGS += -Werror -Wall -Wextra -Wconversion -pedantic -fno-builtin
LDFLAGS += -lcurl

all: $(BIN)

run: $(BIN)
	./$(BIN)

clean:
	rm -rf $(BIN) $(BIN).dSYM curled
