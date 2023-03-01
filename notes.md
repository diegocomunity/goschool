bbolt cli
bbolt [command] -h

# Tutorials GO
Una lista de ejemplo de funciones y paquetes en go, fàcil de entender 

en la carpeta commons estàn algunos scripts ya sea para cliente http o para rpc etc


token github:
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


ejemplo de codigo de servidor
	directorio := "./static"
	http.Handle("/", http.FileServer(http.Dir(directorio)))
	direccion := ":8000"
	log.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))