#include <stdio.h>

#include "threads.h"

#define ANSI_COLOR_RED "\x1b[31m"
#define ANSI_COLOR_GREEN "\x1b[32m"
#define ANSI_COLOR_YELLOW "\x1b[33m"
#define ANSI_COLOR_BLUE "\x1b[34m"
#define ANSI_COLOR_MAGENTA "\x1b[35m"
#define ANSI_COLOR_CYAN "\x1b[36m"
#define ANSI_COLOR_RESET "\x1b[0m"

#define INVALID_ARGS 2
/*
 * 1 - file flag
 * 2 - type 
 * 3 - threads cnt
 */

int main(int argc, char *argv[])
{
    setbuf(stdout, NULL);

    fprintf(stdout, ANSI_COLOR_MAGENTA "Параллельное умножение матриц\n" ANSI_COLOR_RESET);

    if (argc < 3)
    {
        return INVALID_ARGS;
    }

    args_t *args = create_args(N, M, K, atoi(argv[1]));

    if (!args)
    {
        return ALLOCATE_ERROR;
    }

    int type = atoi(argv[2]);

    if (3 == type)
    {

        uint64_t start = tick();
        base_multiplication(args);
        uint64_t end = tick();

        fprintf(stdout, ANSI_COLOR_YELLOW "Время исполнения: %lu (тиков процессора)\n" ANSI_COLOR_RESET, end - start);
    }
    else
    {
        uint64_t start = tick();

        if (start_threading(args, atoi(argv[3]), type))
        {
            return ALLOCATE_ERROR;
        }

        uint64_t end = tick();

        fprintf(stdout, ANSI_COLOR_YELLOW "Время исполнения: %lu (тиков процессора)\nКоличество потоков: %s\n" ANSI_COLOR_RESET, end - start, argv[3]);
    }

    /*#ifdef __DEBUG__*/
    if (atoi(argv[1]))
    {
        print_matrix(args->res);
    }
    /*#endif*/

    return OK;
}
