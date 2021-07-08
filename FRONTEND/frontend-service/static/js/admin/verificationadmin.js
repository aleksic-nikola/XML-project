const verification_body = $('#verification_admin_body')
const verphotobody = $('#verPhotoBody')
var verification_list = []

$(document).ready(function() {

	getVerificationRequests()

})

function getVerificationRequests() {

	$.ajax({
		type: 'GET',
		crossDomain: true,
		url: REQUEST_SERVICE_URL + '/verificationReqs',
		contentType: 'application/json',
		dataType: 'JSON',
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success : function(data) {
            verification_list = data
		    fillVerificationRequestsTable(data)

		},
		error : function() {
		    console.log("Erorr at ajax call!")
		}
	    })

}

function fillVerificationRequestsTable(data) {

	verification_body.html('')
	var html = ''
	data.forEach(function(vr) {

		html += `
            <tr>
                <td> <a href="profile.html?${vr.request.sentby}"> ${vr.request.sentby} </a> </td>
                <td> ${vr.name}</td>
                <td> ${vr.lastname}</td>
                <td> ${convertIntegerToCategory(vr.verifiedtype)} </td>
                <td> <button class="btn btn-info" id="button-${vr.request.ID}" data-toggle="modal" onclick="editVerPhotoModal(this.id)" data-target="#verPhotoModal"> Show photo </button></td>
                <td> <button class="btn btn-success" id="accept-${vr.request.ID}" onclick="acceptVerRequest(this.id)"> Accept </button></td>
                <td> <button class="btn btn-warning" id="reject-${vr.request.ID}" onclick="rejectVerRequest(this.id)"> Reject </button></td>
			</tr>`


	})

	verification_body.html(html)
	
}

function editVerPhotoModal(id) {

    var split = id.split("-")[1]

    verification_list.forEach(function(v) {

        if (v.request.ID == split) {
            //alert('hey')
            $('#verPhotoBody').html(`<img src="${v.image}" style="max-width: 100%; max-height: 100%">`)
        }

    })

}

function acceptVerRequest(id) {

    var split = id.split("-")[1]

    var obj = {
        "id" :  parseInt(split),
        "new_status" : 1
    }

    $.ajax({
		type: 'POST',
		crossDomain: true,
		url: REQUEST_SERVICE_URL + '/verificationReqs/update',
		contentType: 'application/json',
		dataType: 'JSON',
        data : JSON.stringify(obj),
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success : function(data) {
            alert('Successfully accepted request')
            getVerificationRequests()
		},
		error : function() {
		    console.log("Failed to accept request!")
		}
	    })

}

function rejectVerRequest(id) {

    var split = id.split("-")[1]

    var obj = {
        "id" : parseInt(split), 
        "new_status" : 2
    }

    $.ajax({
		type: 'GET',
		crossDomain: true,
		url: REQUEST_SERVICE_URL + '/verificationReqs/update',
		contentType: 'application/json',
		dataType: 'JSON',
        data : JSON.stringify(obj),
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success : function(data) {
            alert('Successfully rejected request')
            getVerificationRequests()
		},
		error : function() {
		    console.log("Failed to reject request!")
		}
	    })

}

function convertIntegerToCategory(num) {

	if (num == '0') {
		return 'INFLUENCER'
	} else if (num == 1) {
        return 'SPORTS'
	} else if (num == 2) {
		return 'MEDIA'
	} else if (num == 3) {
		return 'BUSINESS'
	} else if (num == 4) {
		return 'BRAND'
	} else {
        return 'ORGANIZATION'
	}

}