/* icurl.c */

#include "fetch.h"

int main() {
    fetch_init();
    fetch("https://apple.com/favicon.ico",  "apple.ico");
    fetch("https://google.com/favicon.ico", "google.ico");
    fetch_cleanup();
}

/* vim: set ft=c ts=8 sts=4 sw=4 expandtab: */
