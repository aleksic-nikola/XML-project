$(document).ready(function() {
    //$("#deleteCollection").hide()
})

const select_group = $("#select_group")
var currentPill = "posts"

function getAllCollections() {

    $("#postsHereSaved").html("")
    currentPill = "saved"
   
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
                fillDropdown(data)
            } 
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })
}

function fillDropdown(collections) {

    select_group.html("")
    var insertHtml = `<option id="select_none"></option>`

    collections.name.forEach(function(col) {
        insertHtml += `<option id="${col}">${col}</option>`
    })

    select_group.html(insertHtml)
}


select_group.change(function(){

    var collection = select_group.children(":selected").attr("id");

    if (collection == "select_none") {
        $("#deleteCollection").hide()

    } else {
        $("#deleteCollection").show()
        $("#deleteCollection").html("Delete collection")
    }

    findPostsByCollection(collection)

}) 

function findPostsByCollection(collection) {
    
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/getFavouritePosts/' + collection,
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            postListSaved = data
            showPhotos(data, "postsHereSaved")
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })

}

function setPostsActive() {
    currentPill = "posts"
}

function removePostFromCollection() {
    var parts = $("#removeFromCollection").html().split(" ")
    var post_id_parts = parts[1].split("-")
    var collection = parts[3]

    var postID = post_id_parts[1]

    var obj = {
        collection_name : collection,
        post_id : parseInt(postID)
    }

    $.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: PROFILE_SERVICE_URL + '/deletepostfromcollection',
	    contentType: 'application/json',
		data: JSON.stringify(obj),
	    beforeSend: function (xhr) {
			xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function () {
			alert('Successfully deleted post from collection')
            window.location.reload()
	    },
	    error: function () {
			alert('Error in deleting post from collection')
	    }
	})



}

function deleteCollection() {
    var collection = select_group.children(":selected").attr("id");

    var obj = {
        collection_name : collection,
        post_id : 0
    }

    $.ajax({
	    type: 'POST',
	    crossDomain: true,
	    url: PROFILE_SERVICE_URL + '/deletecollection',
	    contentType: 'application/json',
		data: JSON.stringify(obj),
	    beforeSend: function (xhr) {
			xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
	    },
	    success: function () {
			alert('Successfully deleted collection')
            window.location.reload()
	    },
	    error: function () {
			alert('Error in deleting collection')
	    }
	})
}