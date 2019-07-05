package context

/**
  context包工作机制：
   第一个创建Context的goroutine被称为root节点，负责创建一个实现Context接口的具体对象，并将该对象作为参数传递到其新拉起的goroutine，
下游的goroutine可以继续封装该对象，再传递到更下游的goroutine。
   Context对象在传递的过程中最终形成一个树状的数据结构，这样通过root节点的Context对象就能遍历整个Context对象树，
通知和消息也可以通过root节点传递出去，实现消息传递，
 */
