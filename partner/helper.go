package partner

import (
	restgostruct "github.com/incompetent-developer/restgo-struct"
)

//CalculateOnlineWithoutVat is calculate online partner without VAT 's amount
// price-(fee/100*price)
// return result,fee
func CalculateOnlineWithoutVat(price, fee float64, isValue bool) (float64, float64) {
	if !isValue {
		fee = restgostruct.ToFixed((fee / 100 * price), 2)
	}

	return restgostruct.ToFixed(price-fee, 2), fee
}

//CalculateOnlineWithVat is calculate online partner with VAT 's amount
// price+(vat/100*price)-(fee/100*(price+(vat/100*price)))
// return result,fee,vat
func CalculateOnlineWithVat(price, vatPercent, fee float64, isValue bool) (float64, float64, float64) {
	vat := restgostruct.ToFixed((vatPercent / 100 * price), 2)
	priceWithVat := price + vat

	if !isValue {
		fee = restgostruct.ToFixed((fee / 100 * priceWithVat), 2)
	}

	return restgostruct.ToFixed(priceWithVat-fee, 2), fee, vat
}

//CalculateOfflineWithoutVat is calculate offline partner without VAT 's amount
// (markup/100*price)+(vat/100*(markup/100*price))
// return result,vat
func CalculateOfflineWithoutVat(price, markup, vatPercent float64) (float64, float64) {
	vat := restgostruct.ToFixed((vatPercent / 100 * markup), 2)
	return restgostruct.ToFixed(markup+vat, 2), vat
}

//CalculateOfflineWithVat is calculate offline partner with VAT 's amount
// total=(markup/100*price)+(vat/100*(markup/100*price))
// total-(withholding/100*total)
// return result,vat,withholding
func CalculateOfflineWithVat(price, markup, vatPercent, withholdingPercent float64) (float64, float64, float64) {
	vat := restgostruct.ToFixed((vatPercent / 100 * markup), 2)
	total := restgostruct.ToFixed(markup+vat, 2)
	withholding := restgostruct.ToFixed((withholdingPercent / 100 * total), 2)
	return restgostruct.ToFixed(total-withholding, 2), vat, withholding
}
