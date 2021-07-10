var userStoriesMap = {}
var stories_here
var globalCloseFriendFlag


function fillStories(data) {
    var isOnProfile = false
    stories_here = $('#stories_here_plz')
    if (window.location.href.includes("profile")) {
        stories_here = $('#highlightedStoriesHere')
        isOnProfile = true
    }

    var html = ''
    var story_url
    console.log("PODACI IZ STORija: ")
    console.log(data)

    //var userStoriesMap ={} // key = username, value = [story1, story2...]
    if (data == undefined) {
        return 
    }
    data.forEach(function (s) {

        //story_url = '../' + s.media.path
        story_url = s.media.path
        postedBy = s.postedby
        //console.log("SVE IZ S:")
        //console.log(s)
        if (window.location.href.includes("profile") == false) {

            var dontshow = checkifPostedByIsInBlackListOrGrayList(s.postedby)
    
            if(dontshow == true) {
                return;
            }
        }

        if (s.isforclosefriendsonly == true) {
            checkIfCloseFriends(s.postedby)

            if (globalCloseFriendFlag == false) {
                return
            }
        }

        if (!(postedBy in userStoriesMap)) {
            userStoriesMap[postedBy] = [];
        }
        userStoriesMap[postedBy].push(s)

        //html+= `<img class="card-img-center story-css" onClick="showStory('${story_url}', '${postedBy}')" height="80px" src="${story_url}" alt="Card image cap"></img>`
    })

    console.log("MAPAAA:")
    console.log(userStoriesMap)

    for ([postedBy, s] of Object.entries(userStoriesMap)) {
        story_url = "./img/avatar.png"
        if (isOnProfile) {
            story_url = "./img/highlights.png"

        }


        html += `<img class="card-img-center story-css" onClick="showStoryModal('${postedBy}')" height="80px" src="${story_url}" alt="Card image cap"></img>`



    }
    console.log("**********************************")



    stories_here.html(html)
}


function showStoryModal(postedBy) {
    //console.log('showing image ' + id)
    var post = userStoriesMap[postedBy]
    console.log("DOBIO NA POCETKU:")
    console.log(post)

    console.log("==============> length of post (num of content): " + post.length)

    if (post[0].media.path.split(".")[1] != 'jpg' && post[0].media.path.split(".")[1] != 'jpeg' && post[0].media.path.split(".")[1] != 'jfif' && post[0].media.path.split(".")[1] != 'png' && post.length == 1) {
        // only one VIDEO


        /*
        $("#insertStory").replaceWith(`
        <div id="insertStory">
        <video width="100%" height="100%" controls alt="Card image cap" class="card-img-top modal-img">
            <source src="${post[0].media.path}" type="video/mp4">                       
            Your browser does not support the video tag.
        </video>
        <p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>
        </div>
    	
        `)
        */

        divId = "insertStory"
        if (window.location.href.includes("profile")) {
            divId = "insertStoryHighlight"
        }

        var htmlOneVideo = `
            <div id="${divId}">
            <video width="100%" height="100%" controls alt="Card image cap" class="card-img-top modal-img">
                <source src="${post[0].media.path}" type="video/mp4">                       
                Your browser does not support the video tag.
            </video>
            <p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>
            </div>
        
        `

        if (window.location.href.includes("profile")) {
            $("#insertStoryHighlight").replaceWith(htmlOneVideo)
            $('#highlightedStoryModal').modal('show');
        }
        else {
            $("#insertStory").replaceWith(htmlOneVideo)
            $("#btnTriggerStory").click();
        }






    }
    else {
        //                           
        if (post.length > 1) {
            // MORE VIDEOS or PICTURES (albums)

            console.log("There is more then 1 PICTURE or VIDEO")

            // carousel
            var postsHTML = '<div id="insertStory">'


            postsHTML += `<div id="demo-${post[0].postedby}" class="carousel slide" data-ride="carousel">`
            postsHTML += `<ul class="carousel-indicators">`

            for (var j = 0; j < post.length; j++) {

                if (j == 0) {
                    postsHTML += `<li data-target="#demo-'${post[0].postedby}'" data-slide-to="${j}" class="active"></li>`
                } else {
                    postsHTML += `<li data-target="#demo-'${post[0].postedby}'" data-slide-to="${j}" ></li>`
                }

            }


            postsHTML += `</ul>`

            // SLIDESHOW:

            postsHTML += `<div class="carousel-inner" style=" width:100%; height:auto !important; padding-left:13%; padding-right:13%">`

            for (var g = 0; g < post.length; g++) {
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

                if (g == 0) {
                    postsHTML += `<div class="carousel-item active">`
                    postsHTML += media_string
                    postsHTML += `</div>`
                } else {
                    postsHTML += `<div class="carousel-item">`
                    postsHTML += media_string
                    postsHTML += `</div>`
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



            //stories_here.replaceWith(postsHTML)

            if (window.location.href.includes("profile")) {
                $("#insertStoryHighlight").replaceWith(postsHTML)
                $('#highlightedStoryModal').modal('show');
            }
            else {
                $("#insertStory").replaceWith(postsHTML)
                $("#btnTriggerStory").click();
            }




            //which_image.html(postsHTML)

        } else {
            // Only 1 picture
            /*
            $("#insertStory").replaceWith(`
            <div id="insertStory">
            <img 
                class="card-img-top modal-img" 
                src="${post[0].media.path}" 
                alt="Card image cap">
            <p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>	
            "</div>"
            `)
            */

            $("#btnTriggerStory").click();

            divId = "insertStory"
            if (window.location.href.includes("profile")) {
                divId = "insertStoryHighlight"
            }

            htmlOneStory = `<div id="${divId}">
                <img 
                class="card-img-top modal-img" 
                src="${post[0].media.path}" 
                alt="Card image cap">
                <p> <i> Posted by</i>:  <a href="profile.html?${post[0].postedby}"><b> ${post[0].postedby}</b></a>  </p>	
                "</div>"
            `



            if (window.location.href.includes("profile")) {
                $("#insertStoryHighlight").replaceWith(htmlOneStory)
                $('#highlightedStoryModal').modal('show');
            }
            else {
                $("#insertStory").replaceWith(htmlOneStory)
                $("#btnTriggerStory").click();
            }






        }

    }



    // like button id looks like  like-ID e.g. like-1 current user likes current post



    //$('#gallery-modal').modal('show');
}


function checkIfCloseFriends(postedby) {

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/checkIfCloseFriends/' + postedby,
        contentType: 'application/json',
        //dataType: 'JSON',
        async: false,
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data, textStatus, xhr) {
            console.log("OVO U CHECKA")

            if (xhr.status == 200) {
                globalCloseFriendFlag = false
            }
            else {
                globalCloseFriendFlag = true
            }
        },
        error: function () {
            console.log("Erorr at ajax call!")
            alert("Error ajax")
            globalCloseFriendFlag = false

        }
    })

}