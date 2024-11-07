package sprint

func ShiftBy(r rune, step int) rune {
	step_int := step % 26
	res := r
	if int(r)+step_int <= int('z') {
		res = rune(int(r) + step_int)
	} else {
		res = rune(int(r) + step_int - 26)
	}
	return res
}