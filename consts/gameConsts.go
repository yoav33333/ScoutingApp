package consts

func GetGameTemp() []string {
	return append(GetGameNonLimit(), GetGamePark()...)
}
func GetGamePark() []string {
	return []string{"ascent", "netZone", "observation zone"}
}
func GetGameNonLimit() []string {
	return []string{"highBasket", "lowBasket", "highChamber", "lowChamber"}
}
