
var myName = ""
var myLastName = ""
var myUsername = ""
var allFollowers = []
var allFollowing = []
var generatedFollowers = false
var generatedFollowing = false
var loggedIn
var un
var postList
var postListSaved    // menja se u zavisnosti od selektovane kolekcije
var postListLiked    // lista za prikaz lajkovanih postova
var postListDisliked // lista za prikaz dislajkovanih postova
var this_is_me // username / role
var this_is_my_profile // profile of currently logged in user
var user_on_page
var whoCanISee
var currentOpenedPost

$(document).ready(function () {


    //whoAmI()
    editprofmodal()
    fetchCurrentPageUser()




    printOriginVariables()
    //checkUserPublicity()

    getMyDatas()

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
        success: function (data) {
            loggedIn = true;
            //console.log(data)
            //console.log('FETCHED POSTS FOR USER')
            postList = data
            //console.log('post list is ')
            //console.log(data)
            $('#num_of_posts').html(data.length)
            // checkUserPublicity()
        },
        error: function () {
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
        success: function (data) {
            loggedIn = true;
            //console.log(data)
            this_is_me = data

            //prebaceno ovde zbog sinhronizacije
            getMyDatas1();
            setFollowers();


            if (user_on_page.username == this_is_me.username) {
                modifyForMyProfile()
            }
            else {
                modifyForNotMyProfile()
            }
            fetchCurrentUserPosts()

        },
        error: function () {
            loggedIn = false
            fetchCurrentUserPosts()
            //IamNotLoggedIn


        }
    })
}

function modifyForMyProfile() {
    //console.log('hey this is my profile')
    //console.log('hide follow button')
    $('#follow_button').hide()
}

function modifyForNotMyProfile() {
    //console.log('not my profile')
    $('#edit_profile_button').hide()
    // disable liked disliked saved pills
    $('#liked-tab').addClass('disabled')
    $('#disliked-tab').addClass('disabled')
    $('#saved-tab').addClass('disabled')

    var visitedUsername = location.href.split("?")[1];

    checkIfUserBlockedMeAndRestrict();

    if(currentprofileObj.blacklist == undefined) {
        return
    } else {

    currentprofileObj.blacklist.forEach(function(b) {
        if(visitedUsername == b.username) {
            //console.log("profile is blocked - do not show the profile");
            showBlockedUserModal();
        }
    })
    }


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
        success: function (data) {
            loggedIn = true;
            //console.log('succesfully fetched user')
            //console.log(data)
            user_on_page = data
            if (data.biography == "")
                $('#bio').html('This user has no biography')
            else {
                $('#bio').html(data.biography)
            }
            whoAmI()
            //fetchCurrentUserPosts()

        },
        error: function () {
            //alert('Could not fetch user')
            $("#maincontainer").css("visibility", "hidden")
            $("#userdoesntexist").html("User with that username doesn't exists")

        }
    })
}

function fetchCurrentPageUser() {

    const currurl = window.location.href

    un = currurl.split('?')[1]
    //console.log(un);
    if (un == "") {
        countdownToRedirect(3, "back")
    } else {
        getAuthInfoForPageUser(un)
        fetchUser(un)
    }
}

function getAuthInfoForPageUser() {
    //console.log('getting auth info for ' + un)
    // getdata from user
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/getuser/' + un,
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            //console.log(data)
            $('#current_name').html(data.name + ' ' + data.lastname)
        },
        error: function (xhr, status, data) {
            //console.log(xhr)
            //console.log('Cant get profile data');
        }
    })

}




document.addEventListener("click", function (e) {
    if (e.target.classList.contains("gallery-item")) {

        const src = e.target.getAttribute("src");
        //console.log(src);

        document.querySelector(".modal-img").src = src;
        const myModal = new bootstrap.Modal(document.getElementById('gallery-modal'));
        const post_id = e.target.getAttribute("id")
        console.log(post_id)
        //console.log(post_id)
        //showImageModal(post_id.split("-")[1], postList)


        if (checkIfShowingPostIsAllowed()) {
            myModal.show();
        }
        else {
            alert("You must be logged in to view photo")
        }

        //Rendering unsaved button

        // todo: za sve pillove posebne liste showovati
        if (currentPill == "saved") {
            $("#removeFromCollection").show()
            var c = select_group.children(":selected").attr("id");
            $("#removeFromCollection").html("Remove " + post_id + " from " + c)
            currentOpenedPost = post_id
            showImageModal(post_id.split("-")[1], postListSaved)
        } else if (currentPill == "posts") {
            $("#removeFromCollection").hide()
            currentOpenedPost = post_id
            showImageModal(post_id.split("-")[1], postList)
        } else if (currentPill == "liked") {
            $("#removeFromCollection").hide()
            currentOpenedPost = post_id
            showImageModal(post_id.split("-")[1], postListLiked)
        } else if (currentPill == "disliked") {
            $("#removeFromCollection").hide()
            currentOpenedPost = post_id
            showImageModal(post_id.split("-")[1], postListDisliked)
        }
    }
})

