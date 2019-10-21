BIN = icurl
LIB_PATH = ../lib
INCLUDE_PATH = ../include
CFLAGS += -Werror -Wall -Wextra -Wconversion -pedantic -fno-builtin
CFLAGS += -I $(INCLUDE_PATH)
LDFLAGS += -Wl,-rpath=$(LIB_PATH) -L $(LIB_PATH) -lfetch

all: $(BIN)

run: $(BIN)
	./$(BIN)

clean:
	rm -rf $(BIN) $(BIN).dSYM *.ico
