package goschool

import (
	"git.mills.io/prologic/bitcask"

	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	bolt "go.etcd.io/bbolt"
	"golang.org/x/oauth2"

	"github.com/diegocomunity/goschool/examples/cms"
	"github.com/diegocomunity/goschool/examples/mvcexample"
	"github.com/diegocomunity/goschool/examples/rpcserver"
	"github.com/diegocomunity/goschool/examples/socialfit"
	"github.com/google/go-github/github"
)

//go:embed testdata
var content embed.FS

//go:embed ext
var ext embed.FS

type goSchool struct {
}

func New() *goSchool {
	return &goSchool{}
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// interface para ejecutar otros ejemplos
type i interface {
	Run()
}

// ejecuta otros ejemplos de otras carpetas
func (*goSchool) RunOther_(name string) {
	var example i
	switch name {
	case "rpc":
		example = rpcserver.NewRPCServer()
	case "mvc":
		example = mvcexample.Mvc{}
	case "socialfit":
		example = socialfit.New()
	case "cms":
		example = cms.New()
	default:
		log.Fatalf("enter name of example")
	}
	example.Run()
}

// example buffio split para determinar si hay espacios en un texto
func (g *goSchool) Split() {
	const input = `three spaces in text`
	scanner := bufio.NewScanner(strings.NewReader(input))
	count := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		check(err)
	}
	fmt.Println("El texto tiene: ", count, "espacio")
}

func (*goSchool) Final_token() {
	const input = "amor0=1=1=2=3=4,="
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '=' {
				return i + 1, data[:1], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	})
	for scanner.Scan() {
		fmt.Printf("%q", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		check(err)
	}

}

/*
buffio writer
*/
func (*goSchool) Writer() {
	f, err := os.Create("file.txt")
	check(err)

	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "Hello world\n")
	fmt.Fprintf(w, "Te amo\n")
	fmt.Fprintf(w, " No se que hacer\n")
	w.Flush()

	r, err := os.Open("file.txt")
	check(err)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	fmt.Println("lineas: ", count, "espacios: ")
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error al crear els acn", err)
	}

}

func (*goSchool) Writer_append() {
	f, err := os.Create("file.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, i := range []int64{1, 2, 3, 4, 5, 6} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
	reader, err := os.Open("file.txt")
	check(err)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}
}

//writer zip

func (*goSchool) Writer_zip() {
	tests := []struct {
		Name, Body string
	}{
		{
			"title.txt",
			strings.Repeat("Lorem ipsum", 50),
		},
		{
			"content.txt",
			strings.Repeat("Hola mundo ", 150),
		},
		{
			"secret.txt",
			strings.Repeat("Diego es el mejor", 200),
		},
	}
	f, err := os.Create("archivo.zip")
	if err != nil {
		check(err)
	}
	//buf := new(bytes.Buffer)
	w := zip.NewWriter(f)
	for _, test := range tests {
		f, err := w.Create(test.Name)
		if err != nil {
			check(err)
		}
		_, err = f.Write([]byte(test.Body))
		if err != nil {
			check(err)
		}
	}
	err = w.Close()
	if err != nil {
		check(err)
	}
}

// como leer un zip
func (*goSchool) Reader_zip() {
	r, err := zip.OpenReader("archivo.zip")
	if err != nil {
		check(err)
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Println("Estos son los archivos: ", f.Name)
		rc, err := f.Open()
		if err != nil {
			check(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			check(err)
		}
		rc.Close()
	}
}

/*
bytes examples
*/

func (*goSchool) AvailableBuffer() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "hello world\n")
	b := strconv.AppendBool([]byte("amor"), false)
	b = append(b, ' ')
	b = append(b, '\n')
	w.Write(b)
	w.Flush()
}

func (*goSchool) ExampleBuffer() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprint(&b, "world")
	b.WriteTo(os.Stdout)
}

// example bytes
func (goSchool) ExampleBuffer_Read() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	rdbuf := make([]byte, 1)
	n, err := b.Read(rdbuf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(b.String())
	fmt.Println(string(rdbuf))
	// Output
	// 1
	// bcde
	// a
}

