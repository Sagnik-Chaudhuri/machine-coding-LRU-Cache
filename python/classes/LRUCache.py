from collections import OrderedDict


class LRUCache():

    def __init__(self, capacity) -> None:
        self.__capacity = capacity
        self.__cache = OrderedDict()

    def get(self, key):
        if key not in self.__cache:
            return -1
        else:
            self.__cache.move_to_end(key)
            return self.__cache[key]

    def put(self, key, value):
        self.__cache[key] = value
        self.__cache.move_to_end(key)
        if len(self.__cache) > self.__capacity:
            self.__cache.popitem(last=False)

    def display(self):
        print(self.__cache)
