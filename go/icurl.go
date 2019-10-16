package main

// #cgo CFLAGS: -g -O0 -Wconversion -pedantic
// #cgo LDFLAGS: -lcurl
// #include <stdlib.h>
// #include <string.h>
// #include <curl/curl.h>
//
// char *go2cstr(_GoString_ s) {
//     size_t len = _GoStringLen(s);
//     char *p = malloc(len + 1U);
//     memcpy(p, _GoStringPtr(s), len);
//     p[len] = '\0';
//     return p;
// }
//
// size_t fetch(_GoString_ goURL, _GoString_ goFilename) {
//     char *filename = go2cstr(goFilename);
//     FILE *fp = fopen(filename, "w");
//     free(filename);
//     if (NULL == fp) {
//         perror("Failed to open file");
//         return 0;
//     }
//
//     char *url = go2cstr(goURL);
//     CURL* curl = curl_easy_init();
//     curl_easy_setopt(curl, CURLOPT_URL, url);
//     curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
//     curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);
//     CURLcode res = curl_easy_perform(curl);
//     if (res != CURLE_OK) {
//         fprintf(stderr, "request for <%s> failed: %s\n",
//                 url, curl_easy_strerror(res));
//     }
//     curl_easy_cleanup(curl);
//     free(url);
//
//     long bytes_written = ftell(fp);
//     fprintf(stderr, "Wrote %ld bytes.\n", bytes_written);
//     fclose(fp);
//
//     if (bytes_written < 0) {
//         perror("Failed to write file");
//         return 0;
//     }
//
//     return (size_t)bytes_written;
// }
import "C"

func main() {
	urls := map[string]string{
		"https://penryu.app/":            "index.html",
		"https://penryu.app/favicon.ico": "favicon.ico",
	}

	C.curl_global_init(C.CURL_GLOBAL_ALL)

	for url, filename := range urls {
		C.fetch(url, filename)
	}

	C.curl_global_cleanup()
}