// example buffer
func (goSchool) ExampleBuffer_v2() {
	var w = bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "Hello")
	fmt.Fprintf(w, "world\n")
	//b := w.AvailableBuffer()
	//w.Write(b)
	w.Flush()

}
func (goSchool) ExampleBytes_v2() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, "world! from bytes \n")
	b.WriteTo(os.Stdout)
}

func (goSchool) Example_Valida_pount_comma() {
	var input = `b = 10;
println(a+b);
hola;
a;
`
	//var w bytes.Buffer
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); s != nil {
		check(err)
	}

	for _, line := range lines {
		size_line := bytes.NewReader([]byte(line)).Len() //numero de caracteres de una linea
		pos_pount_comma := strings.IndexRune(line, ';')  //posiciòn del punto y coma

		if pos_pount_comma != (size_line - 1) { //validar si el punto y coma està al final
			fmt.Fprint(os.Stderr, "failed ;")
			os.Exit(1)
		}

	}
}

// ejercicio de urls amigables
func (goSchool) Example_http_v1() {

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.Host)

	})
	var hanlder_u = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json")
		fmt.Fprintf(os.Stdout, r.Host)
		fmt.Fprintf(w, r.Host)
	})
	var mux = http.NewServeMux()
	mux.Handle("/", handler)
	mux.Handle("/u", hanlder_u)
	fmt.Println("Server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

/*
start example http_v2
*/
func containsDotFile(name string) bool {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

type dotFileHidingFile struct {
	http.File
}

type dotFileHidingFileSystem struct {
	http.FileSystem
}

func (fsys dotFileHidingFileSystem) Open(name string) (http.File, error) {
	if containsDotFile(name) { // If dot file, return 403 response
		return nil, fs.ErrPermission
	}

	file, err := fsys.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return dotFileHidingFile{file}, err
}
func (goSchool) Example_http_with_reader() {
	fsys := dotFileHidingFileSystem{http.Dir("./testdata")}
	http.Handle("/", http.FileServer(fsys))
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		var scanner = bufio.NewScanner(r.Body)
		for scanner.Scan() {
			go fmt.Printf(">_ %v\n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	})
	fmt.Println("running server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
end example_http_v2
*/

// other example

func (goSchool) Example_http_Zip() {
	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		fmt.Fprintf(w, `
		<form id="exampleForm">
			<input type="text" name="title">
			<input type="text" name="body">
			<input type="submit">
		</form>
		<script>
		const get_exampleFormWithId = document.getElementById("exampleForm");
		get_exampleFormWithId.addEventListener("submit", (e) => {
			e.preventDefault()
			const data = new FormData(document.getElementById('exampleForm'));
			fetch('/foo', {
				body: data,
				method: 'POST'
			}).then(res => res.text())
			.then(res => {
					console.log(res)
				})
		})
		</script>
		`)
	})
	f, err := os.Create("testdata/bd.zip")
	if err != nil {
		check(err)
	}
	wr := zip.NewWriter(f)
	http.Handle("/", handler)
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		var title, body string = r.FormValue("title"), r.FormValue("body")
		f, err := wr.Create(title)
		if err != nil {
			check(err)
		}
		_, err = f.Write([]byte(body))
		if err != nil {
			check(err)
		}
		defer wr.Close()
	})
	fmt.Println("listen server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// example with template engine
type dotContent struct {
	Title, Content string
}

func (goSchool) Example_http_htmlEngine() {
	var tpl = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
	<style>
	body, html {
		background-color: #333;
	}
	</style>
</head>
<body>
	<header claass="header">
		<nav class="nav">
			<a href="/">/</a>
			<br>
			<a href="/foo">Foo</a>
			<br>
			<a href="/bar">Bar</a>
		</nav>
	</header>
	<main>
		{{.Content}}
	</main>
	<footer>
		<p>FOOTER</p>
	</footer>
</body>
</html>
	
	`
	t, err := template.New("application").Parse(tpl)

	if err != nil {
		check(err)
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data = &dotContent{
			Title:   "root",
			Content: "Root",
		}
		w.Header().Add("Content-Type", "text/html; charset=utf-8")

		err := t.Execute(w, data)
		if err != nil {
			check(err)
		}
	})
	handlerFoo := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data = &dotContent{
			Title:   "foo",
			Content: `<h1>Foo</h1>`,
		}
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		err = t.Execute(w, data)
		if err != nil {
			check(err)
		}
	})
	handlerBar := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data = &dotContent{
			Title:   "bar",
			Content: `<h1>Bar</h1>`,
		}

		w.Header().Add("Content-Type", "text/plain")

		err = t.Execute(w, data)
		if err != nil {
			check(err)
		}

	})
	http.Handle("/", handler)
	http.Handle("/foo", handlerFoo)
	http.Handle("/bar", handlerBar)
	fmt.Println("running server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//START example http

func hFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	}
}
func hPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "End point")
	}
}
func (goSchool) ExampleHttp() {
	http.Handle("/", hFound())
	http.Handle("/endpoint", hPoint())
	println("running server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//ENXD example http

/*
START
example httpnotfound
*/
func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Route: /")
}
func homeHandler() http.Handler {
	return http.HandlerFunc(home)
}
func newPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.NotFound(w, r)
		return
	}
	io.WriteString(w, "New People...")
}
func newPeopleHandler() http.Handler {
	return http.HandlerFunc(newPeople)
}
func (goSchool) ExampleHttpNotFound() {
	mux := http.NewServeMux()
	mux.Handle("/", homeHandler())
	mux.Handle("/resource", http.NotFoundHandler())
	mux.Handle("/resource/people", newPeopleHandler())
	println("running server port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

/*
END
example not found
*/

/*
example httpv3
*/
func common(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "/common")
}

type apiHandler struct {
}

func (apih apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "/api")
}

func Home(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/common" {
		common(w, req)
		return
	}
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	io.WriteString(w, "Welcome from home page")
}
func HomeHandler() http.Handler {
	return http.HandlerFunc(Home)
}
func (goSchool) ExampleHttpV3() {
	fsys := dotFileHidingFileSystem{http.Dir("./testdata")}
	mux := http.NewServeMux()
	mux.Handle("/", HomeHandler())
	mux.Handle("/api/", apiHandler{})
	mux.Handle("/ui", http.FileServer(fsys))
	println("running server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// end example httpV3

// st Tutorial http

type ApiV2Handler struct{}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It is message from other controller")
}
func (ApiV2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		{
			PostMessage(w, r)
			return
		}
	}

	io.WriteString(w, "api...")
}

func (goSchool) TutorialHttp() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(content)))
	mux.Handle("/api/", ApiV2Handler{})
	println("run server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

//end tutorial http

//=====================================

// sync example-> httpclient/main.go
func (goSchool) ExampleHttp_Sync() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Home 1")
	})
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "foo 2")
	})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "bar 3")
	})
	http.HandleFunc("/say", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "say 4")
	})
	http.HandleFunc("/fobar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "fobar 5")
	})

	println("run server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (goSchool) ExampleSyncOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Output:
	// Only once
}

// example interface
type adapterFtp interface {
	foo()
	say()
}
type adapter interface {
	foo()
}
type service struct{}

func (service) foo() {
	println("foo")
}
func (service) say() {
	println("SAY!")
}
func (goSchool) ExampleInterface() {
	var api any = service{}
	var api2 any = service{}
	var srvc = api.(adapter) //tampbien -> api.(adapterFtp) y srvc.say()
	var srvc2 = api2.(adapterFtp)
	srvc2.say()
	srvc.foo()
}

// other sync
func (goSchool) Sync() {
	var wg sync.WaitGroup
	sl := func() {
		for i := 0; i < 5; i++ {
			println("sl: ", i)
			time.Sleep(time.Second)
		}
	}
	for i := 0; i < 7; i++ {
		wg.Add(1)
		println("executando... ciclo", i)
		time.Sleep(time.Millisecond * 70)
		go func() {
			sl()
			defer wg.Done()
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)
		println("Ultimo ciclo", i)
	}
	println("exit program")
	wg.Wait()
}

/*
se encuentra contenida en la carpeta go el paquete ast
*/
func (goSchool) Ast() {
	/*src := `
	package main
	func main(){
		println("hello world")
	}
	`*/

	file, err := ioutil.ReadFile("testdata/foo.go")
	if err != nil {
		check(err)
	}
	src := string(file)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		check(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			//s = x.Name
			if x.Name == "foo" {
				fmt.Println("esta variable es igual a foo")
			}
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

}

func (goSchool) T() {

	swap := func(in []reflect.Value) []reflect.Value {
		println("Se està ejecutando una funcòn")
		switch in[0].Kind() {
		case reflect.String:
			fmt.Println("...")
		}

		return []reflect.Value{in[1], in[0]}
	}
	makeSwap := func(fptr any) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	intSwap(10, 20)

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)

	makeSwap(&floatSwap)
	floatSwap(1.2, 2.3)

	var strSwap func(string, string) (string, string)
	makeSwap(&strSwap)
	strSwap("foo", "fobar")
}

type f func()

func (f1 f) So() {
	f1()
}
func (goSchool) FuncTy() {

	//http.HandlerFunc
	//http.HandleFunc()
	//http.Handle()
	//var s = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	println("Inicia")
	var f = f(func() {
		println("Fobar")
	})
	f.So()

}

func (goSchool) Gzipv2() {
	var buf, err = os.Create("testdata/test.gzip")
	if err != nil {
		log.Fatal("error al crear el archivo ", err)
	}
	zw := gzip.NewWriter(buf)

	// Setting the Header fields is optional.
	zw.Name = "README.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err = zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	if err != nil {
		log.Fatal("error al escribir los bytes ", err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal("error al cerrar el archvo", err)
	}

	zr, err := gzip.NewReader(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}
}

func (goSchool) ExampleEncode() {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (goSchool) ExampleGithub() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	// create a new private repository named "foo"
	/*repo := &github.Repository{
		Name:    github.String("foo"),
		Private: github.Bool(false),
	}
	client.Repositories.Create(ctx, "", repo)
	*/
	//var message *string
	/*c := &github.RepositoryCommit{
		Commit: &github.Commit{
			Message: message,
		},
	}*/
	var serv github.GitService = *client.Git
	_, res, _ := serv.CreateCommit(ctx, "", "foo", &github.Commit{
		Message: github.String("Hola mundo"),
		Author: &github.CommitAuthor{
			Name:  github.String("Diego"),
			Email: github.String("diegoalejandrogutierrezs@gmail.com"),
		},
	})

	fmt.Println("status: ", res.StatusCode)
	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, r := range repos {
		println(*r.Name)
	}
}
func (goSchool) U_example() {
	go func() {
		for range time.Tick(time.Second) {
			println("fo bar")
		}
	}()
	fmt.Println("runtime")
	time.Sleep(time.Second * 4)
}

func (goSchool) Example_Bbolt() {

	println("abriendo db")
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func (goSchool) ExamplebitCask() {
	db, _ := bitcask.Open("/tmp/db")
	defer db.Close()
	db.Put([]byte("Hello"), []byte("World"))
	val, _ := db.Get([]byte("Hello"))
	log.Printf(string(val))

}

/*
func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}
*/
/*
	func (goSchool) GoPython() {
		gostr := "foo"
		pystr := python.PyString_FromString(gostr)
		str := python.PyString_AsString(pystr)
		fmt.Println("hello [", str, "]")
	}
*/
func a(c chan int, i int) {
	if i < 10 {
		i++
		a(c, i)
		time.Sleep(time.Second)
		println("a ", i)
	}
}
func b(c chan int, i int) {
	if i < 10 {
		i++
		c <- i
		b(c, i)
		time.Sleep(time.Second * 2)
		println("b ", i)
	}
}
func (goSchool) Proteus() {
	var c = make(chan int)
	var i int = 0
	a(c, i)
	go b(c, i)
	fmt.Println("recived ", <-c)
	for range time.Tick(time.Second) {
		println("runtime")
	}
}

/**


service.html
<route="/" controller="landing" middleware="islgoin">
<route="/foo" controller="login">

*/
