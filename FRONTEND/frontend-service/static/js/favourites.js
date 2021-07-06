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
            fillDropdown(data)
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
    var post_id = parts[1]
    var collection = parts[3]

    alert(post_id + "  " +collection)
}