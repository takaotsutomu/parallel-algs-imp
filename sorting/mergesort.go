package sorting

func MergeSort(data []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(data, left, mid)
		MergeSort(data, mid+1, right)
		merge(data, left, mid, right)
	}
}

type rank struct {
	x int
	r int
}

type rankM struct {
	x  int
	rM int
}

func merge(data []int, left, mid, right int) {
	rACh := make(chan *rank, mid-left+1)
	rBCh := make(chan *rank, right-mid)
	rAMCh := make(chan *rankM, mid-left+1)
	rBMCh := make(chan *rankM, right-mid)
	for i := 0; i < mid-left+1; i++ {
		go biSearchA(data, mid+1, right, rACh, rAMCh)
	}
	for i := 0; i < right-mid; i++ {
		go biSearchB(data, left, mid, rBCh, rBMCh)
	}

	for i := left; i <= mid; i++ {
		rACh <- &rank{data[i], i}
	}
	close(rACh)

	for i := mid + 1; i <= right; i++ {
		rBCh <- &rank{data[i], i - mid - 1}
	}
	close(rBCh)

	temp := make([]int, right-left+1)
	for i := 0; i < len(temp); i++ {
		select {
		case rankA := <-rAMCh:
			temp[rankA.rM-left] = rankA.x
		case rankB := <-rBMCh:
			temp[rankB.rM-left] = rankB.x
		}
	}

	i := left
	for j := 0; j < len(temp); j++ {
		data[i] = temp[j]
		i++
	}
}

func biSearchA(data []int, left, right int, rACh chan *rank, rMCh chan *rankM) {
	whence := left
	rA := <-rACh
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == rA.x {
			rMCh <- &rankM{rA.x, rA.r + mid - whence}
			return
		} else if data[mid] < rA.x {
			left = mid + 1
		} else {
			right = mid - 1

		}
	}
	rMCh <- &rankM{rA.x, rA.r + left - whence}

}

func biSearchB(data []int, left, right int, rBCh chan *rank, rMCh chan *rankM) {
	rB := <-rBCh
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == rB.x {
			rMCh <- &rankM{rB.x, rB.r + mid + 1}
			return
		} else if data[mid] < rB.x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	rMCh <- &rankM{rB.x, rB.r + left}
}
