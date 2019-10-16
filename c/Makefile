# Macros
BIN = icurl
PRODUCTS = $(BIN) $(BIN).dSYM
CFLAGS += -g -O0 -Werror -Wall -Wextra -Wconversion -pedantic -fno-builtin
LIBS = -lcurl

# Targets
all: $(BIN)

install: $(BIN)
	install -s $(BIN) /usr/local/bin/$(BIN)

test: $(BIN)
	./$(BIN)

clean:
	rm -rf $(PRODUCTS)

$(BIN): $(BIN).c
	$(CC) $(CFLAGS) -o $(BIN) $(BIN).c $(LIBS)
