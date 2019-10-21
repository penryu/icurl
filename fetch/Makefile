NAME = fetch

CFLAGS += -g -Werror -Wall -Wextra -Wconversion -pedantic -fno-builtin

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
  LIB := lib$(NAME).dylib
  CFLAGS += -dynamiclib
else
  LIB := lib$(NAME).so
  CFLAGS += -shared -fpic
endif

LIB_VERSIONED = $(LIB).1
LIBS = $(LIB) $(LIB_VERSIONED)
LDFLAGS = -Wl,-soname,$(LIB_VERSIONED) -o $(LIB_VERSIONED) -lcurl

INCLUDE_PATH = ../include
LIB_PATH = ../lib

all: $(LIB_VERSIONED)

install: $(NAME).h $(LIB_VERSIONED)
	ln -s $(LIB_VERSIONED) $(LIB)
	mkdir -p $(INCLUDE_PATH) $(LIB_PATH)
	install -t $(INCLUDE_PATH) $(NAME).h
	install -s -t $(LIB_PATH) $(LIB_VERSIONED)
	ln -s $(LIB_VERSIONED) $(LIB_PATH)/$(LIB)

uninstall:
	rm -f ../include/$(NAME).h
	cd ../lib/ && rm -f $(LIBS)

clean:
	rm -rf *.o *.dylib $(LIBS)

$(LIB_VERSIONED): $(NAME).c
	$(CC) $(CFLAGS) $(LDFLAGS) $<

