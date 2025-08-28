package main

import "github.com/LX/base-go/blog"

var mainVar string = getMainVar()

func init() {
	//fmt.Println("main init function has been called")
}

func main() {
	//fmt.Println("Hello, world!")
	//fmt.Println("main function has been called")
	// 1.2 基本数据类型
	//exercise.BaseDataType()
	// 1.3 定义变量
	//exercise.DefineVariable()
	// 1.4 指针
	//exercise.Pointer()
	// 1.4.3 修改指针指向的值
	//exercise.ModifyPointer()
	// 1.4.4 unsafe unitptr
	//exercise.UnsafeUnitPtr()
	// 1.5 结构体
	//exercise.DefineAnonymousStruct()
	// 1.5.3 嵌套结构体
	//exercise.NestedStruct()
	// 1.5.4 定义结构体方法
	//exercise.StructMethod()
	// 1.6 常量与枚举
	//exercise.DefineConstant()
	// 1.7 运算操作符
	//exercise.ArithmeticOperator()
	// 1.8 流程控制
	//exercise.ProcessControl()
	//exercise.SwitchCase()
	// 1.9 循环控制
	//exercise.LoopControl()
	// 1.9.4 循环控制语句
	//exercise.LoopBreakControl()
	// 1.9.4.2 continue语句
	//exercise.ContinueControl()
	// 1.9.4.3 goto语句
	//exercise.GotoControl()
	// 1.10 函数
	//exercise.FunctionExercixe()
	// 1.11 变量的作用域
	//exercise.DefineScope()
	// 1.12 数组
	//exercise.ArrayExercise()
	// 1.12.2 访问数组
	//exercise.ArrayAccess()
	// 1.12.3 多维数组
	//exercise.MultiDimensionalArray()
	// 1.12.4 数组作为参数
	//exercise.ArrayAsParam()
	// 1.13 切片
	//exercise.SliceExercise()
	// 1.13.2 使用切片
	//exercise.UseSlice()
	// 1.13.2.3 复制切片
	//exercise.CopySlice()
	// 1.13.3 切片原理
	//exercise.SlicePrinciple()
	// 没有触发切片扩容
	//exercise.SliceNoExpand()
	// 触发切片扩容
	//exercise.SliceExpand()
	// 1.14 Map集合
	//exercise.MapExercise()
	// 1.14.2 使用Map集合
	//exercise.UseMap()
	// 1.14.3 map作为参数
	//exercise.MapAsParam()
	// 1.14.4 并发使用Map集合
	//exercise.ConcurrentUseMap()
	// 1.15.1 range迭代
	//exercise.RangeExercise()
	// 1.15.2 对数组与切片迭代
	//exercise.ArrayAndSliceIteration()
	// 1.15.3 对通道的迭代
	//exercise.ChannelIteration()
	// 1.15.4 对Map集合迭代
	//exercise.MapIteration()
	// 1.16 类型转换
	//exercise.TypeConversion()
	// 字符串和数字转换
	//exercise.StringToNumber()
	// 1.16.3 接口类型转换
	//exercise.InterfaceTypeConversion()
	// 1.16.4 结构体类型转换
	//exercise.StructTypeConversion()
	// 任务一
	//task.TaskOne()
	// 2.1 接口
	//exercise.InterfaceExercise()
	// 2.2 并发
	//exercise.GoroutineChannel()
	// Go并发线程安全
	//exercise.ConcurrentThreadSafe()
	// Channel
	//exercise.ChannelExercise()
	// 并发 select
	//exercise.GoroutineSelect()
	// 编译
	//platformSpecificFunction()
	// 任务二
	//task2.Pointer()
	//task2.Goroutine()
	//task2.ObjectExercise()
	//task2.TaskTwoChannel()
	//task2.LockExercise()
	//gorm.GormExercise()
	// 任务三
	//sqlx.SqlxExercise()
	//gormadvance.GormAdvanceExercise()
	// Gin
	//gin.GinExercise()
	// Gin中间件
	//gin.GinMiddleware()
	// Gin验证
	//gin.GinValidator()
	blog.BlogExercise()
}

func getMainVar() string {
	//fmt.Println("main.getMainVar has been called")
	//fmt.Println("pkg1.PkgNameVar: " + pkg1.PkgNameVar)
	return "main"
}
