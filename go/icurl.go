package main

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -pedantic
// #cgo LDFLAGS: -lcurl
// #include <curl/curl.h>
// #include <stdlib.h>
// #include <sys/utsname.h>
//
// const char *file_mode = "w";
//
// CURLcode curl_fetch(CURL *easy, const char *url, FILE *fp) {
//     curl_easy_setopt(easy, CURLOPT_URL, url);
//     curl_easy_setopt(easy, CURLOPT_FOLLOWLOCATION, 1L);
//     curl_easy_setopt(easy, CURLOPT_WRITEDATA, fp);
//     return curl_easy_perform(easy);
// }
import "C"

func display_uname() {
	var buf C.struct_utsname

	status, err := C.uname(&buf)
	if status < 0 {
		// Go can't access errno itself.
		// Use a helper function to get error as a C string.
		fmt.Println("failed to get system data", err)
		return
	}

	nodename := C.GoString(&buf.nodename[0])
	sysname := C.GoString(&buf.sysname[0])
	release := C.GoString(&buf.release[0])

	fmt.Printf("Host %s is running %s %s\n",
		nodename, sysname, release)
}

func fetch_file() {
	url := "https://curl.haxx.se/favicon.ico"
	c_url := C.CString(url)

	filename := "curled"
	c_filename := C.CString(filename)

	fmt.Println("Opening file", filename, "...")

	// open file for libcurl to write to it
	fp, err := C.fopen(c_filename, C.file_mode)
	if nil == fp {
		fmt.Println("failed to open file:", filename, ":", err)
		return
	}

	fmt.Println("Fetching url", url, "...")

	// init and configure curl runtime
	C.curl_global_init(C.CURL_GLOBAL_ALL)
	curl := C.curl_easy_init()
	C.curl_fetch(curl, c_url, fp)
	C.curl_easy_cleanup(curl)
	C.curl_global_cleanup()

	C.free(unsafe.Pointer(c_filename))
	C.free(unsafe.Pointer(c_url))
}

func main() {
	display_uname()
	fetch_file()
}
