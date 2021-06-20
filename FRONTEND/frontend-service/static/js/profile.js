
var myName = ""
var myLastName = ""
var myUsername = ""
var allFollowers = []
var allFollowing = []
var generatedFollowers = false
var generatedFollowing = false
var loggedIn

var postList
var this_is_me
var user_on_page

$(document).ready(function() {

    setFollowers()
    getMyDatas1();
  
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
        //countdownToRedirect(3, "back")
        console.log("....nestoo")
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


async function setFollowers(){

    try{
        const res = await getMyDatas1()

    }catch(err){
        console.log(err)
    }

     var obj = {
        username : myUsername
    }

    console.log("USAOOO OVDEE")
    console.log(JSON.stringify(obj))
    $.ajax({
        type:'POST',
        crossDomain: true,
        url: 'http://localhost:3030/getAllFollowers',
        data : JSON.stringify(obj),
        contentType : 'application/json',
        //dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
           // console.log("UKUPNO FOLLOWERA: ")
            $("#numberOfFollowers").text(data.length) 
            allFollowers = data
           // console.log(data.length)
        },
        error : function(xhr, status, data) {
            console.log(status)
            console.log(xhr)
            console.log('Cant get followers data');
        }
    })


    $.ajax({
        type:'POST',
        crossDomain: true,
        url: 'http://localhost:3030/getAllFollowing',
        data : JSON.stringify(obj),
        contentType : 'application/json',
        //dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            //console.log("UKUPNO FOLLOWERA: ")
           // console.log(data)
            $("#numberOfFollowing").text(data.length) 
            allFollowing = data;
            //console.log(data.length)
        },
        error : function(xhr, status, data) {
            console.log(status)
            console.log(xhr)
            console.log('Cant get followers data');
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







function getMyDatas1() {
    // getdata from user
    return $.ajax({
        type:'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/getdata',
        contentType : 'application/json',
        dataType: 'JSON',
        //async: false,
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            console.log(data)
            myUsername = data.username
            $("#myFullName").text(data.name + " " + data.lastname) 
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Cant get profile data');
        }
    })

}
function showFollowers(){
    if(generatedFollowers)
        return

    var allFollowersHtml = ''
    allFollowersHtml += '<ul class="list-unstyled">' + '\n'

    var n = allFollowers.length
    for(let i = 0; i<n; i++){
        console.log("Brojac:" + i)
        var username = allFollowers[i].username
        //var fullName = allFollowers[i].name + " " + allFollowers[i].lastname
        var fullName = ""
        console.log(username)

        allFollowersHtml += '<li class="media">' + '\n'
        allFollowersHtml += '<img src="./img/avatar.png" width="35" height="35" class="mr-3" alt="...">' + '\n'
        allFollowersHtml += '<div class="media-body">' + '\n'
        allFollowersHtml += '<h5 class="mt-0 mb-1">' + username + '</h5>' + '\n'
        allFollowersHtml += '<p>' + fullName + '</p>' + '\n'
        allFollowersHtml +=  '</div>' + '\n'
        allFollowersHtml += '<span class="btnFollow">' + '\n'
        allFollowersHtml += '<button type="button"' + '\n'
        allFollowersHtml += 'class="btn btn-primary btn-lg btnFollow">Follow</button>' + '\n'
        allFollowersHtml +=  '</span></li><hr>'
        
    }

    if(n!=0)
        allFollowersHtml = allFollowersHtml.slice(0, -4)

    console.log(allFollowersHtml)   

    allFollowersHtml += '</ul>'

    $("#insertAllFollowers").after(allFollowersHtml)
    generatedFollowers = true;

}

function showFollowing(){
    if(generatedFollowing)
        return

    var allFollowersHtml = ''
    allFollowersHtml += '<ul class="list-unstyled">' + '\n'

    var n = allFollowing.length
    for(let i = 0; i<n; i++){
        console.log("Brojac:" + i)
        var username = allFollowing[i].username
        //var fullName = allFollowing[i].name + " " + allFollowing[i].lastname
        var fullName = ""
        console.log(username)

        allFollowersHtml += '<li class="media">' + '\n'
        allFollowersHtml += '<img src="./img/avatar.png" class="mr-3" width="35" height="35" alt="...">' + '\n'
        allFollowersHtml += '<div class="media-body">' + '\n'
        allFollowersHtml += '<h5 class="mt-0 mb-1">' +  username + '</h5>' + '\n'
        allFollowersHtml += '<p>' + fullName + '</p>' + '\n'
        allFollowersHtml +=  '</div>' + '\n'
        allFollowersHtml += '<span class="btnFollow">' + '\n'
        allFollowersHtml += '<button type="button"' + '\n'
        allFollowersHtml += 'class="btn btn-primary btn-lg btnFollow">Follow</button>' + '\n'
        allFollowersHtml +=  '</span></li><hr>'
        
    }
    if(n!=0)
        allFollowersHtml = allFollowersHtml.slice(0, -4)

    console.log(allFollowersHtml)   

    allFollowersHtml += '</ul>'

    $("#insertAllFollowing").after(allFollowersHtml)
    generatedFollowing = true;

}


