$(document).ready(function() {
    //alert("CONNECTED")
    loadFeedContent();
})


function loadFeedContent(){

     $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/current/getAllPostsForFeed',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
           generatePostsHTML(data)
        },
        error : function() {
            console.log("Erorr at ajax call!")
            
            
        }
    })

}


function generatePostsHTML(allPosts){
    console.log(allPosts);
    var postsHTML = ''
    for(var i=0; i<allPosts.length;i++){
        postsHTML += '<div class="one-feed text-center">' + '\n';
        postsHTML += '<div class="form-group">' + '\n';
        postsHTML += '<img src="img/avatar.png" class="profile_img">' + '\n';
                                            //dodati posle href
        postsHTML += '<a class="postedby" href="">' + allPosts[i].postedby + '</a>\n';
        postsHTML += '</div>' + '\n';

        postsHTML += '<div class="card">' + '\n';
            //          ovde treba ici allPosts[i].medias[j].path,  ubaciti jos jednu for petlju, ali trenutno putanje nisu sredjene
        postsHTML += '<img class="card-img-center" src="img/test.png" alt="Card image cap">' + '\n';
        postsHTML += '<div class="card-body"></div>' + '\n';
        postsHTML += '<p class="card-text text-left description_part">' + allPosts[i].description + '</p>' + '\n';
        postsHTML += '<p class="text-left"></p>'
        postsHTML += '<a href=""> Like &nbsp&nbsp&nbsp&nbsp</a></p><hr>' + '\n';
        postsHTML += '<h5 class="text-left">Comments:</h5><hr>' + '\n';
        postsHTML += '<table class="table" id="mydatatable">'
        postsHTML += '<thead class="hideheader">'
        postsHTML += '<tr><th scope="col" id="commented_by">CommentedBy</th><th scope="col" id="comment">Comment</th></tr></thead>'
        postsHTML += '<div class="bodycontainer scrollable">'
        postsHTML += '<tbody class="text-left">'

        for(var j=0; j<allPosts[i].comments.length; j++){
            var comms = allPosts[i].comments
            postsHTML += '<tr>'
                                                //dodati posle href za profil
            postsHTML +='<td class="first_td"> <a class="commentedby" href="">' +  comms[j].postedby + '</a>'
            postsHTML += '</td>'
            postsHTML += '<td>' + comms[j].text +'</td>'
            postsHTML += '</tr>'
        }

        postsHTML += '</tbody>'
        postsHTML += '</div>'
        postsHTML += '</table><hr>'
        postsHTML += '<div class="writecomment_section"> <div class="form-group">'
                                                                    //ovaj id ce se verovatno menjati
        postsHTML += '<textarea class="form-control comment_textarea" id="insert_comment" rows="2"'
        postsHTML += 'placeholder="Add a comment..."></textarea>'
        postsHTML += '<button class="btn btn-primary postcommentbtn" id="post_comment_btn"> Post </button>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'

    }

    $("#insertFeedPosts").after(postsHTML)

    



}