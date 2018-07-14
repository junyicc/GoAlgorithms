package designpattern

// WaterWork interface
type WaterWork interface {
	OnDemand(m WaterMediator, amount float64)
}

// WaterTreatmentWork struct
type WaterTreatmentWork struct {
	PotableWater float64
}

// OutputToStorage for WaterTreatmentWork
func (wtw *WaterTreatmentWork) OutputToStorage(amount float64) float64 {
	if amount >= wtw.PotableWater {
		amount = wtw.PotableWater
	}
	wtw.PotableWater -= amount
	return amount
}

// WaterStorageWork struct
type WaterStorageWork struct {
	MaxStorage float64
	Storage    float64
}

// InputFromWTW for WaterStorageWork
func (wsw *WaterStorageWork) InputFromWTW(amount float64) {
	wsw.Storage += amount
}

// OnDemand for WaterStorageWork
func (wsw *WaterStorageWork) OnDemand(m WaterMediator, amount float64) {
	if amount+wsw.Storage > wsw.MaxStorage {
		amount = wsw.MaxStorage - wsw.Storage
	}
	m.TransferWater(amount)
}

// WaterMediator for Mediator pattern
type WaterMediator interface {
	TransferWater(amount float64)
}

// WTWWSWMediator is a mediator for WTW and WSW
type WTWWSWMediator struct {
	WTW *WaterTreatmentWork
	WSW *WaterStorageWork
}

// TransferWater transfers water from WTW to WSW
func (m *WTWWSWMediator) TransferWater(amount float64) {
	amount = m.WTW.OutputToStorage(amount)
	m.WSW.InputFromWTW(amount)
}

// NewWTWWSWMediator return WTWWSWMediator
func NewWTWWSWMediator(wtw *WaterTreatmentWork, wsw *WaterStorageWork) *WTWWSWMediator {
	m := new(WTWWSWMediator)
	m.WTW = wtw
	m.WSW = wsw
	return m
}
