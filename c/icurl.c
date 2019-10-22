/* icurl.c */

#include <curl/curl.h>          /* libcurl suite */
#include <stdlib.h>             /* EXIT_FAILURE, EXIT_SUCCESS */
#include <sys/utsname.h>        /* uname */

int display_uname();
int fetch_favicon();

int main() {
    int status = 0;
    status += display_uname();
    status += fetch_favicon();
    return status;
}


/*
 * perform system call to get kernel info
 */

int display_uname() {
    // allocate buffer for kernel data
    struct utsname buf;

    // uname(2) returns < 0 on failure
    if (uname(&buf) < 0) {
        perror("failed to get uname data");
        return EXIT_FAILURE;
    }

    // dump hostname, kernel name, and version
    printf("Host %s is running %s %s\n",
        buf.nodename, buf.sysname, buf.release);

    return EXIT_SUCCESS;
}

/*
 * fetch favicon.ico from libcurl's homepage
 */

int fetch_favicon() {
    const char *url = "https://curl.haxx.se/favicon.ico";
    const char *filename = "curled";

    printf("Opening file '%s' ...\n", filename);

    // open file for libcurl to write to it; returns NULL on failure
    FILE *fp = fopen(filename, "w");
    if (NULL == fp) {
        perror("Failed to open file");
        return EXIT_FAILURE;
    }

    printf("Fetching url %s ...\n", url);

    // create and configure easy curl handle for the fetch
    curl_global_init(CURL_GLOBAL_ALL);                  // init curl
    CURL* curl = curl_easy_init();                      // get easy handle
    curl_easy_setopt(curl, CURLOPT_URL, url);           // set url
    curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L); // follow redirects
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);      // write to open file
    CURLcode res = curl_easy_perform(curl);             // Engage
    curl_easy_cleanup(curl);                            // cleanup easy handle
    curl_global_cleanup();                              // global curl cleanup

    // check for failure
    if (res != CURLE_OK) {
        printf("request for <%s> failed: %s\n", url, curl_easy_strerror(res));
    }

    // read bytes written from file cursor
    fprintf(stderr, "Wrote %ld bytes.\n", ftell(fp));
    fclose(fp);

    return EXIT_SUCCESS;
}

/* vim: set ft=c ts=8 sts=4 sw=4 expandtab: */
