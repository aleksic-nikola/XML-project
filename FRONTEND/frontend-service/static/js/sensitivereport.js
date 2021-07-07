var postForReport

function setGlobalPostToReport(id) {
	postForReport = id
    document.getElementById('note').value = "";

    if (window.location.href.includes("profile")) {
		postForReport = currentOpenedPost.split("-")[1]
	}
}

function createSensitiveReport() {

	reason = $("#note").val()

	var obj = {
        post_id: postForReport,
        note : reason
    }

	console.log(obj)
	
	$.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: REQUEST_SERVICE_URL + '/sensitiveContentReqs/create',
	    contentType: 'application/json',
		data: JSON.stringify(obj),
	    beforeSend: function (xhr) {
			xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function () {
			alert('successfully sent report request')
	    },
	    error: function () {
			alert('cant send report request')
	    }
	})
}