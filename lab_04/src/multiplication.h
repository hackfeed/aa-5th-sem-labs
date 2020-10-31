#ifndef __PRINT_MATRIX__
#define __PRINT_MATRIX__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "thread_structs.h"

#define N 1024
#define M N
#define K N

void print_matrix(int **matrix);

void read_matrix(int **matrix, FILE *file);

void init_matrix(int **matrix);

int **create_matrix(int n, int m);

args_t *create_args(int n, int m, int k, int read_file);

void base_multiplication(args_t *args);

void *parallel_multiplication1(void *args);

void *parallel_multiplication2(void *args);

#endif
