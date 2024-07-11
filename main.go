package main

import (
	. "DesignPattern/behavioral-pattern/command-pattern"
	. "DesignPattern/behavioral-pattern/iterator-pattern"
	. "DesignPattern/behavioral-pattern/mediator-pattern"
	. "DesignPattern/behavioral-pattern/memento-pattern"
	. "DesignPattern/behavioral-pattern/null-object-pattern"
	. "DesignPattern/behavioral-pattern/visitor-pattern"
	. "DesignPattern/chain-of-responsibility-pattern"
	. "DesignPattern/creational-pattern/objectpool-pattern"
	. "DesignPattern/creational-pattern/prototype-pattern"
	. "DesignPattern/observer-pattern"
	. "DesignPattern/state-pattern"
	. "DesignPattern/structural-pattern/adapter-pattern"
<<<<<<< Updated upstream
	cp "DesignPattern/structural-pattern/composite-pattern"
	"reflect"

	dp "DesignPattern/structural-pattern/decorator-pattern"
	dfp "DesignPattern/structural-pattern/facade-pattern"
	fly "DesignPattern/structural-pattern/flyweight-pattern"

=======
	bp "DesignPattern/structural-pattern/bridge-pattern"
>>>>>>> Stashed changes
	"fmt"
	"log"
	"strconv"
)

func main() {
<<<<<<< Updated upstream
	flyweight()
}

func flyweight() {
	factory := fly.NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("A")
	flyweight1.Operation("First call")

	flyweight2 := factory.GetFlyweight("B")
	flyweight2.Operation("Second call")

	flyweight3 := factory.GetFlyweight("A")
	flyweight3.Operation("Third call")

	fmt.Printf("flyweight1 and flyweight3 are the same instance: %v\n", flyweight1 == flyweight3)
	fmt.Printf("flyweight1 and flyweight3 are the same instance: %v\n", reflect.DeepEqual(flyweight1, flyweight3))
}

func facade() {
	mediaFacade := dfp.NewMediaFacade()
	mediaFacade.Play()
}

func decorator() {

	notifier := &dp.ConcreteNotifier{}
	emailNotifier := dp.NewEmailNotifier(notifier)
	smsNotifier := dp.NewSMSNotifier(emailNotifier)

	smsNotifier.Send("Hello, World!")
}

func composite() {
	file1 := &cp.File{Name: "File1"}
	file2 := &cp.File{Name: "File2"}
	file3 := &cp.File{Name: "File3"}
	folder1 := &cp.Folder{
		Name: "Folder1",
	}
	folder1.Add(file1)

	folder2 := &cp.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)
	folder2.Search("rose")
=======
	bridge()
}

func bridge(){
	red := &bp.Red{}
	blue :=&bp.Blue{}

	redSquare := bp.NewSquare(red)
	fmt.Println(redSquare.Draw())
	blueCircle := bp.NewCircle(blue)
	fmt.Println(blueCircle.Draw())

>>>>>>> Stashed changes
}

func visitor() {
	areaCalculator := NewAreaCalculator("A面积计算器")
	square := NewSquare(5)
	circle := NewCircle(4)

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)

}

func state() {

	trafficLight := NewTrafficLight("世纪大道1号交通灯", &RedLight{})
	trafficLight.SwitchToRed()    // 输出: 已经是红灯
	trafficLight.SwitchToYellow() // 输出: 红灯变为黄灯
	trafficLight.SwitchToGreen()  // 输出: 黄灯变为绿灯

	// 更改状态
	trafficLight.SetState(&GreenLight{})
	trafficLight.SwitchToRed() // 输出: 绿灯变为红灯

}

func nullObject() {
	qingHua := NewCollege("清华大学")
	scienceDepartment := NewScienceDepartment("核科学系", 2)
	musicDepartment := NewMusicDepartment("民族音乐系", 2)
	qingHua.AddDepartment(scienceDepartment)
	qingHua.AddDepartment(musicDepartment)

	csDepartment := qingHua.FindDepartment("计算机")
	fmt.Println(csDepartment.GetName())

<<<<<<< Updated upstream
=======
	adapter()
>>>>>>> Stashed changes
}

