<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Site Creator</title>
    <link rel="stylesheet" href="/css/global">
</head>
<body>
    <header class="header" id="header">
        <nav class="nav menu">
            <ul class="menu--items">
                <li class="item"><span class="icon">LOGO</span></a></li>
                <li class="item"><a href="/" class="link">HOME</a></li>
                <li class="item"><a href="#" class="link">Projects</a></li>
            </ul>
        </nav>
    </header>
    <main class="main" id="workspace">
        <section class="site--creator">
            <form id="siteCreatorForm" class="site-creator">
                <fieldset>
                    <legend title="site creator">Sites creator</legend>
                </fieldset>
                <label for="lblName">
                    Nombre del proyecto *
                    <input type="text" name="name">
                </label>
                <label for="lblTitle">
                    Title
                    <input type="text" name="title">
                </label>
                <label for="lblCode">
                    Code
                    <textarea name="code"cols="30" rows="10"></textarea>
                </label>
                <input type="submit">
            </form>
            <form id="runSite">
                <fieldset title="run project">
                    <legend>Correr proyecto</legend>
                </fieldset>
                <input type="submit">
            </form>
        </section>
        <!--salida del proyecto-->
        <div id="output"></div>
    </main>
    <footer class="footer">
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Ut iure, consectetur ratione accusantium eaque, tempore quod placeat numquam quia harum reprehenderit blanditiis ipsam facere non enim maiores aliquam, at aspernatur!</p>
    </footer>
    <script src="/js/sitecreator"></script>
    <script src="/js/run"></script>
</body>
</html>