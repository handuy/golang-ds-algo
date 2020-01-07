package main

import (
	"log"
)

type Order struct {
	priority int
	quantity int
	name string
	customerName string
}

// Danh sách các Order sắp xếp theo trường priority của Order: cao lên đầu, thấp xuống dưới
type PriorityQueue []*Order

// Tạo Order mới
func(order *Order) New(priority int, quantity int, name string, customer string) {
	order.priority = priority
	order.quantity = quantity
	order.name = name
	order.customerName = customer
}

// Thêm Order vào hàng đợi theo đúng nguyên tắc: Order nào có priority cao thì xếp trước, priority thấp bị ủn xuống dưới
func(queue *PriorityQueue) Add(order *Order) {
	// Trong hàng chưa có Order nào thì cứ thế nhét Order vào hàng
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended = false
		var orderInQueue *Order
		var i int

		// So sánh order được thêm vào với từng order trong hàng
		for i, orderInQueue = range *queue {
			// Nếu phát hiện order trong hàng có priority < order được thêm vào 
			// thì order mới sẽ được nhét lên trước, ủn các order còn lại xuống dưới
			// Ví dụ trong 1 hàng có các order với priority theo thứ tự: [6, 3, 2]
			// Giờ ta thêm một order với priority = 4 vào
			// Do 4 > 3 --> order 4 sẽ được chèn vào giữa 6 và 3
			// Hàng lúc này sẽ là: [6, 4, 3, 2]
			if order.priority >= orderInQueue.priority {
				var newQueue PriorityQueue = []*Order{order}
				*queue = append( (*queue)[:i], append( newQueue, (*queue)[i:]... )... )
				appended = true
				break
			}
		}

		// Nếu appended vẫn bằng false chứng tỏ order mới có priority thấp hơn tất cả các order trong hàng
		// Lúc đó ta nhét nó vào đáy hàng
		if !appended {
			*queue = append( (*queue), order )
		}
	}
}

func(queue *PriorityQueue) ListItemInQueue() {
	for _, item := range *queue {
		log.Println(*item)
	}
	log.Println("---------------------------")
}

func main() {
	var queue PriorityQueue

	var order1 Order
	order1.New(1, 10, "apple", "Steve Job")
	queue.Add(&order1)
	queue.ListItemInQueue()

	var order2 Order
	order2.New(3, 20, "amazon", "Voldemort")
	queue.Add(&order2)
	queue.ListItemInQueue()

	var order3 Order
	order3.New(2, 15, "facebook", "Mark")
	queue.Add(&order3)
	queue.ListItemInQueue()

	var order4 Order
	order4.New(9, 30, "google", "Lisa")
	queue.Add(&order4)
	queue.ListItemInQueue()

	var order5 Order
	order5.New(7, 30, "ibm", "Kid")
	queue.Add(&order5)
	queue.ListItemInQueue()

	var order6 Order
	order6.New(2, 19, "samsung", "Quill")
	queue.Add(&order6)
	queue.ListItemInQueue()

	var order7 Order
	order7.New(1, 12, "panasonic", "Milk")
	queue.Add(&order7)
	queue.ListItemInQueue()

	var order8 Order
	order8.New(0, 12, "disney", "Gotham")
	queue.Add(&order8)
	queue.ListItemInQueue()
}