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
        $("#muteandblock").removeClass("active");

        $('#editprofbody').show();
        $('#privacysettingsbody').hide();
        $('#notificationpropertiesbody').hide();
        $('#muteblockbody').hide();
    });

    $( "#editprivacy" ).click(function() {
        whoCanISee = "privacy"

        $(this).addClass("active");
        $("#editprofdata").removeClass("active");
        $("#editnotif").removeClass("active");
        $("#muteandblock").removeClass("active");

        $('#editprofbody').hide();
        $('#privacysettingsbody').show();
        $('#notificationpropertiesbody').hide();
        $('#muteblockbody').hide();
    });

    $( "#editnotif" ).click(function() {
        whoCanISee = "notification"

        $(this).addClass("active");
        $("#editprivacy").removeClass("active");
        $("#editprofdata").removeClass("active");
        $("#muteandblock").removeClass("active");

        $('#editprofbody').hide();
        $('#privacysettingsbody').hide();
        $('#muteblockbody').hide();
        $('#notificationpropertiesbody').show();
 
    });

    $( "#muteandblock").click(function() {
        whoCanISee = "muteblock"
        $(this).addClass("active");
        $("#editprivacy").removeClass("active");
        $("#editprofdata").removeClass("active");
        $("#editnotif").removeClass("active");

        $('#editprofbody').hide();
        $('#privacysettingsbody').hide();
        $('#notificationpropertiesbody').hide();

        $('#muteblockbody').show();

    })

}

function getMyDatas() {
    
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: 'http://localhost:9090/getdata',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            console.log(data)
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Cant get profile data');
        }
    })


    $.ajax({
        type:'GET',
        crossDomain: true,
        url: 'http://localhost:3030/getdata',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            console.log(data)
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Cant get user data');
        }
    })

    
}