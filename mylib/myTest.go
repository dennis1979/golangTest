package mytest

// this is the entry point of test cases
// you have to call each test function here
func ExecTest() {
	HelloTest()

	VarTest()
	EnumTest()

	OpTestArithmetic()
	OpTestRelation()
	OpTestBoolean()
	OpTestBit()
	OpTestAssign()
	OpTestPriority()
	OpTestOther()

	LoopTest()

	ArrayOneDimTest()
	ArrayTwoDimTest()

	PointerTest()
	PtrArrayTest()
	Ptr2PtrTest()
	PtrParamTest()

	FuncTest()
	FuncArrayParamTest()
	FuncMultiReturnTest()
	FuncRecurTest()
	FuncClosureTest()

	RangeTest()
	StructTest()
	SliceTest()
	TypeCastTest()
	MapTest()

	InterfaceTest()

	MethodTest()
	ErrorTest()
}
