package structs

/*
Constructor for a country
Parameters in the constructor are:
Name, Cca2, Region, Capital, Borders, Languages, Maps*/
type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Cca2      string   `json:"cca2"`
	Borders   []string `json:"borders"`
	Languages struct {
		Cnr string `json:"cnr,omitempty"`
		Smo string `json:"smo,omitempty"`
		Tkl string `json:"tkl,omitempty"`
		Spa string `json:"spa,omitempty"`
		Ell string `json:"ell,omitempty"`
		Khm string `json:"khm,omitempty"`
		Zho string `json:"zho,omitempty"`
		Ara string `json:"ara,omitempty"`
		Hrv string `json:"hrv,omitempty"`
		Deu string `json:"deu,omitempty"`
		Kaz string `json:"kaz,omitempty"`
		Rus string `json:"rus,omitempty"`
		Dan string `json:"dan,omitempty"`
		Cal string `json:"cal,omitempty"`
		Cha string `json:"cha,omitempty"`
		Ita string `json:"ita,omitempty"`
		Afr string `json:"afr,omitempty"`
		Nbl string `json:"nbl,omitempty"`
		Nso string `json:"nso,omitempty"`
		Sot string `json:"sot,omitempty"`
		Ssw string `json:"ssw,omitempty"`
		Tsn string `json:"tsn,omitempty"`
		Tso string `json:"tso,omitempty"`
		Ven string `json:"ven,omitempty"`
		Xho string `json:"xho,omitempty"`
		Zul string `json:"zul,omitempty"`
		Kin string `json:"kin,omitempty"`
		Por string `json:"por,omitempty"`
		Run string `json:"run,omitempty"`
		Ukr string `json:"ukr,omitempty"`
		Rar string `json:"rar,omitempty"`
		Pov string `json:"pov,omitempty"`
		Eng string `json:"eng,omitempty"`
		Fra string `json:"fra,omitempty"`
		Nld string `json:"nld,omitempty"`
		Isl string `json:"isl,omitempty"`
		Grn string `json:"grn,omitempty"`
		Hun string `json:"hun,omitempty"`
		Ron string `json:"ron,omitempty"`
		Kal string `json:"kal,omitempty"`
		Nau string `json:"nau,omitempty"`
		Sin string `json:"sin,omitempty"`
		Tam string `json:"tam,omitempty"`
		Heb string `json:"heb,omitempty"`
		Dzo string `json:"dzo,omitempty"`
		Kor string `json:"kor,omitempty"`
		Fao string `json:"fao,omitempty"`
		Pap string `json:"pap,omitempty"`
		Arc string `json:"arc,omitempty"`
		Ckb string `json:"ckb,omitempty"`
		Ber string `json:"ber,omitempty"`
		Mkd string `json:"mkd,omitempty"`
		Pol string `json:"pol,omitempty"`
		Slv string `json:"slv,omitempty"`
		Tha string `json:"tha,omitempty"`
		Sag string `json:"sag,omitempty"`
		Hye string `json:"hye,omitempty"`
		Nor string `json:"nor,omitempty"`
		Ltz string `json:"ltz,omitempty"`
		Srp string `json:"srp,omitempty"`
		Swe string `json:"swe,omitempty"`
		Mfe string `json:"mfe,omitempty"`
		Swa string `json:"swa,omitempty"`
		Kat string `json:"kat,omitempty"`
		Mlt string `json:"mlt,omitempty"`
		Bwg string `json:"bwg,omitempty"`
		Kck string `json:"kck,omitempty"`
		Khi string `json:"khi,omitempty"`
		Ndc string `json:"ndc,omitempty"`
		Nde string `json:"nde,omitempty"`
		Nya string `json:"nya,omitempty"`
		Sna string `json:"sna,omitempty"`
		Toi string `json:"toi,omitempty"`
		Zib string `json:"zib,omitempty"`
		Bjz string `json:"bjz,omitempty"`
		Lat string `json:"lat,omitempty"`
		Bel string `json:"bel,omitempty"`
		Her string `json:"her,omitempty"`
		Hgm string `json:"hgm,omitempty"`
		Kwn string `json:"kwn,omitempty"`
		Loz string `json:"loz,omitempty"`
		Ndo string `json:"ndo,omitempty"`
		Mya string `json:"mya,omitempty"`
		Gle string `json:"gle,omitempty"`
		Lit string `json:"lit,omitempty"`
		Kon string `json:"kon,omitempty"`
		Lin string `json:"lin,omitempty"`
		Lua string `json:"lua,omitempty"`
		Fil string `json:"fil,omitempty"`
		Msa string `json:"msa,omitempty"`
		Mon string `json:"mon,omitempty"`
		Mey string `json:"mey,omitempty"`
		Tet string `json:"tet,omitempty"`
		Fin string `json:"fin,omitempty"`
		Bar string `json:"bar,omitempty"`
		Amh string `json:"amh,omitempty"`
		Bos string `json:"bos,omitempty"`
		Zdj string `json:"zdj,omitempty"`
		Nno string `json:"nno,omitempty"`
		Nob string `json:"nob,omitempty"`
		Smi string `json:"smi,omitempty"`
		Ben string `json:"ben,omitempty"`
		Jam string `json:"jam,omitempty"`
		Cat string `json:"cat,omitempty"`
		Mlg string `json:"mlg,omitempty"`
		Gsw string `json:"gsw,omitempty"`
		Roh string `json:"roh,omitempty"`
		Tvl string `json:"tvl,omitempty"`
		Vie string `json:"vie,omitempty"`
		Nep string `json:"nep,omitempty"`
		Gil string `json:"gil,omitempty"`
		Som string `json:"som,omitempty"`
		Mah string `json:"mah,omitempty"`
		Aze string `json:"aze,omitempty"`
		Hmo string `json:"hmo,omitempty"`
		Tpi string `json:"tpi,omitempty"`
		Tur string `json:"tur,omitempty"`
		Tuk string `json:"tuk,omitempty"`
		Lao string `json:"lao,omitempty"`
		Crs string `json:"crs,omitempty"`
		Uzb string `json:"uzb,omitempty"`
		Nrf string `json:"nrf,omitempty"`
		Ces string `json:"ces,omitempty"`
		Slk string `json:"slk,omitempty"`
		Mri string `json:"mri,omitempty"`
		Nzs string `json:"nzs,omitempty"`
		Tir string `json:"tir,omitempty"`
		Jpn string `json:"jpn,omitempty"`
		Sqi string `json:"sqi,omitempty"`
		Bis string `json:"bis,omitempty"`
		Pau string `json:"pau,omitempty"`
		Ton string `json:"ton,omitempty"`
		Fij string `json:"fij,omitempty"`
		Hif string `json:"hif,omitempty"`
		Niu string `json:"niu,omitempty"`
		Aym string `json:"aym,omitempty"`
		Que string `json:"que,omitempty"`
		Ind string `json:"ind,omitempty"`
		Nfr string `json:"nfr,omitempty"`
		Kir string `json:"kir,omitempty"`
		Div string `json:"div,omitempty"`
		Pih string `json:"pih,omitempty"`
		Prs string `json:"prs,omitempty"`
		Pus string `json:"pus,omitempty"`
		Tgk string `json:"tgk,omitempty"`
		Hin string `json:"hin,omitempty"`
		Urd string `json:"urd,omitempty"`
		Lav string `json:"lav,omitempty"`
		Fas string `json:"fas,omitempty"`
		Bul string `json:"bul,omitempty"`
		Est string `json:"est,omitempty"`
		Hat string `json:"hat,omitempty"`
		Glv string `json:"glv,omitempty"`
	} `json:"languages,omitempty"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
}

/*
Constructor for university information.
Parameter in the constructor is:
Webpage, NameUni, Domains, CountryName, Cca2, Capital, Maps, Languages*/
type UniInfo struct {
	Webpage     []string `json:"web_pages"`
	NameUni     string   `json:"name"`
	Domains     []string `json:"domains"`
	CountryName string   `json:"country"`
	Cca2        string   `json:"cca2"`
	Languages   struct {
		Cnr string `json:"cnr,omitempty"`
		Smo string `json:"smo,omitempty"`
		Tkl string `json:"tkl,omitempty"`
		Spa string `json:"spa,omitempty"`
		Ell string `json:"ell,omitempty"`
		Khm string `json:"khm,omitempty"`
		Zho string `json:"zho,omitempty"`
		Ara string `json:"ara,omitempty"`
		Hrv string `json:"hrv,omitempty"`
		Deu string `json:"deu,omitempty"`
		Kaz string `json:"kaz,omitempty"`
		Rus string `json:"rus,omitempty"`
		Dan string `json:"dan,omitempty"`
		Cal string `json:"cal,omitempty"`
		Cha string `json:"cha,omitempty"`
		Ita string `json:"ita,omitempty"`
		Afr string `json:"afr,omitempty"`
		Nbl string `json:"nbl,omitempty"`
		Nso string `json:"nso,omitempty"`
		Sot string `json:"sot,omitempty"`
		Ssw string `json:"ssw,omitempty"`
		Tsn string `json:"tsn,omitempty"`
		Tso string `json:"tso,omitempty"`
		Ven string `json:"ven,omitempty"`
		Xho string `json:"xho,omitempty"`
		Zul string `json:"zul,omitempty"`
		Kin string `json:"kin,omitempty"`
		Por string `json:"por,omitempty"`
		Run string `json:"run,omitempty"`
		Ukr string `json:"ukr,omitempty"`
		Rar string `json:"rar,omitempty"`
		Pov string `json:"pov,omitempty"`
		Eng string `json:"eng,omitempty"`
		Fra string `json:"fra,omitempty"`
		Nld string `json:"nld,omitempty"`
		Isl string `json:"isl,omitempty"`
		Grn string `json:"grn,omitempty"`
		Hun string `json:"hun,omitempty"`
		Ron string `json:"ron,omitempty"`
		Kal string `json:"kal,omitempty"`
		Nau string `json:"nau,omitempty"`
		Sin string `json:"sin,omitempty"`
		Tam string `json:"tam,omitempty"`
		Heb string `json:"heb,omitempty"`
		Dzo string `json:"dzo,omitempty"`
		Kor string `json:"kor,omitempty"`
		Fao string `json:"fao,omitempty"`
		Pap string `json:"pap,omitempty"`
		Arc string `json:"arc,omitempty"`
		Ckb string `json:"ckb,omitempty"`
		Ber string `json:"ber,omitempty"`
		Mkd string `json:"mkd,omitempty"`
		Pol string `json:"pol,omitempty"`
		Slv string `json:"slv,omitempty"`
		Tha string `json:"tha,omitempty"`
		Sag string `json:"sag,omitempty"`
		Hye string `json:"hye,omitempty"`
		Nor string `json:"nor,omitempty"`
		Ltz string `json:"ltz,omitempty"`
		Srp string `json:"srp,omitempty"`
		Swe string `json:"swe,omitempty"`
		Mfe string `json:"mfe,omitempty"`
		Swa string `json:"swa,omitempty"`
		Kat string `json:"kat,omitempty"`
		Mlt string `json:"mlt,omitempty"`
		Bwg string `json:"bwg,omitempty"`
		Kck string `json:"kck,omitempty"`
		Khi string `json:"khi,omitempty"`
		Ndc string `json:"ndc,omitempty"`
		Nde string `json:"nde,omitempty"`
		Nya string `json:"nya,omitempty"`
		Sna string `json:"sna,omitempty"`
		Toi string `json:"toi,omitempty"`
		Zib string `json:"zib,omitempty"`
		Bjz string `json:"bjz,omitempty"`
		Lat string `json:"lat,omitempty"`
		Bel string `json:"bel,omitempty"`
		Her string `json:"her,omitempty"`
		Hgm string `json:"hgm,omitempty"`
		Kwn string `json:"kwn,omitempty"`
		Loz string `json:"loz,omitempty"`
		Ndo string `json:"ndo,omitempty"`
		Mya string `json:"mya,omitempty"`
		Gle string `json:"gle,omitempty"`
		Lit string `json:"lit,omitempty"`
		Kon string `json:"kon,omitempty"`
		Lin string `json:"lin,omitempty"`
		Lua string `json:"lua,omitempty"`
		Fil string `json:"fil,omitempty"`
		Msa string `json:"msa,omitempty"`
		Mon string `json:"mon,omitempty"`
		Mey string `json:"mey,omitempty"`
		Tet string `json:"tet,omitempty"`
		Fin string `json:"fin,omitempty"`
		Bar string `json:"bar,omitempty"`
		Amh string `json:"amh,omitempty"`
		Bos string `json:"bos,omitempty"`
		Zdj string `json:"zdj,omitempty"`
		Nno string `json:"nno,omitempty"`
		Nob string `json:"nob,omitempty"`
		Smi string `json:"smi,omitempty"`
		Ben string `json:"ben,omitempty"`
		Jam string `json:"jam,omitempty"`
		Cat string `json:"cat,omitempty"`
		Mlg string `json:"mlg,omitempty"`
		Gsw string `json:"gsw,omitempty"`
		Roh string `json:"roh,omitempty"`
		Tvl string `json:"tvl,omitempty"`
		Vie string `json:"vie,omitempty"`
		Nep string `json:"nep,omitempty"`
		Gil string `json:"gil,omitempty"`
		Som string `json:"som,omitempty"`
		Mah string `json:"mah,omitempty"`
		Aze string `json:"aze,omitempty"`
		Hmo string `json:"hmo,omitempty"`
		Tpi string `json:"tpi,omitempty"`
		Tur string `json:"tur,omitempty"`
		Tuk string `json:"tuk,omitempty"`
		Lao string `json:"lao,omitempty"`
		Crs string `json:"crs,omitempty"`
		Uzb string `json:"uzb,omitempty"`
		Nrf string `json:"nrf,omitempty"`
		Ces string `json:"ces,omitempty"`
		Slk string `json:"slk,omitempty"`
		Mri string `json:"mri,omitempty"`
		Nzs string `json:"nzs,omitempty"`
		Tir string `json:"tir,omitempty"`
		Jpn string `json:"jpn,omitempty"`
		Sqi string `json:"sqi,omitempty"`
		Bis string `json:"bis,omitempty"`
		Pau string `json:"pau,omitempty"`
		Ton string `json:"ton,omitempty"`
		Fij string `json:"fij,omitempty"`
		Hif string `json:"hif,omitempty"`
		Niu string `json:"niu,omitempty"`
		Aym string `json:"aym,omitempty"`
		Que string `json:"que,omitempty"`
		Ind string `json:"ind,omitempty"`
		Nfr string `json:"nfr,omitempty"`
		Kir string `json:"kir,omitempty"`
		Div string `json:"div,omitempty"`
		Pih string `json:"pih,omitempty"`
		Prs string `json:"prs,omitempty"`
		Pus string `json:"pus,omitempty"`
		Tgk string `json:"tgk,omitempty"`
		Hin string `json:"hin,omitempty"`
		Urd string `json:"urd,omitempty"`
		Lav string `json:"lav,omitempty"`
		Fas string `json:"fas,omitempty"`
		Bul string `json:"bul,omitempty"`
		Est string `json:"est,omitempty"`
		Hat string `json:"hat,omitempty"`
		Glv string `json:"glv,omitempty"`
	} `json:"languages,omitempty"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
}

