/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 03.12.2020  10:40
   Notes   :
*/
package main

import (
	"encoding/json"
	"fmt"
	"strings"
  "time"
)

type TextValue struct {
	Text  string `json:"Text"`
	Value string `json:"Value"`
}


type City struct {
	Country string `gorm:"primary_key" json:"country"`
	// City Name
	Name string `gorm:"primary_key" json:"name"`
	// Created At
	CreatedAt time.Time `json:"created_at"`
	// Updated At
	UpdatedAt time.Time `json:"updated_at"`
}


func main() {

	var cities = []models.City{

	}
	cities = getCities("33", "ABD", cities)
	cities = getCities("166", "AFGANISTAN", cities)
	cities = getCities("13", "ALMANYA", cities)
	cities = getCities("17", "ANDORRA", cities)
	cities = getCities("140", "ANGOLA", cities)
	cities = getCities("125", "ANGUILLA", cities)
	cities = getCities("90", "ANTIGUA VE BARBUDA", cities)
	cities = getCities("199", "ARJANTIN", cities)
	cities = getCities("25", "ARNAVUTLUK", cities)
	cities = getCities("153", "ARUBA", cities)
	cities = getCities("59", "AVUSTRALYA", cities)
	cities = getCities("35", "AVUSTURYA", cities)
	cities = getCities("5", "AZERBAYCAN", cities)
	cities = getCities("54", "BAHAMALAR", cities)
	cities = getCities("132", "BAHREYN", cities)
	cities = getCities("177", "BANGLADES", cities)
	cities = getCities("188", "BARBADOS", cities)
	cities = getCities("208", "BELARUS", cities)
	cities = getCities("11", "BELCIKA", cities)
	cities = getCities("182", "BELIZE", cities)
	cities = getCities("181", "BENIN", cities)
	cities = getCities("51", "BERMUDA", cities)
	cities = getCities("93", "BIRLESIK ARAP EMIRLIGI", cities)
	cities = getCities("83", "BOLIVYA", cities)
	cities = getCities("9", "BOSNA HERSEK", cities)
	cities = getCities("167", "BOTSVANA", cities)
	cities = getCities("146", "BREZILYA", cities)
	cities = getCities("97", "BRUNEI", cities)
	cities = getCities("44", "BULGARISTAN", cities)
	cities = getCities("91", "BURKINA FASO", cities)
	cities = getCities("154", "BURMA (MYANMAR)", cities)
	cities = getCities("65", "BURUNDI", cities)
	cities = getCities("155", "BUTAN", cities)
	cities = getCities("156", "CAD", cities)
	cities = getCities("43", "CECENISTAN", cities)
	cities = getCities("16", "CEK CUMHURIYETI", cities)
	cities = getCities("86", "CEZAYIR", cities)
	cities = getCities("160", "CIBUTI", cities)
	cities = getCities("61", "CIN", cities)
	cities = getCities("26", "DANIMARKA", cities)
	cities = getCities("180", "DEMOKRATIK KONGO CUMHURIYETI", cities)
	cities = getCities("176", "DOGU TIMOR", cities)
	cities = getCities("123", "DOMINIK", cities)
	cities = getCities("72", "DOMINIK CUMHURIYETI", cities)
	cities = getCities("139", "EKVATOR", cities)
	cities = getCities("63", "EKVATOR GINESI", cities)
	cities = getCities("165", "EL SALVADOR", cities)
	cities = getCities("117", "ENDONEZYA", cities)
	cities = getCities("175", "ERITRE", cities)
	cities = getCities("104", "ERMENISTAN", cities)
	cities = getCities("6", "ESTONYA", cities)
	cities = getCities("95", "ETYOPYA", cities)
	cities = getCities("145", "FAS", cities)
	cities = getCities("197", "FIJI", cities)
	cities = getCities("120", "FILDISI SAHILI", cities)
	cities = getCities("126", "FILIPINLER", cities)
	cities = getCities("204", "FILISTIN", cities)
	cities = getCities("41", "FINLANDIYA", cities)
	cities = getCities("21", "FRANSA", cities)
	cities = getCities("79", "GABON", cities)
	cities = getCities("109", "GAMBIYA", cities)
	cities = getCities("143", "GANA", cities)
	cities = getCities("111", "GINE", cities)
	cities = getCities("58", "GRANADA", cities)
	cities = getCities("48", "GRONLAND", cities)
	cities = getCities("171", "GUADELOPE", cities)
	cities = getCities("169", "GUAM ADASI", cities)
	cities = getCities("99", "GUATEMALA", cities)
	cities = getCities("67", "GUNEY AFRIKA", cities)
	cities = getCities("128", "GUNEY KORE", cities)
	cities = getCities("62", "GURCISTAN", cities)
	cities = getCities("82", "GUYANA", cities)
	cities = getCities("70", "HAITI", cities)
	cities = getCities("187", "HINDISTAN", cities)
	cities = getCities("30", "HIRVATISTAN", cities)
	cities = getCities("4", "HOLLANDA", cities)
	cities = getCities("66", "HOLLANDA ANTILLERI", cities)
	cities = getCities("105", "HONDURAS", cities)
	cities = getCities("113", "HONG KONG", cities)
	cities = getCities("15", "INGILTERE", cities)
	cities = getCities("124", "IRAK", cities)
	cities = getCities("202", "IRAN", cities)
	cities = getCities("32", "IRLANDA", cities)
	cities = getCities("23", "ISPANYA", cities)
	cities = getCities("205", "ISRAIL", cities)
	cities = getCities("12", "ISVEC", cities)
	cities = getCities("49", "ISVICRE", cities)
	cities = getCities("8", "ITALYA", cities)
	cities = getCities("122", "IZLANDA", cities)
	cities = getCities("119", "JAMAIKA", cities)
	cities = getCities("116", "JAPONYA", cities)
	cities = getCities("161", "KAMBOCYA", cities)
	cities = getCities("184", "KAMERUN", cities)
	cities = getCities("52", "KANADA", cities)
	cities = getCities("34", "KARADAG", cities)
	cities = getCities("94", "KATAR", cities)
	cities = getCities("92", "KAZAKISTAN", cities)
	cities = getCities("114", "KENYA", cities)
	cities = getCities("168", "KIRGIZISTAN", cities)
	cities = getCities("57", "KOLOMBIYA", cities)
	cities = getCities("88", "KOMORLAR", cities)
	cities = getCities("18", "KOSOVA", cities)
	cities = getCities("162", "KOSTARIKA", cities)
	cities = getCities("209", "KUBA", cities)
	cities = getCities("206", "KUDUS", cities)
	cities = getCities("133", "KUVEYT", cities)
	cities = getCities("1", "KUZEY KIBRIS", cities)
	cities = getCities("142", "KUZEY KORE", cities)
	cities = getCities("134", "LAOS", cities)
	cities = getCities("174", "LESOTO", cities)
	cities = getCities("20", "LETONYA", cities)
	cities = getCities("73", "LIBERYA", cities)
	cities = getCities("203", "LIBYA", cities)
	cities = getCities("38", "LIECHTENSTEIN", cities)
	cities = getCities("47", "LITVANYA", cities)
	cities = getCities("42", "LUBNAN", cities)
	cities = getCities("31", "LUKSEMBURG", cities)
	cities = getCities("7", "MACARISTAN", cities)
	cities = getCities("98", "MADAGASKAR", cities)
	cities = getCities("100", "MAKAO", cities)
	cities = getCities("28", "MAKEDONYA", cities)
	cities = getCities("55", "MALAVI", cities)
	cities = getCities("103", "MALDIVLER", cities)
	cities = getCities("107", "MALEZYA", cities)
	cities = getCities("152", "MALI", cities)
	cities = getCities("24", "MALTA", cities)
	cities = getCities("87", "MARTINIK", cities)
	cities = getCities("164", "MAURITIUS ADASI", cities)
	cities = getCities("157", "MAYOTTE", cities)
	cities = getCities("53", "MEKSIKA", cities)
	cities = getCities("85", "MIKRONEZYA", cities)
	cities = getCities("189", "MISIR", cities)
	cities = getCities("60", "MOGOLISTAN", cities)
	cities = getCities("46", "MOLDAVYA", cities)
	cities = getCities("3", "MONAKO", cities)
	cities = getCities("147", "MONTSERRAT (U.K.)", cities)
	cities = getCities("106", "MORITANYA", cities)
	cities = getCities("151", "MOZAMBIK", cities)
	cities = getCities("196", "NAMBIYA", cities)
	cities = getCities("76", "NEPAL", cities)
	cities = getCities("84", "NIJER", cities)
	cities = getCities("127", "NIJERYA", cities)
	cities = getCities("141", "NIKARAGUA", cities)
	cities = getCities("178", "NIUE", cities)
	cities = getCities("36", "NORVEC", cities)
	cities = getCities("80", "ORTA AFRIKA CUMHURIYETI", cities)
	cities = getCities("131", "OZBEKISTAN", cities)
	cities = getCities("77", "PAKISTAN", cities)
	cities = getCities("149", "PALAU", cities)
	cities = getCities("89", "PANAMA", cities)
	cities = getCities("185", "PAPUA YENI GINE", cities)
	cities = getCities("194", "PARAGUAY", cities)
	cities = getCities("69", "PERU", cities)
	cities = getCities("183", "PITCAIRN ADASI", cities)
	cities = getCities("39", "POLONYA", cities)
	cities = getCities("45", "PORTEKIZ", cities)
	cities = getCities("68", "PORTO RIKO", cities)
	cities = getCities("112", "REUNION", cities)
	cities = getCities("37", "ROMANYA", cities)
	cities = getCities("81", "RUANDA", cities)
	cities = getCities("207", "RUSYA", cities)
	cities = getCities("64", "S. ARABISTAN", cities)
	cities = getCities("198", "SAMOA", cities)
	cities = getCities("102", "SENEGAL", cities)
	cities = getCities("138", "SEYSEL ADALARI", cities)
	cities = getCities("200", "SILI", cities)
	cities = getCities("179", "SINGAPUR", cities)
	cities = getCities("27", "SIRBISTAN", cities)
	cities = getCities("14", "SLOVAKYA", cities)
	cities = getCities("19", "SLOVENYA", cities)
	cities = getCities("150", "SOMALI", cities)
	cities = getCities("74", "SRI LANKA", cities)
	cities = getCities("129", "SUDAN", cities)
	cities = getCities("172", "SURINAM", cities)
	cities = getCities("191", "SURIYE", cities)
	cities = getCities("163", "SVALBARD", cities)
	cities = getCities("170", "SVAZILAND", cities)
	cities = getCities("101", "TACIKISTAN", cities)
	cities = getCities("110", "TANZANYA", cities)
	cities = getCities("137", "TAYLAND", cities)
	cities = getCities("108", "TAYVAN", cities)
	cities = getCities("71", "TOGO", cities)
	cities = getCities("130", "TONGA", cities)
	cities = getCities("96", "TRINIDAT VE TOBAGO", cities)
	cities = getCities("118", "TUNUS", cities)
	cities = getCities("159", "TURKMENISTAN", cities)
	cities = getCities("2", "TÜRKİYE", cities)
	cities = getCities("75", "UGANDA", cities)
	cities = getCities("40", "UKRAYNA", cities)
	cities = getCities("29", "UKRAYNA-KIRIM", cities)
	cities = getCities("173", "UMMAN", cities)
	cities = getCities("192", "URDUN", cities)
	cities = getCities("201", "URUGUAY", cities)
	cities = getCities("56", "VANUATU", cities)
	cities = getCities("10", "VATIKAN", cities)
	cities = getCities("186", "VENEZUELA", cities)
	cities = getCities("135", "VIETNAM", cities)
	cities = getCities("148", "YEMEN", cities)
	cities = getCities("115", "YENI KALEDONYA", cities)
	cities = getCities("193", "YENI ZELLANDA", cities)
	cities = getCities("144", "YESIL BURUN", cities)
	cities = getCities("22", "YUNANISTAN", cities)
	cities = getCities("158", "ZAMBIYA", cities)
	cities = getCities("136", "ZIMBABVE", cities)

	err := config.DB.Create(cities).Error
	if err != nil {
		fmt.Println(err.Error())
	}

}

func getCities(countryCode string, countryName string, cities []models.City) []models.City {
	result := []TextValue{}

	res, err := config.RClient.R().
		Get("http://namaz.diyanet.gov.tr/PrayerTime/FillMainPageState?countryCode=" + countryCode)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		for _, city := range result {
			cities = append(cities, models.City{
				Country: strings.Title(strings.ToLower(countryName)),
				Name:    strings.Title(strings.ToLower(city.Text)),
			})
		}
	}
	return cities
}
