package main

import (
	"errors"
	"fmt"
	"sync"
)

type Bank struct {
	sync.RWMutex
	balance map[string]float64
}

var xingxiaoli = "xingxiaoli"

// 收入
func (b *Bank) Income(account string, in float64)  {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.balance[account]; !ok {
		b.balance[account] = 0.0
	}
	
	b.balance[account] += in
}

// 支出
func (b *Bank) Expenditure(account string, out float64) error {
	b.Lock()
	defer b.Unlock()

	v, ok := b.balance[account];
	if !ok || v < out {
		return errors.New("account not enough balance~")
	}

	b.balance[account] -= out
	return nil
}

// 查询
func (b *Bank) Query(account string) (float64, error) {
	b.RLock()
	defer b.RUnlock()

	v, ok := b.balance[account];
	if !ok {
		return 0.0, errors.New(fmt.Sprintf("don't exist account: %s", account))
	}

	return v, nil
}

func NewBank() *Bank {
	return &Bank{
		RWMutex:   sync.RWMutex{},
		balance: map[string]float64{},
	}
}

// 初始化一个账户
func InitBank() *Bank {
	b := NewBank()

	b.balance[xingxiaoli] = 1000
	return b
}

func main() {
	// 获取账户
	newB := InitBank()
	fmt.Println("New Bank account: ", newB.balance)
	// 今日收入
	newB.Income(xingxiaoli, 100)
	balance, err := newB.Query(xingxiaoli)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("xingxiaoli new income balance: ", balance)
	// 近日支出
	newB.Expenditure(xingxiaoli, 50)
	balance, err = newB.Query("xingxiaoli1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("xingxiaoli new expenditure balance: ", balance)
}
