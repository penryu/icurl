/* -*- mode: c -*- */

/* icurl.c */


#include <curl/curl.h>
#include <string.h>                 /* strncat */


size_t fetch(const char *url, const char *filename) {
    FILE *fp = fopen(filename, "w");
    if (NULL == fp) {
        perror("Failed to open file");
        return 0;
    }

    const curl_version_info_data *curl_data = curl_version_info(CURLVERSION_NOW);
    char user_agent[128] = "curl/";
    snprintf(user_agent, sizeof(user_agent), "curl/%s", curl_data->version);

    CURL* curl = curl_easy_init();
    curl_easy_setopt(curl, CURLOPT_URL, url);
    curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);
    curl_easy_setopt(curl, CURLOPT_USERAGENT, user_agent);
    CURLcode res = curl_easy_perform(curl);
    if (res != CURLE_OK) {
        fprintf(stderr, "request for <%s> failed: %s\n",
                url, curl_easy_strerror(res));
    }
    curl_easy_cleanup(curl);

    long bytes_written = ftell(fp);
    fprintf(stderr, "Wrote %ld bytes.\n", bytes_written);
    fclose(fp);

    if (bytes_written < 0) {
        perror("Failed to write file");
        return 0;
    }

    return (size_t)bytes_written;
}

int main() {
    curl_global_init(CURL_GLOBAL_ALL);

    fetch("https://penryu.app/", "index");
    fetch("https://penryu.app/favicon.ico", "favicon.ico");

    curl_global_cleanup();

    return 0;
}

/* vim: set ft=c ts=8 sts=4 sw=4 expandtab: */