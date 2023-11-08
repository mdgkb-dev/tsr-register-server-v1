// You can edit this code!
// Click here and start typing.
package validators

import (
	"errors"
	"strconv"
	"strings"
)

func SnilsCheckControlSum(snils string) error {
	//убираем пробелы и дефисы
	snils = strings.Replace(snils, "-", "", -1)
	snils = strings.Replace(snils, " ", "", -1)
	//проверка на длину
	if len(snils) != 11 {
		return errors.New("снилс неправильной длины")
	}

	//отделяем от снилса контрольную сумму
	checksum := snils[9:]
	//fmt.Println(checksum)
	snils = snils[0:9]
	//fmt.Println(snils)

	//проверяем подходит ли снилс для проверки
	q, err := strconv.Atoi(snils)
	//fmt.Println(q)
	if q <= 1001998 {
		return errors.New("снилс не подходит для проверки")
	}
	if err != nil {
		return err
	}
	//fmt.Println(q, err)

	//проверка на три одинаковые цифры подряд
	cnt := len(snils)
	//fmt.Println(cnt)
	for i := 2; i < cnt; i++ {
		if snils[i-2] == snils[i-1] && snils[i-1] == snils[i] {
			return errors.New("в снилс три одинаковые цифры подряд")
		}
	}

	//подсчет контрольной суммы
	sum := 0

	for i := 0; i < cnt; i++ {
		j := int(snils[i]) - 48 // берем по цифре из снилса
		k := cnt - i            //позиция с конца
		//fmt.Println(j, k)
		sum += j * k
		//fmt.Println(reflect.TypeOf(snils[i]))
	}
	//fmt.Println(sum)

	//доп проверки
	ch := ""
	if sum > 99 {
		tmp := sum / 101
		sum = sum - tmp*101
		//fmt.Println(sum)

		if sum == 100 || sum == 101 {
			ch = "00"
		} else {
			ch = strconv.Itoa(sum)
		}
	} else {
		ch = strconv.Itoa(sum)
	}
	//fmt.Println(ch)

	if ch == checksum {
		//fmt.Printf("СНИЛС корректный! Подсчитанная сумма равна %s \n", ch)
		return nil
	} else {
		return errors.New("Ошибка! СНИЛС некорректный! Подсчитанная сумма равна " + ch + ", контрольная сумма равна " + checksum)
	}
}
