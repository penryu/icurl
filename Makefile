LIBS = fetch
BINS = c go rs

all:
	for lib in $(LIBS); do make -C $$lib install; done
	for bin in $(BINS); do make -C $$bin run; done
run: all

test:
	for 
uninstall: clean
	make -C fetch uninstall

clean:
	for dir in $(LIBS) $(BINS); do make -C $$dir clean; done

