package shared

import (
	"strconv"
	"strings"
)

func (q Queue) String() string {
	var ids []string
	for id, itmCnt := range q {
		ids = append(ids, "[("+strconv.Itoa(itmCnt)+")\t"+strconv.Itoa(id)+"]")
	}
	return strings.Join(ids, "\n")
}

func (s Session) ItemCount() int {
	var cnt int
	for _, ord := range s.Cubbies {
		cnt += ord.ItemCount()
	}
	return cnt
}

func (o Order) ItemCount() int {
	var cnt int
	for _, itm := range o.Items {
		cnt += itm.Quantity
	}
	return cnt
}

func (o Order) Item(upc string) Item {
	for _, itm := range o.Items {
		if itm.UPC != upc && itm.SKU != upc {
			println(`itm.UPC vs upc |`, itm.UPC, upc)
			continue
		}
		return itm
	}
	println("Item: failed to find UPC")
	return Item{}
}

func (s Session) String() string {
	var buf []string
	for idx, ord := range s.Cubbies {
		buf = append(buf, strconv.Itoa(idx+1)+" "+ord.String())
	}
	return strings.Join(buf, "\n")
}

func (o Order) String() string {
	var itms []string
	for _, itm := range o.Items {
		skupc := itm.UPC
		if len(skupc) == 0 {
			skupc = itm.SKU
		}
		itms = append(itms, "["+strconv.Itoa(itm.Quantity)+"x "+skupc+"]")
	}
	return o.OrderNumber + ":" + strings.Repeat(" ", 24-len(o.OrderNumber)) + strings.Join(itms, ", ")
}
