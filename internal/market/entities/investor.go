package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}
type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func (investor *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	investor.AssetPosition = append(investor.AssetPosition, assetPosition)
}

func (investor *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := investor.GetAssetPosition(assetID)
	if assetPosition == nil {
		investor.AssetPosition = append(investor.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (investor *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPositoon := range investor.AssetPosition {
		if assetPositoon.AssetID == assetID {
			return assetPositoon
		}
	}
	return nil
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}
