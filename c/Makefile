BIN = icurl
FETCH_DIR = ../fetch
CFLAGS += -Werror -Wall -Wextra -Wconversion -pedantic -fno-builtin -I $(FETCH_DIR)
LDFLAGS += -lfetch -L $(FETCH_DIR)

all: $(BIN)
	cp $(FETCH_DIR)/libfetch* .

run: $(BIN)
	LD_LIBRARY_PATH=$(FETCH_DIR) ./$(BIN)

clean:
	rm -rf $(BIN) $(BIN).dSYM *.ico
