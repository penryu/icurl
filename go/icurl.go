package main

// #cgo CFLAGS: -pedantic -I ../include
// #cgo LDFLAGS: -Wl,-rpath=../lib -L ../lib -lfetch
// #include <stdlib.h>
// #include <string.h>
// #include "fetch.h"
//
// char *go2cstr(_GoString_ s) {
//     size_t len = _GoStringLen(s);
//     char *p = malloc(len + 1U);
//     memcpy(p, _GoStringPtr(s), len);
//     p[len] = '\0';
//     return p;
// }
//
// size_t fetch_wrapper(_GoString_ goURL, _GoString_ goFilename) {
//     char *url = go2cstr(goURL);
//     char *filename = go2cstr(goFilename);
//     size_t bytesRead = fetch(url, filename);
//     free(url);
//     free(filename);
//     return bytesRead;
// }
import "C"

func main() {
	urls := map[string]string{
		"https://apple.com/favicon.ico":  "apple.ico",
		"https://google.com/favicon.ico": "google.ico",
	}

	C.fetch_init()

	for url, filename := range urls {
		C.fetch_wrapper(url, filename)
	}

	C.fetch_cleanup()
}
