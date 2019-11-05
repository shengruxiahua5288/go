package main

type MyCircularDeque struct {
	size   int
	length int
	head   *DListNode
	rear   *DListNode
}
type DListNode struct {
	value int
	next  *DListNode
	pre   *DListNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	var myCircularDeque MyCircularDeque
	myCircularDeque.size = k
	myCircularDeque.length = 0
	myCircularDeque.head = nil
	myCircularDeque.rear = nil
	return myCircularDeque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {

	if this.length == this.size {
		return false
	}

	var temp *DListNode = new(DListNode)
	temp.value = value
	temp.pre = this.head
	temp.next = this.head.next
	this.head.next.pre = temp
	head.next = temp

	this.length = this.length + 1

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {

	if this.length == this.size {
		return false
	} else {
		var temp *DListNode = new(DListNode)
		temp.value = value

		if this.head == nil && this.rear == nil {

			this.rear = temp
			this.head = this.rear

		} else {
			this.rear.next = temp
			temp.next = this.head
			temp.pre = this.rear
			this.rear = temp
		}

		this.length = this.length + 1
		return true
	}
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	var length int = this.length
	if length == 0 {
		return false
	}
	if this.head == this.rear {
		this.head = nil
		this.rear = nil
		return true
	}
	var temp *DListNode = this.head.next
	this.head.next = nil
	this.head.pre = nil
	temp.pre = this.rear
	this.head = temp
	this.length = this.length - 1
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	var length int = this.length
	if length == 0 {
		return false
	} else {
		if this.head == this.rear {
			this.head = nil
			this.rear = nil
			return true
		}
		var temp *DListNode = this.rear.pre
		this.rear.pre = nil
		this.rear.next = nil
		temp.next = this.head
		this.rear = temp
		this.length = this.length - 1
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.head != nil {
		return this.head.value
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {

	if this.rear != nil {
		return this.rear.value
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}

}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.length == this.size {
		return true
	} else {
		return false
	}
}

// 打印出队列中的值
func (this *MyCircularDeque) printDeque() {
	var temp *DListNode = this.head
	for temp != nil {
		print(temp.value)
		temp = temp.next
	}
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {

	example := [...]string{"MyCircularDeque", "insertFront", "getRear", "insertLast", "getFront", "insertLast", "getFront", "insertFront", "getRear", "getRear", "deleteFront", "deleteLast", "isFull", "getRear", "getRear", "getFront", "getRear", "deleteLast", "insertLast", "getFront", "isEmpty", "insertLast", "insertLast", "getRear", "insertFront", "insertLast", "deleteFront", "getRear", "getFront", "isFull", "isFull", "insertLast", "getRear", "getFront", "insertLast", "getRear", "deleteLast", "getRear", "getFront", "getRear", "insertFront", "getFront", "getFront", "getRear", "getRear", "insertFront", "getRear", "insertLast", "insertFront", "getRear", "getFront", "getFront", "insertLast", "getFront", "deleteFront", "getFront", "deleteLast", "getRear", "deleteLast", "getRear", "getRear", "getFront", "isEmpty", "getRear", "deleteLast", "insertFront", "insertFront", "getFront", "deleteFront", "insertLast", "getRear", "insertFront", "insertLast", "insertFront", "insertLast", "insertFront", "getFront", "getRear", "insertFront", "deleteLast", "getRear", "isFull", "insertLast", "getRear", "getFront", "getFront", "insertFront", "getRear", "getRear", "deleteFront", "isEmpty", "isFull", "deleteLast", "insertFront", "getFront", "insertFront", "deleteLast", "insertLast", "getRear", "insertFront", "getFront", "insertLast"}
	nums := [...]int{96, 72, 0, 34, 0, 38, 0, 91, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 27, 0, 0, 68, 42, 0, 68, 17, 0, 0, 0, 0, 0, 19, 0, 0, 94, 0, 0, 0, 0, 0, 85, 0, 0, 0, 0, 59, 0, 40, 60, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 59, 99, 0, 0, 20, 0, 75, 12, 46, 72, 0, 0, 0, 5, 0, 0, 0, 74, 0, 0, 0, 93, 0, 0, 0, 0, 0, 0, 37, 0, 37, 0, 11, 0, 66, 0, 63}

	obj := Constructor(52)
	for i := 1; i < len(example); i++ {

		if "insertFront" == example[i] {
			print(obj.InsertFront(nums[i]))
		}
		if "insertLast" == example[i] {
			print(obj.InsertLast(nums[i]))
		}
		if "getFront" == example[i] {
			print(obj.GetFront())
		}
		if "getRear" == example[i] {
			print(obj.GetRear())
		}
		if "deleteFront" == example[i] {
			print(obj.DeleteFront())
		}
		if "deleteLast" == example[i] {
			print(obj.DeleteLast())
		}
		if "isEmpty" == example[i] {
			print(obj.IsEmpty())
		}
		if "isFull" == example[i] {
			print(obj.IsFull())
		}
	}
}

// obj := Constructor(52)
// var param_1 = obj.InsertFront(80)
// print(param_1)
// println()
// obj.printDeque()

// println()
// var param_11 = obj.InsertFront(5)
// print(param_11)
// obj.printDeque()

// println()
// var param_2 = obj.InsertLast(8)
// print(param_2)
// obj.printDeque()
// var param_21 = obj.InsertLast(9)
// print(param_21)
// obj.printDeque()

// println()
// var param_3 = obj.DeleteFront()
// print(param_3)
// println()
// obj.printDeque()

// println()
// var param_4 = obj.DeleteLast()
// print(param_4)
// println()
// obj.printDeque()

// println()
// var param_5 = obj.GetFront()
// print(param_5)
// println()
// obj.printDeque()

// println()
// var param_6 = obj.GetRear()
// print(param_6)
// println()
// obj.printDeque()

// println()
// var param_7 = obj.IsEmpty()
// print(param_7)
// obj.printDeque()

// println()
// var param_8 = obj.IsFull()
// print(param_8)
// obj.printDeque()
