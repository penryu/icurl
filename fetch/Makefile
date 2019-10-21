NAME = fetch

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
  LIB := lib$(NAME).dylib
else
  LIB := lib$(NAME).so
endif

CFLAGS += -g -Werror -Wall -Wextra -Wconversion -pedantic -fno-builtin
LDFLAGS = -lcurl


all: $(LIB)
run: $(LIB)
clean:
	rm -rf *.o *.dylib $(LIB)

$(LIB):
	$(CC) -dynamiclib -o $(LIB) $(NAME).c -lcurl

