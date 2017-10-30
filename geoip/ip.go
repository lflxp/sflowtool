package models

/**
将geoip的数据解析成ip2region数据
 */
import (
	"encoding/csv"
	"fmt"
	"bufio"
	"io"
	"os"
	"github.com/lflxp/cidr"
	"strings"
	"strconv"
	"errors"
	"time"
	"flag"
)

var Data *[]Origin
var Locations *map[string]CityLocations
var Asn *[]AsnBlocks
var path = flag.String("path", "./data", "GeoIP2 文件目录")

func init() {
	flag.Parse()
	Data, Locations, Asn = NewOrigin(*path)
}

type CityLocations struct {
	GeonameId	string
	LocaleCode	string
	ContinentCode	string
	ContinentName	string
	CountryIsoCode	string
	CountryName	string
	S1IsoCode	string
	S1Name		string
	S2IsoCode	string
	S2Name		string
	CityName	string
	MetroCode	string
	TimeZone 	string
}

type Origin struct {
	Start 				int64
	End 				int64
	FirstIp				string
	EndIp				string
	Network 			string
	Geoname_id 			string
	Registered_country_geoname_id 	string
	Represented_country_geoname_id 	string
	Is_anonymous_proxy		string
	Is_satellite_provider 		string
	Postal_code 			string
	Latitude			string
	Longitude			string
	Accuracy_radius			string
}

type AsnBlocks struct {
	Start 				int64
	End 				int64
	FirstIp				string
	EndIp				string
	Network 			string
	Autonomous_system_number 	string
	Autonomous_system_organization 	string
}

type JSON struct {
	Ip 		string
	Status 		string
	Time 		string
	Blocks  	Origin
	Locations 	CityLocations
	Asn 		AsnBlocks
}

func Reader(path string) {
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record,err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:",err)
			return
		}
		fmt.Println(record)
	}
}

//读取csv文件,初始化信息
func NewOrigin(path string) (*[]Origin,*map[string]CityLocations,*[]AsnBlocks) {
	Blocks := GetCityBlocksIpv4(path+"/GeoLite2-City-Blocks-IPv4.csv")
	Locations := GetCityLocations(path+"/GeoLite2-City-Locations-zh-CN.csv")
	Asn := GetAsnBlocks(path+"/GeoLite2-ASN-Blocks-IPv4.csv")
	return Blocks,Locations,Asn
}

//根据给定的IP解析所有信息
func ParseIp(ip string) *JSON {
	begin := time.Now()
	rs := JSON{}
	rs.Ip = ip

	id := BinarySearchCityBlocksIPv4(ip)
	if id == -1 {
		rs.Time = time.Since(begin).String()
		rs.Status = fmt.Sprintf("%s 查无此IP,请确认是否为内网IP",ip)
		return &rs
	} else {
		rs.Blocks = (*Data)[id]
		rs.Locations = (*Locations)[rs.Blocks.Geoname_id]
	}

	asnId := BinarySearchAsnIPv4(ip)
	if asnId == -1 {
		rs.Status = fmt.Sprintf("%s 无法查询网络运营商,请确认是否为内网IP",ip)
	} else {
		rs.Asn = (*Asn)[asnId]
	}
	rs.Status = "Successfully"
	rs.Time = time.Since(begin).String()
	return &rs
}

//二分法查询Asn
func BinarySearchAsnIPv4(ip string) int {
	lo,hi := 0,len(*Asn)-1
	k,err := ip2long(ip)
	//fmt.Println("ip ",ip,k)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for lo <= hi {
		m := (lo+hi)>>1
		//fmt.Println(m,k,(*data)[m].Start,(*data)[m].End,(*data)[m].FirstIp,(*data)[m].EndIp)
		if (*Data)[m].Start < k {
			if (*Data)[m].End < k {
				lo = m + 1
			} else if (*Data)[m].End > k {
				return m
			}
		} else if (*Data)[m].Start > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

//通过二分法解析ip对应的ip段
//切片s是升序的
//k为待查找的整数
//如果查到有就返回对应角标,
//没有就返回-1
func BinarySearchCityBlocksIPv4(ip string) int {
	lo,hi := 0,len(*Data)-1
	k,err := ip2long(ip)
	//fmt.Println("ip ",ip,k)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for lo <= hi {
		m := (lo+hi)>>1
		//fmt.Println(m,k,(*data)[m].Start,(*data)[m].End,(*data)[m].FirstIp,(*data)[m].EndIp)
		if (*Data)[m].Start < k {
			if (*Data)[m].End < k {
				lo = m + 1
			} else if (*Data)[m].End > k {
				return m
			}
		} else if (*Data)[m].Start > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

func GetAsnBlocks(path string) *[]AsnBlocks {
	fmt.Println("开始读取AsnBlocks文件"+path)
	data := []AsnBlocks{}

	//locations := map[string]string{}

	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
		    "Does the file exist?\n" +
		    "Have you got acces to it?\n")
		return nil
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	//读取数据到内存
	//踢重 国家 城市 省份
	for {

		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		} else if readerError != nil {
			fmt.Println("Error:",readerError)
			return nil
		}
		record := strings.Split(strings.Replace(inputString,"\"","",-1),",")
		if record[0] != "network" {
			//fmt.Println(record)
			count := cidr.NewCidr(record[0]).GetCidrIpRange()
			tmp := AsnBlocks{}
			tmp.Start,_ = ip2long(count.Min)
			tmp.End,_ = ip2long(count.Max)
			tmp.FirstIp = count.Min
			tmp.EndIp = count.Max
			tmp.Network = record[0]
			tmp.Autonomous_system_number = record[1]
			tmp.Autonomous_system_organization = record[2]

			if record[1] != "" {
				//locations[record[5]] = record[7]
				data = append(data,tmp)
			}
		}
	}
	fmt.Println("读取AsnBlocks文件完毕"+path)
	return &data
}

