import threading

print("Welcome to Python")

def worker():
    for x in range(5):
        print("我是干活的，第 1 句")
        raise Exception("我定义异常")
    print("线程结束了")

# worker()    # 普通函数调用
# 线程是操作系统才能创建的，调用系统调用syscall 帮助我们构造出一个thread

print("=" * 50)
# Python中的顶层代码定义的都是包的全局变量
t1 = threading.Thread(target=worker, name="w1")    # 类（类名称为大驼峰），封装一个可以被Python管理的线程对象
# 类()   类的初始化，创建一个该类的实例
# __init__() 方法method，Thread类型
# 特别注意，上面这行只是创建了一个可以使用的对象
# 其他线程崩，主线程不受影响

t1.start()  # 底层调用syscall创建os的线程，w1线程运行指令（worker编译后的指令），相当于w1线程调用worker()
