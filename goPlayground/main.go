package main
import "fmt"
import "strconv"
import "time"
import "sync"
import "goPlayground"

//Enums & consts
const {
    enumSimilar = 1
    enumSImilar2 = 2
}

//OOP:
type aStructType struct {
    firstVariable string
    secondVariable int
    thirdPointerVariable *int
}

//Find unique use of new...
func New(firstVariable string, secondVariable int) aStructType {
    newType := aStructType{firstVariable, secondVariable, nil}
    return newType
}

//Struct functions
func (instance aStructType) structFunction() {
    fmt.Println("Actioning function of struct: ")
    fmt.Println(instance.firstVariable)
    fmt.Println(instance.secondVariable)
}

//Functions
func functionWithType(inputType aStructType) aStructType {
    newSecond, err := strconv.Atoi(inputType.firstVariable)
    if err != nil {
        panic(err)
    }

    returnType := aStructType { strconv.Itoa(inputType.secondVariable), newSecond, nil}
    return returnType
}

//Routine function (threadding)
func goRoutineFunction(inputType* aStructType, mutexVar sync.Mutex, returnChannel chan int) { //No return

    mutexVar.Lock()

    fmt.Println(inputType.secondVariable)

    returnChannel <- inputType.secondVariable //Return via channel

    mutexVar.Unlock()
}


type InterfaceType interface {
    interfaceFunction() int
}

type InterfaceFloater int
func (f InterfaceFloater) interfaceFunction() int {
    return int(f)
}

func main() {

    //loops
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    //Containers - Slice
    arrayVariable := [2]int{1,2}
    sliceVariable := []int{1,2}
    sliceA := []int{1,2,3,4,5}
    sliceB := []int{10,9,8,7,6}

    //Printing
    fmt.Println(sliceVariable)
    fmt.Println(arrayVariable)
    fmt.Println("hello world")

    //Maps

    //"Coding templates":
    //Diff cross comms
    //Diff int-int
    //Diff order/order
    //Diff combos
    //Diff un-inclusions
    //Diff mutations on all of the above
    //Diff summarisations of mutations of referecnes
    permutationsMap := make(map[int][]int)
    counter := 0

    //Loops
    for i, s1 := range sliceA {
        for j, s2 := range sliceB {


            permutationsMap[counter] = []int{s1, s2}

            counter++
            fmt.Println(i)
            fmt.Println(j)
        }
    }

    fmt.Println(permutationsMap)

    //Container manipulation
    delete(permutationsMap, 7) 
    clear(permutationsMap)


    //Structs
    myType := aStructType{"4321", 1234, nil}

    fmt.Println(myType.firstVariable)
    fmt.Println(myType.secondVariable)

    invertedType := functionWithType(myType)
    fmt.Println(invertedType.firstVariable)
    fmt.Println(invertedType.secondVariable)

    invertedType.structFunction()


    // Pointers
    var pointer *int
    i := 42
    pointer = &i
    fmt.Println(*pointer) //No delete (Garbage collection)


    // Threadding
    var mutexVar sync.Mutex
    returnChannel := make(chan int)
    go goRoutineFunction(&myType, mutexVar, returnChannel)
    go goRoutineFunction(&invertedType, mutexVar, returnChannel)

    returnMyType, returnInvertedType := <-returnChannel, <-returnChannel

    fmt.Println(returnMyType, returnInvertedType)
    //defer mutexVar.Unlock()


    //Select switch w/ channels
    go func() {
        tick := time.Tick(100 * time.Millisecond)
        boom := time.After(500 * time.Millisecond)
        for {
            select {
            case <-tick:
                fmt.Println("tick.")
            case <-boom:
                fmt.Println("BOOM!")
                return
            default:
                fmt.Println("    .")
                time.Sleep(50 * time.Millisecond)
            }
        }
    }()
    time.Sleep(time.Second)


    //Interfaces
    var interfaceVariable InterfaceType
    randomInterfaceFloater := InterfaceFloater(4)
    interfaceVariable = randomInterfaceFloater
    fmt.Println(interfaceVariable.interfaceFunction())



    //External files & file handling
    fileManipulator.fileInOutAndAbout()
}

