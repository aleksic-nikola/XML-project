var blokcmutebtn = $("#blockmutePart"); 

$(document).ready(function () {
    whoAmI2();

    // po defaultu je sakriven, pa kad je korisnik blokiran onda se prikazuje ovo dugme (checkIfUserIsMutedOrBlocked() -->u editprof.js je poziv)
    $("#umuteuserBtn").hide();
})

function whoAmI2() {
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
            doNotShowBlockMuteOnMyProfile(data)
        },
        error: function () {
            console.log("no one is logged in")
            blokcmutebtn.hide();
        }
    })
}

function doNotShowBlockMuteOnMyProfile(data) {
    var visitedprofUsername = location.href.split("?")[1]

    //console.log(visitedprofUsername)
    //console.log(data.username)
    if(data.username == visitedprofUsername) {
        blokcmutebtn.hide();
    } else {
        blokcmutebtn.show();
    }

}

function blockUser() {
    var usernameToBlock = location.href.split("?")[1]
    var obj = {
        "usernametoblockmute" : usernameToBlock
    }

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/blockuser',
        contentType: 'application/json',
        data: JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            alert('Succesfully blocked user: ' + usernameToBlock)
            window.location.reload();

        },
        error: function () {
            alert('Error while blocking user: ' + usernameToBlock)
        }
    })

}

function muteUser() {
    var usernameToMute = location.href.split("?")[1]

    var obj = {
        "usernametoblockmute" : usernameToMute
    }

    console.log("=======================")
    console.log(obj)

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/muteuser',
        contentType: 'application/json',
        data: JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            alert('Succesfully blocked user: ' + usernameToMute)
            window.location.reload();

        },
        error: function () {
            alert('Error while blocking user: ' + usernameToMute)
        }
    })

}


function checkIfUserIsMutedOrBlocked() {

    var visitedUser = location.href.split("?")[1]

    var mutedlist = currentprofileObj.graylist
    var blockedlist = currentprofileObj.blacklist

    console.log("MUTEDS: ")
    console.log(currentprofileObj.graylist)
    
    console.log("BLOCKLIST: ")
    console.log(currentprofileObj.blacklist)


    mutedlist.forEach(function(u) {
        
        if(visitedUser == u.username) {
            console.log("USER IS ALREADY MUTED")

            $("#muteuserBtn").hide();
            $("#umuteuserBtn").show();

        } else {

            $("#umuteuserBtn").hide();
            $("#muteuserBtn").show();

        }

    })

    blockedlist.forEach(function(u) {
        if(visitedUser == u.username) {
            console.log("USER IS ALREADY BLOCKED")
        }

    })

}

function unmuteUser() {
    var usernameToUnmute = location.href.split("?")[1]

    var obj = {
        "usernametoblockmute" : usernameToUnmute
    }

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/unmuteuser',
        contentType: 'application/json',
        data: JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            alert('Succesfully umuted user: ' + usernameToUnmute)
            window.location.reload();

        },
        error: function () {
            alert('Error while unmuting user: ' + usernameToUnmute)
        }
    })
    
}
