package item

type Item struct {
	ID                 uint16
	Count              uint8
	ChargeCount        uint16
	ActionID           int64
	UniqueID           int64
	FluidType          rune
	SpecialDescription string
	Text               string
}

func NewItem() *Item {
	return &Item{}
}

func (i *Item) GetUniqueID() int64    { return i.ActionID }
func (i *Item) GetActionID() int64    { return i.ActionID }
func (i *Item) GetItemCharge() uint16 { return i.ChargeCount }
func (i *Item) GetFluidType() rune    { return i.FluidType }
