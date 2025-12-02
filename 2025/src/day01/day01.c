#include "../main.h"

#include <stdio.h>
#include <stdlib.h>

#define MOD(a, b) ((((a) % (b)) + (b)) % (b))

void day01a(char lines[][LINE_LENGTH], int num_lines) {
  int password = 0;
  int position = 50;

  for (int i = 0; i < num_lines; ++i) {
    char *line = lines[i];

    int direction = line[0];
    int distance = atoi(line + 1);

    switch (direction) {
    case 'R':
      position = (position + distance) % 100;
      break;
    case 'L':
      position = (position - distance) % 100;
      break;
    default:
      fprintf(stderr, "Unknown direction");
      return;
    }

    if (position == 0) {
      ++password;
    }
  }

  printf("Password: %d\n", password);
}

void day01b(char lines[][LINE_LENGTH], int num_lines) {
  int password = 0;
  int position = 50;

  for (int i = 0; i < num_lines; ++i) {
    char *line = lines[i];

    int direction = line[0];
    int distance = atoi(line + 1);
    int increment = 0;
    int diff = 0;

    switch (direction) {
    case 'R':
      diff = position + distance;
      increment = diff / 100;
      position = diff % 100;
      break;
    case 'L':
      diff = position - distance;
      increment = abs(diff) / 100;
      if (position != 0 && diff <= 0) {
        increment += 1;
      }
      position = MOD(diff, 100);
      break;
    default:
      fprintf(stderr, "Unknown direction");
      return;
    }
    password += increment;
  }

  printf("Password: %d\n", password);
}
