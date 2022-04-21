from classes.KVPair import KVPair
from classes.DoublyLinkedList import DoublyLinkedList


class LRUCacheCustom():
    def __init__(self, max_size=10):
        if max_size <= 0:
            raise Exception('Max size must be larger than zero')
        self.max_size = max_size
        self.list = DoublyLinkedList()
        self.nodes = {}

    def set(self, key, value):
        node = self.nodes.get(key, None)
        if node != None:
            node.data.value = value
            self.list.move_front(node)
            return

        if self.list.len == self.max_size:
            expired = self.list.remove_tail()
            del self.nodes[expired.data.key]

        self.nodes[key] = self.list.create_node_and_move_front(
            KVPair(key, value))

    def get(self, key):
        node = self.nodes.get(key, None)
        if node is None:
            return None

        self.list.move_front(node)
        return node.data.value
