function showPhotos() {
	console.log('I am printing the images')
	var posts = "";
	console.log('post number is ' + postList.length)
	postList.forEach(function(post) {
	    console.log(post)
	    posts += `<div class="col-md-6 col-lg-4">
			    <div class="card border-0 transform-on-hover">
				<a class="lightbox">
				    <img src="${post.medias[0].path}" id="post-${post.ID}" class="card-img-top gallery-item">
				</a>
			    </div>
		      </div>`
	    console.log('show me image ' + post.medias[0])	      
	})
	console.log(posts)
	$("#postsHere").html(posts)
}


