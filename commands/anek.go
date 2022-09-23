package commands

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type M map[string]interface{}
type anek struct {
	// Content string `json:"content"`
	pid     string
	method  string
	format  string
	charset string
	uts     int64
	markup  int
	note    int
	wlist   int
	censor  int
	genre   int
}

func Make_request(method string) string {
	api_url := "http://anecdotica.ru/api"
	// api_params := anek{
	// 	pid:     "k2rp6o52g1wd17ly432o",
	// 	method:  method,
	// 	charset: "cp1251",
	// 	uts:     time.Now().Unix(),
	// 	markup:  1,
	// 	note:    0,
	// 	wlist:   0,
	// 	censor:  0,
	// 	genre:   1,
	// }
	// M{"pid": "k2rp6o52g1wd17ly432o",
	// 	"method":  method,
	// 	"format":  "json",
	// 	"charset": "cp1251",
	// 	"uts":     time.Now().UTC(),
	// 	"markup":  1,
	// 	"note":    0,
	// 	"wlist":   0,
	// 	"censor":  0,
	// 	"genre":   1,
	// }
	skey := "e741bcd7a1b58396d2dcf115088a72c3c8d4b32940204b5db61b7b10ee1f4f8c"

	q := url.Values{}
	q.Set("pid", "k2rp6o52g1wd17ly432o")
	q.Set("method", method)
	q.Set("uts", fmt.Sprint(time.Now().Unix()))
	// q.Set("category", "json")
	q.Set("genre", "1")
	q.Set("lang", "1")
	q.Set("format", "txt")
	q.Set("charset", "utf-8")
	q.Set("markup", "1")
	// q.Set("note", "0")
	// q.Set("wlist", "0")
	// q.Set("censor", "0")
	q.Set("hash", GetMD5Hash(q.Encode()+skey))

	final_url := api_url + "?" + q.Encode()
	resp, err := http.Get(final_url)
	if err != nil {
		fmt.Println("error getting anek", err)
		return ""
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "Не удалось получить анек. Сервис поломался("
		} else {
			return string(body)
		}
	}
	// '?'.$url_params.'&hash='.$signature;
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// # Функция GET-запроса к API
// function make_request($method, $params) {
//    # $method - имя метода API
//    # $params - массив с параметрами запроса

//    $api_url='http://anecdotica.ru/api';            // URL API
//    # предустановленные параметры запроса
//    $api_params=['pid' => '••••••••••••••••••••',   // имя профиля
//                 'method' => $method,               // имя метода API
//                 'format' => 'xml',                 // формат ответа (json, xml, html, txt)
//                 'charset' => 'cp1251',             // кодировка (utf-8, cp1251, koi8-r)
//                 'uts' => time(),                   // метка времени
//                 'markup' => 1,                     // html-разметка (0 - откл., 1 - вкл.)
//                 'note' => 0,                       // примечания (0 - откл., 1 - вкл.)
//                 'wlist' => 0,                      // флаг «white_list» (0 - откл., 1 - вкл.)
//                 'censor' => 1,                     // цензура (0 - откл., 1 - вкл.)
//                 'genre'=>2                         // код жанра (2 - анекдоты-загадки)
//                ];

//    $skey = '••••••••••••••••••••••••••••••••';	   // секретный ключ
//    $prms = array_replace($api_params, $params);    // переопределение предустановленных параметров

//    $url_params = http_build_query($prms);          // формируем строку параметров
//    $signature = md5($url_params.$skey);            // получаем хеш строки параметров
//    $url. = '?'.$url_params.'&hash='.$signature;    // формируем запрос, дописывая хеш к строке параметров

//    return file_get_contents($url);                 // отправляем GET-запрос
// }

// # Пример вызова функции запроса к API (для тега «автомашина» в жанре «анекдот»)
// $result=make_request('getRandItemP', ['tag'=>110, 'genre'=>1]);
