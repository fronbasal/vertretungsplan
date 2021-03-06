package lookup

import (
	"bytes"
	"encoding/csv"
	"github.com/lunny/log"
	"strings"
)

type Lookup struct {
	RawRecords    [][]string
	Lookup        map[string]string
	ReverseLookup map[string]string
	Title         map[string]string
	path          string
}

//TeacherLookup is the default Lookup table, created on New()
var TeacherLookup *Lookup

const LookupTable = `title   ,name                ,initials,subjects       ,special
Frau    ,Anker               ,Ank     ,D E PP         ,Individuelle Förderung; Betreuung der Studierenden im Praxissemester
Frau    ,Arendt              ,Ar      ,D Ge Kr        ,Lehrerausbildung
Frau    ,Avagyan             ,Ava     ,E Rus          ,Austausch St. Petersburg
Herr Dr.,Baumann             ,Bau     ,F Ge           ,Austausch Calais und Rochefort; DELF; Homepage (redaktionell); Schularchiv
Herr    ,Becker              ,Bec     ,Sp Ch          ,"NRW-Sportschule, Sport: Rudern"
Herr    ,Bernhardt           ,Bdt     ,Ge Sp          ,
Herr    ,Buchthal            ,Bu      ,SW Bi          ,Schulorganisation (Schulleiter)
Frau    ,van den Daele       ,Da      ,Ek L KR        ,Methodenkonzept SII
Frau    ,Deboße-Stenger      ,Db      ,M Ph           ,Laufbahnberatung Oberstufe (Leiterin)
Herr    ,Debski              ,Deb     ,Sp             ,"Sport: Basketball, Judo, Schwimmen, Tischtennis"
Frau    ,Demberger           ,Dem     ,Mu             ,
Frau    ,Diederichs          ,Die     ,Ek Ph          ,"Laufbahnberatung Oberstufe, Streitschlichter"
Frau    ,Fastje              ,Fas     ,E L            ,Lehrerfortbildung
Frau    ,Freudenhammer       ,Fre     ,Ku PL          ,
Frau    ,Finis               ,Fin     ,E KR           ,
Frau    ,Fritz               ,Fri     ,D Pa           ,"Schülerzeitung ""Steinbartblätter"""
Herr    ,van Fürden          ,Für     ,Sp Bi          ,
Herr Dr.,Gaida               ,Ga      ,Bi Ek          ,
Frau    ,Görgen              ,Gör     ,D KR           ,"Schule ohne Rassismus, SV-Verbindungslehrerin"
Frau    ,Graby-Meimers       ,Mei     ,E F            ,"Individuelle Förderung, Förderpläne"
Frau    ,Grätz               ,Gz      ,M CH           ,"Begabtenförderung ""Mathematik"""
Frau    ,Haase               ,Haa     ,Bi E           ,"Suchtprävention, Berufswahlorientierung; KaoA"
Frau    ,Haßlinghaus         ,Hhs     ,D KR           ,Kooperation Theater der Stadt
Frau    ,Heikaus-Loske       ,Hk      ,E Bi           ,Betreuung der Mentoren/Schülerhelfer
Herr Dr.,Henke               ,Hen     ,Ch M           ,"Erste Hilfe, Sicherheit in der Schule, Schulsanitäter"
Frau    ,Heß                 ,Hes     ,Mu             ,Pop-Chor
Frau    ,Heydari             ,Hey     ,Bi Ch          ,"Suchtprävention, SV-Verbindungslehrerin"
Herr    ,Hötten              ,Höt     ,E M            ,
Frau    ,Hofer               ,Ho      ,F Ge           ,
Herr    ,Ilgen               ,Ilg     ,Bi Sp KR       ,"Laufbahnberatung Erprobungsstufe (Leiter); Sport: Fußball, Golf"
Frau    ,Jost                ,Jos     ,F D Sp         ,
Herr    ,Jäger               ,Jäg     ,Mu Ev. Religion,Big-Band
Frau    ,Kaiser              ,Kai     ,M Ph           ,
Frau    ,Kedzierski          ,Ked     ,D Ku           ,"Lehrerausbildung, Praktikanten"
Herr    ,Kirschner           ,Kir     ,Bi Ch          ,Sport: Handball
Herr    ,Kizilaslan          ,Kiz     ,SW Inf         ,Börsenspiel
Frau    ,Kleene              ,Kln     ,L E            ,Schulbücher
Herr    ,Klein               ,Kl      ,Ge ER          ,Schulbücher / Steinbart-Journal (Redaktion)
Frau    ,Kloos               ,Kls     ,Sp Ek          ,
Frau    ,Kohlrausch-Schneider,Koh     ,Ph If          ,Schulnetzwerk
Frau    ,Kriegs              ,Kri     ,M Ek           ,
Herr    ,Kunze               ,Kun     ,D Pl           ,
Frau    ,Leschczyk           ,Les     ,Ku D           ,Kooperation Filmforum
Frau    ,Letzner             ,Letz    ,D SW           ,Laufbahnberatung Mittelstufe
Frau Dr.,Lilie               ,Lil     ,Ph Ch Pa       ,"Begabtenförderung ""Naturwissenschaften"""
Frau    ,Lohbeck             ,Loh     ,M Ku           ,
Herr    ,Loiberzeder         ,Loi     ,Ku L           ,
Frau    ,Mählck              ,Mäh     ,SW M           ,
Frau    ,Mai-Buchholz        ,Mai     ,Bi S           ,
Frau    ,Mainka              ,Mnk     ,Sonderpäd.'    ,Beratung Inklusion
Frau    ,Matsuo              ,Mat     ,Jap            ,Schüleraustausch Japan
Frau    ,Mißler              ,Mis     ,SW D           ,"Beratung LRS, Lernberatung"
Herr Dr.,Mönkemeier          ,Mön     ,ER             ,Laufbahnberatung Oberstufe
Herr    ,Musleh              ,Mus     ,D              ,
Frau    ,Neumann             ,Nm      ,M F            ,"Begabtenförderung ""Mathematik"""
Frau    ,Pietsch             ,Pie     ,E S            ,"Sprachförderung / Begabtenförderung ""Sprachen"""
Frau    ,Plorin              ,Pl      ,Ch F           ,Laufbahnberatung Mittelstufe (Leiterin)
Frau    ,Rödiger             ,Röd     ,E Ge           ,Inklusion / Schüleraustausch USA
Herr    ,Rosenthal           ,Rt      ,ER SW          ,Sozialpraktikum / Soziales Lernen
Frau    ,Schädlich           ,Säl     ,D Pl           ,Kooperation Stadtbibliothek
Frau Dr.,Schmahl             ,Shma    ,D Ge           ,
Herr Dr.,Schmidt             ,Sci     ,Bi Ch          ,Biologie-Olympiade
Herr    ,Schostok            ,Stk     ,M Ph           ,"Laufbahnberatung Erprobungsstufe, Energiesparen, SV-Verbindungslehrer"
Frau    ,Schubert            ,Shu     ,D ER           ,
Frau    ,Schroeder           ,Sdr     ,M E            ,
Frau    ,Schwanke            ,Swa     ,E Bi           ,
Frau    ,Söller              ,Söl     ,E S            ,
Frau    ,Spinne              ,Spi     ,Sp F           ,Sport: Geräteturnen
Herr    ,Spliethoff          ,Sph     ,KR E D         ,Laufbahnberatung Oberstufe
Frau    ,Stabel              ,Stab    ,S              ,Spanisch Konversation
Frau    ,Steindor            ,Ste     ,Ge Sp eR       ,"Schulnetzwerk, Vertretungsplan; Sport: Fußball, Hockey, Leichtathletik"
Frau    ,Stuhrmann           ,Stu     ,D S            ,Organisation Erprobungsstufe
Frau    ,Sultan              ,Sul     ,E S            ,
Herr    ,Thamm               ,Tha     ,Ge SW          ,Laufbahnberatung Mittelstufe (Leiter)
Herr    ,Thieler             ,Thi     ,Sp SW          ,
Herr    ,Thurow              ,Thu     ,Bi Ek          ,"Suchtprävention, Lernberatung"
Frau    ,Uhlendorff          ,Uff     ,F EK           ,
Herr    ,Uhlendorff          ,Uhl     ,D Pl           ,"Suchtprävention, Schülerzeitung ""Steinbart-Blätter"""
Herr    ,Unseld              ,Un      ,M Sp           ,Sport: Volleyball
Frau    ,Unseld              ,Uns     ,E Sp Bi        ,"Sport: Badminton, Drachenboot, Schülermarathon"
Herr    ,Volke               ,Vk      ,Mu L           ,"Kooperation Deutsche Oper am Rhein, Chor"
Frau    ,Walter              ,Wal     ,Ek Pa M        ,Vertretungsplan
Frau    ,Weber               ,Web     ,D Pa           ,Didaktische u. pädagogische Schulentwicklung
Frau    ,Weppner             ,Wep     ,D L            ,
Frau    ,Werner              ,We      ,D Pa           ,Pädagogische Arbeit der Internationalen Vorbereitungsklassen
Herr    ,Weyer               ,Wey     ,Sp             ,"Sport: Schwimmen, Förderschwimmen, Schülermarathon"
Herr    ,Wissing             ,Wis     ,M E            ,Raum- und Stundenplan / Schulorganisation
Frau    ,Wißmann-Gennert     ,Wm      ,D E            ,"Beratung LRS, Päd. Austauschdienst (PAD)"
Frau    ,Wollenweber         ,Wol     ,D Sp Bi        ,"Sport: Drachenboot, Leichtathletik, Schülermarathon"
Frau    ,Wülfingen           ,Wül     ,Ge F           ,Methodenkonzept SI
Frau    ,Zaree-Parsi         ,Zar     ,D ER           ,
Frau    ,Schädlich           ,Sdl     ,Pl             ,
Frau    ,Mälck               ,Mlk     ,M              ,`

