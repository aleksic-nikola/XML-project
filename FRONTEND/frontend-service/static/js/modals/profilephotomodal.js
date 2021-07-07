// which_image - koja slika
// like_button_div like button
// comments_body
// post_comment_button_div
// image_description
const which_image = $('#which_image')
const like_button_div = $('#like_button_div')
const comments_body = $('#comments_body')
const post_comment_button_div = $('#post_comment_button_div')
const image_description = $('#image_description')


function showImageModal(id, postList) {
	console.log('showing image ' + id)
	var post = undefined
	postList.forEach(function(p) {
		console.log('post loop ' + p.ID)
		console.log(id==p.ID)
		console.log(id)
		console.log(p.ID)
		if (p.ID == id) {
			post = p
			console.log('post is ' + post)
		}
	})
	var media = post.medias[0].path

	$("#user_posted_by").text(post.postedby)
	$("#user_posted_by").attr("href", "profile.html?" + post.postedby)

	console.log("==============> length of post (num of content): " + post.medias.length)

	if (post.medias[0].path.split(".")[1] != 'jpg' && post.medias[0].path.split(".")[1] != 'png' && post.medias.length == 1) {
		// only one VIDEO

		which_image.html(`

		<video width="100%" height="100%" controls alt="Card image cap" class="card-img-top modal-img">
			<source src="${media}" type="video/mp4">                       
			Your browser does not support the video tag.
		</video>
		
		`)

	} else {
		
		if(post.medias.length > 1) {
			// MORE VIDEOS or PICTURES (albums)

			console.log("There is more then 1 PICTURE or VIDEO")

			// carousel
			var postsHTML = ''


				postsHTML +=  `<div id="demo-${post.ID}" class="carousel slide" data-ride="carousel">`
				postsHTML +=  `<ul class="carousel-indicators">`
		
				for(var j=0; j < post.medias.length ; j++) {

					if(j==0) {
						postsHTML += `<li data-target="#demo-${post.ID}" data-slide-to="${j}" class="active"></li>`
					} else {
						postsHTML += `<li data-target="#demo-${post.ID}" data-slide-to="${j}" ></li>`    
					}
		
				}
				
		
				postsHTML +=  `</ul>`
		
				// SLIDESHOW:
		
				postsHTML += `<div class="carousel-inner" style=" width:100%; height:auto !important; padding-left:13%; padding-right:13%">`
				
				for(var g=0; g < post.medias.length; g++) {
					img_url = '../' + post.medias[g].path
					media_string = `<img width="100%" height="auto; !important" class="card-img-top modal-img"  src="${img_url}" alt="Card image cap">`
			
					if (post.medias[g].path.split(".")[1] != 'jpg' && post.medias[g].path.split(".")[1] != 'png') {
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
									<a class="carousel-control-prev" href="#demo-${post.ID}" data-slide="prev">
										<span class="carousel-control-prev-icon"  style="background-color: black; border: 1px white;"></span>
									</a>
									<a class="carousel-control-next" href="#demo-${post.ID}" data-slide="next" >
										<span class="carousel-control-next-icon" style="background-color: black; border: 1px white;"></span>
									</a>`
		
				postsHTML += `</div>`

			

			which_image.html(postsHTML)

		} else {
			// Only 1 picture

			which_image.html(`
	
			<img 
				class="card-img-top modal-img" 
				src="${media}" 
				alt="Card image cap">	
	
			`)
		}


	}


	console.log('give me image ' + post.medias[0])
	// like button id looks like  like-ID e.g. like-1 current user likes current post

	like_button_div.html(`<a href="" id="like-${post.ID}> Like &nbsp&nbsp&nbsp&nbsp</a>`)
	
	var comment_html = ''

	post.comments.forEach(function(c) {

		comment_html += `
			<tr>
			<td class = "first_td">
			<a class ="commentedby" href=""> ${c.postedby}: </a> 
			</td>
			<td>
			${c.text}
			</td>
			</tr>
		`
	})

	comments_body.html(comment_html)

	post_comment_button_div.html(`
	<button class="btn btn-primary postcommentbtn"
	 id="post_comment_btn-${post.ID}"
	  onclick="postComment(this.id)"> Post </button>
	`)

	image_description.html(post.description)

	//$('#gallery-modal').modal('show');
}