var postForSave
const tableCollection = $("#tableInsert")

function likePost(id) {

	var split = id.split("-")[1]
	console.log("I like post " + split)

	if (window.location.href.includes("profile")) {
		// na profilu smo
		split = currentOpenedPost.split("-")[1]
	}
    
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
			createNotification(split, "LIKE")
	    },
	    error: function () {
			alert('you have already liked this post')    
	    }
	})

	console.log('here')
}

function createNotification(postId, type) {
	var obj = {
		post_id : parseInt(postId),
		type : type
	}

	$.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: INTERACTION_SERVICE_URL + '/postnotification/add',
	    contentType: 'application/json',
		//dataType: 'json',
		data : JSON.stringify(obj),
	    beforeSend: function (xhr) {
		xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function () {
			alert('notification created')
	    },
	    error: function () {
			alert('error creating like notif')    
	    }
	})

}

function dislikePost(id) {

	console.log('Dislike')

	var split = id.split("-")[1]
	console.log("I like post " + split)

	if (window.location.href.includes("profile")) {
		// na profilu smo
		split = currentOpenedPost.split("-")[1]
	}
    
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
			createNotification(split, "DISLIKE")
	    },
	    error: function () {
			alert('you have already disliked this post')   
	    }
	})

	console.log('here')
}


function postComment(id) {

	var split = id.split("-")[1]
	var field = $('#insert_comment-' + split)
    
	console.log('post is ' + split)
	console.log(field.val())
    
	$.ajax({
		type: 'POST',
		crossDomain: true,
		url: CONTENT_SERVICE_URL + '/comment/' + split,
		contentType: 'application/json',
		data : JSON.stringify({"text" : field.val()}),
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success: function (data) {
		    alert('successfully posted comment')
			createNotification(split, "COMMENT")
		},
		error: function () {
		    alert('cant post this tntntntn')
	
	
		}
	})    
}

function setGlobalPostToSave(id) {
	postForSave = id

	if (window.location.href.includes("profile")) {
		// na profilu smo
		postForSave = currentOpenedPost.split("-")[1]
	}

	$.ajax({
        type: 'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getallcollections',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
			if (data.name != null) {
            	fillTable(data)
			}
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })
}


function savePost() {

	collectionName = $("#collection_field").val()

	var obj = {
        collection_name: collectionName,
        post_id : parseInt(postForSave)
    }

	console.log(obj)
	
	$.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: PROFILE_SERVICE_URL + '/addposttofavourites',
	    contentType: 'application/json',
		data: JSON.stringify(obj),
	    beforeSend: function (xhr) {
			xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function () {
			alert('successfully saved post to favourites')
	    },
	    error: function () {
			alert('cant save post to favourites')
	    }
	})
}

function savePostOnProfile() {

	collectionName = $("#collection_field_profile").val()

	var obj = {
        collection_name: collectionName,
        post_id : parseInt(postForSave)
    }

	console.log(obj)
	
	$.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: PROFILE_SERVICE_URL + '/addposttofavourites',
	    contentType: 'application/json',
		data: JSON.stringify(obj),
	    beforeSend: function (xhr) {
			xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function () {
			alert('successfully saved post to favourites')
	    },
	    error: function () {
			alert('cant save post to favourites')
	    }
	})
}

function fillTable(collections) {
	
	tableCollection.html("")
    var insertHtml = ""

    collections.name.forEach(function(col) {
        insertHtml += `<tr><td class="save_td_th">${col}</td></tr>`
    })

    tableCollection.html(insertHtml)
}
    