package utils

import (
	"fmt"
	"math"
	"strconv"
)

func Chaos(hash string, height int64, limit int) int {
	var hash256 = [256]byte{}
	for i := 0; i < len(hash); i++ {
		base, _ := strconv.ParseUint(hash[i:i+1], 16, 8)
		b := fmt.Sprintf("%04b", base)
		copy(hash256[i*4:(i+1)*4], b[:])
	}
	
	var rand = random{height}
	var iMin float64 = 3
	var iMax float64 = 13
	n_sat := int(math.Floor(rand.seededRandom()*(iMax+1-iMin)) + iMin)
	x0_sat := [3]float64{rand.seededRandom(), rand.seededRandom(), rand.seededRandom()}
	x_re_sat := newChaosMap(x0_sat, n_sat, 1)
	
	for i := 0; i < len(x_re_sat); i++ {
		x_re_sat[i] = x_re_sat[i] * 10000
	}
	
	val_sat := sum(x_re_sat)
	flag_sat := getResVal(val_sat, 3)
	
	idx_slt := []int{1, 2, 3}
	idx_slt = append(idx_slt[:flag_sat-1], idx_slt[flag_sat:]...)
	
	mx_flag := [3]int{}
	
	idx_par := idx_slt[0]
	val_par := x_re_sat[idx_par-1]
	flag_par := getResVal(val_par, 12)
	mx_flag[0] = flag_par
	
	idx_ite := idx_slt[1]
	val_ite := x_re_sat[idx_ite-1]
	flag_ite := getResVal(val_ite, 12)
	mx_flag[1] = flag_ite
	
	flag_ini := 32 - (flag_par + flag_ite)
	mx_flag[2] = flag_ini
	
	resArr := getParmAlg(hash256, mx_flag)
	val_par_scale := resArr[0]
	val_N := int(resArr[1])
	val_ini := resArr[2]
	
	x0_mid := [3]float64{float64(val_ini), float64(val_ini), float64(val_ini)}
	x_final := newChaosMap(x0_mid, val_N, val_par_scale)
	
	x_chnl_out := x_final[flag_sat-1] * 10000
	val_out := getResVal(x_chnl_out, limit)
	
	return val_out - 1;
}

func newChaosMap(x0 [3]float64, N int, scale float64) []float64 {
	var a float64 = 1.65 * (1 + scale*0.01)
	var b float64 = 1.5
	var c float64 = -1.2
	var d float64 = 0.2
	var e float64 = 1.5
	x, y, z := make([]float64, N+1), make([]float64, N+1), make([]float64, N+1)
	x[0] = x0[0] / 10
	y[0] = x0[1] / 10
	z[0] = x0[2] / 10
	
	for n := 0; n < N; n++ {
		x[n+1] = a*y[n] + b*y[n]*y[n];
		y[n+1] = c*x[n] + d*y[n] + d*x[n]*z[n];
		z[n+1] = x[n]*x[n] + e*y[n]*x[n];
	}
	
	re := make([]float64, 3)
	re[0] = x[N-1]
	re[1] = y[N-1]
	re[2] = z[N-1]
	
	return re;
}

func sum(arr []float64) float64 { //数组求和
	var total float64
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

func mean(arr []int) int { //数组求平均值
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / len(arr)
}

func toDec(b256 [256]byte, arr [8]int) int {
	var val int
	for i := 0; i < len(arr); i++ {
		val = val * 2
		val += int(b256[arr[i]-1] - byte('0'))
	}
	return val
}

func bit2Dec(bits []byte) int64 {
	var val int64
	for i := 0; i < len(bits); i++ {
		val = val * 2
		val += int64(bits[i] - byte('0'))
	}
	return val
}

func getResVal(x float64, gap int) int {//绝对值取模+1
	y := int(math.Floor(math.Abs(x)))
	z := y%gap + 1
	return z
}

func getParmAlg(b256 [256]byte, parts [3]int) [3]float64 {
	int256 := [32][8]int{}
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			int256[i][j] = i*8 + j + 1
		}
	}
	
	idx_par := int256[0:parts[0]]
	idx_ite := int256[parts[0] : parts[0]+parts[1]]
	idx_ini := int256[parts[0]+parts[1] : parts[0]+parts[1]+parts[2]]
	
	val_par := getModelPar(b256, idx_par)
	val_ite := getModelIte(b256, idx_ite)
	val_ini := getModelIni(b256, idx_ini)
	
	return [3]float64{val_par, val_ite, val_ini}
}

func getModelPar(b256 [256]byte, par [][8]int) float64 {
	ite := getModelIte(b256, par)
	max := math.Pow(2, 8)
	nor := float64(ite) / max
	return nor
}

func getModelIte(b256 [256]byte, ite [][8]int) float64 {
	size := len(ite)
	var vals = make([]int, size)
	for i := 0; i < size; i++ {
		tmp := ite[i]
		vals[i] = toDec(b256, tmp)
	}
	val := mean(vals)
	if val == 0 {
		val = 1
	}
	return float64(val)
}

func getModelIni(b256 [256]byte, ini [][8]int) float64 {
	beg := (32 - len(ini)) * 8
	end := beg + 53
	
	if end >= 256 {
		end = 256
	}
	
	bin := b256[beg:end]
	dec := bit2Dec(bin)
	max := math.Pow(2, 53)
	nor := float64(dec) / max
	return nor
}

type random struct {
	seed int64
}

func (r *random) seededRandom(args ...float64) float64 {
	var max, min float64 = 1, 0
	if len(args) > 0 {
		max = args[0]
	}
	if len(args) > 1 {
		min = args[1]
	}
	r.seed = (r.seed*9301 + 49297) % 233280
	rnd := float64(r.seed) / 233280
	return min + (max-min)*rnd
}
