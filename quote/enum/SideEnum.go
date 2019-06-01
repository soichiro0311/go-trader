package enum

type SideEnum int

const (
	BUY SideEnum = iota
	SELL
)

func (buySell SideEnum) String() string {
	switch buySell {
	case BUY:
		return "BUY"

	case SELL:
		return "SELL"
	default:
		return "-"
	}
}
