
window.onscroll = function () { window.scrollTo(0, 0); }; // disable scrolling

async function lock() {
    try {
        var img = document.getElementById('img');
        const response = await fetch('lock');
        if (response.status == 500) {
            response.text().then(function (text) {

                console.log(text)
                document.getElementById("Error").innerHTML = text;
                document.getElementById("Error").style.display = '';

                throw new Error(text)
            });

        } else {
            document.getElementById("Error").innerHTML = "";
            document.getElementById("Error").style.display = 'none';
        }
        img.src = 'dynamic/img.png?rand=' + Math.random(); //rand to prevent caching
        console.log(new Date())
    } catch (err) {
        console.log(err);
    }
    finally {
        setTimeout(lock, 1)

    }
}


document.addEventListener("visibilitychange", (event) => {
    if (document.visibilityState == "visible") {
        location.reload(); //realod when returning to the tab
    } else {
        // console.log("tab is inactive")
    }
});

lock()

