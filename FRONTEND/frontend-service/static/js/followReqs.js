$(document).ready(function () {
    //alert("CONNECTED")
    loadAllFollowReqs()
    setMyProfileHREF();
})

var myUserName = ""
var isGeneratedFollowReqs = false;

function loadAllFollowReqs() {
    console.log("USAOO")
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: REQUEST_SERVICE_URL + '/followReqs/getMy',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
            console.log("OVO SALJEM: " + localStorage.getItem('myToken'))
        },
        success: function (data) {
            console.log(data)
            generateAllFollowReqsHTML(data)
        },
        error: function () {
            console.log("Erorr at ajax call!")


        }
    })

}


function generateAllFollowReqsHTML(allReqs) {
    if (isGeneratedFollowReqs)
        return

    var allReqsHtml = ''
    allReqsHtml += '<ul class="list-unstyled">' + '\n'


    var n = allReqs.length
    var numActiveReqs = n
    for (let i = 0; i < n; i++) {

        if (allReqs[i].request.RequestStatus != 0) {
            numActiveReqs--;
            continue
        }

        if (myUserName == "") {
            myUserName = allReqs[i].forWho
        }

        var username = allReqs[i].request.sentby
        //var fullName = allFollowers[i].name + " " + allFollowers[i].lastname
        var fullName = ""
        console.log(username)

        allReqsHtml += '<li class="media">' + '\n'
        allReqsHtml += '<img src="./img/avatar.png" width="40" height="40" class="mr-3" alt="...">' + '\n'
        allReqsHtml += '<div class="media-body">' + '\n'
        allReqsHtml += '<h5 class="mt-0 mb-1">' + username + '</h5>' + '\n'
        allReqsHtml += '<p>' + fullName + '</p>' + '\n'
        allReqsHtml += '</div>' + '\n'
        allReqsHtml += '<span class="btnFollow">' + '\n'
        allReqsHtml += '<button type="button"' + '\n'
        allReqsHtml += 'id="button-' + username + '" onClick="acceptFollow(this.id)" class="btn btn-primary btn-lg btnFollow">Accept</button>' + '\n'
        allReqsHtml += '</span>&nbsp&nbsp&nbsp'
        allReqsHtml += '<span class="btnFollow">'
        allReqsHtml += '<button type="button" class="btn btn-danger btn-lg btnFollow">Delete</button>'
        allReqsHtml += '</span></li><hr>'
    }

    if (n != 0)
        allReqsHtml = allReqsHtml.slice(0, -4)

    //console.log(allReqsHtml)   

    allReqsHtml += '</ul>'

    $("#insertAllFollowReqs").after(allReqsHtml)

    $("#numOfFollowReqs").text('(' + numActiveReqs + ')')
    isGeneratedFollowReqs = true;
}


function acceptFollow(buttonID) {
    var buttonIDParts = buttonID.split('-')
    var ReqUsername = buttonIDParts[1]

    var obj = {
        sentBy: ReqUsername,
        forWho: myUserName
    }

    console.log("OVO SALJEM ACCEPT")
    console.log(obj)

    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: REQUEST_SERVICE_URL + '/followReqs/accept',
        contentType: 'application/json',
        // dataType: 'JSON',
        data: JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            $("#" + buttonID).text("Done")
        },
        error: function (err) {
            console.log("Erorr at ajax ACCEPT call!")


        }
    })

    $("#" + buttonID).attr("disabled", 'true')


}

function setMyProfileHREF() {
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
            console.log("POSTAVLJAM HREF:")
            console.log(data)
            $("#myProfileHref").attr('href', "profile.html?" + data.username);
        },
        error: function (xhr) {
            console.log(xhr)
            redirectToLoginPage()
            console.log('ERROR IN WHOAMI');
        }
    })

}



