from classes.Node import Node


class DoublyLinkedList(object):
    def __init__(self):
        self.root = Node(None)
        self.len = 0

        self.root.next = self.root
        self.root.prev = self.root

    def create_node_and_move_front(self, data):
        node = Node(data)

        self.move_front(node)
        self.inc_len()
        return node

    def move_front(self, node):
        if node is None:
            return None
        elif node.prev is not None and node.next is not None:
            self.isolate(node)

        node.prev = self.root
        node.next = self.root.next
        self.root.next.prev = node
        self.root.next = node

        return node

    def remove_tail(self):
        return self.remove(self.root.prev)

    def remove(self, node):
        if self.len == 0:
            return None
        self.dec_len()
        return self.isolate(node)

    def print_cap(self):
        print(self.len)

    # @staticmethod
    def isolate(self, node):
        node.next.prev = node.prev
        node.prev.next = node.next
        node.next = None
        node.prev = None
        return node

    def inc_len(self):
        self.len += 1

    def dec_len(self):
        self.len -= 1
