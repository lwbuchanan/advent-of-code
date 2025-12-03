#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#include "main.h"

int main(int argc, char *argv[]) {
  if (argc < 2) {
    fprintf(stderr, "usage: ./advent <day>\n");
    return EXIT_FAILURE;
  }
  int day = atoi(argv[1]);
  if (day == 0) {
    fprintf(stderr, "Invalid day number\n");
    return EXIT_FAILURE;
  }

  char input_filename[32];
  sprintf(input_filename, "src/day%02d/day%02d.txt", day, day);

  FILE *input_fd = fopen(input_filename, "r");
  if (input_fd == NULL) {
    fprintf(stderr, "Error opening input file\n");
    return EXIT_FAILURE;
  }

  char lines[MAX_LINES][LINE_LENGTH];
  int num_lines = 0;
  while (fgets(lines[num_lines], LINE_LENGTH, input_fd)) {
    ++num_lines;
  }

  switch (day) {
  case 1:
    day01a(lines, num_lines);
    day01b(lines, num_lines);
    break;
  default:
    fprintf(stderr, "No implementation for this day\n");
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
