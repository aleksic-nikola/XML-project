
var post1 = {postedby:"dparip",timestamp:"2020-01-01T00:00:00Z",description:"opdaddsis",comments:[],
            medias:[{type:"IMAGE",path:"../img/avatar.png",location:{country:"Srbija",city:"BT"}}],
            likes:[{Username:"gfgf"}],
            dislikes:[{Username:"cas"},{Username:"aaaa"}]}

var post2 = {postedby:"dparip",timestamp:"2020-01-01T00:00:00Z",description:"opdaddsis",comments:[],
            medias:[{type:"IMAGE",path:"../img/phones_image.png",location:{country:"Srbija",city:"BT"}}],likes:
            [{Username:"gfgf"}],
            dislikes:[{Username:"cas"},{Username:"aaaa"}]}

var postList = [post1, post2]

var loggedIn

$(document).ready(function() {

    editprofmodal();
    myProfileOrOther();

})

function myProfileOrOther() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const name = urlParams.get('user')
    console.log(name);
    if (name == null) {
        getMyDatas()
    } else {
        getOtherProfile(name)
    }
}

document.addEventListener("click", function(e) {
    if(e.target.classList.contains("gallery-item")) {

        if (loggedIn == true) {

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
        else {
            alert("You must be logged in to view photo")
        }
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

function getOtherProfile(name) {
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: 'http://localhost:9090/whoami',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            loggedIn = true;
            console.log(data)
        },
        error : function() {
            loggedIn = false
            //IamNotLoggedIn
            $.ajax({
                type: 'GET',
                crossDomain: true,
                url: 'http://localhost:3030/isuserpublic/' + name,
                contentType: 'application/json',
                dataType: 'JSON',
                success : function(user) {
                    alert(user.ispublic)
                    if (user.ispublic == true) {
                        showPhotos()    //AJAX POZIV ZA DOBIJANJE POSTOVA OD USERA
                    } else {
                        $("#photos").css("visibility", "hidden")
                        $("#userprivate").html("User is private")
                    }
                }, 
                error : function() {                    
                    $("#maincontainer").css("visibility", "hidden")
                    $("#userdoesntexist").html("User with that username doesn't exists")
                    countdownToRedirect(3, "back")
                }
            })
        }
    })
}

function showPhotos() {

    var posts = "";

    postList.forEach(function(post) {

        posts += `<div class="col-md-6 col-lg-4">
                        <div class="card border-0 transform-on-hover">
                            <a class="lightbox">
                                <img src="${post.medias[0].path}" alt="testimg" class="card-img-top gallery-item">
                            </a>
                        </div>
                  </div>`
    })

    $("#postsHere").html(posts) 
}

/*
function getMyDatas() {
    // getdata from user
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

    // getdata from profile
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
*/

