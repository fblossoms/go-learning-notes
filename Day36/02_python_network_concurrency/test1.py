import time

def counter(c=1):   # c=1表示为c的默认值为1，如果不给就是1
    f = open("logs/test.log", "a")
    while True:
        time.sleep(1)
        print(c, 100, file=f, sep="##")
        c += 1

if __name__ == '__main__':  # 简单认为就是main函数
    counter()
