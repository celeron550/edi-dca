#pragma once

typedef struct ArrayList ArrayList;
/// @brief inicializa o arraylist com um espaço
/// @param list &arraylist
/// @param size tamanho para ser alocado
void ArrayList_Init(ArrayList *list, size_t size);

/// @brief retorna tamanho do arraylist
/// @param list &arraylist
/// @return size_t contendo tamanho de list->v
size_t ArrayList_Size(ArrayList *list);

/// @brief retorna o elemento contido no indice do arraylist
/// @param list arraylist
/// @param index indice
/// @return -1 se o index for invalido ou o elemento correspondente no indice
int ArrayList_Get(ArrayList *list, size_t index);

/// @brief adiciona um elemento ao final do array
/// @param list &arraylist
/// @param e elemento para ser adicionado
void ArrayList_Add(ArrayList *list, int e);

/// @brief dobra a capacidade do arraylist
/// @param list &arraylist
void ArrayList_doubleV(ArrayList *list);

/// @brief remove elemento contido no indice
/// @param list &arraylist
/// @param index indice 
void ArrayList_RemoveOnIndex(ArrayList *list, size_t index);

/// @brief remove o ultimo elemento e o retorna
/// @param list &arraylist
/// @return elemento removido
int ArrayList_Pop(ArrayList *list);

/// @brief acessa o indice do arraylist, substituindo-o pelo e
/// @param list &arraylist
/// @param e elemento
/// @param index incice
void ArrayList_Set(ArrayList *list, int e, size_t index);


 
void ArrayList_Destroy(ArrayList *list);


typedef struct ArrayList
{
    int* v;
    size_t inserted;
    size_t capacity;

}ArrayList;