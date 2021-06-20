// which_image - koja slika
// like_button_div like button
// comments_body
// post_comment_button_div
const which_image = $('#which_image')
const like_button_div = $('#like_button_div')
const comments_body = $('#comments_body')
const post_comment_button_div = $('#post_comment_button_div')


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
	which_image.html(`
	
	<img 
	class="card-img-center modal-img" 
	src="${media}" 
	alt="Card image cap">	
	`)
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

	//$('#gallery-modal').modal('show');
}