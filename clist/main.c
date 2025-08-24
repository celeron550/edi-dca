#include <stdio.h>
#include <stdlib.h>
#include "arraylist.h"

int main(int argc, char const *argv[])
{
    printf("=== Teste de ArrayList ===\n");

    // --- Teste com struct local ---
    ArrayList al;
    ArrayList_Init(&al, 3);
    printf("Estrutura local criada com capacidade %zu\n", al.capacity);

    ArrayList_Add(&al, 10);
    ArrayList_Add(&al, 20);
    ArrayList_Add(&al, 30);

    printf("Depois de Add: tamanho = %zu\n", ArrayList_Size(&al));
    for (size_t i = 0; i < ArrayList_Size(&al); i++) {
        printf("Elemento %zu = %d\n", i, ArrayList_Get(&al, i));
    }

    ArrayList_Set(&al, 25, 1);
    printf("Depois de Set(25,1): elemento 1 = %d\n", ArrayList_Get(&al, 1));

    ArrayList_RemoveOnIndex(&al, 0);
    printf("Depois de RemoveOnIndex(0): tamanho = %zu\n", ArrayList_Size(&al));
    for (size_t i = 0; i < ArrayList_Size(&al); i++) {
        printf("Elemento %zu = %d\n", i, ArrayList_Get(&al, i));
    }

    ArrayList_Pop(&al);
    printf("Depois de Pop: tamanho = %zu\n", ArrayList_Size(&al));
    for (size_t i = 0; i < ArrayList_Size(&al); i++) {
        printf("Elemento %zu = %d\n", i, ArrayList_Get(&al, i));
    }

    ArrayList_Destroy(&al);


    printf("\nTodos os testes concluidos!\n");
    return 0;

    return 0;
}
