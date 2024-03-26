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

//will be used later
async function getItems() {
    let items = []

    document.getElementById("item").src = "/files/images/iphone 15 pro natural titanium.jpg"
    await fetch('http://127.0.0.1:8080/items', {
        headers: {
            'Accept': 'application/json'
        }
    })
    .then(response => response.text())
    .then(response => items = JSON.parse(response))

    console.log("All items:", items)

    // Assuming items is an array and each item has an id property
    items.forEach(item => {
        // Create a new HTML element for each item
        var itemElement = document.createElement('div');
        itemElement.innerHTML = item.name; // Replace with your item property
        itemElement.onclick = function() {
            // Redirect to the item details page with the item ID as a URL parameter
            window.location.href = "/files/html/itemDetails.html?itemId=" + item.id;
        }

        // Add the new element to the DOM
        document.body.appendChild(itemElement);
    });
}



// will be used later
async function checkout() {
    let data = {
        'item': document.getElementById("item").src,
        'quantity': document.getElementById("quantity").value
    }

// update the url to the correct one
    try {
        let response = await fetch("http://example.com/checkout", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json;"
        }
        }).then((response => response.json()))
        alert("Checkout successful.")
    } catch {
        alert("Couldn't complete checkout.")
    }
}



// this will run when the itemDetails page loads
window.onload = function() {
    // Get the item ID from the URL
    var urlParams = new URLSearchParams(window.location.search);
    var itemId = urlParams.get('itemId');
  
    // Fetch the item details (you will need to implement this)
    var itemDetails = fetchItemDetails(itemId);
  
    // Display the item details on the page
    var itemDetailsDiv = document.getElementById('itemDetails');
    itemDetailsDiv.innerHTML = JSON.stringify(itemDetails);
}