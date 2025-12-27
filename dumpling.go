package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type number uint

type Data struct{
	value float64
	unit string
}

func main() {
	//Входные данные
	Q_day := Data{-1, "т"};
	t := Data{-1, "ч"};
	a_t := Data{-1, "%"};
	p_dumpM := Data{-1, "т/ч"};
	p_doughM := Data{-1, "т/ч"};
	p_cutter := Data{-1, "т/ч"};

	//Ввод данных
	reader := bufio.NewReader(os.Stdin);
	Q_day.InputDataValue(reader, "Введите суточную выработку готовой продукции (пельменей), т: ");
	t.InputDataValue(reader, "Введите продолжительность рабочей смены, ч: ");
	for{
		a_t.InputDataValue(reader, "Введите массовую долю теста в готовой продукции, %: ");
		if a_t.value > 100 {
			fmt.Println("Данное значение не может превышать 100");
			continue
		}
		break
	}
	p_dumpM.InputDataValue(reader, "Введите производительность пельменного автомата, т/ч: ");
	p_doughM.InputDataValue(reader, "Введите производительность тестомесильной машины, т/ч: ");
	p_cutter.InputDataValue(reader, "Введите производительность куттера, т/ч: ");
	

	//Результат
	P_tl, n_dumpM := n_dumpCalculate(Q_day, t, p_dumpM);
	n_doughM := n_doughCalculate(P_tl, a_t, p_doughM);
	n_cutter := n_cutterCalculate(P_tl, a_t, p_cutter);

	//Вывод
	fmt.Printf("Исходные данные\nСуточная выработка готовой продукции (пельменей): %s;\nПродолжительность рабочей смены: %s;\nМассовая доля теста в готовой продукции: %s;\nПроизводительность пельменного автомата: %s;\nПроизводительность тестомесильной машины: %s;\nПроизводительность куттера: %s;\n", formatedData(&Q_day), formatedData(&t), formatedData(&a_t), formatedData(&p_dumpM), formatedData(&p_doughM), formatedData(&p_cutter));

	fmt.Printf("Результаты расчетов\nПроизводительность технологической линии изготовления пельменей: %v\nКоличество пельменных автоматов: %v\nКоличество тестомесильных машин: %v\nКоличество куттеров: %v", P_tl, n_dumpM, n_doughM, n_cutter);
}

func (d *Data) InputDataValue(reader *bufio.Reader, printableString string) {
	for {
		fmt.Println(printableString);
		input, err := reader.ReadString('\n');
		if err != nil {
			fmt.Println("Ошибка ввода");
			continue
		}
		input = strings.TrimSpace(input);
		value, err := strconv.ParseFloat(input, 64);
		if err != nil {
			fmt.Println("Вы ввели неверное значение для этого поля");
			continue
		}
		if value <= 0 {
			fmt.Println("Значение этого поля не может быть <= 0");
			continue
		}
		d.value = value;
		break;
	}
}

func formatedData(d *Data) string {
	return fmt.Sprintf("%v %s", d.value, d.unit)
}

	//Расчеты

func n_dumpCalculate(Q_day, t, p_dumpM Data) (P_tl float64, n_dumpM float64) {
	P_tl  = Q_day.value / 2 * t.value;
	n_dumpM = math.Ceil(P_tl / p_dumpM.value);
	return	
}

func n_doughCalculate(P_tl float64, a_t, p_doughM Data) (n_doughM float64) {
	var P_tl2 float64 = (a_t.value * P_tl) / 100;
	n_doughM = math.Ceil(P_tl2 / p_doughM.value);
	return 
}

func n_cutterCalculate(P_tl float64, a_t, p_cutter Data) (n_cutter float64){
	var P_tl2 = ((100 - a_t.value) * P_tl) / 100;
	n_cutter = math.Ceil(P_tl2 / p_cutter.value);
	return
}