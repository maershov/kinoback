#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <math.h>
#include <stdbool.h>
#include <ctype.h>
#include <stdint.h>

void reverse(char*, int );
void swap(char*, char*);
char* itoa(size_t , char*, int);
size_t toten(char*, int);
size_t todegree(int, int);

int main() {
    int num = 100;
    int position = 0;
    size_t valueten = 0;
    int P,Q,r;
    P = Q = r = 0;
    char* buf;
    char* str = (char*) malloc(sizeof(char)*(num + 1));

    r = scanf("%d %d %100s", &P, &Q, str);

    if (r != 3)
    {
        free(str);
        printf("%s", "[error]");    // как проверить на другое количество введенных данных? 5 4 1 2 читает только первые 3 и все окей
        return 0;
    }

    if (!((2 <= Q) && (Q <= P) && (P <= 36)))
    {
        printf("%s", "[error]");
        free(str);
        return 0;
    }

    while ((strlen(str) == (num)))
    {
        num += 100;
        buf = (char *) realloc(str, sizeof(char) * (num + 1));
        if (!buf) {

            free(str);
            printf("%s", "[error]");
            return 0;
        }
        str = buf;
        char *tbuf = (char *) malloc(11);
        position += 100;
        scanf("%100s", tbuf);
        if (strlen(tbuf)) {
            memcpy((str + position), tbuf, 11);

        }
        free(tbuf);
    }

    valueten = toten(str, P);;
    if (valueten == -10)
    {
        printf("%s", "[error]");
        free(str);
        return 0;
    }

    char* str1 = (char*) malloc(sizeof(char) * (num + 1));
    itoa(valueten, str1, Q);
    if (!(strcmp(str1, "-1"))){
        printf("%s", "[error]");
        free(str1);
        free(str);
        return 0;
    }
    printf("%s", str1);

    free(str);
    free(str1);
    return 0;
}

size_t todegree(int value, int degree)
{
    size_t newvalue = 0;

    if (degree == 0)
        return 1;
    else
        newvalue = todegree(value, --degree) * value;

    return newvalue;
}

void swap(char *a, char *b)
{
    char temp;
    temp = *a;
    *a = *b;
    *b = temp;
}

void reverse(char str[], int length)
{
    int start = 0;
    int end = length -1;

    while (start < end)
    {
        swap(&(str[start]), &(str[end]));
        start++;
        end--;
    }
}

char* itoa(size_t num, char* str, int base)
{
    int i = 0;
    int capacity = 0;

    capacity = strlen(str) - 1;

    if (num == 0)
    {
        str[i++] = '0';
        str[i] = '\0';
        return str;
    }

    while (num != 0)
    {
        int rem = num % base;
        if (i == capacity){
            char *tstr;
            capacity = strlen(str) * 2;
            tstr = (char *) realloc(str, sizeof(char) * (capacity + 1));
            if (!tstr) {

                free(str);
                printf("%s", "[error]");
                return "-1";
            }
            str = tstr;
        }
        str[i++] = (rem > 9)? (rem-10) + 'a' : rem + '0';
        num = num/base;
    }

    str[i] = '\0';
    reverse(str, i);

    return str;
}

size_t toten(char* str, int base)
{
    size_t newvalue = 0;
    size_t buf = 0;
    int position = 0;
    size_t temp = 0;
    int length = strlen(str);
    if (!length)
        return 0;
    while (length != 0)
    {

        buf = (str[length - 1] > '9') ? ((char)(tolower(str[length - 1])) - 'a' + 10) : ((char)tolower((str[length - 1])) - '0');

        if  (buf > (base - 1))
        {
            return -10;
        }
        temp = buf * todegree(base, position);

        if (!((SIZE_MAX - temp) > newvalue)) {
            return -10;
        }
        newvalue += temp;
        position++;
        length--;
    }
    return newvalue;
}
