((c) => {
    const get_exampleFormWithId = document.getElementById("exampleForm");
    get_exampleFormWithId.addEventListener("submit", (e) => {
        e.preventDefault()
        const data = new FormData(document.getElementById('exampleForm'));
        fetch('/foo', {
            body: data,
            method: 'POST'
        }).then(res => res.text())
        .then(res => {
                c(res)
            })
    })
})(console.log)