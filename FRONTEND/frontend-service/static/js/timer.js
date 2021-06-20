function countdownToRedirect(x, page) {
    var counter = x;
    var interval = setInterval(function() {
        counter--;
        // Display 'counter' wherever you want to display it.
        if (counter < 1) {
            // Display a login box
            if (page == "back") {
                history.back();
            }
            window.location.replace(page);
        }

    }, 1000);
}