package commons

import (
	"fmt"
	"net/http"
)

func CSS_Global(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	fmt.Fprint(w, `
	*{
		box-sizing: border-box;
	}
	:root{
		--bg-color: #666;
		--link-color:  rgb(173, 39, 113);
		--bg-seccundary: #333;
		--text-color: white;
	}
	fieldset, legend, button {
		border: 0;
	}
	ul li {
		list-style: none;
	}
	a {
		font-size: 1.3rem;
		text-decoration: none;
	}

	.link {
		color:var(--link-color);
	}
	html, body {
		margin: 0;
		font-family: sans-serif;
	}
	#header {
		padding: 5px;
		background-color: #333;
		margin-bottom: 20px;
	}
	.menu--items {
		display: flex;
		justify-content: space-evenly;

	}
	#workspace {
		display: grid;
		grid-template-columns: 1fr 1fr;
	}
	.site--creator {
		background-color: var(--bg-seccundary);
	}
	.site-creator {
		display: flex;
		flex-direction: column;
	}
	#output{
		background-color: white;
	}
	`)
}
func CSS_GlobalHandler() http.Handler {
	return http.HandlerFunc(CSS_Global)
}
