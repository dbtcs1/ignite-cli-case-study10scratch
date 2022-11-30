package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BuyOrderBookKeyPrefix is the prefix to retrieve all BuyOrderBook
	BuyOrderBookKeyPrefix = "BuyOrderBook/value/"
)

// BuyOrderBookKey returns the store key to retrieve a BuyOrderBook from the index fields
func BuyOrderBookKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func (b *BuyOrderBook) FillSellOrder(order Order) (
    remainingSellOrder Order,
    liquidated []Order,
    gain int32,
    filled bool,
) {
    var liquidatedList []Order
    totalGain := int32(0)
    remainingSellOrder = order

    // Liquidate as long as there is match
    for {
        var match bool
        var liquidation Order
        remainingSellOrder, liquidation, gain, match, filled = b.LiquidateFromSellOrder(
            remainingSellOrder,
        )
        if !match {
            break
        }

        // Update gains
        totalGain += gain

        // Update liquidated
        liquidatedList = append(liquidatedList, liquidation)

        if filled {
            break
        }
    }

    return remainingSellOrder, liquidatedList, totalGain, filled
}
