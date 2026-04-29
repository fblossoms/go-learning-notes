import string
import threading
import time


# 协程
def count():
    c = 1
    for i in range(5):
        print(c, "#")
        yield c # Go 1.22 实验性已经出现了yield，能够暂停函数执行，出去一个数据
        print("#####")
        c += 1
    return c

def char():
    s = string.ascii_lowercase
    for c in s:
        print(c, "c")
        yield c

t1 = count()
t2 = char()

# 交替运行
tasks = {"t1": t1, "t2": t2}

def add_task():
    time.sleep(10)
    tasks["t300"] = count()

def event_loop():
    while True:
        if len(tasks) == 0:
            print("我大循环睡一会")
            time.sleep(1)
            continue

        pops = []
        for name, task in tasks.items(): # items() -> 一个个entry kv对
            if next(task, None) is None:
                # tasks.pop(name) # 移除当前kv对，但是不能这样写，因为不能一边遍历，一边减少kv对数
                pops.append(name)
                print("task {} finished".format(name))
        for name in pops:
            tasks.pop(name)

threading.Thread(target=event_loop, name="eventloop").start()

time.sleep(10)
tasks["t400"] = char()

# x = count() # 要么函数调用异常，要么正常return
# print(type(x), x)
# print(next(x)) # 挤，才会有代码的执行
# print(next(x, None)) # 除非给默认值（区分），否则挤爆了就报错
# print('~' * 50)
# for t in x: # 遇到stopiteration错误时，自动停止。但是有些情况会有无限值，挤不完
#     print(t, "%%%")