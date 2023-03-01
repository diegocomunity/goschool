<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <header class="header">
        <nav class="nav">
           <!-- {{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}-->
            <ul>
                {{range .Items}}
                <li class="nav--items">
                    <a href="#">{{.}}</a>
                </li>
                {{else}}
                <div></div>
                {{end}}
            </ul>
        </nav>
    </header>
</body>
</html>