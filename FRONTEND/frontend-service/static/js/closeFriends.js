function getCloseFriends() {
    currentPill = "closeFriends"
    var username = location.href.split("?")[1];


    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getCloseFriendsByUsername/' + username,
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            console.log("DOBIO::")
            console.log(data)
            fillCloseFriends(data)
        },
        error: function () {
            console.log("Erorr at ajax call!")
        }
    })



}

function fillCloseFriends(data) {

    var htmlCloseFriends = ''
    data.forEach(function (f) {
        htmlCloseFriends += `
            <tr>
               <td class="cf_td_th><a href="profile.html?${f.username}>${f.username}</a></td>
                <td class="cf_td_th">
                <button type="button" id="${f.username}" onclick="removeFromCloseFriend(this.id)" class="btn btn-danger">Delete</button>

                </td>
            </tr>
        
        `
    })

    $("#close_body").html("")
    $("#close_body").html(htmlCloseFriends)


}


function removeFromCloseFriend(username) {

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/removeProfileFromCloseFriends/' + username,
        contentType: 'application/json',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            alert("Successfully removed from close friends!")
            getCloseFriends()
        },
        error: function () {
            console.log("Erorr at ajax call!")
            alert("cant remove")
        }
    })


}

function addToCloseFriends() {
    var myUsername = location.href.split("?")[1];
    var usernameForCF = $("#input_cf").val()

    if (usernameForCF == "") {
        return
    }
    else if (myUsername == usernameForCF) {
        alert("Cant add yourself to close friend")
        return
    }

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/addProfileToCloseFriends/' + usernameForCF,
        contentType: 'application/json',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            alert("Successfully added to close friends!")
            $("#input_cf").val("")
            getCloseFriends()
        },
        error: function () {
            console.log("Erorr at ajax call!")
            alert("User not found")
        }
    })





}