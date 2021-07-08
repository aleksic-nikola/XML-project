function setLikedPostsActive() {
    currentPill = "liked"

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/current/likedposts',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            postListLiked = data
            showPhotos(data, "postsHereLiked")
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })
}

function setDislikedPostsActive() {
    currentPill = "disliked"

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/current/dislikedposts',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            postListDisliked = data
            showPhotos(data, "postsHereDisliked")
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })
}