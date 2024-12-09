package main

import (
	"fmt"
)

type block struct {
	id, size int
}

func toBlocks(parts []int) []block {
	id := 0
	blocks := make([]block, 0)
	for i, part := range parts {
		if part == 0 { continue }
		if i%2 == 0 {
			blocks = append(blocks, block{id: id, size: part})
			id++
			} else {
				blocks = append(blocks, block{id: -1, size: part})
			}
	}
	return blocks
}

func toString(blocks []block) string {
	res := ""
	for _, block := range blocks {
		for range block.size {
			if block.id == -1 {
				res = fmt.Sprintf("%s%s", res, ".")
				continue
			}
			res = fmt.Sprintf("%s%d", res, block.id)
		}
		res = fmt.Sprintf("%s", res)
	}
	return res
}

func collapseBlocks(blocks []block) []block {
	for i := len(blocks)-1; i > 0; i-- {
		rBlock := blocks[i]
		if rBlock.id == -1 { continue }
		for j := 0; j < i; j++ {
			lBlock := blocks[j]
			if lBlock.id != -1 {
				continue
			}

			if lBlock.size >= rBlock.size {
				if lBlock.size == rBlock.size {
					blocks[i] = lBlock
					blocks[j] = rBlock
					
				} else if lBlock.size > rBlock.size {
					diff := lBlock.size - rBlock.size

					beforeRef := blocks[:j]; 	before := make([]block, len(beforeRef))
					afterRef := blocks[j+1:]; 	after := make([]block, len(afterRef))
					
					copy(before, beforeRef)
					copy(after, afterRef)

					before = append(before, rBlock, block{id: -1, size: diff})
					blocks = append(before, after...)
					i++ //only by one, since free space is moved to end
					blocks[i] = block{id: -1, size: rBlock.size}
				}
				break
			}
		}
	}
	return blocks
}

func checksumBlocks(blocks []block) int {
	total := 0
	pos := 0
	for _, block := range blocks {
		if block.id == -1 {
			pos += block.size
			continue
		}
		for range block.size {
			total += block.id * pos
			pos++
		}
	}
	return total
}