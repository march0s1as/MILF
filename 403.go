package main

import (
	"os"
	"fmt"
	"time"
	"flag"
	"bufio"
	"strings"
	"strconv"
	"net/http"
	"io/ioutil"
	"crypto/tls"

	"github.com/fatih/color"
)

func eu_tenho_depressao_severa(url, size, ctype, linhas, delay, header string){
	_ = header
	fmt.Print("[Status: ")
	frescura("200")
	if header != "Connection:Close"{
		fmt.Print(" :: Content-Length: " + size +  " :: Content-Type: " + ctype + " :: Linhas: " + linhas + " :: Threads: " + delay + "] ======> [" + header + "]\n")
	} else {
		fmt.Print(" :: Content-Length: " + size +  " :: Content-Type: " + ctype + " :: Linhas: " + linhas + " :: Threads: " + delay + "] ======> [" + url + "]\n")
	}
}

func eu_tenho_depressao_severa2(url, headers, size, ctype, linhas, delay string){
	fmt.Println("[" + url + "] " + "					[Content-Length: " + size +  " :: Content-Type: " + ctype + " :: Linhas: " + linhas + " :: Threads: " + delay + "]" + " 		[" + headers + "]")
}

func ascii(){
	color.Set(color.FgGreen)
	fmt.Println(`
┌┬┐  ┬  ┬    ┌─┐ t(-_-t)
│││  │  │    ├┤   
┴ ┴  ┴  ┴─┘  └  
`)
	color.Set(color.FgWhite)
	fmt.Println("=======================")
}

func frescura(status string){
    fmt.Print("")
    color.Set(color.FgGreen)
    fmt.Print(status)
    color.Set(color.FgWhite)
    fmt.Print("")
}

func frescura2(status string){
    fmt.Print("[")
    color.Set(color.FgYellow)
    fmt.Print(status)
    color.Set(color.FgWhite)
    fmt.Print("] ")
}

func frescura3(status string){
    fmt.Print("[")
    color.Set(color.FgRed)
    fmt.Print(status)
    color.Set(color.FgWhite)
    fmt.Print("] ")	
}

func error(err interface{}) {
    if err != nil{
        panic(err)
    }
}

func lambda(){
	color.Set(color.FgRed)
	fmt.Print("λ")
	color.Set(color.FgWhite)
}

func Between(str, starting, ending string) string {
    s := strings.Index(str, starting)
    if s < 0 {
        return ""
    }
    s += len(starting)
    e := strings.Index(str[s:], ending)
    if e < 0 {
        return ""
    }
    return str[s : s+e]
}

func exibir(url, wordlist string, filtro, thread int){
	lambda()
	fmt.Println("  URL                  " + url)
	if filtro == 0{
		lambda()
		fmt.Println("  MATCHER              200,204,301,302,307,401,403,405")
	} else if filtro != 0 {
		lambda()
		fmt.Print("  MATCHER              ")
		fmt.Println(filtro)
	}
	lambda()
	fmt.Println("  THREADS              " + strconv.Itoa(thread))
	lambda()
	fmt.Println("  WORDLIST             " + wordlist)
	fmt.Println("=======================\n")
}

func argumentos(escolha, frase string){
	_ = frase
	if len(escolha) == 0 {

	url_vazia, err := os.OpenFile(escolha, os.O_RDWR, 0000)
	_ = url_vazia
	if err != nil {
		fmt.Println(`./403 -u http://fbi-honeypot.gov/ -w /usr/share/wordlists/headers.txt -i 192.168.0.1

~~~~
-u ----------- URL (http://example.com.br/)
-w ----------- wordlist
-i ----------- IP nos headers (DEFAULT: 127.0.0.1)
-t ----------- threads para as requests (DEFAULT: nenhuma)
-c ----------- filtrar os status code
-h ----------- cookie (DEFAULT: nenhum.)
~~~~`)
		os.Exit(0)
		}
	}
}

func countRune(s string, r rune) int {
    count := 0
    for _, c := range s {
        if c == r {
            count++
        }
    }
    return count
}

