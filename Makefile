UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
  LIB := lib$(NAME).dylib
else
  LIB := lib$(NAME).o
endif

SUBS = fetch c go rs

all:
	for dir in $(SUBS); do make -C $$dir all; done
run:
	for dir in $(SUBS); do make -C $$dir run; done
clean:
	for dir in $(SUBS); do make -C $$dir clean; done
