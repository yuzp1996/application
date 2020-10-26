通过  https://segmentfault.com/a/1190000022896706  来了解gomock解决了之前什么样难以实现的问题

task作为一个interface接口被传入到taskpool中， 而且被调用起来

被测试的是 taskpool 中的run方法  这个方式是被测试的 但是其中会依赖task中的方法来处理事情

mock中的那个结构体MinimalTask 完全是手写的  用来mock使用的 
这里就造成了便利性大大降低  需要自己写个 最小实现(minimal implement) 的结构体来实现某个接口 达到mock的目的