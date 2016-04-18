//Package ordercancelreplacerequest msg type = G.
package ordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
	"time"
)

//NoAllocs is a repeating group in OrderCancelReplaceRequest
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties is a non-required component for NoAllocs.
	NestedParties *nestedparties.NestedParties
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string)                       { m.AllocAccount = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                  { m.IndividualAllocID = &v }
func (m *NoAllocs) SetNestedParties(v nestedparties.NestedParties) { m.NestedParties = &v }
func (m *NoAllocs) SetAllocQty(v float64)                          { m.AllocQty = &v }

//NoTradingSessions is a repeating group in OrderCancelReplaceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//NewNoTradingSessions returns an initialized NoTradingSessions instance
func NewNoTradingSessions() *NoTradingSessions {
	var m NoTradingSessions
	return &m
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a OrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"G"`
	fix43.Header
	//OrderID is a non-required field for OrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//Parties is a non-required component for OrderCancelReplaceRequest.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for OrderCancelReplaceRequest.
	TradeOriginationDate *string `fix:"229"`
	//OrigClOrdID is a required field for OrderCancelReplaceRequest.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for OrderCancelReplaceRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderCancelReplaceRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for OrderCancelReplaceRequest.
	ClOrdLinkID *string `fix:"583"`
	//ListID is a non-required field for OrderCancelReplaceRequest.
	ListID *string `fix:"66"`
	//OrigOrdModTime is a non-required field for OrderCancelReplaceRequest.
	OrigOrdModTime *time.Time `fix:"586"`
	//Account is a non-required field for OrderCancelReplaceRequest.
	Account *string `fix:"1"`
	//AccountType is a non-required field for OrderCancelReplaceRequest.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for OrderCancelReplaceRequest.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for OrderCancelReplaceRequest.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for OrderCancelReplaceRequest.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for OrderCancelReplaceRequest.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for OrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//CashMargin is a non-required field for OrderCancelReplaceRequest.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for OrderCancelReplaceRequest.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a required field for OrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for OrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for OrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for OrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for OrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//Instrument is a required component for OrderCancelReplaceRequest.
	instrument.Instrument
	//Side is a required field for OrderCancelReplaceRequest.
	Side string `fix:"54"`
	//TransactTime is a required field for OrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//QuantityType is a non-required field for OrderCancelReplaceRequest.
	QuantityType *int `fix:"465"`
	//OrderQtyData is a required component for OrderCancelReplaceRequest.
	orderqtydata.OrderQtyData
	//OrdType is a required field for OrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for OrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for OrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for OrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData is a non-required component for OrderCancelReplaceRequest.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for OrderCancelReplaceRequest.
	YieldData *yielddata.YieldData
	//PegDifference is a non-required field for OrderCancelReplaceRequest.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for OrderCancelReplaceRequest.
	DiscretionOffset *float64 `fix:"389"`
	//ComplianceID is a non-required field for OrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
	SolicitedFlag *bool `fix:"377"`
	//Currency is a non-required field for OrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for OrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for OrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for OrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//CommissionData is a non-required component for OrderCancelReplaceRequest.
	CommissionData *commissiondata.CommissionData
	//OrderCapacity is a non-required field for OrderCancelReplaceRequest.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for OrderCancelReplaceRequest.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for OrderCancelReplaceRequest.
	CustOrderCapacity *int `fix:"582"`
	//Rule80A is a non-required field for OrderCancelReplaceRequest.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for OrderCancelReplaceRequest.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for OrderCancelReplaceRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelReplaceRequest.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for OrderCancelReplaceRequest.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for OrderCancelReplaceRequest.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for OrderCancelReplaceRequest.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for OrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//LocateReqd is a non-required field for OrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//CancellationRights is a non-required field for OrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for OrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for OrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for OrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//AccruedInterestRate is a non-required field for OrderCancelReplaceRequest.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for OrderCancelReplaceRequest.
	AccruedInterestAmt *float64 `fix:"159"`
	//NetMoney is a non-required field for OrderCancelReplaceRequest.
	NetMoney *float64 `fix:"118"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized OrderCancelReplaceRequest instance
func New(origclordid string, clordid string, handlinst string, instrument instrument.Instrument, side string, transacttime time.Time, orderqtydata orderqtydata.OrderQtyData, ordtype string) *Message {
	var m Message
	m.SetOrigClOrdID(origclordid)
	m.SetClOrdID(clordid)
	m.SetHandlInst(handlinst)
	m.SetInstrument(instrument)
	m.SetSide(side)
	m.SetTransactTime(transacttime)
	m.SetOrderQtyData(orderqtydata)
	m.SetOrdType(ordtype)
	return &m
}

func (m *Message) SetOrderID(v string)                         { m.OrderID = &v }
func (m *Message) SetParties(v parties.Parties)                { m.Parties = &v }
func (m *Message) SetTradeOriginationDate(v string)            { m.TradeOriginationDate = &v }
func (m *Message) SetOrigClOrdID(v string)                     { m.OrigClOrdID = v }
func (m *Message) SetClOrdID(v string)                         { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)                { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)                     { m.ClOrdLinkID = &v }
func (m *Message) SetListID(v string)                          { m.ListID = &v }
func (m *Message) SetOrigOrdModTime(v time.Time)               { m.OrigOrdModTime = &v }
func (m *Message) SetAccount(v string)                         { m.Account = &v }
func (m *Message) SetAccountType(v int)                        { m.AccountType = &v }
func (m *Message) SetDayBookingInst(v string)                  { m.DayBookingInst = &v }
func (m *Message) SetBookingUnit(v string)                     { m.BookingUnit = &v }
func (m *Message) SetPreallocMethod(v string)                  { m.PreallocMethod = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)                    { m.NoAllocs = v }
func (m *Message) SetSettlmntTyp(v string)                     { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)                     { m.FutSettDate = &v }
func (m *Message) SetCashMargin(v string)                      { m.CashMargin = &v }
func (m *Message) SetClearingFeeIndicator(v string)            { m.ClearingFeeIndicator = &v }
func (m *Message) SetHandlInst(v string)                       { m.HandlInst = v }
func (m *Message) SetExecInst(v string)                        { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                         { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                       { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                   { m.ExDestination = &v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions)  { m.NoTradingSessions = v }
func (m *Message) SetInstrument(v instrument.Instrument)       { m.Instrument = v }
func (m *Message) SetSide(v string)                            { m.Side = v }
func (m *Message) SetTransactTime(v time.Time)                 { m.TransactTime = v }
func (m *Message) SetQuantityType(v int)                       { m.QuantityType = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData) { m.OrderQtyData = v }
func (m *Message) SetOrdType(v string)                         { m.OrdType = v }
func (m *Message) SetPriceType(v int)                          { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                          { m.Price = &v }
func (m *Message) SetStopPx(v float64)                         { m.StopPx = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData)                { m.YieldData = &v }
func (m *Message) SetPegDifference(v float64)                        { m.PegDifference = &v }
func (m *Message) SetDiscretionInst(v string)                        { m.DiscretionInst = &v }
func (m *Message) SetDiscretionOffset(v float64)                     { m.DiscretionOffset = &v }
func (m *Message) SetComplianceID(v string)                          { m.ComplianceID = &v }
func (m *Message) SetSolicitedFlag(v bool)                           { m.SolicitedFlag = &v }
func (m *Message) SetCurrency(v string)                              { m.Currency = &v }
func (m *Message) SetTimeInForce(v string)                           { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)                      { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                            { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)                         { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)                            { m.GTBookingInst = &v }
func (m *Message) SetCommissionData(v commissiondata.CommissionData) { m.CommissionData = &v }
func (m *Message) SetOrderCapacity(v string)                         { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)                     { m.OrderRestrictions = &v }
func (m *Message) SetCustOrderCapacity(v int)                        { m.CustOrderCapacity = &v }
func (m *Message) SetRule80A(v string)                               { m.Rule80A = &v }
func (m *Message) SetForexReq(v bool)                                { m.ForexReq = &v }
func (m *Message) SetSettlCurrency(v string)                         { m.SettlCurrency = &v }
func (m *Message) SetText(v string)                                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                           { m.EncodedText = &v }
func (m *Message) SetFutSettDate2(v string)                          { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)                            { m.OrderQty2 = &v }
func (m *Message) SetPrice2(v float64)                               { m.Price2 = &v }
func (m *Message) SetPositionEffect(v string)                        { m.PositionEffect = &v }
func (m *Message) SetCoveredOrUncovered(v int)                       { m.CoveredOrUncovered = &v }
func (m *Message) SetMaxShow(v float64)                              { m.MaxShow = &v }
func (m *Message) SetLocateReqd(v bool)                              { m.LocateReqd = &v }
func (m *Message) SetCancellationRights(v string)                    { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)                 { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                              { m.RegistID = &v }
func (m *Message) SetDesignation(v string)                           { m.Designation = &v }
func (m *Message) SetAccruedInterestRate(v float64)                  { m.AccruedInterestRate = &v }
func (m *Message) SetAccruedInterestAmt(v float64)                   { m.AccruedInterestAmt = &v }
func (m *Message) SetNetMoney(v float64)                             { m.NetMoney = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX43, "G", r
}
