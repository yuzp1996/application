# client-go informer  [study golang with k8s]

## Informer 代码剖析

Informer 用来保证消息的实时性，可靠性和顺序性。 k8s 的很多组件都是通过 client-go 的 Informer 机制跟 k8s 的 apiserver 沟通的

1. Reflector
监控 k8s 指定资源变化， 主要是通过资源本身的 list 和 watch 方法

```cassandraql

可以在 这个函数中 进行大量的处理，也可以避免拆分函数

if err := func() error {return nil}(); err != nil {
    return err
}

通过gorountinue 来通信  通过 select 实现了可以同时判断多个条件

go func() {
    defer func() {
        if r := recover(); r != nil {
            panicCh <- r
        }
    }()
    list, err = r.listerWatcher.List(options)
    close(listCh)
}()
select {
case <-stopCh:
    return nil
case r := <-panicCh:
    panic(r)
case <-listCh:
}
if err != nil {
    return fmt.Errorf("%s: Failed to list %v: %v", r.name, r.expectedType, err)
}
```

loop 这种东西  不知道是不是应该常用




2. DeltaFIFO



## 文件总结
上下间隔明确，上下逻辑清晰

看完的代码或者无论什么东西 都要总结一些 否则日后就跟没看一样了