package partner

import (
	"testing"

	restgostruct "github.com/incompetent-developer/restgo-struct"
)

func TestCalculateOnlineWithoutVat(t *testing.T) {
	result, _ := CalculateOnlineWithoutVat(1000, 2.75, false)
	if result != 972.5 {
		t.Errorf("Expect %v got %v", 972.5, result)
	}

	result, _ = CalculateOnlineWithoutVat(8000, 2.75, false)
	if result != 7780 {
		t.Errorf("Expect %v got %v", 7780, result)
	}

	result, _ = CalculateOnlineWithoutVat(1000, 12, true)
	if result != 988 {
		t.Errorf("Expect %v got %v", 988, result)
	}
}

func TestCalculateOnlineWithVat(t *testing.T) {
	result, _, _ := CalculateOnlineWithVat(1000, 7, 2.75, false)
	if result != 1040.57 {
		t.Errorf("Expect %v got %v", 1040.57, result)
	}

	result, _, _ = CalculateOnlineWithVat(8000, 7, 2.75, false)
	if result != 8324.6 {
		t.Errorf("Expect %v got %v", 8324.6, result)
	}

	result, _, _ = CalculateOnlineWithVat(1000, 7, 12, true)
	if result != 1058 {
		t.Errorf("Expect %v got %v", 1058, result)
	}

	// result, _, _ = CalculateOnlineWithVat(2000, 7, 2.75, false)
	// if result != 2081.15 {
	// 	t.Errorf("Expect %v got %v", 2081.15, result)
	// }

	// result, fee, vat := CalculateOnlineWithVat(2500, 7, 2.75, false)
	// if result != 2601.44 {
	// 	t.Errorf("Expect %v got %v %v %v", 2601.44, result, fee, vat)
	// }

	// result, _, _ = CalculateOnlineWithVat(2600, 7, 2.75, false)
	// if result != 2705.49 {
	// 	t.Errorf("Expect %v got %v", 2705.49, result)
	// }
}

func TestCalculateOfflineWithoutVat(t *testing.T) {
	price := 1000.00
	markupPercent := 17.0
	markup := restgostruct.ToFixed((markupPercent / 100 * price), 2)
	result, _ := CalculateOfflineWithoutVat(price, markup, 7)
	if result != 181.9 {
		t.Errorf("Expect %v got %v", 181.9, result)
	}
}

func TestCalculateOfflineWithVat(t *testing.T) {
	price := 1000.00
	markupPercent := 17.0
	markup := restgostruct.ToFixed((markupPercent / 100 * price), 2)
	result, _, _ := CalculateOfflineWithVat(price, markup, 7, 3)
	if result != 176.44 {
		t.Errorf("Expect %v got %v", 176.44, result)
	}
}