// 客户端使用现有接口target使用适配器，而不需要知道适配器内部细节
func adapter() {
	// 创建 specificRequest 对象
	specificRequest := SpecificRequest{}

	// 创建适配器，包含 specificRequest 的引用
	adapter := NewAdapter(specificRequest)
	// 使用适配器，通过 Request 接口调用 specificRequest 的方法
	fmt.Println(adapter.Request())

}

func memento() {
	//笔记管理员
	caretaker := &Caretaker{
		MementoList: make([]*Memento, 0),
	}
	//用户A  A只有一个状态  一个人有多份笔记
	originatorA := &Originator{
		OriginatorState: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originatorA.GetState())
	firstMementoOfOriginatorA := originatorA.CreateMemento()
	caretaker.AddMemento(firstMementoOfOriginatorA)

	originatorA.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originatorA.GetState())

	secondMementoOfOriginatorA := originatorA.CreateMemento()
	caretaker.AddMemento(secondMementoOfOriginatorA)
	originatorA.SetState("C")

	fmt.Printf("Originator Current State: %s\n", originatorA.GetState())

	thirdMementoOfOriginatorA := originatorA.CreateMemento()
	caretaker.AddMemento(thirdMementoOfOriginatorA)

	originatorA.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originatorA.GetState())
	//重置状态只能到初始状态

	//originatorA.RestoreMemento(caretaker.GetMemento(0))
	//fmt.Printf("Restored to State: %s\n", originatorA.GetState())
}

func mediator() {
	stationManager := NewStationManger()
	passengerTrain := &PassengerTrain{
		Mediator: stationManager,
	}
	goodsTrain := &GoodsTrain{
		Mediator: stationManager,
	}
	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}

func iterator() {
	// 创建具体聚合对象并添加元素
	aggregate := &ConcreteAggregate{
		Items: []interface{}{"Item1", "Item2", "Item3"},
	}

	// 创建迭代器
	iterator := aggregate.CreateIterator()

	// 使用迭代器遍历聚合对象中的元素
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}

func command() {
	// 创建接收者
	receiver := &Receiver{}

	// 创建具体命令并设置接收者
	command1 := &ConcreteCommand1{Receiver: receiver}
	command2 := &ConcreteCommand2{Receiver: receiver}

	// 创建调用者并设置命令
	invoker := &Invoker{}

	invoker.SetCommand(command1)
	invoker.ExecuteCommand() // 输出: Receiver: Executing Action1

	invoker.SetCommand(command2)
	invoker.ExecuteCommand() // 输出: Receiver: Executing Action2
}

func chainOfResponsibility() {

	medical := &Medical{}

	//Set next for cashier department
	//收银
	cashier := &Cashier{}
	cashier.SetNext(medical)
	//Set next for doctor department
	doctor := &Doctor{}
	doctor.SetNext(cashier)
	//Set next for reception department
	//挂号
	//链条为主，病人是处理对象
	registrationOffice := &RegistrationOffice{}
	registrationOffice.SetNext(doctor)

	patient := &Patient{Name: "abc"}
	//Patient visiting
	//挂号执行病人
	registrationOffice.Execute(patient)
}

func objectPoolPattern() {
	connections := make([]IConnection, 0)
	for i := 0; i < 3; i++ {
		c := &Connection{Id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := InitPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	_ = pool.Recycle(conn1)
	_ = pool.Recycle(conn2)
}
func prototypePattern() {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		Name:     "Folder1",
		Children: []IInode{file1},
	}
	folder2 := &Folder{
		Children: []IInode{folder1, file2, file3},
		Name:     "Folder2",
	}
	//层级
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}

func observerPatter() {
	xiaGuRiBao := NewSubject("峡谷日报")

	mingMing := NewObserver(xiaGuRiBao, "小明")
	xiaGuRiBao.RegisterObserver(mingMing)
	huaHua := NewObserver(xiaGuRiBao, "小花")
	xiaGuRiBao.RegisterObserver(huaHua)

	xiaGuRiBao.SetNewMessage(true, "孙悟空全面加强")

}
