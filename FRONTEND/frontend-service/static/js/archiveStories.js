
function setArchiveStories() {
    
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/current/archiveStories',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            console.log("OVO SU MOJI STORIJI>>")
            console.log(data)
            storiesArchive = data

            
            showArchiveStories(data, "archiveStoriesHere")
            
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })
}
//getAllStoriesByUser/{username}

function setHighlightedStoriesProfile(username) {
    
    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/getAllStoriesByUser/' + username,
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            for(var i=0; i<data.length; i++){
                if(data[i].ishighlighted){
                    storiesHighlighted.push(data[i])
                }
            }
            fillStories(storiesHighlighted)
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })
}




function setArchivePillActive(){
    currentPill = "archiveStories"
}


function setAsHighlightStory(){
    var id = currentOpenedPost.split("-")[1]
    var obj = {
        story_id : parseInt(id)
    }

     $.ajax({
        type: 'POST',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/story/setHighlightedOn',
        contentType: 'application/json',
        //dataType: 'JSON',
        data : JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function() {
            alert("Successfully added story to highlight")
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })

}

function setAsHighlightStoryOff(){
    var id = currentOpenedPost.split("-")[1]
    var obj = {
        story_id : parseInt(id)
    }

     $.ajax({
        type: 'POST',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/story/setHighlightedOff',
        contentType: 'application/json',
        //dataType: 'JSON',
        data : JSON.stringify(obj),
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function() {
            alert("Successfully removed story from highlight")
        },
        error : function() {
            console.log("Erorr at ajax call!")
        }
    })   



}