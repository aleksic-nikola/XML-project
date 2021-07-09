function showPhotos(postList, divName) {
	console.log('I am printing the images')
	var posts = "";
	//console.log('post number is ' + postList.length)
	if (postList == undefined) {
		$("#" + divName).html("")
		return
	}

	postList.forEach(function (post) {
		//console.log(post)
		/*posts += `<div class="col-md-6 col-lg-4">
			    <div class="card border-0 transform-on-hover">
				<a class="lightbox">
				    <img src="${post.medias[0].path}" id="post-${post.ID}" class="card-img-top gallery-item">
				</a>
			    </div>
		      </div>`
		console.log('show me image ' + post.medias[0])

		*/

		if (post.medias[0].path.split(".")[1] != 'jpg' && post.medias[0].path.split(".")[1] != 'png') {
			// video
			posts += 
			`
			<div class="col-md-6 col-lg-4">
				<div class="card border-0 transform-on-hover">
					<a class="lightbox">
						<video width="100%" height="100%" id="post-${post.ID}" class="card-img-top gallery-item">
							<source src="${post.medias[0].path}"  type="video/mp4">                       
							Your browser does not support the video tag.
						</video>
					</a>
				</div>
		 	</div>	
			
			`

		} else {
			// img
			posts += 
				`<div class="col-md-6 col-lg-4">
					<div class="card border-0 transform-on-hover">
						<a class="lightbox">
							<img src="${post.medias[0].path}" id="post-${post.ID}" class="card-img-top gallery-item">
						</a>
					</div>
		 	 	</div>`
		}



	})
	//console.log(posts)
	//$("#postsHere").html(posts)
	$("#" + divName).html(posts)

}

function showArchiveStories(postList, divName) {
	console.log('I am printing the stories...........')
	var stories = "";
	//console.log('post number is ' + postList.length)
	if (postList == undefined) {
		$("#" + divName).html("")
		return
	}

	postList.forEach(function (post) {
		//console.log(post)
		/*posts += `<div class="col-md-6 col-lg-4">
			    <div class="card border-0 transform-on-hover">
				<a class="lightbox">
				    <img src="${post.medias[0].path}" id="post-${post.ID}" class="card-img-top gallery-item">
				</a>
			    </div>
		      </div>`
		console.log('show me image ' + post.medias[0])

		*/

		if (post.media.path.split(".")[1] != 'jpg' && (post.media.path.split(".")[1] != 'png')) {
			// video
			stories += 
			`
			<div class="col-md-6 col-lg-4">
				<div class="card border-0 transform-on-hover">
					<a class="lightbox">
						<video width="100%" height="100%" id="post-${post.ID}" class="card-img-top gallery-item archive-stories">
							<source src="${post.media.path}"  type="video/mp4">                       
							Your browser does not support the video tag.
						</video>
					</a>
				</div>
		 	</div>	
			
			`

		} else {
			// img
			stories += 
				`<div class="col-md-6 col-lg-4">
					<div class="card border-0 transform-on-hover">
						<a class="lightbox">
							<img src="${post.media.path}" id="post-${post.ID}" class="card-img-top gallery-item archive-stories">
						</a>
					</div>
		 	 	</div>`
		}



	})
	//console.log(posts)
	//$("#postsHere").html(posts)
	$("#" + divName).html(stories)

}


