package utils

import "9.TaskSW/data"

func CalcNumOfRoots() {

	for _, v := range data.DB {

		//Неполное квадратное уравнение ax*x = 0, тогда один корень = 0
		if v.B == 0 && v.C == 0 {
			v.Nroots = 1
			v.AlreadyCalculated = true
			return
		}

		//неполное квадратное уравнение вида ax2 + bx = 0. В этом случае уравнение имеет 2 корня, один из которых равен нулю
		if v.A != 0 && v.B != 0 && v.C == 0 {
			v.Nroots = 2
			v.AlreadyCalculated = true
			return
		}

		//неполное квадратное уравнение вида ax2 + с = 0. Здесь, если а и с одного знака, то уравнение корней не имеет.
		//Если же а и с разных знаков, то уравнение имеет два корня.
		if v.B == 0 && v.A != 0 && v.C != 0 {
			if (v.A > 0 && v.C > 0) || (v.A < 0 && v.C < 0) {
				v.Nroots = 0
				v.AlreadyCalculated = true
			} else {
				v.Nroots = 2
				v.AlreadyCalculated = true

			}
			return
		}

		// полное квадратное уравнение ax2 + bx + с = 0
		// D = b*b – 4ac
		// Если дискриминант D > 0, то уравнение имеет ровно два корня,
		// если дискриминант D = 0, то уравнение имеет ровно один корень,
		// если дискриминант D < 0, то уравнение не имеет корней.

		D := v.B*v.B - 4*v.A*v.C

		if D > 0 {
			v.Nroots = 2
			v.AlreadyCalculated = true
		} else if D == 0 {
			v.Nroots = 1
			v.AlreadyCalculated = true
		} else {
			v.Nroots = 0
			v.AlreadyCalculated = true
		}
	}
	return
}
