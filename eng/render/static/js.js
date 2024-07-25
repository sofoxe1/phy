
window.onscroll = function () { window.scrollTo(0, 0); }; // disable scrolling

async function show_error(response) {
    if (response.status == 500) {
        response.text().then(function (text) {

            console.log(text)
            document.getElementById("Error").innerHTML = text;
            document.getElementById("Error").style.display = '';

            throw new Error(text)
        });
    }
}


async function update_parameters() {
    await fetch('config', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "Width": window.innerWidth,
            "Height": window.innerHeight

        })
    }).then((response) => response.json())
        .then((json) => console.log(json));
    await show_error(response)
}

async function update() {
    try {

        var img = document.getElementById('img');
        img.src = '/img.png?rand=' + Math.random(); //rand to prevent caching
    } catch (err) {
        console.log(err);
    }
    finally {
        setTimeout(update, 30)

    }
}


document.addEventListener("visibilitychange", (event) => {
    if (document.visibilityState == "visible") {
        location.reload(); //realod when returning to the tab
    } else {
        // console.log("tab is inactive")
    }
});
update_parameters()
update()

