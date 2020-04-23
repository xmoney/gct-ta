package indicators

import (
	"testing"
)

func TestBBands(t *testing.T) {
	expectedOutput := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8137.972061598561, 8206.497688069518, 8307.553011967851, 8380.302366610505, 8607.924596039536, 8790.639794903474, 8922.97584775253, 9074.553649922997, 9209.262648069798, 9276.787021065087, 9311.474709090533, 9340.636632407995, 9298.968962574285, 9262.465609190078, 9207.787048681046, 9120.056492273194, 9098.901859827967, 9148.152796874698, 9293.138560116806, 9348.174458357891, 9477.676866692716, 9523.689388419927, 9579.319782730048, 9595.34066418091, 9639.524703006888, 9663.544491167446, 9760.491432987159, 9872.963008218368, 9982.006418886825, 10095.836842762104, 10247.357524575194, 10316.808043571815, 10447.759471158159, 10539.708535283153, 10579.610902542756, 10590.58713155943, 10537.489213822213, 10510.244560504765, 10509.619778329386, 10544.884611467787, 10543.117759495839, 10530.571291094613, 10520.060804806923, 10502.099131149433, 10483.544634486074, 10420.759146362172, 10456.471552645977, 10587.978598073507, 10663.882789407226, 10723.66286621575, 10750.143784993797, 10783.5642176388, 10707.117903810558, 10608.0568183109, 10512.029459146883, 10357.327523222557, 10285.032097186097, 10200.747186874085, 10216.056530542268, 10109.251348926653, 10094.95408731182, 10044.360451697066, 10670.367688032984, 10781.194557256686, 10814.776197706524, 10774.95640440927, 10757.75890127151, 10718.399519385186, 10625.869414899349, 10481.923776724903, 10342.712893206946, 10186.228058697918, 9949.199740891447, 9715.38208712414, 9471.900407784302, 9128.704383312024, 8720.326507994307, 8295.899609658098, 8041.685398904088, 7762.110271930646, 7450.378620957321}
	expectedOutput2 := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7446.008, 7471.334, 7514.563499999999, 7557.238, 7638.1005000000005, 7719.144, 7792.3515, 7872.170999999999, 7948.3205, 8022.138500000001, 8094.978999999999, 8171.939500000002, 8257.207999999999, 8309.1845, 8363.067000000001, 8411.707, 8453.094000000001, 8489.8705, 8557.648500000001, 8631.768500000002, 8696.768500000002, 8762.35, 8822.1705, 8883.3895, 8907.142, 8924.652, 8970.0655, 9013.1655, 9057.91, 9118.18, 9195.62, 9252.226, 9332.878500000003, 9431.3955, 9521.7225, 9623.4495, 9688.914500000003, 9739.832, 9754.6895, 9799.489000000001, 9804.275000000001, 9818.5205, 9834.4685, 9851.4055, 9885.2615, 9909.922, 9894.270999999999, 9845.507499999998, 9796.161499999998, 9736.253499999999, 9653.941999999997, 9587.889, 9520.307, 9440.6405, 9366.871000000001, 9302.0015, 9264.702000000001, 9213.753999999999, 9130.5815, 9018.1355, 8932.645500000002, 8848.736, 8606.013000000003, 8403.7435, 8163.951499999998, 7948.5045, 7734.905500000001, 7562.187, 7391.708000000001, 7265.575, 7149.6950000000015, 7032.6365000000005, 6877.427000000001, 6764.329999999999, 6665.098, 6546.376499999999, 6426.723, 6300.8279999999995, 6211.796999999999, 6108.529499999999, 6011.583999999999}
	expectedOutput3 := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6754.043938401439, 6736.170311930483, 6721.573988032147, 6734.173633389495, 6668.276403960464, 6647.648205096527, 6661.727152247469, 6669.788350077002, 6687.3783519302015, 6767.489978934915, 6878.483290909466, 7003.242367592009, 7215.447037425713, 7355.903390809921, 7518.346951318955, 7703.357507726808, 7807.286140172034, 7831.588203125304, 7822.1584398831965, 7915.362541642112, 7915.860133307288, 8001.010611580073, 8065.021217269951, 8171.438335819089, 8174.759296993112, 8185.759508832553, 8179.639567012843, 8153.36799178163, 8133.813581113174, 8140.523157237896, 8143.882475424807, 8187.643956428186, 8217.997528841846, 8323.082464716848, 8463.834097457244, 8656.31186844057, 8840.339786177792, 8969.419439495236, 8999.759221670614, 9054.093388532216, 9065.432240504164, 9106.469708905388, 9148.87619519308, 9200.711868850569, 9286.978365513927, 9399.084853637829, 9332.07044735402, 9103.036401926489, 8928.440210592771, 8748.844133784249, 8557.740215006197, 8392.2137823612, 8333.496096189443, 8273.224181689098, 8221.71254085312, 8246.675476777444, 8244.371902813906, 8226.760813125913, 8045.1064694577335, 7927.019651073348, 7770.336912688184, 7653.111548302936, 6541.65831196702, 6026.292442743315, 5513.126802293473, 5122.052595590731, 4712.052098728491, 4405.974480614814, 4157.546585100655, 4049.2262232750954, 3956.6771067930563, 3879.0449413020833, 3805.6542591085536, 3813.2779128758584, 3858.2955922156984, 3964.048616687975, 4133.119492005693, 4305.756390341901, 4381.908601095909, 4454.948728069351, 4572.789379042677}

	ret, ret2, ret3 := BBANDS(testClose, 20, 2.0, 2.0, Sma)
	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
		if ret2[x] != expectedOutput2[x] {
			t.Fatalf("unexpected value returned %v", ret2[x])
		}
		if ret3[x] != expectedOutput3[x] {
			t.Fatalf("unexpected value returned %v", ret3[x])
		}
	}
}