/*
Setter function.
The fields it will edit is:
Cca2, Capital, Language, Maps
o *UniInfo -> says that it is a UniInfo object that the function will work on
country []Country -> country slice needed for the function to do it's functions*/
func (o *UniInfo) SetCountryInfo(country []Country) {
	o.Cca2 = country[0].Cca2
	o.Languages = country[0].Languages
	o.Maps = country[0].Maps
}

/*
Constructor for Diag information
Parameters in the constructor are:
Universitiapi, Countriesapi, Version, Uptime*/
type Diag struct {
	Universitiapi int
	Counteriesapi int
	Version       string
	Uptime        float64
}

/*
Function that will run a algorithm to remove duplicated universeties added to the slice.
u []UniInfo -> slice that will be check for duplications.
return u -> return the slice after it have been check and edited*/
func RemoveDup(u []Country) []Country {
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			if u[i].Name.Common == u[j].Name.Common {
				u = append(u[:j], u[j+1:]...)
			}
		}
	}

	return u
}

/*
Function to remove an element from a slice
u []UniInfo -> slice that will be work on
i int -> number how wich element in the slice that will be removed
return []UniInfo -> returns a new slice*/
func RemoveElement(u []UniInfo, i int) []UniInfo {
	u[i] = u[len(u)-1]
	return u[:len(u)-1]
}

/*
Function to remove an element from a slice
c []Country -> slice that will be work on
i int -> number how wich element in the slice that will be removed
return []Country -> returns a new slice*/
func RemoveElementCountry(c []Country, i int) []Country {
	c[i] = c[len(c)-1]
	return c[:len(c)-1]
}
