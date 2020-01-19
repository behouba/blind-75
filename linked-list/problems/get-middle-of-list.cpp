// Source https://practice.geeksforgeeks.org/problems/finding-middle-element-in-a-linked-list/1

#include <stdio.h>
#include <stdlib.h>
#include <iostream>
using namespace std;

/* Link list Node */
struct Node
{
    int data;
    struct Node *next;

    Node(int x)
    {
        data = x;
        next = NULL;
    }
};

void append(struct Node **head_ref, struct Node **tail_ref, int new_data)
{
    struct Node *new_node = new Node(new_data);

    if (*head_ref == NULL)
        *head_ref = new_node;
    else
        (*tail_ref)->next = new_node;
    *tail_ref = new_node;
}

/* Function to get the middle of the linked list*/
int getMiddle(struct Node *head);

/* Driver program to test above function*/
int main()
{
    int T, i, n, l;

    cin >> T;

    while (T--)
    {
        struct Node *head = NULL, *tail = NULL;

        cin >> n;
        for (i = 1; i <= n; i++)
        {
            cin >> l;
            append(&head, &tail, l);
        }

        printf("%d\n", getMiddle(head));
    }
    return 0;
}

int getLength(Node *head)
{
    Node *node = head;
    int count = 0;
    while (node != NULL)
    {
        count = count + 1;
        node = node->next;
    }
    return count;
}

int getDataAt(Node *head, int target)
{
    Node *node = head;
    int index = 1;
    while (node != NULL)
    {
        if (index == target)
        {
            return node->data;
        }
        index = index + 1;
        node = node->next;
    }
}

/* Should return data of middle node. If linked list is empty, then  -1*/
int getMiddle(Node *head)
{
    int length = getLength(head);
    if (length == 0)
    {
        return -1;
    }
    int targetIndex = (length / 2) + 1;
    return getDataAt(head, targetIndex);
}