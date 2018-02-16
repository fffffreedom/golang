# context - 管理goroutine

## context简介

golang中的创建一个新的goroutine，并不会返回像c语言类似的pid，所以我们不能从外部杀死某个goroutine，所以我们就得让它自己结束。之前我们用channel+select的方式，来解决这个问题，但是有些场景实现起来比较麻烦。  
例如，由一个请求衍生出的各个goroutine之间需要满足一定的约束关系，以实现一些诸如有效期，中止routine树，传递请求全局变量之类的功能。  

google就为我们提供一个解决方案，开源了context包；使用context实现上下文功能约定，需要将context.Context类型的变量作为函数的第一个参数（见package的介绍）。  

**GO1.7之后，新增了context.Context这个package，实现goroutine的管理。**  

context包是用来管理goroutine的，它给goroutine提供一个运行环境；context可以被取消、会超时、有deadline、还可以在各goroutine之间传递值，这些不同功能的context是通过不同函数创建的！这些版本的函数是：  
```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```
这四个函数接受相应的参数，再返回context，以及后面用来取消Context的CancelFunc。  
- WithCancel
WithCancel对应的是cancelCtx，其中，返回一个cancelCtx，同时返回一个CancelFunc，CancelFunc是context包中定义的一个函数类型：type CancelFunc func()。调用这个CancelFunc时，关闭对应的c.done，也就是让他的后代goroutine退出。  
- WithDeadline 和 WithTimeout
WithDeadline和WithTimeout对应的是timerCtx，WithDeadline和WithTimeout是相似的，WithDeadline是设置具体的deadline时间，到达deadline的时候，后代goroutine退出，而WithTimeout简单粗暴，直接return WithDeadline(parent, time.Now().Add(timeout))。  
- WithValue
WithValue对应valueCtx，WithValue是在Context中设置一个map，拿到这个Context以及它的后代的goroutine都可以拿到map里的值。  

从函数的声明可以看出，前面三个都是可以取消的，而WithValue则没有返回CancelFunc！  

## 实际案例

可以参见开源的网络插件flannel的实现，在main函数中，会创建context，并在后面的函数中使用：  
```
	// This is the main context that everything should run in.
	// All spawned goroutines should exit when cancel is called on this context.
	// 在对此context调用cancel时，所有衍生出的goroutines都应该退出。
	//
	// Go routines spawned from main.go coordinate using a WaitGroup. 
	// This provides a mechanism to allow the shutdownHandler goroutine to block
	// until all the goroutines return. 
	//
	// If those goroutines spawn other goroutines then they are responsible for
	// blocking and returning only when cancel() is called.
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		shutdownHandler(ctx, sigs, cancel)
		wg.Done()
	}()
	......
```

## 源码剖析

### package介绍
```
// Package context defines the Context type, which carries deadlines,
// cancelation signals, and other request-scoped values across API boundaries
// and between processes.
//
// Incoming requests to a server should create a Context, and outgoing calls to
// servers should accept a Context.  The chain of function calls between must
// propagate the Context, optionally replacing it with a modified copy created
// using WithDeadline, WithTimeout, WithCancel, or WithValue.
//
// Programs that use Contexts should follow these rules to keep interfaces
// consistent across packages and enable static analysis tools to check context
// propagation:
//
// Do not store Contexts inside a struct type; instead, pass a Context
// explicitly to each function that needs it.  The Context should be the first
// parameter, typically named ctx:
//
// 	func DoSomething(ctx context.Context, arg Arg) error {
// 		// ... use ctx ...
// 	}
//
// Do not pass a nil Context, even if a function permits it.  Pass context.TODO
// if you are unsure about which Context to use.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The same Context may be passed to functions running in different goroutines;
// Contexts are safe for simultaneous use by multiple goroutines.
//
// See http://blog.golang.org/context for example code for a server that uses
// Contexts.
```

