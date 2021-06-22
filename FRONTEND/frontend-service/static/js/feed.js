$(document).ready(function() {
    //alert("CONNECTED")
    loadFeedContent();
    loadStories()
})

function loadStories() {

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: CONTENT_SERVICE_URL + '/current/getallstoriesforfeed',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
           console.log('stories')
           console.log(data)
           fillStories(data)
        },
        error : function() {
            console.log("Erorr at ajax call!")
            
            
        }
    })
}

function fillStories(data) {

    var stories_here = $('#stories_here_plz')
    var html = ''
    var story_url

    data.forEach(function(s) {

        story_url = '../' + s.media.path

        html+= `<img class="card-img-center story-css" height="50px" src="${story_url}" alt="Card image cap"></img>`
    })
    stories_here.html(html)
}

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
           //generatePostsHTML(data)
           generatePostsHTML1(data)
        },
        error : function() {
            console.log("Erorr at ajax call!")
            
            
        }
    })

}

/*
<video width="320" height="240" controls>
  <source src="movie.mp4" type="video/mp4">
  <source src="movie.ogg" type="video/ogg">
Your browser does not support the video tag.
</video>
*/


function generatePostsHTML(allPosts){
    console.log(allPosts);
    var postsHTML = ''
    var img_url
    var media_string
    for(var i=0; i<allPosts.length;i++){
        img_url = '../' + allPosts[i].medias[0].path
        media_string = `<img class="card-img-center" src="${img_url}" alt="Card image cap">` + '\n'

        if (allPosts[i].medias[0].path.split(".")[1] != 'jpg' && allPosts[i].medias[0].path.split(".")[1] != 'png') {
            // then its a video
            media_string = `
                        <video width="100%" height="100%" controls>
                            <source src="${img_url}" type="video/mp4">                       
                            Your browser does not support the video tag.
                        </video>`
        }

        console.log(img_url)
        var img_url1 = '../../../BACKEND/content-service/temp/id-3/3d-render-banner-with-network-communications-low-poly-design.png'
        //console.log(img_url)
        // BACKEND/content-service/temp/id-3/3d-render-banner-with-network-communications-low-poly-design.png
        postsHTML += '<div class="one-feed text-center">' + '\n';
        postsHTML += '<div class="form-group">' + '\n';
        postsHTML += '<img src="img/avatar.png" class="profile_img">' + '\n';
                                            //dodati posle href
        postsHTML += '<a class="postedby" href="">' + allPosts[i].postedby + '</a>\n';
        postsHTML += '</div>' + '\n';

        postsHTML += '<div class="card">' + '\n';
            //          ovde treba ici allPosts[i].medias[j].path,  ubaciti jos jednu for petlju, ali trenutno putanje nisu sredjene
        postsHTML += media_string
        postsHTML += '<div class="card-body"></div>' + '\n';
        postsHTML += '<p class="card-text text-left description_part">' + allPosts[i].description + '</p>' + '\n';
        postsHTML += '<p class="text-left"></p>'
        postsHTML += `<button class="btn btn-info" id="like-${allPosts[i].ID}" onclick="likePost(this.id)">&nbsp&nbsp&nbsp&nbsp Like &nbsp&nbsp&nbsp&nbsp</button></p><hr>` + '\n';
        postsHTML += `<button class="btn btn-info" id="dislike-${allPosts[i].ID}" onclick="dislikePost(this.id)">&nbsp&nbsp&nbsp&nbsp Dislike &nbsp&nbsp&nbsp&nbsp</button></p><hr>` + '\n';
        
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
        postsHTML += `<button class="btn btn-primary postcommentbtn" id="post_comment_btn-${allPosts[i].ID}"> Post </button>`
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'

    }

    $("#insertFeedPosts").after(postsHTML)

}



function generatePostsHTML1(allPosts){
    console.log(allPosts);
    var postsHTML = ''
    var img_url
    var media_string

    for(var i=0; i<allPosts.length;i++){


        console.log(img_url)
        //var img_url1 = '../../../BACKEND/content-service/temp/id-3/3d-render-banner-with-network-communications-low-poly-design.png'
        //console.log(img_url)
        // BACKEND/content-service/temp/id-3/3d-render-banner-with-network-communications-low-poly-design.png
        postsHTML += '<div class="one-feed text-center">' + '\n';
        postsHTML += '<div class="form-group">' + '\n';
        postsHTML += '<img src="img/avatar.png" class="profile_img">' + '\n';
                                            //dodati posle href
        postsHTML += '<a class="postedby" href="">' + allPosts[i].postedby + '</a>\n';
        postsHTML += '</div>' + '\n';

        postsHTML += '<div class="card">' + '\n';
            //          ovde treba ici allPosts[i].medias[j].path,  ubaciti jos jednu for petlju, ali trenutno putanje nisu sredjene
            //CAROUSEL deo
        //postsHTML += media_string

        postsHTML +=  `<div id="demo-${i}" class="carousel slide" data-ride="carousel">`
        postsHTML +=  `<ul class="carousel-indicators">`

        console.log(allPosts[i].medias)
        for(var j=0; j < allPosts[i].medias.length ; j++) {

            if(j==0) {
                postsHTML += `<li data-target="#demo-${i}" data-slide-to="${j}" class="active"></li>`
            } else {
                postsHTML += `<li data-target="#demo-${i}" data-slide-to="${j}" ></li>`    
            }

        }

        postsHTML +=  `</ul>`

        // SLIDESHOW:

        postsHTML += `<div class="carousel-inner" style=" width:100%; height:auto !important;">`
        
        for(var g=0; g < allPosts[i].medias.length; g++) {
            img_url = '../' + allPosts[i].medias[g].path
            media_string = `<img width="100%" height="auto; !important" src="${img_url}" alt="Card image cap">`
    
            if (allPosts[i].medias[g].path.split(".")[1] != 'jpg' && allPosts[i].medias[g].path.split(".")[1] != 'png') {
                // then its a video
                media_string = `
                            <video width="70%" height="100%" controls>
                                <source src="${img_url}" type="video/mp4">                       
                                Your browser does not support the video tag.
                            </video>`
            }


            if(g==0) {
                postsHTML += `<div class="carousel-item active">`
                postsHTML +=  media_string
                postsHTML +=  `</div>`
            } else {
                postsHTML += `<div class="carousel-item">`
                postsHTML +=  media_string
                postsHTML +=  `</div>`
            }

        }

        postsHTML += `</div>`

        postsHTML += `
                            <a class="carousel-control-prev" href="#demo-${i}" data-slide="prev">
                                <span class="carousel-control-prev-icon"  style="background-color: black; border: 1px white;"></span>
                            </a>
                            <a class="carousel-control-next" href="#demo-${i}" data-slide="next" >
                                <span class="carousel-control-next-icon" style="background-color: black; border: 1px white;"></span>
                            </a>`

        postsHTML += `</div>`
            // CAROUSEL END --> card-body mi treba isti
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

        for(var h=0; h<allPosts[i].comments.length; h++){
            var comms = allPosts[i].comments
            postsHTML += '<tr>'
                                                //dodati posle href za profil
            postsHTML +='<td class="first_td"> <a class="commentedby" href="">' +  comms[h].postedby + '</a>'
            postsHTML += '</td>'
            postsHTML += '<td>' + comms[h].text +'</td>'
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