func (l *Lookup) ReadFile() {
	csvReader := csv.NewReader(bytes.NewBuffer([]byte(LookupTable)))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read file - invalid CSV? ", err)
	}
	l.RawRecords = records
	for _, record := range records {
		l.Lookup[removeSpaces(record[2])] = removeSpaces(record[1])
		l.ReverseLookup[removeSpaces(record[1])] = removeSpaces(record[2])
		l.Title[removeSpaces(record[2])] = removeSpaces(record[0])
	}
	log.Debug("Loaded teachers")
}

func New() *Lookup {
	l := &Lookup{Lookup: map[string]string{}, ReverseLookup: map[string]string{}, Title: map[string]string{}}
	TeacherLookup = l
	l.ReadFile()
	return TeacherLookup
}

func removeSpaces(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func (l *Lookup) Get(s string) string {
	if len(s) >= 2 && len(s) <= 3 {
		l.GetRaw(s)
	}
	// Split
	if strings.Contains(s, "=>") {
		return l.GetRaw(strings.Split(s, " => ")[0]) + " => " + l.GetRaw(strings.Split(s, " => ")[1])
	} else {
		return l.GetRaw(s)
	}
}

func (l *Lookup) GetRaw(s string) string {
	if l.Lookup[s] == "" {
		return s
	}
	return l.Lookup[s]
}

func (l *Lookup) GetFull(s string) string {
	return l.Title[l.Get(s)] + " " + l.Get(s)
}
