const insertNotifications = $("#insertNotifications")
const numOfPostNotifs = $("#numOfPostNotifs")
var notificationList

function getNotifications() {
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: INTERACTION_SERVICE_URL + '/getunreadpostnotif',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            notificationList = data
			fillNotificationTable()
        },
        error : function(xhr, status, data) {
            alert('Error in getting my notifications')
        }
    })
}

function fillNotificationTable() {

    insertNotifications.html("")
    var text = ""
    var tableHTML = ""

    notificationList.forEach(function(n) {
        if (n.type == "LIKE") {
            text = `${n.notification.fromUser} liked your post.`
        } else if (n.type == "DISLIKE") {
            text = `${n.notification.fromUser} disliked your post.`
        } else {
            text = `${n.notification.fromUser} commented your post.`
        }

        tableHTML += `<tr><td class="save_td_th">${text}</td></tr>`;
    })

    insertNotifications.html(tableHTML)
    numOfPostNotifs.html("(" + notificationList.length + ")")
}

function readNotif() {
    $.ajax({
        type: 'POST',
        crossDomain: true,
        url: INTERACTION_SERVICE_URL + '/readpostnotif',
        contentType: 'application/json',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function () {
            alert("SUCCESS")
            insertNotifications.html("")
            numOfPostNotifs.html("(0)")
        },
        error: function () {
            alert("You already sent follow request to this account!")
        }
    })  
}