func request(url, wordlist, cookie, ip, wordlist_header string, sleep, filtro int){

    url_final := strings.Replace(url, "MILF", wordlist, -1)

    var (
      duration int = sleep
    )

    cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

    req, err := http.NewRequest("GET", url_final, nil)
    error(err)

    p := cookie

    index := 0
    cookie_primeiro := p[:index] + "HI" + p[index:]
    retirar_cookie := Between(string(cookie_primeiro),"HI",":")

    res := strings.ReplaceAll(p, retirar_cookie, "")
    primeiro_cookie := retirar_cookie
    segundo_cookie := res[2:]

    req.Header.Add(primeiro_cookie, segundo_cookie)
    req.Header.Add(wordlist_header, ip)
    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")
    req.Header.Add("Accept-Encoding", "gzip, deflate")

    request, err := cli.Do(req)
    tempo_inicial := time.Now().Unix()
    time.Sleep(time.Duration(duration) * time.Second)
    if err != nil{
        color.Red("[!] Ocorreu um erro ao estabelecer a conexão com o site.")
        os.Exit(0)
    }
    defer func() {
        _ = request.Body.Close()
    }()

    body, err := ioutil.ReadAll(request.Body)
    error(err)

    header_verificar := wordlist_header + ":" + ip
    tempo_final := time.Now().Unix()
    corpo := string(body)
    header := wordlist_header + ":" + ip
    _ = header

    size := request.Header.Get("Content-Length")
    content_type := request.Header.Get("Content-Type")
    linhas := countRune(corpo, '\n')
    linhas_string := strconv.Itoa(linhas)
    delay := tempo_final - tempo_inicial
    delay_string := strconv.FormatInt(delay, 10)

    if filtro == 0 {

        if request.Status == "200 OK"{
            eu_tenho_depressao_severa(url_final, size, content_type, linhas_string, delay_string, header_verificar)
        } else if request.Status == "403 Forbidden"{
            return
        } else if request.Status == "503 Service Temporarily Unavailable"{
            frescura2("503")
            eu_tenho_depressao_severa(url_final, size, content_type, linhas_string, delay_string, header_verificar)
        } else if request.Status == "400 Bad Request"{
            return
        } else if request.Status == "301 Moved Permanently"{
            frescura2("301")
            eu_tenho_depressao_severa(url_final, size, content_type, linhas_string, delay_string, header_verificar)
        } else if request.Status == "404 File not found"{
        	return
        } else if request.Status == "404 Not Found"{
            return
        } else if request.Status != "403 Forbidden"{
            frescura2(request.Status)
            eu_tenho_depressao_severa(url_final, size, content_type, linhas_string, delay_string, header_verificar)
        }

    } else if request.Status == "200 OK"{
        if filtro == 200{
            frescura("200")
            eu_tenho_depressao_severa(url_final, size, content_type, linhas_string, delay_string, header_verificar)
        }

    } else if request.Status == "403 Forbidden"{
        if filtro == 403 {
            frescura2("403")
            eu_tenho_depressao_severa(url_final, size, content_type, linhas_string, delay_string, header_verificar)
        }
    }
}

func ler_arquivo(wordlist, url, ip, cookie string, thread int, filtro int){

	arquivo, err := os.Open(wordlist) 
	error(err)
	 
	scanner := bufio.NewScanner(arquivo)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	
	arquivo.Close()

	exibir(url, wordlist, filtro, thread)
	 
	for _, linhas_arquivo := range txtlines {

		fuzz_verificar := strings.Contains(url, "MILF")
	    if fuzz_verificar == false{
	    	if len(cookie) == 0{
	    		request(url, linhas_arquivo, "Accept: */*", ip, linhas_arquivo, thread, filtro)
	    	} else{
	    		request(url, linhas_arquivo, cookie, ip, linhas_arquivo, thread, filtro)
	    	}
	    } else if fuzz_verificar == true {
	    	if len(cookie) == 0{
	    		request(url, linhas_arquivo, "Accept: */*", "Close", "Connection", thread, filtro)
	    	} else{
	    		request(url, linhas_arquivo, cookie, "Close", "Connection", thread, filtro)
	    	}
	    }
	}
}

func main(){
	ascii()
	var url,wordlist,ip, filtro, cookie string
	var thread string = "2" 

	flag.StringVar(&thread, "t", "", "")
	flag.StringVar(&cookie, "h", "", "")
	flag.StringVar(&filtro, "c", "", "")
	flag.StringVar(&url, "u", "", "")
	flag.StringVar(&wordlist, "w", "", "")
	flag.StringVar(&ip, "i", "", "")
	flag.CommandLine.Usage = func() { fmt.Println(`
./403 -u http://fbi-honeypot.gov/assets/ -w /usr/share/wordlists/headers.txt -i 192.168.0.1

~~~~
-u ----------- URL (http://example.com.br/)
-w ----------- wordlist
-i ----------- IP nos headers (DEFAULT: 127.0.0.1)
-t ----------- threads para as requests (DEFAULT: nenhuma)
-c ----------- filtrar os status code
-h ----------- cookie (DEFAULT: nenhum.)
~~~~`) }
	flag.Parse()

	argumentos(url, "url")
	argumentos(wordlist, "wordlist")

	st, _ := strconv.Atoi(filtro)
	tt, _ := strconv.Atoi(thread)

	if len(ip) == 0 {

	url_vazia, err := os.OpenFile(ip, os.O_RDWR, 0000)
	_ = url_vazia
	if err != nil {
		ler_arquivo(wordlist, url, "127.0.0.1", cookie, tt, st)
		}
	}

	if len(filtro) == 0 {

	url_vazia, err := os.OpenFile(ip, os.O_RDWR, 0000)
	_ = url_vazia
	if err != nil {
		ler_arquivo(wordlist, url, ip, cookie, tt, 0)
		}
	}

	if len(cookie) == 0 {

	cookie_vazio, err := os.OpenFile(cookie, os.O_RDWR, 0000)
	_ = cookie_vazio
	if err != nil {
		ler_arquivo(wordlist, url, ip, "nada", tt, st)
		}
	}

	fmt.Println(cookie)
	ler_arquivo(wordlist, url, ip, cookie, tt, st)
}