func GetCityLocations(path string) *map[string]CityLocations {
	fmt.Println("开始读取CityLocations文件"+path)
	data := map[string]CityLocations{}

	//locations := map[string]string{}

	//locations := map[string]string{}

	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
		    "Does the file exist?\n" +
		    "Have you got acces to it?\n")
		return nil
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	//读取数据到内存
	//踢重 国家 城市 省份
	for {

		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		} else if readerError != nil {
			fmt.Println("Error:",readerError)
			return nil
		}
		record := strings.Split(strings.Replace(inputString,"\"","",-1),",")
		if record[0] != "geoname_id" {
			//fmt.Println(record)
			tmp := CityLocations{
				GeonameId:record[0],
				LocaleCode:record[1],
				ContinentCode:record[2],
				ContinentName:record[3],
				CountryIsoCode:record[4],
				CountryName:record[5],
				S1IsoCode:record[6],
				S1Name:record[7],
				S2IsoCode:record[6],
				S2Name:record[9],
				CityName:record[10],
				MetroCode:record[11],
				TimeZone:record[12],
			}
			if record[5] != "" {
				//locations[record[5]] = record[7]
				data[tmp.GeonameId] = tmp
			}
		}
	}
	fmt.Println("读取CityLocations完毕")
	return &data
}


//解析数据
func GetCityBlocksIpv4(path string) *[]Origin {
	result := []Origin{}
	fmt.Println("开始读取CityBlocks文件"+path)
	//locations := map[string]string{}

	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
		    "Does the file exist?\n" +
		    "Have you got acces to it?\n")
		return nil
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	//读取数据到内存
	//踢重 国家 城市 省份
	for {

		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		} else if readerError != nil {
			fmt.Println("Error:",readerError)
			return nil
		}
		record := strings.Split(strings.Replace(inputString,"\"","",-1),",")
		if record[0] != "network" {
			tmp := Origin{}
			count := cidr.NewCidr(record[0]).GetCidrIpRange()
			tmp.Start,_ = ip2long(count.Min)
			tmp.End,_ = ip2long(count.Max)
			tmp.FirstIp = count.Min
			tmp.EndIp = count.Max
			tmp.Network = record[0]
			tmp.Geoname_id = record[1]
			tmp.Registered_country_geoname_id = record[2]
			tmp.Represented_country_geoname_id = record[3]
			tmp.Is_anonymous_proxy = record[4]
			tmp.Is_satellite_provider = record[5]
			tmp.Postal_code = record[6]
			tmp.Latitude = record[7]
			tmp.Longitude = record[8]
			tmp.Accuracy_radius = record[9]
			//fmt.Println("start ",record[0],count.Min,count.Max,tmp.Start)
			result = append(result,tmp)
		}
	}
	fmt.Println("读取完毕CityBlocks文件")
	return &result
}

func WriteFile(path string,info []string) {
	fmt.Println("开始写入文件...")
	file,err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _,x := range info {
		file.WriteString(x+"\n")
	}
	file.Close()
	fmt.Println("写入完毕")
}



func getLong(b []byte, offset int64) int64 {

	val := (int64(b[offset]) |
		int64(b[offset+1])<<8 |
		int64(b[offset+2])<<16 |
		int64(b[offset+3])<<24)

	return val

}

func ip2long(IpStr string) (int64, error) {
	bits := strings.Split(IpStr, ".")
	if len(bits) != 4 {
		return 0, errors.New("ip format error")
	}

	var sum int64
	for i, n := range bits {
		bit, _ := strconv.ParseInt(n, 10, 64)
		sum += bit << uint(24-8*i)
	}

	return sum, nil
}