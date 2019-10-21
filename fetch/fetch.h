#pragma once

#include <stdio.h>

void fetch_init();
void fetch_cleanup();
size_t fetch(const char *url, const char *filename);
