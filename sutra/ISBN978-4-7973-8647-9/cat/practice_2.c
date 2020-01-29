#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
  int line = 0;

  for (int i = 1; i < argc; i++) {
    FILE *f;
    int c;

    f = fopen(argv[i], "r");
    if (!f) {
      perror(argv[i]);
      exit(1);
    }
    while ((c = fgetc(f)) != EOF) {
//      if (putchar(c) < 0) {
//        exit(1);
//      }
      if (c == '\n') {
        line += 1;
      }
    }
    fclose(f);
  }
  printf("%d lines\n", line);
  exit(0);
}
