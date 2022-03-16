from classes.LRUCache import LRUCache


def main():
    cache = LRUCache(2)
    cache.put(1, "asd")
    cache.display()
    cache.put(2, 2)
    cache.display()
    cache.get(1)
    cache.display()
    cache.put(3, 3)
    cache.display()
    cache.get(2)
    cache.display()
    cache.put(4, 4)
    cache.display()
    cache.get(1)
    cache.display()
    cache.get(3)
    cache.display()
    cache.get(4)
    cache.display()

    return None


main()
