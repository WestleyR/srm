// created by: WestleyR
// email: westleyr@nym.hush.com
// https://github.com/WestleyR/srm
// date: Dec 18, 2019
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

#define SCRIPT_VERSION "v1.1.0, Dec 18, 2019"

#ifndef COMMIT_HASH
#define COMMIT_HASH "unknown"
#endif

int force_flag = 0;
int recursive_flag = 0;

void help_menu(const char* script_name) {
  printf("Copyright (c) 2019 WestleyR, All rights reserved.\n");
  printf("This software is licensed under a Clear BSD License.\n");
  printf("\n");
  printf("Usage:\n");
  printf("  %s [option] <file...>\n", script_name);
  printf("\n");
  printf("Options:\n");
  printf("  -c, --cache        print the cache (comming soon)\n");
  printf("  -f, --force        remove with force\n");
  printf("  -r, --recursive    remove a directory\n");
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

int mkdir_all(const char *dirs) {
  int dirs_len = strlen(dirs);

#ifdef DEBUG
  printf("DEBUG: max dir len: %d\n", dirs_len);
#endif

  char dir_path[dirs_len+1];
  char *p = NULL;
  struct stat st;

  strcpy(dir_path, dirs);

  if (dir_path[dirs_len - 1] == '/') {
    dir_path[dirs_len - 1] = '\0';
  }

  for (p = dir_path + 1; *p != '\0'; p++) {
    if (*p == '/') {
      *p = '\0';
#ifdef DEBUG
      printf("Creating: %s\n", dir_path);
#endif
      if (stat(dir_path, &st) == -1) {
        if (mkdir(dir_path, 0700) != 0) {
          fprintf(stderr, "Failed to create: %s\n", dir_path);
          return(1);
        }
      }
      *p = '/';
    }
  }

#ifdef DEBUG
  printf("Creating: %s\n", dir_path);
#endif
  if (stat(dir_path, &st) == -1) {
    if (mkdir(dir_path, 0700) != 0) {
      fprintf(stderr, "Failed to create: %s\n", dir_path);
      return(1);
    }
  }

  return(0);
}

int init_cache() {
  char* cache_path = get_srm_cachedir();
  if (cache_path == NULL) {
    return(1);
  }

  struct stat st;

  if (stat(cache_path, &st) == -1) {
    if (mkdir_all(cache_path) != 0) {
      fprintf(stderr, "Failed to create: %s\n", cache_path);
      return(1);
    }
  }

  free(cache_path);

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
  sprintf(current_date, "Y%d/M%d/D%d/H%d-M%d-S%d", tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec);

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

  struct stat st;

  for (int i = 1; i < 20; i++) {
    if (stat(srm_file, &st) == 0) {
      char b[10];
      b[0] = '\0';
      sprintf(b, ".%d", i);
      strcat(srm_file, b);
    } else {
      break;
    }
  }

  return(srm_file);
}

int srm(const char* file) {
  struct stat st;

  if (stat(file, &st) != 0) {
    fprintf(stderr, "%s: No such file or directory\n", file);
    return(1);
  }

  if (S_ISREG(st.st_mode) == 0) {
    if (recursive_flag == 0) {
      fprintf(stderr, "%s: is a directory\n", file);
      return(1);
    }
  }

  if (access(file, R_OK) != 0) {
    if (force_flag == 0) {
      fprintf(stderr, "%s: is write-protected file\n", file);
      return(1);
    }
  }

  char* trash_dir = gen_cachedir(file);
  if (trash_dir == NULL) {
    fprintf(stderr, "Failed to make trash dir\n");
    return(1);
  }

#ifdef DEBUG
  printf("TRASH DIR+FILE: %s\n", trash_dir);
#endif

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
    {"force", no_argument, 0, 'f'},
    {"recursive", no_argument, 0, 'r'},
    {"cache", no_argument, 0, 'c'},
    {"version", no_argument, 0, 'V'},
    {"commit", no_argument, 0, 'C'},
    {NULL, 0, 0, 0}
  };

  while ((opt = getopt_long(argc, argv,"crfVhC", long_options, 0)) != -1) {
    switch (opt) {
      case 'c':
        list_srm_cache = 1;
        break;
      case 'r':
        recursive_flag = 1;
        break;
      case 'f':
        force_flag = 1;
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

  int success = 0;

  if (optind < argc) {
    for (int i = optind; i < argc; i++) {
#ifdef DEBUG
      printf("Removing: %s\n", argv[i]);
#endif
      if (srm(argv[i]) != 0) {
        success = 1;
      }
    }
  } else {
    printf("Nothing to do...\n");
    return(1);
  }

  return(success);
}

// vim: tabstop=2 shiftwidth=2 expandtab autoindent softtabstop=0
