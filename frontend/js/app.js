//post request for loginpage.html
async function login() {
    let data = {
        'username': document.getElementById("username").value,
        'password': document.getElementById("password").value
    }

    try {
        let response = await fetch("http://127.0.0.1:8080/login", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json;"
        }
        }).then((response => response.json()))
        alert("Logged in successfully.")
    } catch {
        alert("Couldn't login. Make sure you type in the correct credentials.")
    }
}

//post request for registerpage.html
async function registerAccount() {
    let data = {
        'username': document.getElementById("username").value,
        'password': document.getElementById("password").value
    }

    try {
        let response = await fetch("http://127.0.0.1:8080/register", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json;"
        }
        }).then((response => response.json()))
        alert("Registered account successfully.")
    } catch {
        alert("Couldn't register account.")
    }
}