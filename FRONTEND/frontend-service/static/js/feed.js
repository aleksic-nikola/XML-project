var currentUserFeed;

$(document).ready(function() {
    //alert("CONNECTED")
    //loadFeedContent();
    //loadStories()

    getCurrentUserInformation()  // loadFeedContent i loadStories su prebaceni u success-u ove fje
    getNotifications()
    // pozovem fju gde uzimam trenutnog usera - izvadim mu liste (blacklist, graylist)
    // varijable globalna za usera --> currentUserFeed
    // kad to uzmem, u success --> loadStories, LoadFeedContent --> postedby na contentu i continue ako je u nekoj listi(mute/block) i tjt
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



var userStoriesMap ={} 

function fillStories(data) {

    var stories_here = $('#stories_here_plz')
    var html = ''
    var story_url
    console.log("PODACI IZ STORija: ")
    console.log(data)

    //var userStoriesMap ={} // key = username, value = [story1, story2...]

    data.forEach(function(s) {

        //story_url = '../' + s.media.path
        story_url = s.media.path
        postedBy = s.postedby
        console.log("SVE IZ S:")
        console.log(s)

        if(!(postedBy in userStoriesMap)){
            userStoriesMap[postedBy] = [];
        }
        userStoriesMap[postedBy].push(s)
        
        //html+= `<img class="card-img-center story-css" onClick="showStory('${story_url}', '${postedBy}')" height="80px" src="${story_url}" alt="Card image cap"></img>`
    })

    console.log("MAPAAA:")
    console.log(userStoriesMap)

    for ( [postedBy, s] of Object.entries(userStoriesMap)) {
        story_url = "./img/avatar.png"
        html+= `<img class="card-img-center story-css" onClick="showStoryModal('${postedBy}')" height="80px" src="${story_url}" alt="Card image cap"></img>`



    }
    console.log("**********************************")



    stories_here.html(html)
}


function showStory(imgPath, postedBy){
    console.log("Pogodjena funckija!!")
    console.log(imgPath)
    console.log(postedBy)
    //$("#insertStory").remove();
   // $("#forAddingDiv").after('<div id="insertStory"></div>')

    htmlStory = '<div id="insertStory">'
    htmlStory += `<img class="img-responsive" src="./img/avatar.png" style="max-height:500px; width=100%;" alt="Not found" >`
    htmlStory += `<p> <i> Posted by</i>:  <a href="profile.html?${postedBy}"><b> ${postedBy}</b></a>  </p>`
    htmlStory += '</div>'

    $("#insertStory").replaceWith(htmlStory)

    $("#btnTriggerStory").click();

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

        var dontshow = checkifPostedByIsInBlackListOrGrayList(allPosts[i].postedby)

        if(dontshow == true) {
            continue;
        }

        console.log(img_url)
        //var img_url1 = '../../../BACKEND/content-service/temp/id-3/3d-render-banner-with-network-communications-low-poly-design.png'
        //console.log(img_url)
        // BACKEND/content-service/temp/id-3/3d-render-banner-with-network-communications-low-poly-design.png
        postsHTML += '<div class="one-feed text-center">' + '\n';
        postsHTML += '<div class="form-group">' + '\n';
        postsHTML += '<a href="profile.html?'+ allPosts[i].postedby +'"><img src="img/avatar.png" class="profile_img"></a>' + '\n';
                                            //dodati posle href
        postsHTML += '<a class="postedby" href="profile.html?'+ allPosts[i].postedby +'">' + allPosts[i].postedby + '</a>';
        postsHTML += `<button class="btn btn-secondary" style="float: right;" id="${allPosts[i].ID}" onclick="setGlobalPostToReport(this.id)" data-toggle="modal" data-target="#reportModal">&nbsp&nbsp&nbsp&nbsp Report &nbsp&nbsp&nbsp&nbsp</button></p>`;

        postsHTML += '</div>' + '\n';

        postsHTML += '<div class="card" style="width: 100%;">' + '\n';
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

        if(allPosts[i].medias.length > 1) {

        postsHTML += `
                            <a class="carousel-control-prev" href="#demo-${i}" data-slide="prev">
                                <span class="carousel-control-prev-icon"  style="background-color: black; border: 1px white;"></span>
                            </a>
                            <a class="carousel-control-next" href="#demo-${i}" data-slide="next" >
                                <span class="carousel-control-next-icon" style="background-color: black; border: 1px white;"></span>
                            </a>`

        postsHTML += `</div>`
        }
        
            // CAROUSEL END --> card-body mi treba isti
        postsHTML += '<div class="card-body"></div>' + '\n';
        postsHTML += '<p class="card-text text-left description_part">' + allPosts[i].description + '</p>' + '\n';
        postsHTML += '<p class="text-left"></p>'

        postsHTML += `<div class="row">
            <div class="col">
                <button class="btn btn-success" id="like-${allPosts[i].ID}" onclick="likePost(this.id)">&nbsp&nbsp&nbsp&nbsp Like &nbsp&nbsp&nbsp&nbsp</button></p><hr>
            </div>
            <div class="col">
                <button class="btn btn-danger" id="dislike-${allPosts[i].ID}" onclick="dislikePost(this.id)">&nbsp&nbsp&nbsp&nbsp Dislike &nbsp&nbsp&nbsp&nbsp</button></p><hr>
            </div>
            <div class="col">
                <button class="btn btn-warning" id="${allPosts[i].ID}" onclick="setGlobalPostToSave(this.id)" data-toggle="modal" data-target="#saveModal">&nbsp&nbsp&nbsp&nbsp Save &nbsp&nbsp&nbsp&nbsp</button></p><hr>
            </div>
        </div>`

        //postsHTML += `<button class="btn btn-info" id="like-${allPosts[i].ID}" onclick="likePost(this.id)">&nbsp&nbsp&nbsp&nbsp Like &nbsp&nbsp&nbsp&nbsp</button></p><hr>` + '\n';
        //postsHTML += `<button class="btn btn-info" id="dislike-${allPosts[i].ID}" onclick="dislikePost(this.id)">&nbsp&nbsp&nbsp&nbsp Dislike &nbsp&nbsp&nbsp&nbsp</button></p><hr>` + '\n';
        //postsHTML += `<button class="btn btn-warning" id="${allPosts[i].ID}" onclick="setGlobalPostToSave(this.id)" data-toggle="modal" data-target="#saveModal">&nbsp&nbsp&nbsp&nbsp Save &nbsp&nbsp&nbsp&nbsp</button></p><hr>` + '\n';
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
            postsHTML +='<td class="first_td"> <a class="commentedby" href="profile.html?' + comms[h].postedby +'">' +  comms[h].postedby + '</a>'
            postsHTML += '</td>'
            postsHTML += '<td>' + comms[h].text +'</td>'
            postsHTML += '</tr>'
        }

        postsHTML += '</tbody>'
        postsHTML += '</div>'
        postsHTML += '</table><hr>'
        postsHTML += '<div class="writecomment_section"> <div class="form-group">'
                                                                    //ovaj id ce se verovatno menjati
        postsHTML += `<textarea class="form-control comment_textarea" id="insert_comment-${allPosts[i].ID}" rows="2"`
        postsHTML += 'placeholder="Add a comment..."></textarea>'
        postsHTML += `<button class="btn btn-primary postcommentbtn" id="post_comment_btn-${allPosts[i].ID}" onclick="postComment(this.id)"> Post </button>`
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'
        postsHTML += '</div>'

    }

    $("#insertFeedPosts").after(postsHTML)
}

