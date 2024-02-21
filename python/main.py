from classes.LRUCacheCustom import LRUCacheCustom
from classes.LRUCache import LRUCache


def main():
    # cache = LRUCache(2)
    # cache.put(1, "asd")
    # cache.display()
    # cache.put(2, 2)
    # cache.display()
    # cache.get(1)
    # cache.display()
    # cache.put(3, 3)
    # cache.display()
    # cache.get(2)
    # cache.display()
    # cache.put(4, 4)
    # cache.display()
    # cache.get(1)
    # cache.display()
    # cache.get(3)
    # cache.display()
    # cache.get(4)
    # cache.display()
    cache = LRUCacheCustom(3)
    print(cache.get(5))
    cache.set(1, "asd")
    cache.set(2, "bcd")
    cache.set(3, "efg")
    cache.set(4, "5")
    print(cache.get(1))
    print(cache.get(2))
    cache.list.print_cap()
    cache.print_cache()
    return None


main()
