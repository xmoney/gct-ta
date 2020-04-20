package indicators

import (
	"testing"
)

// nolint dupl false positive
func TestSMA(t *testing.T) {
	expectedOutput := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7256.463571428571, 7245.534999999999, 7277.432857142857, 7342.2535714285705, 7402.871428571429, 7447.0485714285705, 7515.191428571428, 7566.61857142857, 7623.419999999999, 7686.617857142856, 7804.094999999999, 7920.558571428571, 8046.24357142857, 8157.223571428571, 8268.941428571428, 8364.648571428572, 8426.250714285716, 8466.170714285716, 8509.848571428573, 8550.68142857143, 8566.547857142858, 8588.925000000001, 8618.190714285714, 8674.585714285713, 8716.511428571428, 8750.797857142856, 8807.215714285714, 8838.302857142857, 8871.777142857141, 8916.976428571428, 8964.327857142856, 8995.862142857142, 9064.611428571428, 9162.910714285714, 9261.574285714285, 9373.895714285713, 9487.002142857144, 9555.25, 9617.320000000002, 9693.005714285715, 9745.05357142857, 9819.06714285714, 9856.396428571428, 9898.074999999999, 9927.29357142857, 10000.565, 9998.94357142857, 9988.81, 9981.146428571428, 9964.441428571428, 9949.529999999999, 9935.427857142857, 9866.600714285716, 9754.499285714284, 9653.630000000001, 9535.044285714286, 9437.041428571427, 9338.18857142857, 9282.487857142856, 9180.365, 9120.190714285714, 9081.182857142856, 9042.444285714286, 8987.144285714285, 8849.179285714283, 8726.584285714283, 8625.192857142856, 8564.850714285712, 8280.934285714286, 8060.855714285715, 7820.972142857144, 7593.45, 7316.027142857143, 7070.994285714286, 6831.9685714285715, 6625.863571428571, 6415.602857142859, 6222.306428571428, 6063.695714285713, 5960.642857142856, 5880.837142857142, 5792.097857142857, 5929.172857142857, 5982.979285714285, 6060.387142857143, 6098.0485714285705, 6163.304999999999}

	ret := SMA(testClose, 14)
	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}

	ret = SMA([]float64{10, 5.0}, 15)
	if ret[0] != 0 {
		t.Fatalf("unexpected value returned %v", ret[0])
	}
}

// nolint dupl false positive
func TestEMA(t *testing.T) {
	expectedOutput := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7256.463571428571, 7269.828428571429, 7335.549971428572, 7445.344641904762, 7525.224022984127, 7562.764819919577, 7647.729510596967, 7696.861575850705, 7761.3813657372775, 7807.1985169723075, 7941.310714709333, 8057.025286081423, 8144.108581270566, 8243.911437101156, 8332.527245487669, 8381.070279422645, 8413.616908832959, 8454.186654321898, 8481.08576707898, 8467.399664801782, 8461.363709494877, 8443.781881562227, 8463.341630687262, 8520.834746595627, 8638.056780382876, 8724.87320966516, 8828.223448376471, 8894.840321926275, 8959.150945669438, 9008.516152913513, 9045.792665858378, 9060.936310410594, 9135.267469022516, 9217.849139819513, 9295.925254510244, 9376.735220575545, 9483.033191165472, 9531.96343234341, 9630.232308030954, 9726.038666960161, 9793.05617803214, 9869.187354294521, 9873.266373721919, 9878.54952389233, 9854.362920706686, 9898.470531279128, 9858.103793775244, 9825.391287938544, 9808.15644954674, 9789.48625627384, 9812.91608877066, 9791.537943601239, 9726.719551121074, 9600.404944304932, 9495.884285064274, 9390.278380389038, 9275.273263003834, 9175.763494603323, 9141.307028656214, 9089.859424835386, 9044.982168190667, 9047.871212431912, 9061.99371744099, 9039.58522178219, 8905.467192211232, 8776.238233249735, 8657.529135483102, 8561.134584085356, 8065.205972873975, 7739.543843157445, 7396.853997403119, 7123.092131082703, 6844.469180271676, 6641.871956235453, 6477.211695404059, 6437.740802683518, 6407.394028992382, 6378.112158460064, 6302.783870665389, 6328.214021243338, 6386.937485077559, 6427.641820400551, 6472.052244347145, 6459.249278434192, 6431.760041309633, 6356.978702468349, 6302.3162088059025}

	ret := EMA(testClose, 14)
	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}

	ret = EMA([]float64{10, 5.0}, 15)
	if ret[0] != 0 {
		t.Fatalf("unexpected value returned %v", ret[0])
	}

	ret = MA(testClose, 14, Ema)
	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}
}

func TestMA(t *testing.T) {
	ret := MA([]float64{10, 0}, 1, Sma)
	if ret[0] != 10 && ret[1] != 0 {
		t.Fatalf("unexpected values returned %v | %v", ret[0], ret[1])
	}
}