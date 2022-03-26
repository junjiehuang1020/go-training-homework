package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	err := foo()
	return errors.Wrap(err, "bar failed")
}

/* 如果开发者是在开发应用和具体服务的过程中，如果遇到了类似于sql.ErrNoRows这种基本类型的error，
我们可以讲基本错误类型包装以后向上抛出，便于开发者去追踪错误和相应的堆栈。应该注意的是，如果可以在error出现的时候可以直接处理掉，例如重试、降级等
措施，我们就不再应该向上抛出error，避免错误信息的冗余以及错误的二次处理。

如果开发者是在开发软件库或者一些公用组件的时候，比较好的做法应该是定义一些基础的错误类型，发生error的时候，直接向上抛出，不需要去包装这些基本错误类型。
*/
func main() {

	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
		return
	}

}
