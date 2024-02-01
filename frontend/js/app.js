//testing login post request
async function login() {
    let data = {
        'username': 'useri41024',
        'password': "n09cqnncqnwd"
    }

    let response = await fetch("http://127.0.0.1:8080/login", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json;"
        }
    }).then((response => response.json()))

    //display data to frontend
    document.getElementById("username").innerText = response.username
}