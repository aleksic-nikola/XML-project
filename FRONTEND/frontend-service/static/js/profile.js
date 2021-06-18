$(document).ready(function() {

    editprofmodal();

})




document.addEventListener("click", function(e) {
    if(e.target.classList.contains("gallery-item")) {
        const src = e.target.getAttribute("src");
        console.log(src);

        document.querySelector(".modal-img").src = src;
        const myModal = new bootstrap.Modal(document.getElementById('gallery-modal'));

        /*
        NAPOMENA:  ****************
            - dodati da se i opis, komentari... se menjaju a ne samo slika.
        */

        myModal.show();
    }
})


var whoCanISee = "profile";

function editprofmodal() {

    $( "#editprofdata" ).click(function() {
        whoCanISee = "profile"

        $(this).addClass("active");
        $("#editprivacy").removeClass("active");
        $("#editnotif").removeClass("active");

        $('#editprofbody').show();
        $('#privacysettingsbody').hide();
        $('#notificationpropertiesbody').hide();
    });

    $( "#editprivacy" ).click(function() {
        whoCanISee = "privacy"

        $(this).addClass("active");
        $("#editprofdata").removeClass("active");
        $("#editnotif").removeClass("active");

        $('#editprofbody').hide();
        $('#privacysettingsbody').show();
        $('#notificationpropertiesbody').hide();
    });

    $( "#editnotif" ).click(function() {
        whoCanISee = "notification"

        $(this).addClass("active");
        $("#editprivacy").removeClass("active");
        $("#editprofdata").removeClass("active");

        $('#editprofbody').hide();
        $('#privacysettingsbody').hide();
        $('#notificationpropertiesbody').show();
    });

}