function checkIfShowingPostIsAllowed() {

    return this_is_me != undefined
}

whoCanISee = "profile";

function editprofmodal() {

    $("#editprofdata").click(function () {
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

    $("#editprivacy").click(function () {
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

    $("#editnotif").click(function () {
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

    $("#muteandblock").click(function () {
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
        success: function (data) {
            loggedIn = true;
            //console.log(data)
            checkUserPublicity(name)
        },
        error: function () {
            loggedIn = false
            //IamNotLoggedIn
            checkUserPublicity(name)

        }
    })
}


async function setFollowers() {

    try {
        const res = await getMyDatas1()

    } catch (err) {
        //console.log(err)
    }

    var obj = {
        username: user_on_page.username
    }

    //console.log("USAOOO OVDEE")
    //console.log(JSON.stringify(obj))
    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getAllFollowers',
        data: JSON.stringify(obj),
        contentType: 'application/json',
        //dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            //console.log("PRVO SE OVO ZAVRSILO SET FOLLOWERS")
            // //console.log("UKUPNO FOLLOWERA: ")
            $("#numberOfFollowers").text(data.length)
            allFollowers = data
            // //console.log(data.length)
            checkIfAlreadyFollowed()    //this change Follow button to disable if user already follow some profile
            checkUserPublicity()
        },
        error: function (xhr, status, data) {
            //console.log(status)
            //console.log(xhr)
            //console.log('Cant get followers data');
        }
    })


    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getAllFollowing',
        data: JSON.stringify(obj),
        contentType: 'application/json',
        //dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            //console.log("PRVO SE OVO ZAVRSILO SET FOLLOWING")
            ////console.log("UKUPNO FOLLOWERA: ")
            // //console.log(data)
            $("#numberOfFollowing").text(data.length)
            allFollowing = data;
            ////console.log(data.length)
        },
        error: function (xhr, status, data) {
            //console.log(status)
            //console.log(xhr)
            //console.log('Cant get followers data');
        }
    })

}



function checkUserPublicity() {
    var user = user_on_page
    //alert(user.privacy_setting.is_public)
    //console.log(user)
    //console.log(this_is_me)
    if (user.username == this_is_me.username) {
        showPhotos(postList, "postsHere")
        return
    }

    if (user.privacy_setting.is_public == true) {
        console.log('profile is public')
        showPhotos(postList, "postsHere")    //AJAX POZIV ZA DOBIJANJE POSTOVA OD USERA

    } else {
        //console.log("NA DUGMETU TRENUTNO: ")
        //console.log($("#follow_button").text())
        if ($("#follow_button").text() === "Following") {   //znaci da ga vec prati
            console.log("USPEO DA UDJEM U IF")
            showPhotos(postList, "postsHere")    //AJAX POZIV ZA DOBIJANJE POSTOVA OD USERA
            return
        }
        else {
            $("#photos").css("visibility", "hidden")
            $("#userprivate").html("User is private")
        }

    }
    //$("#maincontainer").css("visibility", "hidden")
    //$("#userdoesntexist").html("User with that username doesn't exists")

}







function getMyDatas1() {
    // getdata from user
    return $.ajax({
        type: 'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/getdata',
        contentType: 'application/json',
        dataType: 'JSON',
        //async: false,
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            //console.log("PRVO SE OVO ZAVRSILO getMyDatas")

            //console.log(data)
            myUsername = data.username
            $("#myFullName").text(data.name + " " + data.lastname)
        },
        error: function (xhr, status, data) {
            //console.log(xhr)
            //console.log('Cant get profile data');
        }
    })

}
function showFollowers() {
    if (generatedFollowers)
        return

    var allFollowersHtml = ''
    allFollowersHtml += '<ul class="list-unstyled">' + '\n'

    var n = allFollowers.length
    for (let i = 0; i < n; i++) {
        //console.log("Brojac:" + i)
        var username = allFollowers[i].username
        //var fullName = allFollowers[i].name + " " + allFollowers[i].lastname
        var fullName = ""
        //console.log(username)

        allFollowersHtml += '<li class="media">' + '\n'
        allFollowersHtml += '<img src="./img/avatar.png" width="35" height="35" class="mr-3" alt="...">' + '\n'
        allFollowersHtml += '<div class="media-body">' + '\n'
        allFollowersHtml += '<h5 class="mt-0 mb-1"><a href="profile.html?' + username + '">' + username + '</a>' + '</h5>' + '\n'
        allFollowersHtml += '<p>' + fullName + '</p>' + '\n'
        allFollowersHtml += '</div>' + '\n'
        allFollowersHtml += '<span class="btnFollow">' + '\n'
        allFollowersHtml += '<button type="button"' + '\n'
        allFollowersHtml += 'class="btn btn-primary btn-lg btnFollow">Follow</button>' + '\n'
        allFollowersHtml += '</span></li><hr>'

    }

    if (n != 0)
        allFollowersHtml = allFollowersHtml.slice(0, -4)

    //console.log(allFollowersHtml)

    allFollowersHtml += '</ul>'

    $("#insertAllFollowers").after(allFollowersHtml)
    generatedFollowers = true;

}

