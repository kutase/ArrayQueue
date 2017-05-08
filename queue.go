package ArrayQueue

import "log"

type ArrayQueue struct {
	maxIndex     int32
	array        []*interface{}
	writePointer int32
	readPointer  int32
	length       int32
	wasFull      bool
}

func New(length int32) *ArrayQueue {
	queue := new(ArrayQueue)
	queue.array = make([]*interface{}, length)
	queue.maxIndex = length - 1
	queue.length = length
	queue.writePointer = 0
	queue.readPointer = -1
	return queue
}

func (queue *ArrayQueue) incWritePointer() {
	val := queue.writePointer + 1

	if val == queue.length {
		val = 0
		queue.wasFull = true
	}

	if val == queue.readPointer {
		log.Panicln("WritePointer: Out of free space in Queue. Index:", val)
	}

	queue.writePointer = val
}

func (queue *ArrayQueue) incReadPointer() {
	val := queue.readPointer + 1

	if val == queue.length {
		val = 0
	}

	if val == queue.writePointer {
		log.Panicln("ReadPointer: Out of free space in Queue. Index:", val)
	}

	queue.readPointer = val
}

func (queue *ArrayQueue) Enqueue(cmd interface{}) {
	queue.array[ queue.writePointer ] = &cmd
	queue.incWritePointer()
}

func (queue *ArrayQueue) Dequeue() *interface{} {
	if queue.readPointer + 1 == queue.writePointer || (queue.readPointer == queue.maxIndex && queue.writePointer == 0) {
		return nil
	}

	queue.incReadPointer()
	return queue.array[ queue.readPointer ]
}

func (queue *ArrayQueue) GetLastElements(length int32) []*interface{} {
	if length > queue.length {
		log.Panicln("Length", length, "cannot be greater then maxLength:", queue.maxIndex + 1)
	}

	var queueElementsLength int32 = queue.writePointer

	if queue.wasFull {
		queueElementsLength = queue.length
	}

	if length > queueElementsLength {
		length = queueElementsLength
	}

	var buffer []*interface{} = make([]*interface{}, length)

	pointer := queue.writePointer - 1
	if pointer == -1 {
		if !queue.wasFull {
			log.Println("Queue is empty")
			return nil
		}

		pointer = queue.maxIndex
	}

	var bufferIndex int32 = 0

	for bufferIndex < length {
		element := queue.array[pointer]

		buffer[bufferIndex] = element
		bufferIndex += 1

		pointer -= 1

		if pointer == -1 {
			if !queue.wasFull {
				break
			}

			pointer = queue.maxIndex
		}

	}

	return buffer
}
