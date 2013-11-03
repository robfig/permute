/* #include <stdlib.h> */
/* #include <stdio.h> */

/* void swap_characters(char *first, char *second) */
/* { */
/*     char tmp = *first; */
/*     *first = *second; */
/*     *second = tmp; */
/* } */

/* void permute(char array[], int N, int start_index) */
/* { */
/*     if (start_index == N - 1) */
/*     { */
/*         printf("%s\n", array); */
/*         return; */
/*     } */

/*     for (int i = start_index; i < N; i++) */
/*     { */
/*         if (start_index != i) */
/*             swap_characters((array+start_index), (array+i)); */
/*         permute(array, N, start_index+1); */
/*         swap_characters(array+start_index, array+i); */
/*     } */
/* } */

/* /\* */
/* #include <stdlib.h> */
/* void permute(char array[], int N, int start_index); */
/* import "C" */

/* func main() { */
/* 	var ( */
/* 		input string  = os.Args[1] */
/* 		cstr  *C.char = C.CString(input) */
/* 		N     C.int   = C.int(len(input)) */
/* 	) */
/* 	C.permute(cstr, N, C.int(0)) */
/* 	C.free(unsafe.Pointer(cstr)) */
/* } */
/* *\/ */
