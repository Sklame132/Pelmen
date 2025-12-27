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


	//Вывод
	fmt.Printf("Исходные данные\nСуточная выработка готовой продукции (пельменей): %s;\nПродолжительность рабочей смены: %s;\nМассовая доля теста в готовой продукции: %s;\nПроизводительность пельменного автомата: %s;\nПроизводительность тестомесильной машины: %s;\nПроизводительность куттера: %s;\n", formatedData(&Q_day), formatedData(&t), formatedData(&a_t), formatedData(&p_dumpM), formatedData(&p_doughM), formatedData(&p_cutter));

	
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


