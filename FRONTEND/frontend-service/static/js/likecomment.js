function likePost(id) {

	var split = id.split("-")[1]
	console.log("I like post " + split)
    
	$.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: CONTENT_SERVICE_URL + '/like/' + split,
	    contentType: 'application/json',
	    beforeSend: function (xhr) {
		xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function (data) {
		alert('successfully liked')
	    },
	    error: function () {
		alert('cant like this tntntntn')
    
    
	    }
	})

	console.log('here')
}

function dislikePost(id) {

	console.log('Dislike')

	var split = id.split("-")[1]
	console.log("I like post " + split)
    
	$.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: CONTENT_SERVICE_URL + '/dislike/' + split,
	    contentType: 'application/json',
	    beforeSend: function (xhr) {
		xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function (data) {
		alert('successfully disliked')
	    },
	    error: function () {
		alert('cant dislike this tntntntn')
    
    
	    }
	})

	console.log('here')
}