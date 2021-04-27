package lib

import (
	"errors"
	"fmt"
)

// GoTickets interface of the pool of tickets of goroutines
type GoTickets interface {
	// to take one ticket
	Take()
	// to return one ticket
	Return()
	// whether the ticket pool has been activated or not
	Active() bool
	// total amount of tickets
	Total() uint32
	// the remainder of tickets
	Remainder() uint32
}

// myGoTickets implementation of gouroutine ticket pool
type myGoTickets struct {
	total    uint32        // the total amount of the tickets
	ticketCh chan struct{} // the container of the tickets
	active   bool          // whether the ticket has been activated or not
}

// NewGoTickets creates a ticket pool of goroutine
func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg := fmt.Sprintf("The goroutine ticket pool cannot be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}

func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take() {
	<-gt.ticketCh
}

func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}

func (gt *myGoTickets) Active() bool {
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
