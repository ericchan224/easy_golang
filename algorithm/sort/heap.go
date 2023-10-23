package sort

// 先自底向上构建最大堆，再移除堆元素实现堆排序
func HeapSort(array []int) {
	// 堆的元素数量
	count := len(array)

	// 最底层的叶子节点下标，该节点位置不定，但是该叶子节点右边的节点都是叶子节点
	start := count/2 + 1

	// 最后的元素下标
	end := count - 1

	// 从最底层开始，逐一对节点进行下沉
	for start >= 0 {
		sift(array, start, count)
		start-- // 表示左偏移一个节点，如果该层没有节点了，那么表示到了上一层的最右边
	}

	// 下沉结束了，现在要来排序了
	// 元素大于2个的最大堆才可以移除
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		// 对堆顶进行下沉操作
		sift(array, 0, end)
		// 一直移除堆顶元素
		end--
	}
}

// 下沉操作，需要下沉的元素时 array[start]，参数 count 只要用来判断是否到底堆底，使得下沉结束
func sift(array []int, start, count int) {
	// 父亲节点
	root := start

	// 左儿子
	child := root*2 + 1

	// 如果有下一代
	for child < count {
		// 右儿子比左儿子大，那么要翻转的儿子改为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}

		// 父亲节点比儿子小，那么将父亲和儿子位置交换
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// 继续往下沉
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}

// 一个最大堆，一棵完全二叉树
// 最大堆要求节点元素都不小于其左右孩子
type Heap struct {
	// 堆的大小
	Size int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

// 初始化一个堆
func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array
	return h
}

// 最大堆插入元素
func (h *Heap) Push(x int) {
	// 堆没有元素时，使元素成为顶点后退出
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	// i 是要插入节点的下标
	i := h.Size

	// 如果下标存在
	// 将小的值 x 一直上浮
	for i > 0 {
		// parent为该元素父亲节点的下标
		parent := (i - 1) / 2

		// 如果插入的值小于等于父亲节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x <= h.Array[parent] {
			break
		}

		// 否则将父亲节点与该节点互换，然后向上翻转，将最大的元素一直往上推
		h.Array[i] = h.Array[parent]
		i = parent
	}

	// 将该值 x 放在不会再翻转的位置
	h.Array[i] = x

	// 堆数量加一
	h.Size++
}

// 最大堆移除根节点元素，也就是最大的元素
func (h *Heap) Pop() int {
	// 没有元素，返回-1
	if h.Size == 0 {
		return -1
	}

	// 取出根节点
	ret := h.Array[0]

	// 因为根节点要被删除了，将最后一个节点放到根节点的位置上
	h.Size--
	x := h.Array[h.Size]  // 将最后一个元素的值先拿出来
	h.Array[h.Size] = ret // 将移除的元素放在最后一个元素的位置上

	// 对根节点进行向下翻转，小的值 x 一直下沉，维持最大堆的特征
	i := 0
	for {
		// a，b为下标 i 左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 左儿子下标超出了，表示没有左子树，那么右子树也没有，直接返回
		if a >= h.Size {
			break
		}

		// 有右子树，拿到两个子节点中较大节点的下标
		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b
		}

		// 父亲节点的值都大于或等于两个儿子较大的那个，不需要向下继续翻转了，返回
		if x >= h.Array[a] {
			break
		}

		// 将较大的儿子与父亲交换，维持这个最大堆的特征
		h.Array[i] = h.Array[a]

		// 继续往下操作
		i = a
	}

	// 将最后一个元素的值 x 放在不会再翻转的位置
	h.Array[i] = x
	return ret
}