function showFollowing() {
    if (generatedFollowing)
        return

    var allFollowersHtml = ''
    allFollowersHtml += '<ul class="list-unstyled">' + '\n'

    var n = allFollowing.length
    for (let i = 0; i < n; i++) {
        //console.log("Brojac:" + i)
        var username = allFollowing[i].username
        //var fullName = allFollowing[i].name + " " + allFollowing[i].lastname
        var fullName = ""
        //console.log(username)

        allFollowersHtml += '<li class="media">' + '\n'
        allFollowersHtml += '<img src="./img/avatar.png" class="mr-3" width="35" height="35" alt="...">' + '\n'
        allFollowersHtml += '<div class="media-body">' + '\n'
        allFollowersHtml += '<h5 class="mt-0 mb-1"><a href="profile.html?' + username + '">' + username + '</a>' + '</h5>' + '\n'
        allFollowersHtml += '<p>' + fullName + '</p>' + '\n'
        allFollowersHtml += '</div>' + '\n'
        allFollowersHtml += '<span class="btnFollow">' + '\n'
        allFollowersHtml += '<button type="button"' + '\n'
        allFollowersHtml += 'class="btn btn-primary btn-lg btnFollow">Follow</button>' + '\n'
        allFollowersHtml += '</span></li><hr>'

    }
    if (n != 0)
        allFollowersHtml = allFollowersHtml.slice(0, -4)

    //console.log(allFollowersHtml)

    allFollowersHtml += '</ul>'

    $("#insertAllFollowing").after(allFollowersHtml)
    generatedFollowing = true;
}

$("#follow_button").click(function () {
    //console.log("TREBA DA ZAPRATIM: ")
    //console.log(user_on_page.username)

    var obj = {
        followToUsername: user_on_page.username
    }

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/follow',
        contentType: 'application/json',
        //dataType: 'JSON',
        data: JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            //console.log("USPESNO ZAPRAĆENO")
            $("#follow_button").text("SENT REQUEST")
        },
        error: function () {
            alert("You already sent follow request to this account!")
            //console.log("Erorr at ajax call FOLLOW!!!")


        }
    })



})



function checkIfAlreadyFollowed() {
    //console.log("USAO U FUNKCIJA CHECK IF ALREADY!!!")
    //console.log("DUZINA: ", allFollowers.length)

    for (var i = 0; i < allFollowers.length; i++) {
        //console.log(allFollowers[i].username)
        //console.log(this_is_me.username)
        //console.log("*******************************")
        if (allFollowers[i].username === this_is_me.username) {
            //console.log("NASAO SAM SEBE!!!")
            $('#follow_button').text("Following")
            $('#follow_button').attr("disabled", 'true')
            return
        }
    }

}

function showBlockedUserModal() {
    $('#blockedUserModal').modal('show'); 
}

function redirectOnPrevPage() {
    history.go(-1);
}

function checkIfUserBlockedMeAndRestrict() {

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getuserwhoblockedme',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            console.log(data)
            checkIfImInHisBlacklist(data)
        },
        error: function () {
            alert("Error in getUsersWhoBlockedMeAndRestricTheirProfile")
        }
    })

}

function checkIfImInHisBlacklist(data) {
    data.listwhoblockedme.forEach(function(p) {

        var visitedUsername = location.href.split("?")[1];
        //console.log(p)

        if(visitedUsername == p) {
            $("#maincontainer").css("visibility", "hidden")
            $("#userdoesntexist").html("User with that username doesn't exists")
        }
    })
}

//Dropdown sa nazivima kolekcija
//Klik na tab, poziva se da se vrati lista naziva kolekcija(profile service), popuniti dropdown (1. vrednost prazna)
//Kad se odabere stavka iz dd, 