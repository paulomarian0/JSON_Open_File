package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cidades struct {
	Afghanistan              []string `json:"Afghanistan"`
	Albania                  []string `json:"Albania"`
	Algeria                  []string `json:"Algeria"`
	Andorra                  []string `json:"Andorra"`
	Angola                   []string `json:"Angola"`
	AntiguaAndBarbuda        []string `json:"Antigua and Barbuda"`
	Argentina                []string `json:"Argentina"`
	Armenia                  []string `json:"Armenia"`
	Aruba                    []string `json:"Aruba"`
	Australia                []string `json:"Australia"`
	Austria                  []string `json:"Austria"`
	Azerbaijan               []string `json:"Azerbaijan"`
	Bahamas                  []string `json:"Bahamas"`
	Bahrain                  []string `json:"Bahrain"`
	Bangladesh               []string `json:"Bangladesh"`
	Barbados                 []string `json:"Barbados"`
	Belarus                  []string `json:"Belarus"`
	Belgium                  []string `json:"Belgium"`
	Belize                   []string `json:"Belize"`
	Bolivia                  []string `json:"Bolivia"`
	BosniaAndHerzegovina     []string `json:"Bosnia and Herzegovina"`
	Botswana                 []string `json:"Botswana"`
	Brazil                   []string `json:"Brazil"`
	Brunei                   []string `json:"Brunei"`
	Bulgaria                 []string `json:"Bulgaria"`
	Cambodia                 []string `json:"Cambodia"`
	Cameroon                 []string `json:"Cameroon"`
	Canada                   []string `json:"Canada"`
	CaymanIslands            []string `json:"Cayman Islands"`
	Chile                    []string `json:"Chile"`
	China                    []string `json:"China"`
	Colombia                 []string `json:"Colombia"`
	Congo                    []string `json:"Congo"`
	CostaRica                []string `json:"Costa Rica"`
	Croatia                  []string `json:"Croatia"`
	Cuba                     []string `json:"Cuba"`
	Cyprus                   []string `json:"Cyprus"`
	CzechRepublic            []string `json:"Czech Republic"`
	Denmark                  []string `json:"Denmark"`
	DominicanRepublic        []string `json:"Dominican Republic"`
	Ecuador                  []string `json:"Ecuador"`
	Egypt                    []string `json:"Egypt"`
	ElSalvador               []string `json:"El Salvador"`
	Estonia                  []string `json:"Estonia"`
	FaroeIslands             []string `json:"Faroe Islands"`
	Finland                  []string `json:"Finland"`
	France                   []string `json:"France"`
	FrenchPolynesia          []string `json:"French Polynesia"`
	Gabon                    []string `json:"Gabon"`
	Georgia                  []string `json:"Georgia"`
	Germany                  []string `json:"Germany"`
	Ghana                    []string `json:"Ghana"`
	Greece                   []string `json:"Greece"`
	Greenland                []string `json:"Greenland"`
	Guadeloupe               []string `json:"Guadeloupe"`
	Guam                     []string `json:"Guam"`
	Guatemala                []string `json:"Guatemala"`
	Guinea                   []string `json:"Guinea"`
	Haiti                    []string `json:"Haiti"`
	HashemiteKingdomOfJordan []string `json:"Hashemite Kingdom of Jordan"`
	Honduras                 []string `json:"Honduras"`
	HongKong                 []string `json:"Hong Kong"`
	Hungary                  []string `json:"Hungary"`
	Iceland                  []string `json:"Iceland"`
	India                    []string `json:"India"`
	Indonesia                []string `json:"Indonesia"`
	Iran                     []string `json:"Iran"`
	Iraq                     []string `json:"Iraq"`
	Ireland                  []string `json:"Ireland"`
	IsleOfMan                []string `json:"Isle of Man"`
	Israel                   []string `json:"Israel"`
	Italy                    []string `json:"Italy"`
	Jamaica                  []string `json:"Jamaica"`
	Japan                    []string `json:"Japan"`
	Kazakhstan               []string `json:"Kazakhstan"`
	Kenya                    []string `json:"Kenya"`
	Kosovo                   []string `json:"Kosovo"`
	Kuwait                   []string `json:"Kuwait"`
	Latvia                   []string `json:"Latvia"`
	Lebanon                  []string `json:"Lebanon"`
	Libya                    []string `json:"Libya"`
	Liechtenstein            []string `json:"Liechtenstein"`
	Luxembourg               []string `json:"Luxembourg"`
	Macedonia                []string `json:"Macedonia"`
	Madagascar               []string `json:"Madagascar"`
	Malaysia                 []string `json:"Malaysia"`
	Malta                    []string `json:"Malta"`
	Martinique               []string `json:"Martinique"`
	Mauritius                []string `json:"Mauritius"`
	Mayotte                  []string `json:"Mayotte"`
	Mexico                   []string `json:"Mexico"`
	Mongolia                 []string `json:"Mongolia"`
	Montenegro               []string `json:"Montenegro"`
	Morocco                  []string `json:"Morocco"`
	Mozambique               []string `json:"Mozambique"`
	MyanmarBurma             []string `json:"Myanmar [Burma]"`
	Namibia                  []string `json:"Namibia"`
	Nepal                    []string `json:"Nepal"`
	Netherlands              []string `json:"Netherlands"`
	NewCaledonia             []string `json:"New Caledonia"`
	NewZealand               []string `json:"New Zealand"`
	Nicaragua                []string `json:"Nicaragua"`
	Nigeria                  []string `json:"Nigeria"`
	Norway                   []string `json:"Norway"`
	Oman                     []string `json:"Oman"`
	Pakistan                 []string `json:"Pakistan"`
	Palestine                []string `json:"Palestine"`
	Panama                   []string `json:"Panama"`
	PapuaNewGuinea           []string `json:"Papua New Guinea"`
	Paraguay                 []string `json:"Paraguay"`
	Peru                     []string `json:"Peru"`
	Philippines              []string `json:"Philippines"`
	Poland                   []string `json:"Poland"`
	Portugal                 []string `json:"Portugal"`
	PuertoRico               []string `json:"Puerto Rico"`
	RepublicOfKorea          []string `json:"Republic of Korea"`
	RepublicOfLithuania      []string `json:"Republic of Lithuania"`
	RepublicOfMoldova        []string `json:"Republic of Moldova"`
	Romania                  []string `json:"Romania"`
	Russia                   []string `json:"Russia"`
	SaintLucia               []string `json:"Saint Lucia"`
	SanMarino                []string `json:"San Marino"`
	SaudiArabia              []string `json:"Saudi Arabia"`
	Senegal                  []string `json:"Senegal"`
	Serbia                   []string `json:"Serbia"`
	Singapore                []string `json:"Singapore"`
	Slovakia                 []string `json:"Slovakia"`
	Slovenia                 []string `json:"Slovenia"`
	SouthAfrica              []string `json:"South Africa"`
	Spain                    []string `json:"Spain"`
	SriLanka                 []string `json:"Sri Lanka"`
	Sudan                    []string `json:"Sudan"`
	Suriname                 []string `json:"Suriname"`
	Swaziland                []string `json:"Swaziland"`
	Sweden                   []string `json:"Sweden"`
	Switzerland              []string `json:"Switzerland"`
	Taiwan                   []string `json:"Taiwan"`
	Tanzania                 []string `json:"Tanzania"`
	Thailand                 []string `json:"Thailand"`
	TrinidadAndTobago        []string `json:"Trinidad and Tobago"`
	Tunisia                  []string `json:"Tunisia"`
	Turkey                   []string `json:"Turkey"`
	USVirginIslands          []string `json:"U.S. Virgin Islands"`
	Ukraine                  []string `json:"Ukraine"`
	UnitedArabEmirates       []string `json:"United Arab Emirates"`
	UnitedKingdom            []string `json:"United Kingdom"`
	UnitedStates             []string `json:"United States"`
	Uruguay                  []string `json:"Uruguay"`
	Venezuela                []string `json:"Venezuela"`
	Vietnam                  []string `json:"Vietnam"`
	Zambia                   []string `json:"Zambia"`
	Zimbabwe                 []string `json:"Zimbabwe"`
}

func LoadConfiguration(filename string) (Cidades, error) {
	var config Cidades
	configFile, err := os.Open(filename)
	defer configFile.Close()

	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}

func main() {
	fmt.Println("COMEÇANDO A APLICAÇÃO...")

	config, _ := LoadConfiguration("config.json")

	fmt.Println(config.Brazil)

}