function getCurrentUserInformation() {
    // getdata from profile
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getdata',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {

            currentUserFeed = data
            loadFeedContent();
            loadStories();

        },
        error : function(xhr, status, data) {
             alert('Error in getCurrentUserInformation')
        }
    })
}

function checkifPostedByIsInBlackListOrGrayList(username) {
    var myflag = false
    currentUserFeed.blacklist.forEach(function(b) {
        if(b.username == username) {
            myflag = true
            return
        }
    })

    currentUserFeed.graylist.forEach(function(g) {
        if(g.username == username) {
            myflag = true
            return
        }
    })

    return myflag
}

function showStoryModal(postedBy) {
	//console.log('showing image ' + id)
	var post = userStoriesMap[postedBy]
    console.log("DOBIO NA POCETKU:")
    console.log(post)
	// postList.forEach(function(p) {
	// 	console.log('post loop ' + p.ID)
	// 	console.log(id==p.ID)
	// 	console.log(id)
	// 	console.log(p.ID)
	// 	if (p.ID == id) {
	// 		post = p
	// 		console.log('post is ' + post)
	// 	}
	// })

	//var media = post.medias[0].path

	console.log("==============> length of post (num of content): " + post.length)

    if(post[0].media.path.split(".")[1] != 'jpg' && post[0].media.path.split(".")[1] != 'jpeg' && post[0].media.path.split(".")[1] != 'jfif' && post[0].media.path.split(".")[1] != 'png' && post.length == 1) {
		// only one VIDEO

		$("#insertStory").replaceWith(`
        <div id="insertStory">
		<video width="100%" height="100%" controls alt="Card image cap" class="card-img-top modal-img">
			<source src="${post[0].media.path}" type="video/mp4">                       
			Your browser does not support the video tag.
		</video>
        <p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>
        </div>
		
		`)

	}
    else{
            //                           
        if(post.length > 1) {
            // MORE VIDEOS or PICTURES (albums)

            console.log("There is more then 1 PICTURE or VIDEO")

            // carousel
            var postsHTML = '<div id="insertStory">'


                postsHTML +=  `<div id="demo-${post[0].postedby}" class="carousel slide" data-ride="carousel">`
                postsHTML +=  `<ul class="carousel-indicators">`
        
                for(var j=0; j < post.length ; j++) {

                    if(j==0) {
                        postsHTML += `<li data-target="#demo-'${post[0].postedby}'" data-slide-to="${j}" class="active"></li>`
                    } else {
                        postsHTML += `<li data-target="#demo-'${post[0].postedby}'" data-slide-to="${j}" ></li>`    
                    }
        
                }
                
        
                postsHTML +=  `</ul>`
        
                // SLIDESHOW:
        
                postsHTML += `<div class="carousel-inner" style=" width:100%; height:auto !important; padding-left:13%; padding-right:13%">`
                
                for(var g=0; g < post.length; g++) {
                    img_url = '../' + post[g].media.path
                    media_string = `<img width="100%" height="auto; !important" class="card-img-top modal-img"  src="${img_url}" alt="Card image cap">`
                    console.log("OVO JE SPLITOVANO: ")
                    console.log(post[g].media.path.split(".")[1])
                    if (post[g].media.path.split(".")[1] != 'jpg' && post[g].media.path.split(".")[1] != 'png' && post[g].media.path.split(".")[1] != 'jfif') {
						console.log("USAO DA JE VIDEO!!!")
                        // then its a video
						media_string = `
									<video width="100px" height="100%" controls class="card-img-top modal-img" >
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
                                    <a class="carousel-control-prev" href="#demo-${post[0].postedby}" data-slide="prev">
                                        <span class="carousel-control-prev-icon"  style="background-color: black; border: 1px white;"></span>
                                    </a>
                                    <a class="carousel-control-next" href="#demo-${post[0].postedby}" data-slide="next" >
                                        <span class="carousel-control-next-icon" style="background-color: black; border: 1px white;"></span>
                                    </a>`
        
                postsHTML += `</div>`
                postsHTML += `<p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>`
                postsHTML += "</div>"

            $("#insertStory").replaceWith(postsHTML)

            $("#btnTriggerStory").click();

            //which_image.html(postsHTML)

        } else {
            // Only 1 picture

            $("#insertStory").replaceWith(`
            <div id="insertStory">
            <img 
                class="card-img-top modal-img" 
                src="${post[0].media.path}" 
                alt="Card image cap">
            <p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>	
            "</div>"
            `)
            $("#btnTriggerStory").click();
        }

    }


    
    
    


	
	// like button id looks like  like-ID e.g. like-1 current user likes current post

	

	//$('#gallery-modal').modal('show');
}