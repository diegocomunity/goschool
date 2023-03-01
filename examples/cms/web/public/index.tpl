<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CMS App</title>
</head>
<style>
    .run {
        margin-left: auto;
        margin-right: auto;
        border-radius: 25%;
        width: 25px;
        height: 25px;
        background-color: black;
        box-shadow: 4px 3px 3px 3px purple;
    }
</style>
<body>
    <header class="header">
        <form id="optionForm">
            <label for="lblOption">
                Option 
                <input type="text" name="option">
            </label>
            <label for="lblId">
                Id proyet
                <input type="text" name="id">
            </label>
            <input type="submit" name="" id="">
        </form>
        <nav class="menu">
            <ul><li><a href="/">/</a></li></ul>
            <ul><li><a href="/projets">/proyets</a></li></ul>
        </nav>
    </header>
    <main class="main">
        <form id="projetForm">
            <label for="lblOption">
                Option
            <input type="search" name="option">
            </label>
            <label for="lblIdProjet">
                Id del proyecto
                <input type="text" name="id">
            </label>
            <label for="lblNameProject">
                Nombre del proyecto web
                <input type="text" name="name">
            </label>
            <label for="lblDescription">
                Descipciòn 
                <textarea name="description" cols="30" rows="10"></textarea>
            </label>
            <section class="desing">
                <label for="lblMenu">
                    Menu
                    <textarea name="menu" cols="30" rows="10"></textarea>
                </label>
            </section>
            <input type="submit" value="BUILD">
        </form>
    </main>
    <footer class="footer">FOOTER</footer>

    <script>
        const projetForm = document.getElementById("projetForm");
	    projetForm.addEventListener("submit", (e) => {
		e.preventDefault()
		const data = new FormData(document.getElementById("projetForm"));
		fetch('/api/projet/build', {
			body: data,
			method: 'POST'
		}).then(res => res.text())
		.then(res => {
			console.log(res)
		})
	})	
    const optionForm = document.getElementById("optionForm")
    optionForm.addEventListener("submit", (e) => {
		e.preventDefault()
		const data = new FormData(document.getElementById("optionForm"));
		fetch('/api/projet/build', {
			body: data,
			method: 'POST'
		}).then(res => res.text())
		.then(res => {
            document.querySelector("body").innerHTML = res
		})
	})
    </script>
</body>
</html>