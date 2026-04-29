import string
import sys
import threading
import time


def counter(c=2):
    while True:
        time.sleep(1)
        print("{1}-{0}\n".format(c, threading.current_thread()))
        c += 1

def char():
    s = string.ascii_lowercase
    for c in s:
        time.sleep(1)   # 模拟阻塞1秒（置为阻塞态），线程卡了，指令暂停到这里
        print(c)

if __name__ == '__main__':
    threading.Thread(target=counter, name="counter").start()
    threading.Thread(target=char, name="char").start()
    while True:
        time.sleep(3)
        a = ["<{}:{}>".format(t.name, t.ident) for t in threading.enumerate()]
        print(a, file=sys.stderr)
        print(*a)

    # threading.Thread(target=counter, name="ccc").start()
    # counter(200)

    # counter(200)
    # threading.Thread(target=counter, name="ccc").start()
    # 若按照该顺序，由于第一条语句会死循环，第二条语句则永远无不可到达
