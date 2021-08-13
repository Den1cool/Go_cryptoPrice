Сервис нужен для получения информации о криптовалюте

Сервис находится по адресу http://localhost:9000/

В качестве базы данных используется PostgreSQL

Данные в базе обновляются при новом запросе, если предыдущее обновление было более 3-х минут назад

В базе данных есть 3 таблицы, таблицы RAW и DISPLAY, соответствуют структурам RAW и DISPLAY cоответственно

таблица pair содержит столбцы id,fsyms, tsyms, raw_id, display_id, updatetime

id - id записи PK

fsyms - содержит значение fsyms

tsyms - содержит значение tsyms

raw_id - содержит значение raw_id пары fsyms и tsyms

display_id - содержит значение display_id пары fsyms и tsyms

updatetime - Время последнего обновления в базе данных

для быстрого запуска можно использовать команду go build main.go а затем запустить файл main.exe

Примеры запросов

http://localhost:9000/price?fsyms=BTC&tsyms=EUR

Ответ

{
	"DISPLAY": {
		"BTC": {
			"EUR": {
				"CHANGE24HOUR": "€ -1,478.51",
				"CHANGEPCT24HOUR": "-3.78",
				"OPEN24HOUR": "€ 39,159.6",
				"VOLUME24HOUR": "Ƀ 8,862.41",
				"VOLUME24HOURTO": "€ 340,856,549.1",
				"LOW24HOUR": "€ 37,468.1",
				"HIGH24HOUR": "€ 39,547.0",
				"PRICE": "€ 37,681.0",
				"LASTUPDATE": "Just now",
				"SUPPLY": "Ƀ 18,784,200.0",
				"MKTCAP": "€ 707.81 B"
			}
		}
	},
	"RAW": {
		"BTC": {
			"EUR": {
				"CHANGE24HOUR": -1478.510000000002,
				"CHANGEPCT24HOUR": -3.775605184431389,
				"OPEN24HOUR": 39159.55,
				"VOLUME24HOUR": 8862.408959470004,
				"VOLUME24HOURTO": 340856549.0699038,
				"LOW24HOUR": 37468.05,
				"HIGH24HOUR": 39547.01,
				"PRICE": 37681.04,
				"LASTUPDATE": 1628807941,
				"SUPPLY": 18784200,
				"MKTCAP": 707808191568
			}
		}
	}
}









{
	"DISPLAY": {
		"BTC,LINK": {
			"EUR,USD": {
				"CHANGE24HOUR": "€ 148.71",
				"CHANGEPCT24HOUR": "0.39",
				"OPEN24HOUR": "€ 38,502.3",
				"VOLUME24HOUR": "Ƀ 7,820.72",
				"VOLUME24HOURTO": "€ 299,367,135.7",
				"LOW24HOUR": "€ 37,311.8",
				"HIGH24HOUR": "€ 38,912.4",
				"PRICE": "€ 38,651.0",
				"LASTUPDATE": "Just now",
				"SUPPLY": "Ƀ 18,784,493.0",
				"MKTCAP": "€ 726.04 B"
			}
		}
	},
	"RAW": {
		"BTC": {
			"EUR": {
				"CHANGE24HOUR": 148.70999999999913,
				"CHANGEPCT24HOUR": 0.38623636543554407,
				"OPEN24HOUR": 38502.33,
				"VOLUME24HOUR": 7820.71832548,
				"VOLUME24HOURTO": 299367135.67421323,
				"LOW24HOUR": 37311.77,
				"HIGH24HOUR": 38912.4,
				"PRICE": 38651.04,
				"LASTUPDATE": 1628833655,
				"SUPPLY": 18784493,
				"MKTCAP": 726040190322.72
			}
		}
	}
}{
	"DISPLAY": {
		"BTC,LINK": {
			"EUR,USD": {
				"CHANGE24HOUR": "$ 0.32",
				"CHANGEPCT24HOUR": "1.26",
				"OPEN24HOUR": "$ 25.37",
				"VOLUME24HOUR": "LINK 4,634,659.9",
				"VOLUME24HOURTO": "$ 115,544,721.4",
				"LOW24HOUR": "$ 24.09",
				"HIGH24HOUR": "$ 25.87",
				"PRICE": "$ 25.69",
				"LASTUPDATE": "Just now",
				"SUPPLY": "LINK 1,000,000,000.0",
				"MKTCAP": "$ 25.69 B"
			}
		}
	},
	"RAW": {
		"LINK": {
			"USD": {
				"CHANGE24HOUR": 0.3200000000000003,
				"CHANGEPCT24HOUR": 1.2613322822230992,
				"OPEN24HOUR": 25.37,
				"VOLUME24HOUR": 4634659.856140221,
				"VOLUME24HOURTO": 115544721.43927631,
				"LOW24HOUR": 24.09,
				"HIGH24HOUR": 25.87,
				"PRICE": 25.69,
				"LASTUPDATE": 1628833664,
				"SUPPLY": 1000000000,
				"MKTCAP": 25690000000
			}
		}
	}
}