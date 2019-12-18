// created by: WestleyR
// email: westleyr@nym.hush.com
// https://github.com/WestleyR/srm
// date: Dec 17, 2019
// version-1.0.0
//
// The Clear BSD License
//
// Copyright (c) 2019 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
//

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <time.h>
#include <getopt.h>

#define SCRIPT_VERSION "v1.0.0-beta-2, Dec 17, 2019"

#ifndef COMMIT_HASH
#define COMMIT_HASH "unknown"
#endif

void help_menu(const char* script_name) {
  printf("Usage:\n");
  printf("  %s [option] <file...>\n", script_name);
  printf("\n");
  printf("Options:\n");
  printf("  -h, --help         print help menu\n");
  printf("  -C, --commit       print the github commit\n");
  printf("  -V, --version      print version\n");
  printf("\n");
  printf("Source code: https://github.com/WestleyR/srm\n");
  return;
}

void version_print() {
  printf("%s\n", SCRIPT_VERSION);
  return;
}

void version_commit() {
  printf("%s\n", COMMIT_HASH);
  return;
}

char* get_srm_cachedir() {
  char* home = getenv("HOME");

  char* cache_path;
  cache_path = (char*) malloc(200 * sizeof(char));
  if (cache_path == NULL) {
    fprintf(stderr, "malloc failed\n");
    return(NULL);
  }

  strcpy(cache_path, home);
  strcat(cache_path, "/.cache/srm");

  return(cache_path);
}

int init_cache() {
  char* cache_path = get_srm_cachedir();
  if (cache_path == NULL) {
    return(1);
  }

  struct stat st;

  if (stat(cache_path, &st) == -1) {
    if (mkdir(cache_path, 0700) != 0) {
      fprintf(stderr, "Failed to create: %s\n", cache_path);
      return(1);
    }
  }

  free(cache_path);

  return(0);
}

int mkdir_all(const char *dir) {
  char tmp[256];
  char *p = NULL;
  size_t len;
  struct stat st;

  strcpy(tmp, dir);
  len = strlen(tmp);

  if (tmp[len - 1] == '/') {
    tmp[len - 1] = '\0';
  }
  for (p = tmp + 1; *p; p++) {
    if (*p == '/') {
      *p = 0;
      if (stat(tmp, &st) == -1) {
        if (mkdir(tmp, 0700) != 0) {
          fprintf(stderr, "Failed to create: %s\n", tmp);
          return(1);
        }
      }
      *p = '/';
    }
  }

  if (stat(tmp, &st) == -1) {
    if (mkdir(tmp, 0700) != 0) {
      fprintf(stderr, "Failed to create: %s\n", tmp);
      return(1);
    }
  }

  return(0);
}

char* gen_cachedir(const char* to_trash) {
  time_t t = time(NULL);
  struct tm tm = *localtime(&t);

  char* current_date;
  current_date = (char*) malloc(50 * sizeof(char));
  if (current_date == NULL) {
    fprintf(stderr, "malloc failed\n");
    return(NULL);
  }

  current_date[0] = '\0';
  sprintf(current_date, "Y%d/M%d/D%d_H%d-M%d-S%d", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);

  char* srm_file;
  srm_file = (char*) malloc(256 * sizeof(char));
  if (srm_file == NULL) {
    fprintf(stderr, "malloc failed\n");
    return(NULL);
  }


  char* file_cache = get_srm_cachedir();
  if (file_cache == NULL) {
    return(NULL);
  }

  strcpy(srm_file, file_cache);
  strcat(srm_file, "/");
  strcat(srm_file, current_date);
  free(current_date);
  free(file_cache);

  if (mkdir_all(srm_file) != 0) {
    return(NULL);
  }

  strcat(srm_file, "/");
  strcat(srm_file, to_trash);

  return(srm_file);
}

int srm(const char* file) {
  char* trash_dir = gen_cachedir(file);
  if (trash_dir == NULL) {
    fprintf(stderr, "Failed to make trash dir\n");
    return(1);
  }

  printf("TRASH DIR+FILE: %s\n", trash_dir);

  if (rename(file, trash_dir) != 0) {
    fprintf(stderr, "Failed to move file: %s\n", file);
    perror(file);
    return(1);
  }

  free(trash_dir);

  return(0);
}

int main(int argc, char** argv) {
  if (init_cache() != 0) {
    fprintf(stderr, "Failed to init cache\n");
    return(1);
  }

  int list_srm_cache = 0;
  int opt = 0;

  static struct option long_options[] = {
    {"help", no_argument, 0, 'h'},
    {"cache", no_argument, 0, 'c'},
    {"version", no_argument, 0, 'V'},
    {"commit", no_argument, 0, 'C'},
    {NULL, 0, 0, 0}
  };

  while ((opt = getopt_long(argc, argv,"cVhC", long_options, 0)) != -1) {
    switch (opt) {
      case 'c':
        list_srm_cache = 1;
        break;
      case 'h':
        help_menu(argv[0]);
        return(0);
        break;
      case 'V':
        version_print();
        return(0);
        break;
      case 'C':
        version_commit();
        return(0);
        break;
      default:
        // Invalid option
        return(22);
    }
  }

  if (list_srm_cache) {
    printf("Listing cache...\n");
    return(0);
  }

  if (optind < argc) {
    for (int i = optind; i < argc; i++) {
      printf("Removing: %s\n", argv[i]);
      srm(argv[i]);
    }
  } else {
    printf("Nothing to do...\n");
    return(1);
  }

  return(0);
}

// vim: tabstop=2 shiftwidth=2 expandtab autoindent softtabstop=0
