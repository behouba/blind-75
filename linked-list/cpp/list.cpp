#include <iostream>

class Node
{
public:
    int value;
    Node *next;
    Node(int v)
    {
        value = v;
        next = NULL;
    };
    Node *InsertAtEnd(int v);
    Node *InsertAtBegining(int v);
    Node *DeleteNode(int index);
    int Length();
    void PrintList();
};

Node *Node::InsertAtBegining(int newData)
{
    Node *head = this;
    if (head == NULL)
    {
        head = new Node(newData);
        return head;
    }
    Node *newHead = new Node(newData);
    newHead->next = head;
    head = newHead;
    return head;
}

Node *Node::InsertAtEnd(int newData)
{
    Node *head = this;
    if (head == NULL)
    {
        head = new Node(newData);
        return head;
    }
    Node *lastNode = head;
    while (lastNode->next != NULL)
    {
        lastNode = lastNode->next;
    }
    lastNode->next = new Node(newData);
    return head;
}

void Node::PrintList()
{
    Node *current = this;
    while (current != NULL)
    {
        std::cout << current->value << "->";
        current = current->next;
    }
    std::cout << "NULL" << std::endl;
    return;
}

Node *Node::DeleteNode(int x)
{
    Node *head = this;
    if (head == NULL)
    {
        return head;
    }
    Node *node = head;
    Node *prevNode = NULL;
    int index = 1;
    while (node != NULL)
    {
        if (x == index)
        {
            if (prevNode == NULL)
            {
                head = node->next;
                delete node;
                break;
            }
            else
            {
                prevNode->next = node->next;
                delete node;
                break;
            }
        }
        prevNode = node;
        node = node->next;
        index = index + 1;
    }
    return head;
}

int Node::Length()
{
    Node *head = this;
    int count = 0;
    Node *node = head;
    while (node != NULL)
    {
        count = count + 1;
        node = node->next;
    }
    return count;
}

int main()
{
    Node *head = new Node(5);
    head->InsertAtEnd(10);
    head->InsertAtEnd(11);
    head->InsertAtEnd(12);
    head->InsertAtEnd(13);
    head->InsertAtEnd(14);
    head->InsertAtEnd(15);
    head->InsertAtEnd(16);
    head = head->InsertAtBegining(27);
    head = head->InsertAtBegining(100);
    head->PrintList();
    head->DeleteNode(2);
    head->PrintList();
    std::cout << "list length = " << head->Length() << std::endl;
    return 0;
}