### Context interface
```
// Context's methods may be called by multiple goroutines simultaneously.
// context的方法是线程安全的，可以被多个goroutine使用
type Context interface {
	// Deadline returns the time when work done on behalf of this context
	// should be canceled.  Deadline returns ok==false when no deadline is
	// set.  Successive calls to Deadline return the same results.
	// 返回Context将要关闭的时间，如果设置了Deadline
	Deadline() (deadline time.Time, ok bool)

	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled.  Done may return nil if this context can
	// never be canceled.  Successive calls to Done return the same value.
	// 当Context被canceled或是timeout的时候，Done返回一个被closed的channel 
	// 
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	//
	// Done is provided for use in select statements:
	//
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
	// See http://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancelation.
	Done() <-chan struct{}

	// Err returns a non-nil error value after Done is closed.  Err returns
	// Canceled if the context was canceled or DeadlineExceeded if the
	// context's deadline passed.  No other values for Err are defined.
	// After Done is closed, successive calls to Err return the same value.
	// 在Done的channel被closed后， Err代表被关闭的原因，Canceled或者DeadlineExceeded
	Err() error

	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key.  Successive calls to Value with
	// the same key returns the same result.
	//
	// Use context values only for request-scoped data that transits
	// processes and API boundaries, not for passing optional parameters to
	// functions.
	//
	// A key identifies a specific value in a Context.  Functions that wish
	// to store values in Context typically allocate a key in a global
	// variable then use that key as the argument to context.WithValue and
	// Context.Value.  A key can be any type that supports equality;
	// packages should define keys as an unexported type to avoid
	// collisions.
	//
	// Packages that define a Context key should provide type-safe accessors
	// for the values stores using that key:
	//
	// 	// Package user defines a User type that's stored in Contexts.
	// 	package user
	//
	// 	import "golang.org/x/net/context"
	//
	// 	// User is the type of value stored in the Contexts.
	// 	type User struct {...}
	//
	// 	// key is an unexported type for keys defined in this package.
	// 	// This prevents collisions with keys defined in other packages.
	// 	type key int
	//
	// 	// userKey is the key for user.User values in Contexts.  It is
	// 	// unexported; clients use user.NewContext and user.FromContext
	// 	// instead of using this key directly.
	// 	var userKey key = 0
	//
	// 	// NewContext returns a new Context that carries value u.
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 		return context.WithValue(ctx, userKey, u)
	// 	}
	//
	// 	// FromContext returns the User value stored in ctx, if any.
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	// 如果存在，Value返回与key相关了值，不存在返回nil
	Value(key interface{}) interface{}
}
```
我们不需要手动实现这个接口，context包已经给我们提供了两个，一个是Background()，一个是TODO()，这两个函数都会返回一个Context的实例，只是返回的这两个实例都是空Context。  

通过Background返回的空context，它不能被取消，不能传递值，没有deadline，主要被main函数使用、初始化、测试等，是incoming requests的顶级context：  
```
// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	return background
}
```

### context是如何被取消的？

context.go文件中有两个结构体继承了context接口：  
```
// A cancelCtx can be canceled. When canceled, it also cancels any children
// that implement canceler.
type cancelCtx struct {
	Context

	done chan struct{} // closed by the first cancel call.

	mu       sync.Mutex
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}
// A valueCtx carries a key-value pair. It implements Value for that key and
// delegates all other calls to the embedded Context.
type valueCtx struct {
	Context
	key, val interface{}
}
```
cancelCtx继承了context，并且实现了canceler接口（实现了cancel和Done方法），这样context就可以被直接取消了，实现context的取消功能。  
```
// A canceler is a context type that can be canceled directly. The
// implementations are *cancelCtx and *timerCtx.
type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

func (c *cancelCtx) Done() <-chan struct{} {
	return c.done
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func (c *cancelCtx) String() string {
	return fmt.Sprintf("%v.WithCancel", c.Context)
}

// cancel closes c.done, cancels each of c's children, and, if
// removeFromParent is true, removes c from its parent's children.
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
	close(c.done)  <==== 关闭cancel context
	for child := range c.children {
		// NOTE: acquiring the child's lock while holding parent's lock.
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

## Reference
golang中context包解读  
https://studygolang.com/articles/9517  
