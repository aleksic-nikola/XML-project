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
	    },
	    error: function () {
		alert('cant dislike this tntntntn')
    
    
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
		},
		error: function () {
		    alert('cant post this tntntntn')
	
	
		}
	})    
}

function setGlobalPostToSave(id) {
	postForSave = id
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

function fillTable(collections) {
	
	tableCollection.html("")
    var insertHtml = ""

    collections.name.forEach(function(col) {
        insertHtml += `<tr><td class="collectionTable">${col}</td></tr>`
    })

    tableCollection.html(insertHtml)
}
    