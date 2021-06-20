var loggedIn

var post1 = {ID:"1",postedby:"dparip",timestamp:"2020-01-01T00:00:00Z",description:"opdaddsis",
	comments:[{
		postedby:  "lucyxz",
		text:      "some text here",
		timestamp : 'dont care',
	},
	{
		postedby:  "ahaheagha",
		text:      "someeayaeyaeyaweyea text here",
		timestamp : 'dont care',
	}],
	medias:[{type:"IMAGE",path:"../img/avatar.png",location:{country:"Srbija",city:"BT"}}],
	likes:[{Username:"gfgf"}],
	dislikes:[{Username:"cas"},{Username:"aaaa"}]}

var post2 = {ID:"2",postedby:"dparip",timestamp:"2020-01-01T00:00:00Z",description:"opdaddsis",comments:[
	{
		postedby:  "wintzyxz",
		text:      "bravo text here",
		timestamp : 'dont care',
	},
	{
		postedby:  "danip",
		text:      "some agaehahazhawe here",
		timestamp : 'dont care',
	}
	],
medias:[{type:"IMAGE",path:"../img/phones_image.png",location:{country:"Srbija",city:"BT"}}],likes:
[{Username:"gfgf"}],
dislikes:[{Username:"cas"},{Username:"aaaa"}]}

var postList
var this_is_me
var user_on_page

$(document).ready(function() {

    whoAmI()
    editprofmodal()
    fetchCurrentPageUser()
    printOriginVariables()
    //checkUserPublicity()
})

function fetchCurrentUserPosts() {

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/getpostsbyuser/' + user_on_page.username,
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            loggedIn = true;
            console.log(data)
            postList = data
            console.log('post list is ')
            console.log(data)
            checkUserPublicity()
        },
        error : function() {
            loggedIn = false
            //IamNotLoggedIn
            
            
        }
    })

}

function whoAmI() {
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/whoami',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            loggedIn = true;
            console.log(data)
            this_is_me = data
        },
        error : function() {
            loggedIn = false
            //IamNotLoggedIn
            
            
        }
    })
}

function fetchUser(username) {
    
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getuser/' + username,
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            loggedIn = true;
            console.log('succesfully fetched user')
            console.log(data)
            user_on_page = data
            fetchCurrentUserPosts()
            
        },
        error : function() {
            //alert('Could not fetch user')
            $("#maincontainer").css("visibility", "hidden")
            $("#userdoesntexist").html("User with that username doesn't exists")
            
        }
    })
}

function fetchCurrentPageUser() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const name = urlParams.get('user')
    console.log(name);
    if (name == null) {
        countdownToRedirect(3, "back")
    } else {
        fetchUser(name)
    }
}


document.addEventListener("click", function(e) {
    if(e.target.classList.contains("gallery-item")) {

        const src = e.target.getAttribute("src");
            console.log(src);

        document.querySelector(".modal-img").src = src;
        const myModal = new bootstrap.Modal(document.getElementById('gallery-modal'));
        const post_id = e.target.getAttribute("id")
        console.log(post_id)
        showImageModal(post_id.split("-")[1], postList)
        

        if (checkIfShowingPostIsAllowed()) {
            myModal.show();
        }
        else {
            alert("You must be logged in to view photo")
        }
    }
})

function checkIfShowingPostIsAllowed() {

    return this_is_me != undefined
}

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
        url: AUTH_SERVICE_URL + '/whoami',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            loggedIn = true;
            console.log(data)
            checkUserPublicity(name)
        },
        error : function() {
            loggedIn = false
            //IamNotLoggedIn
            checkUserPublicity(name)
            
        }
    })
}

function checkUserPublicity() {
            var user = user_on_page
            alert(user.privacy_setting.is_public)
            if (user.privacy_setting.is_public == true) {
                console.log('profile is public')
                showPhotosForNonLoggedInUser(postList)    //AJAX POZIV ZA DOBIJANJE POSTOVA OD USERA
            } else {
                $("#photos").css("visibility", "hidden")
                $("#userprivate").html("User is private")
            }     
            //$("#maincontainer").css("visibility", "hidden")
            //$("#userdoesntexist").html("User with that username doesn't exists")

}

/*
function getMyDatas() {
    // getdata from user
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/getdata',
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
        url: PROFILE_SERVICE_URL + '/getdata',
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

